package service

import (
	"context"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
	"github.com/Mynor2397/virtual-parish-office/internal/storage"
)

var myUserstorage storage.UserStorage
var myPersonStorage storage.PersonStorage
var myDocumentStorage storage.DocumentStorage

type userService struct{}
type personService struct{}
type documentService struct{}

// UserService implementa el conjunto de metodos de servicio para usuario
type UserService interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (models.User, error)

	Update(cxt context.Context, id, rol string) ( error)
	GetManyUsers(ctx context.Context)([]models.User, error)

	Roles(ctx context.Context)([]models.Rol, error)
}

type PersonService interface {
	Create(ctx context.Context, person models.Person) (models.Person, error)
	GetManyPersons(cxt context.Context, ctx string) ([]models.Person, error)
	GetManyPersonByFilter(ctx context.Context, limit int, filter string)([]models.Person,error)
	GetBaptizedPerson(ctx context.Context, limit int) ([]models.Baptism, error)
	GetBaptizedPersonByFilter(ctx context.Context, filter string) ([]models.Baptism, error)
	GetBaptizedPartida(ctx context.Context, id string) (models.Baptism, error)
	DeleteBaptizedPartida(ctx context.Context, id string) error

	CreatePriest(ctx context.Context, priest models.Priest)(int, error)
	GetManyPriest(cxt context.Context) ([]models.Person, error)
	GetPriestByFilter(ctx context.Context, filter string) ([]models.Person, error)
	GetCountPriest(ctx context.Context)(int, error)

	GetPlaces(ctx context.Context) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) (int, error)
	GetCountPlace(ctx context.Context) (int, error)
}

type DocumentService interface {
	GetFolio(ctx context.Context, idBook int) ([]models.Folio, error)
	GetBook(ctx context.Context) ([]models.Book, error)
	CreateBook(ctx context.Context, book models.Book) error
	GetCountBook(ctx context.Context)(int, error)

	Baptism(ctx context.Context, baptism models.Baptism) (string, error)
	GetBaptisms(ctx context.Context, filter string) ([]models.Baptism, error)

	GetAudit(ctx context.Context)([]models.Audit, error)
	CreateAudi(ctx context.Context, id string, audit models.Audit)error

}

// NewUserService retorna un nuevo servicio para los usuarios
func NewUserService(userstorage storage.UserStorage) UserService {
	myUserstorage = userstorage
	return &userService{}
}

// NewPersonService retorna un nuevo servicio para los usuarios
func NewPersonService(personstorage storage.PersonStorage) PersonService {
	myPersonStorage = personstorage
	return &personService{}
}

// NewDocumentService retorna un nuevo servicio para los usuarios
func NewDocumentService(documentstorage storage.DocumentStorage) DocumentService {
	myDocumentStorage = documentstorage
	return &documentService{}
}
