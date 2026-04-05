package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret_key_change_later")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWT(userID int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtSecret, nil
	})
}
func RegisterUser(name, email, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "USER",
	}

	return CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", err
	}

	return GenerateJWT(user.ID, user.Role)
}