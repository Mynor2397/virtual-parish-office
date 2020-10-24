package routers

import (
	"github.com/Mynor2397/virtual-parish-office/internal/middleware"
	"github.com/gorilla/mux"

	"github.com/Mynor2397/virtual-parish-office/internal/controller"
	"github.com/Mynor2397/virtual-parish-office/internal/service"
	"github.com/Mynor2397/virtual-parish-office/internal/storage"
)

var (
	userStorage    storage.UserStorage       = storage.NewUserStorage()
	userService    service.UserService       = service.NewUserService(userStorage)
	userController controller.UserController = controller.NewUserController(userService)
)

// SetUserRoutes registra la rutas a usar para los controladires de usuario
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", userController.Login).Methods("POST")
	user := router.PathPrefix("/users").Subrouter()
	user.Use(middleware.AuthForAmdmin)
	user.HandleFunc("/register", userController.Create).Methods("POST")
	user.HandleFunc("/changerol/{id}", userController.Update).Methods("PUT")
	user.HandleFunc("/many", userController.GetManyUsers).Methods("GET")
	user.HandleFunc("/rols", userController.Rols).Methods("GET")

	return router
}
