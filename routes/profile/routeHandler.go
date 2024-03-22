package profile

import "github.com/go-chi/chi"

// Routes configures the routes for the application.
func Routes() *chi.Mux {
	router := chi.NewRouter()

	// Route for retrieving a user's profile
	router.Get("/get", getProfile)

	// Additional routes for profile can be configured here
	// Example:
	// router.Put("/edit", editProfile)
	// router.Post("/upload-picture", uploadPicture)
	// router.Delete("/delete-account", deleteAccount)

	return router
}
