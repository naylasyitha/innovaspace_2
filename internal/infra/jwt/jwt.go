package jwt

import (
	"fmt"
	"innovaspace/internal/infra/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTI interface {
	GenerateToken(UserId uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
}

type JWT struct {
	SecretKey   string
	ExpiredTime time.Time
}

func NewJWT() *JWT {
	_env, err := env.New()
	if err != nil {
		panic(err)
	}

	secretKey := _env.JwtSecret
	expiredTime := time.Now().Add(time.Hour * time.Duration(_env.JwtExpired))

	return &JWT{
		SecretKey:   secretKey,
		ExpiredTime: expiredTime,
	}
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(UserId uuid.UUID) (string, error) {
	claim := Claims{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, error) {
	var claim Claims
	var id uuid.UUID

	fmt.Println("SECRET", j.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (any, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return id, err
	}

	if !token.Valid {
		return id, err
	}

	id = claim.UserId

	return id, nil
}
