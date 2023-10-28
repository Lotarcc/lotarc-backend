package handler

import (
    "errors"
    "time"
    "os"
    
    "github.com/dgrijalva/jwt-go"
)

var (
    // Define a secret key used to sign the JWT token
    // Uses .env file to store the secret key
    secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

// Claims represents the JWT claims
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given username
func GenerateToken(username string) (string, error) {
    // Create the claims for the JWT token
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        },
    }

    // Create the JWT token with the claims and sign it with the secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// VerifyToken verifies the given JWT token and returns the username if the token is valid
func VerifyToken(tokenString string) (string, error) {
    // Parse the JWT token with the secret key
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return "", err
    }

    // Verify that the token is valid and extract the username from the claims
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return "", errors.New("invalid token")
    }

    return claims.Username, nil
}
