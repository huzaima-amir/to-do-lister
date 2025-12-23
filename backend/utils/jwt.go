package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SUPER_SECRET_KEY_CHANGE_THIS") // move to env later

// GenerateJWT creates a signed token containing the userID
func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), // token expires in 24h
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT verifies the token and extracts the userID
func ValidateJWT(tokenString string) (uint, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        // Ensure token is signed with HS256
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid token claims")
    }

    // Extract userID
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("invalid user_id in token")
    }

    return uint(userIDFloat), nil
}
