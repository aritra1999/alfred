package models

import (
	"albert/utils"
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Role     Role   `gorm:"type:role; default:user; nullable" json:"role"`
}

type Role string

const (
	AdminRole     Role = "admin"
	ModeratorRole Role = "moderator"
	UserRole      Role = "user"
)

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	if !utils.ValidateEmail(u.Email) {
		return errors.New("invalid email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func SingInCheck(email string, password string) (string, error) {
	u := User{}

	if !utils.ValidateEmail(email) {
		return "", errors.New("invalid email")
	}

	if err := DB.Model(User{}).Where("email = ?", email).Take(&u).Error; err != nil {
		return "", err
	}

	if err := VerifyPassword(password, u.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(utils.TokenBody{
		Email: u.Email,
		ID:    int(u.ID),
		Role:  string(u.Role),
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
