package service

import (
	"errors"
	"fmt"

	"github.com/YuraLk/drone_calc/internal/database/postgres"
	"github.com/YuraLk/drone_calc/internal/dtos"
	"github.com/YuraLk/drone_calc/internal/exeptions"
	"github.com/YuraLk/drone_calc/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Tokens struct {
	Access  string
	Refresh string
}

func CreateDTO(userId uint, name string, email string, phone string) (dtos.UserDTO, error) {
	var access models.Access

	// Проверяем роль пользователя и исходя из роли создаем DTO
	if err := postgres.DB.Where("user_id = ?", userId).First(&access).Error; err != nil {
		// Если доступа нет
		if err == gorm.ErrRecordNotFound {
			return dtos.UserDTO{
				Id:    userId,
				Name:  name,
				Email: email,
				Phone: phone,
			}, nil
		} else {
			return dtos.UserDTO{}, err
		}
		// Если доступ есть
	} else {
		return dtos.UserDTO{
			Id:    userId,
			Name:  name,
			Email: email,
			Phone: phone,
			Role:  access.Role,
		}, nil
	}
}

func GetUsers(c *gin.Context) []models.User {
	users := []models.User{}
	if err := postgres.DB.Find(&users).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return []models.User{}
	}
	return users
}

func Auth(c *gin.Context, email string, password string, device string) (dtos.UserDTO, Tokens, error) {
	var user models.User
	// Ищем пользователя в БД
	if err := postgres.DB.Where("email = ?", email).First(&user).Error; err != nil {
		exeptions.BadRequest(c, fmt.Sprintf("Пользователь с email %s не найден!", email), err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Сравнение паролей
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		exeptions.BadRequest(c, "Неверный пароль!", err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Проверяем роль пользователя и исходя из роли создаем DTO
	dto, err := CreateDTO(user.Id, user.Name, user.Email, user.Phone)
	if err != nil {
		exeptions.InternalServerError(c, err)
	}

	// Гененируем новые Access и Refresh токены
	accessToken, refreshToken, err := GenerateTokens(dto)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// По id пользователя и девайсу перезаписываем сессию либо создаем новую
	if err := SaveToken(refreshToken, user.Id, device); err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	return dto, Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func Register(c *gin.Context, name string, email string, phone string, password string, device string) (dtos.UserDTO, Tokens, error) {
	var exist models.User // Сюда помещаем рузультаты поиска
	// Проверка уникальности данных
	if err := postgres.DB.Where("email = ?", email).First(&exist).Error; err == nil {
		err := errors.New("value is not unique")
		exeptions.BadRequest(c, fmt.Sprintf("Пользователь с email %s уже существует!", email), err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	if err := postgres.DB.Where("phone = ?", phone).First(&exist).Error; err == nil {
		err := errors.New("value is not unique")
		exeptions.BadRequest(c, fmt.Sprintf("Пользователь с номером телефона %s уже существует!", phone), err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Создаем структуру пользователя
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Phone:    phone,
	}
	// Регистрируем пользователя
	if err := postgres.DB.Create(&user).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Проверяем роль пользователя и исходя из роли создаем DTO
	dto, err := CreateDTO(user.Id, name, email, phone)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Генерируем токены доступа
	accessToken, refreshToken, err := GenerateTokens(dto)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Сохраняем токен
	if err := SaveToken(refreshToken, user.Id, device); err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	return dto, Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func Refresh(c *gin.Context, refreshToken string) (dtos.UserDTO, Tokens, error) {
	tokenData, err := ValidateRefreshToken(refreshToken)
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}
	// Ищем токен в базе даных
	var session models.Session
	if err := postgres.DB.Where("refresh_token = ?", refreshToken).First(&session).Error; err != nil {
		exeptions.UnauthorizedError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}
	// Ищем пользователя
	var user models.User
	if err := postgres.DB.Where("id = ?", tokenData.UserDTO.Id).First(&user).Error; err != nil {
		exeptions.UnauthorizedError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Проверяем роль пользователя и исходя из роли создаем DTO
	dto, err := CreateDTO(user.Id, user.Name, user.Email, user.Phone)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Обновляем токены
	accessToken, refreshToken, err := GenerateTokens(dto)
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Сохраняем Refresh - токен в БД
	session.RefreshToken = refreshToken
	if err := postgres.DB.Save(&session).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	return dto, Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func UpdateUser(c *gin.Context, userId uint, name string, email string, phone string, device string) (dtos.UserDTO, Tokens, error) {
	var user models.User
	if err := postgres.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		exeptions.NotFound(c, "Пользователь не найден!")
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Формируем объект для сохраниния
	updateUser := models.User{
		Id:       user.Id,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: user.Password,
	}

	if err := postgres.DB.Save(&updateUser).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Проверяем роль пользователя и исходя из роли создаем DTO
	dto, err := CreateDTO(updateUser.Id, updateUser.Name, updateUser.Email, updateUser.Phone)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	// Формируем новые токены исходя из новых данных
	accessToken, refreshToken, err := GenerateTokens(dto)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	if err := SaveToken(refreshToken, updateUser.Id, device); err != nil {
		exeptions.InternalServerError(c, err)
		return dtos.UserDTO{}, Tokens{}, err
	}

	return dto, Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func Logout(c *gin.Context, refreshToken string) error {
	if err := RemoveToken(refreshToken); err != nil {
		exeptions.InternalServerError(c, err)
	}

	return nil
}
