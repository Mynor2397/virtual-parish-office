package storage

import (
	"context"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
	"github.com/Mynor2397/virtual-parish-office/internal/mysql"
)

var db = mysql.Connect()

type repoUser struct{}
type repoPerson struct{}
type repoDocument struct{}

// UserStorage implementa todos lo metodos de usuario
type UserStorage interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Update(ctx context.Context, id, rol string) ( error)

	Login(ctx context.Context, user *models.User) (models.User, error)
	GetManyUsers(ctx context.Context)([]models.User, error)

	Roles(ctx context.Context)([]models.Rol, error)
}

type PersonStorage interface {
	Create(ctx context.Context, person models.Person) (models.Person, error)
	GetManyPersons(ctx context.Context, sex string) ([]models.Person, error)
	GetManyPersonByFilter(ctx context.Context, limit int, filter string)([]models.Person, error)
	GetBaptizedPerson(ctx context.Context, limit int)([]models.Baptism, error)
	GetBaptizedPersonByFilter(ctx context.Context, filter string)([]models.Baptism, error)
	GetBaptizedPartida(ctx context.Context, id string)(models.Baptism, error)
	DeleteBaptizedPartida(ctx context.Context, id string) error

	GetManyPriest(ctx context.Context) ([]models.Person, error)
	GetPriestByFilter(ctx context.Context, filter string)([]models.Person, error)
	CreatePriest(ctx context.Context, priest models.Priest)(int, error)
	GetLastPriest(ctx context.Context)(int, error)

	GetPlaces(ctx context.Context) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place)(int, error)
	GetLastPlace(ctx context.Context) (int, error)
}

type DocumentStorage interface {
	GetFolio(ctx context.Context, idBook int) ([]models.Folio, error)
	GetBook(ctx context.Context)([]models.Book, error)
	CreateBook(ctx context.Context, book models.Book)(error)
	GetLastBook(ctx context.Context) (int, error)

	Baptism(ctx context.Context, baptism models.Baptism) (string, error)
	GetBaptisms(ctx context.Context, filter string)([]models.Baptism, error)

	GetAudit(ctx context.Context)([]models.Audit, error)
	CreateAudi(ctx context.Context, audit models.Audit)error
}

// NewUserStorage  constructor para userStorage
func NewUserStorage() UserStorage {
	return &repoUser{}
}

// NewPersonStorage  constructor para userStorage
func NewPersonStorage() PersonStorage {
	return &repoPerson{}
}

// NewPersonStorage  constructor para userStorage
func NewDocumentStorage() DocumentStorage {
	return &repoDocument{}
}
