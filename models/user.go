package models

import (
	"albert/utils"
	"errors"
	"html"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Role  Role   `gorm:"type:role; default:user; nullable" json:"role"`
}

type Role string

const (
	AdminRole     Role = "admin"
	ModeratorRole Role = "moderator"
	UserRole      Role = "user"
)

func (user *User) SaveUser() (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return &User{}, utils.ParseError(err)
	}
	return user, nil
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	if !utils.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

// TODO: Update this (user *User) Exists() bool
func CheckUser(email string) error {
	u := User{}
	if err := DB.Model(User{}).Where("email = ?", email).Take(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	u := User{}
	if err := DB.Model(User{}).Where("email = ?", email).Take(&u).Error; err != nil {
		return &User{}, err
	}
	return &u, nil
}

func (user *User) GenerateToken() (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
