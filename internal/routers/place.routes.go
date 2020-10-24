package routers

import (
	"github.com/gorilla/mux"

	"github.com/Mynor2397/virtual-parish-office/internal/controller"
	"github.com/Mynor2397/virtual-parish-office/internal/service"
	"github.com/Mynor2397/virtual-parish-office/internal/storage"
)

var (
	placeStorage    storage.PersonStorage       = storage.NewPersonStorage()
	placeService    service.PersonService       = service.NewPersonService(placeStorage)
	placeController controller.PersonController = controller.NewPersonController(placeService)
)

// SetPlacesRoutes registra la rutas a usar para los controladires de usuario
func SetPlacesRoutes(router *mux.Router) *mux.Router {
	person := router.PathPrefix("/place").Subrouter()
	person.HandleFunc("/many", placeController.GetPlaces).Methods("GET")
	person.HandleFunc("/create", placeController.CreatePlace).Methods("POST")
	person.HandleFunc("/all", placeController.GetCountPlace).Methods("GET")
	return router
}
