package handler

import (
	"fmt"
	"time"

	"github.com/Ali-Assar/CashWatch/user-management/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStorer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokens, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok || len(tokens) == 0 {
			return fmt.Errorf("unauthorized")
		}

		token := tokens[0]
		claims, err := validateToken(token)
		if err != nil {
			return fmt.Errorf("invalid credentialsssss")
		}
		expiresStr, ok := claims["expires"].(string)
		if !ok {
			return fmt.Errorf("invalid credentials")
		}
		expires, err := time.Parse(time.RFC3339, expiresStr)
		if err != nil {
			return fmt.Errorf("invalid credentials %s", err)
		}

		if time.Now().After(expires) {
			return fmt.Errorf("unauthorized")
		}

		id := claims["id"]

		user, err := userStore.GetUserByID(c.Context(), id)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		//set the current authenticated user to the context
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}

		secret := "JWTSECRET"
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("field to parse jwt token : ", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
