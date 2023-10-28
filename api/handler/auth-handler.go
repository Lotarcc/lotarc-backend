package handler

import (
    "encoding/json"
    "net/http"
    "github.com/lotarcc/lotarc-backend/internal/models"
    // Assume you have moved your GenerateToken and VerifyToken functions here.
)

// Register registers a new user
func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    // TODO: Hash the password and Save the user to the database.
    // ...

    token, err := GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Login logs in a user
func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    // TODO: Validate the user credentials from the database.
    // ...

    token, err := GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
