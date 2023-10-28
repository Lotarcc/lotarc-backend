package middleware

import (
	"net/http"
	"strings"

	"github.com/lotarcc/lotarc-backend/api/handler"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if user is authenticated
        if !isAuthenticated(r) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}

func isAuthenticated(r *http.Request) bool {
    // Get the JWT token from the request header
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        return false
    }

    // Verify the JWT token
    _, err := handler.VerifyToken(strings.TrimPrefix(tokenString, "Bearer "))
    if err != nil {
        // Log the error or handle it as you like.
        return false
    }

    return true
}

