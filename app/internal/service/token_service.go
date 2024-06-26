package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/YuraLk/drone_calc/backend/internal/configs"
	"github.com/YuraLk/drone_calc/backend/internal/database/postgres"
	"github.com/YuraLk/drone_calc/backend/internal/dtos/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type TokenService struct{}

type Claims struct {
	// Объединение типов в одном
	user.UserDTO
	jwt.RegisteredClaims
}

func (TokenService) Generate(dto user.UserDTO) (string, string, error) {
	// fmt.Print(data)
	accessClaims := Claims{
		UserDTO: dto,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // 15 минут
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(configs.Cfg.JWT.JWTAccessKey))
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}
	refreshClaims := Claims{
		UserDTO: dto,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 14)), // 14 дней
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(configs.Cfg.JWT.JWTRefreshKey))
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

// Проверка подлинности и срока годности токенов
func (TokenService) ValidateRefresh(tokenString string) (*Claims, error) {
	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Cfg.JWT.JWTRefreshKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	// Получение пользовательских данных из токена
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("error of getting user data from token")
	}

	return claims, nil
}

func (TokenService) ValidateAccess(tokenString string) (*Claims, error) {
	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Cfg.JWT.JWTAccessKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	// Получение пользовательских данных из токена
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("error of getting user data from token")
	}

	return claims, nil
}

// Ищем старую сессию или создаем новую
func (TokenService) Save(ctx context.Context, refreshToken string, userId uint, device string) error {
	// Начало транзакции
	tx, err := postgres.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Откат транзакции в случае ошибки

	var existing_session_id uint

	if err := tx.QueryRow(ctx, "SELECT id FROM sessions WHERE user_id = $1 AND device = $2 LIMIT 1", userId, device).Scan(&existing_session_id); err != nil {
		// Если ошибка связана с провалом поиска записи
		if err == pgx.ErrNoRows {
			_, err := tx.Exec(ctx, "INSERT INTO sessions (refresh_token, user_id, device) VALUES ($1, $2, $3)", refreshToken, userId, device)

			if err != nil {
				return fmt.Errorf("failed to insert user: %w", err)
			}
		} else {
			return fmt.Errorf("failed to query user: %w", err)
		}
		// Если запись была найдена, перезаписываем токен и сохраняем
	} else {
		_, err := tx.Exec(ctx, "UPDATE sessions SET refresh_token = $1, updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') WHERE id = $2", refreshToken, existing_session_id)

		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	}

	// Коммит транзакции
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (TokenService) Remove(ctx context.Context, token string) error {

	_, err := postgres.DB.Exec(ctx, "DELETE FROM sessions WHERE refresh_token = $1", token)
	if err != nil {
		return err
	}

	return nil
}
