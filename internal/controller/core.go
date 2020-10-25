package controller

import (
	"net/http"

	"github.com/Mynor2397/virtual-parish-office/internal/service"
)

type userController struct{}
type personController struct{}
type documentController struct{}

var userService service.UserService
var personService service.PersonService
var documentService service.DocumentService

// UserController contiene todos los controladores de usuario
type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	Update(w http.ResponseWriter, r *http.Request)
	GetManyUsers(w http.ResponseWriter, r *http.Request)

	Rols(w http.ResponseWriter, r *http.Request)
}

type PersonController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetManyPersons(w http.ResponseWriter, r *http.Request)
	GetManyPersonByFilter(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)
	GetBaptizedPerson(w http.ResponseWriter, r *http.Request)
	GetBaptizedPersonByFilter(w http.ResponseWriter, r *http.Request)
	GetBaptizedPartida(w http.ResponseWriter, r *http.Request)
	DeleteBaptizedPartida(w http.ResponseWriter, r *http.Request)

	CreatePriest(w http.ResponseWriter, r *http.Request)
	GetManyPriest(w http.ResponseWriter, r *http.Request)
	GetManyPriestByFilter(w http.ResponseWriter, r *http.Request)
	GetCountPriest(w http.ResponseWriter, r *http.Request)

	GetPlaces(w http.ResponseWriter, r *http.Request)
	CreatePlace(w http.ResponseWriter, r *http.Request)
	GetCountPlace(w http.ResponseWriter, r *http.Request)
}

type DocumentController interface {
	GetFolio(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetCountBook(w http.ResponseWriter, r *http.Request)

	Baptism(w http.ResponseWriter, r *http.Request)
	GetBaptisms(w http.ResponseWriter, r *http.Request)

	GetAudit(w http.ResponseWriter, r *http.Request)
	CreateAudi(w http.ResponseWriter, r *http.Request)
}

// NewUserController retorna un nuevo controller de tipo usuario controller
func NewUserController(service service.UserService) UserController {
	userService = service
	return &userController{}
}

func NewPersonController(person service.PersonService) PersonController {
	personService = person
	return &personController{}
}

func NewDocumentController(document service.DocumentService) DocumentController {
	documentService = document
	return &documentController{}
}
