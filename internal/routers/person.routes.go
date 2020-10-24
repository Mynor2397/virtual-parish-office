package routers

import (
	"github.com/gorilla/mux"

	"github.com/Mynor2397/virtual-parish-office/internal/controller"
	"github.com/Mynor2397/virtual-parish-office/internal/service"
	"github.com/Mynor2397/virtual-parish-office/internal/storage"
)

var (
	personStorage    storage.PersonStorage       = storage.NewPersonStorage()
	personService    service.PersonService       = service.NewPersonService(personStorage)
	personController controller.PersonController = controller.NewPersonController(personService)
)

// SetPersonRoutes registra la rutas a usar para los controladires de usuario
func SetPersonRoutes(router *mux.Router) *mux.Router {
	person := router.PathPrefix("/persons").Subrouter()
	person.HandleFunc("/create", personController.Create).Methods("POST")
	person.HandleFunc("/many/{sex}", personController.GetManyPersons).Methods("GET")
	person.HandleFunc("/baptized/by/{limit}", personController.GetBaptizedPerson).Methods("GET")
	person.HandleFunc("/baptized/{filter}", personController.GetBaptizedPersonByFilter).Methods("GET")
	person.HandleFunc("/partida/{id}", personController.GetBaptizedPartida).Methods("GET")
	person.HandleFunc("/partida/{id}", personController.DeleteBaptizedPartida).Methods("DELETE")

	person.HandleFunc("/priest", personController.CreatePriest).Methods("POST")
	person.HandleFunc("/priest", personController.GetManyPriest).Methods("GET")
	person.HandleFunc("/priest/{filter}", personController.GetManyPriestByFilter).Methods("GET")
	person.HandleFunc("/priests/count", personController.GetCountPriest).Methods("GET")
	person.HandleFunc("/by/{limit}/{filter}", personController.GetManyPersonByFilter).Methods("GET")

	// user.HandleFunc("/login", userController.Login).Methods("POST")
	return router
}
