package storage

import (
	"context"
	"database/sql"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*repoPerson) GetManyPriest(ctx context.Context) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	query := "SELECT idPriest, firstName, secondName, lastName, secondLastName FROM VPO_Priest;"

	rows, err := db.QueryContext(ctx, query)
	if err == sql.ErrNoRows {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(
			&person.ID,
			&person.Firstname,
			&person.Secondname,
			&person.Lastname,
			&person.Secondlastname,
		)

		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (*repoPerson) GetPriestByFilter(ctx context.Context, filter string) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	query := "SELECT idPriest, firstName, secondName, lastName, secondLastName FROM VPO_Priest" +
		" WHERE CONCAT(firstName, ' ', secondName, ' ', lastName, ' ', secondLastName)" +
		" like ? OR firstName like ? OR secondName like ?;"

	rows, err := db.QueryContext(ctx, query, filter, filter, filter)

	if err == sql.ErrNoRows {
		return persons, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(
			&person.ID,
			&person.Firstname,
			&person.Secondname,
			&person.Lastname,
			&person.Secondlastname,
		)

		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (*repoPerson) CreatePriest(ctx context.Context, priest models.Priest) (int, error) {
	var count int
	query := "INSERT INTO VPO_Priest (firstName, secondName, lastName, secondLastName, credentials)" +
		" VALUES(?, ?, ?, ?, ?);"

	_, err := db.QueryContext(
		ctx,
		query,
		priest.FirstnamePriest,
		priest.SecondnamePriest,
		priest.LastnamePriest,
		priest.SecondLastnamePriest,
		priest.Credentials,
	)
	if err != nil {
		return count, err
	}

	querycount := "SELECT COUNT(*) FROM VPO_Priest;"
	err = db.QueryRowContext(ctx, querycount).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (*repoPerson) GetLastPriest(ctx context.Context) (int, error) {
	var count int
	query := "SELECT COUNT(*) AS COUNT FROM VPO_Priest;"

	if err := db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
