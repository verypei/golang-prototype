package routes

import (
	"golang-prototype/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/users", controllers.GetUsers)
	r.Get("/user/{id}", controllers.GetUserById)
	r.Post("/user/register", controllers.UserRegister)

	return r
}
