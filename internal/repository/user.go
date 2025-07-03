package repository

import (
	"database/sql"
	"fmt"

	dtoReq "github.com/View-MG/be-project/internal/dto/request"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	CreateUser(c *fiber.Ctx, user dtoReq.User) (string, error)
	GetPasswordHashed(c *fiber.Ctx, username string) (string, error)
}

func (r repository) CreateUser(c *fiber.Ctx, user dtoReq.User) (string, error) {
	query := `
		INSERT INTO public.users(
		id, username, email, password_hash, full_name, created_at)
		VALUES (:ID, :Username, :Email, :Password, :FullName, :CreatedAt)
	`

	args := map[string]interface{}{
		"ID":        user.ID,
		"Username":  user.Username,
		"Email":     user.Email,
		"Password":  user.Password,
		"FullName":  user.FullName,
		"CreatedAt": user.CreatedAt,
	}

	_, err := r.db.NamedExecContext(c.Context(), query, args)
	if err != nil {
		return "", err
	}
	return "Create User Success", nil
}

func (r repository) GetPasswordHashed(c *fiber.Ctx, username string) (string, error) {
	query := `
		SELECT password_hash FROM public.users WHERE username = $1
	`
	var passwordHash string
	err := r.db.GetContext(c.Context(), &passwordHash, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("username or password incorrect")
		}
		return "", err
	}
	fmt.Println(passwordHash)
	return passwordHash, nil
}
