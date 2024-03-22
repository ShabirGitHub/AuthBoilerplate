package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"quecko.com/AuthBoilerplate/routes/authentication"
	"quecko.com/AuthBoilerplate/routes/profile"
	utils "quecko.com/AuthBoilerplate/utils"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "quecko.com/AuthBoilerplate/docs"
)

// @title Auth Boilerplate API
// @version 1.0
// @description This is a sample service for managing AuthBoilerplate
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email shabir@gmail.com
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := Routes()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Routes initializes and configures the router for the application.
// It sets up CORS handling, content type for responses, and timeout,
// and defines routes for API version 1 including authentication and profile endpoints.
func Routes() *chi.Mux {
	// Create a new Chi router
	router := chi.NewRouter()

	// Configure CORS settings
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// Although not all of these methods are currently utilized in this project, the Default value could suffice.
		// However, they are available in case of future extensions or additions.
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		corsHandler.Handler,
		middleware.Timeout(30*time.Second),
	)

	// Define routes for API version 1
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", authentication.Routes())
		r.Mount("/profile", utils.EnforceAuthentication(profile.Routes()))
	})

	// Serve Swagger UI at /swagger URL
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to your API documentation JSON file
	))

	return router
}
