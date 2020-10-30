package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Mynor2397/virtual-parish-office/internal/middleware"
)

// InitRoutes  inicializa las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/images").Handler(http.StripPrefix("/images", http.FileServer(http.Dir("public/"))))

	api := router.PathPrefix("/vpo/v1").Subrouter()
	api.Use(middleware.Auth)
	api = SetUserRoutes(api)
	api = SetPersonRoutes(api)
	api = SetPlacesRoutes(api)
	api = SetDocumentRoutes(api)
	router.Use(middleware.WriteJSONHeader)
	return router
}
