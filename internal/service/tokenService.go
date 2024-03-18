package service

import (
	"errors"
	"time"

	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/dtos/user_dtos"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Claims struct {
	// Объединение типов в одном
	user_dtos.UserDTO
	jwt.StandardClaims
}

func GenerateTokens(dto user_dtos.UserDTO) (string, string, error) {
	// fmt.Print(data)
	accessClaims := Claims{
		UserDTO: dto,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(config.Cfg.JWT.JWTAccessKey))
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}
	refreshClaims := Claims{
		UserDTO: dto,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 дней
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(config.Cfg.JWT.JWTRefreshKey))
	if err != nil {
		// fmt.Print(err)
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

// Проверка подлинности и срока годности токенов
func ValidateRefreshToken(tokenString string) (*Claims, error) {
	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT.JWTRefreshKey), nil
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

func ValidateAccessToken(tokenString string) (*Claims, error) {
	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT.JWTAccessKey), nil
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
func SaveToken(refreshToken string, userId uint, device string) error {
	var session models.Session
	if err := postgres.DB.Where("user_id = ? AND device = ?", userId, device).First(&session).Error; err != nil {
		// Если ошибка связана с провалом поиска записи
		if err == gorm.ErrRecordNotFound {
			newSession := models.Session{
				RefreshToken: refreshToken,
				Device:       device,
				UserID:       userId,
			}
			// Создаем новую сессию
			postgres.DB.Create(&newSession)
		} else {
			return err
		}
		// Если запись была найдена, перезаписываем токен и сохраняем
	} else {
		session.RefreshToken = refreshToken
		if err := postgres.DB.Save(&session).Error; err != nil {
			return err
		}
	}

	return nil
}

func RemoveToken(token string) error {
	if err := postgres.DB.Where("refresh_token = ?", token).Delete(&models.Session{}).Error; err != nil {
		return err
	}

	return nil
}
