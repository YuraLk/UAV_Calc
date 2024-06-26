package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/YuraLk/drone_calc/backend/internal/database/postgres"
	"github.com/YuraLk/drone_calc/backend/internal/dtos/auth/response_properties"
	"github.com/YuraLk/drone_calc/backend/internal/exeptions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	C *gin.Context
}

type Tokens struct {
	Access  string
	Refresh string
}

func create_dto(ctx context.Context, userId uint, name string, surname string, patronymic string, email string, phone string) (response_properties.AuthDTO, error) {

	var access string

	// Проверяем роль пользователя и исходя из роли создаем DTO
	if err := postgres.DB.QueryRow(ctx, "SELECT access FROM accesses WHERE user_id = $1 LIMIT 1", userId).Scan(&access); err != nil {
		// Если доступа нет
		if err == pgx.ErrNoRows {
			return response_properties.AuthDTO{
				Id:         userId,
				Name:       name,
				Surname:    surname,
				Patronymic: patronymic,
				Email:      email,
				Phone:      phone,
			}, nil
		} else {
			return response_properties.AuthDTO{}, err
		}
		// Если доступ есть
	} else {
		return response_properties.AuthDTO{
			Id:         userId,
			Name:       name,
			Surname:    surname,
			Patronymic: patronymic,
			Email:      email,
			Phone:      phone,
			Role:       &access,
		}, nil
	}
}

// func (S AuthService) Auth(email string, password string, device string) (user_dtos.UserDTO, Tokens, error) {
// var user models.User
// // Ищем пользователя в БД
// if err := postgres.DB.Where("email = ?", email).First(&user).Error; err != nil {
// 	exeptions.BadRequest(S.C, fmt.Sprintf("Пользователь с email %s не найден!", email), err)
// 	return user_dtos.UserDTO{}, Tokens{}, err
// }

// // Сравнение паролей
// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 	exeptions.BadRequest(S.C, "Неверный пароль!", err)
// 	return user_dtos.UserDTO{}, Tokens{}, err
// }

// // Проверяем роль пользователя и исходя из роли создаем DTO
// dto, err := create_dto(user.Id, user.Name, user.Email, user.Phone)
// if err != nil {
// 	exeptions.InternalServerError(S.C, err)
// }

// // Гененируем новые Access и Refresh токены
// accessToken, refreshToken, err := TokenService{}.Generate(dto)
// if err != nil {
// 	exeptions.InternalServerError(S.C, err)
// 	return user_dtos.UserDTO{}, Tokens{}, err
// }

// // По id пользователя и девайсу перезаписываем сессию либо создаем новую
// err = TokenService{}.Save(refreshToken, user.Id, device)
// if err != nil {
// 	exeptions.InternalServerError(S.C, err)
// 	return user_dtos.UserDTO{}, Tokens{}, err
// }

// return dto, Tokens{
// 	Access:  accessToken,
// 	Refresh: refreshToken,
// }, nil
// }

func (S AuthService) Register(ctx context.Context, name string, surname string, patronymic string, email string, phone_number string, password string, device string) (response_properties.AuthDTO, Tokens, error) {

	var email_is_exists bool

	// Проверка уникальности данных
	if err := postgres.DB.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1) as is_exists", email).Scan(&email_is_exists); err != nil {
		if err != pgx.ErrNoRows {
			exeptions.InternalServerError(S.C, err)
			return response_properties.AuthDTO{}, Tokens{}, err
		}
	}

	if email_is_exists {
		err := errors.New("value is not unique")
		exeptions.BadRequest(S.C, fmt.Sprintf("Пользователь с электронной почтой %s уже существует!", email), err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	var phone_number_is_exists bool

	if err := postgres.DB.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = $1) as is_exists", phone_number).Scan(&phone_number_is_exists); err != nil {
		if err != pgx.ErrNoRows {
			exeptions.InternalServerError(S.C, err)
			return response_properties.AuthDTO{}, Tokens{}, err
		}
	}

	if phone_number_is_exists {
		err := errors.New("value is not unique")
		exeptions.BadRequest(S.C, fmt.Sprintf("Пользователь с номером телефона %s уже существует!", phone_number), err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	var user_id uint

	// Регистрируем пользователя
	if err := postgres.DB.QueryRow(ctx, `
		INSERT INTO users (name, surname, patronymic, email, phone_number, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, name, surname, patronymic, email, phone_number, hashedPassword).Scan(&user_id); err != nil {
		exeptions.InternalServerError(S.C, err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	// Проверяем роль пользователя и исходя из роли создаем DTO
	dto, err := create_dto(ctx, user_id, name, surname, patronymic, email, phone_number)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	// Генерируем токены доступа
	accessToken, refreshToken, err := TokenService{}.Generate(dto)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	// Сохраняем токен
	err = TokenService{}.Save(ctx, refreshToken, user_id, device)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return response_properties.AuthDTO{}, Tokens{}, err
	}

	return dto, Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

// func (S AuthService) Refresh(refreshToken string) (user_dtos.UserDTO, Tokens, error) {
// 	tokenData, err := TokenService{}.ValidateRefresh(refreshToken)
// 	if err != nil {
// 		exeptions.UnauthorizedError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}
// 	// Ищем токен в базе даных
// 	var session models.Session
// 	if err := postgres.DB.Where("refresh_token = ?", refreshToken).First(&session).Error; err != nil {
// 		exeptions.UnauthorizedError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}
// 	// Ищем пользователя
// 	var user models.User
// 	if err := postgres.DB.Where("id = ?", tokenData.UserDTO.Id).First(&user).Error; err != nil {
// 		exeptions.UnauthorizedError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}

// 	// Проверяем роль пользователя и исходя из роли создаем DTO
// 	dto, err := create_dto(user.Id, user.Name, user.Email, user.Phone)
// 	if err != nil {
// 		exeptions.InternalServerError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}

// 	// Обновляем токены
// 	accessToken, refreshToken, err := TokenService{}.Generate(dto)
// 	if err != nil {
// 		exeptions.UnauthorizedError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}

// 	// Сохраняем Refresh - токен в БД
// 	session.RefreshToken = refreshToken
// 	if err := postgres.DB.Save(&session).Error; err != nil {
// 		exeptions.InternalServerError(S.C, err)
// 		return user_dtos.UserDTO{}, Tokens{}, err
// 	}

// 	return dto, Tokens{
// 		Access:  accessToken,
// 		Refresh: refreshToken,
// 	}, nil
// }

func (S AuthService) Logout(ctx context.Context, refreshToken string) error {
	err := TokenService{}.Remove(ctx, refreshToken)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
	}

	return nil
}
