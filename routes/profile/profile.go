package profile

import (
	"encoding/json"
	"net/http"

	"quecko.com/AuthBoilerplate/models"
)

// getProfile retrieves the profile information for the authenticated user.
// @Summary Retrieves the profile information for the authenticated user
// @Description Retrieves the profile information for the authenticated user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer JWT token"
// @Success 200 {object} models.Profile "Successful response"
// @Success 200 {object} models.Profile "Profile information"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /profile/get [get]
func getProfile(w http.ResponseWriter, r *http.Request) {

	// Fetch profile data from DB or service
	profile := models.Profile{
		Username:    "username",
		Email:       "admin@example.com",
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		Gender:      "Male",
		Address:     "123 Main Street",
		City:        "New York",
		Country:     "USA",
		PhoneNumber: "+1234567890",
		Interests:   []string{"Programming", "Sports", "Music"},
		SocialLinks: map[string]string{"Twitter": "https://twitter.com/johndoe", "LinkedIn": "https://www.linkedin.com/in/johndoe"},
	}

	jsonResponse, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
