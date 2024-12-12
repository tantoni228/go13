package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userIdCtxKey struct{}

type jwtClaims struct {
    jwt.RegisteredClaims
    UserId string `json:"user_id"` // Здесь должно быть "user_id", а не "uuid"
}

func main() {
    ctx := context.Background()

    // Пример UUID (можно заменить на входной параметр)
    uuid := uuid.New()

    // Генерация JWT
    token := generateJWT(uuid.String())

    fmt.Println("Сгенерированный JWT:", token)

    var claims jwtClaims
    _, err := jwt.ParseWithClaims(token, &claims, nil)
    if err != nil && !errors.Is(err, jwt.ErrTokenUnverifiable) {
        fmt.Println(err)
    }

    userId := claims.UserId

    ctx = context.WithValue(ctx, userIdCtxKey{}, userId)

    fmt.Println(UserIdFromCtx(ctx)) // Теперь будет выведен правильный UUID
}

// UserIdFromCtx returns userId associated with context.
// If no userId is associated, the empty string is returned.
func UserIdFromCtx(ctx context.Context) string {
    return ctx.Value(userIdCtxKey{}).(string)
}

// Функция для генерации JWT

func generateJWT(uuid string) string {
    // Определяем секрет для подписи токена
    secret := []byte("my_secret_key")

    // Создаем Claims (дополнительные данные для токена)
    claims := jwt.MapClaims{
        "user_id": uuid, // Здесь должно быть "user_id", а не "uuid"
        "exp":     time.Now().Add(time.Hour * 1).Unix(),
    }

    // Создаем новый токен
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Подписываем токен
    signedToken, err := token.SignedString(secret)
    if err != nil {
        log.Fatal(err)
    }

    return signedToken
}