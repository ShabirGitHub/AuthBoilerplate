package authentication

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/sign_in", signInAndReturnToken)

	// Define route for implementing the signup feature:
	// - Use a POST request to create a new user account at /signup endpoint.

	//   Example: router.Post("/signup", signUserUp)

	return router
}
