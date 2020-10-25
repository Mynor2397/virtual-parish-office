package service

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*personService) Create(ctx context.Context, person models.Person) (models.Person, error) {
	person.ID = uuid.New().String()
	person.IDAddress = uuid.New().String()

	person.Firstname = strings.TrimSpace(strings.Title(person.Firstname))
	person.Secondname = strings.TrimSpace(strings.Title(person.Secondname))
	person.Lastname = strings.TrimSpace(strings.Title(person.Lastname))
	person.Secondlastname = strings.TrimSpace(strings.Title(person.Secondlastname))
	person.DPI = strings.TrimSpace(person.DPI)
	person.Sexo = strings.TrimSpace(person.Sexo)
	person.Address = strings.TrimSpace(strings.Title(person.Address))

	return myPersonStorage.Create(ctx, person)
}

func (*personService) GetManyPersons(cxt context.Context, sex string) ([]models.Person, error) {
	sex = strings.TrimSpace(sex)
	return myPersonStorage.GetManyPersons(cxt, sex)
}

func (*personService) GetManyPersonByFilter(ctx context.Context, limit int, filter string) ([]models.Person, error) {
	if limit == 0 {
		limit = 10
	}

	return myPersonStorage.GetManyPersonByFilter(ctx, limit, filter)
}

func (*personService) GetManyPriest(ctx context.Context) ([]models.Person, error) {
	return myPersonStorage.GetManyPriest(ctx)
}

func (*personService) GetPriestByFilter(ctx context.Context, filter string) ([]models.Person, error) {
	filter = "%" + filter + "%"
	return myPersonStorage.GetPriestByFilter(ctx, filter)
}

func (*personService) GetBaptizedPerson(ctx context.Context, limit int) ([]models.Baptism, error) {
	return myPersonStorage.GetBaptizedPerson(ctx, limit)
}

func (*personService) GetBaptizedPersonByFilter(ctx context.Context, filter string) ([]models.Baptism, error) {
	var filt string = "%" + filter + "%"
	return myPersonStorage.GetBaptizedPersonByFilter(ctx, filt)
}

func (*personService) GetBaptizedPartida(ctx context.Context, id string) (models.Baptism, error) {
	return myPersonStorage.GetBaptizedPartida(ctx, id)
}

func (*personService) DeleteBaptizedPartida(ctx context.Context, id string) error {
	return myPersonStorage.DeleteBaptizedPartida(ctx, id)
}

func (*personService) CreatePriest(ctx context.Context, priest models.Priest) (int, error) {
	priest.IDPriest = uuid.New().String()

	return myPersonStorage.CreatePriest(ctx, priest)
}

func (*personService) GetCountPriest(ctx context.Context) (int, error) {
	return myPersonStorage.GetLastPriest(ctx)
}

func (*personService) UpdatePerson(ctx context.Context, person models.Person, id string) error {
	person.ID = id

	return myPersonStorage.UpdatePerson(ctx, person)
}
