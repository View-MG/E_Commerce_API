package service

import (
	"fmt"
	"time"

	dtoReq "github.com/View-MG/be-project/internal/dto/request"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(c *fiber.Ctx, user dtoReq.User) (string, error)
	Login(c *fiber.Ctx, user dtoReq.LoginRequest) (string, error)
}

func (s service) CreateUser(c *fiber.Ctx, user dtoReq.User) (string, error) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	user.ID = id
	user.CreatedAt = createdAt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)
	res, err := s.repo.User.CreateUser(c, user)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s service) Login(c *fiber.Ctx, user dtoReq.LoginRequest) (string, error) {
	hashedPassword, err := s.repo.User.GetPasswordHashed(c, user.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("compare err")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // token หมดอายุ 24 ชั่วโมง
	secretKey := []byte("secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
