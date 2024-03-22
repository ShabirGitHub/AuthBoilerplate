package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"quecko.com/AuthBoilerplate/models"
	auth "quecko.com/AuthBoilerplate/utils"
)

// signInAndReturnToken signs a user in and returns a session token.
// @Summary Signs in a user and returns a session token
// @Description Signs in a user with the provided credentials and returns a session token if successful
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]interface{} "Successfully signed in"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/sign_in [post]
func signInAndReturnToken(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would typically check the user credentials against your database or authentication service
	// For simplicity, let's assume a hardcoded check
	if user.Username != "shabir" || user.Password != "shabir123" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r,
		map[string]interface{}{
			"message":      "Logged in!",
			"sessionToken": "Bearer " + token,
			// TODO: Validity?
		})
}
