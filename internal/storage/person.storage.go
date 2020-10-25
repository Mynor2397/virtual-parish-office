package storage

import (
	"context"
	"database/sql"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*repoPerson) Create(ctx context.Context, person models.Person) (models.Person, error) {
	var pers models.Person

	query := "call registerperson(?,?,?,?,?,?,?,?,?);"
	_, err := db.QueryContext(ctx,
		query,
		person.ID,
		person.Firstname,
		person.Secondname,
		person.Lastname,
		person.Secondlastname,
		person.DPI,
		person.Sexo,
		person.IDAddress,
		person.Address,
	)

	if err != nil {
		return pers, err
	}

	pers = person

	return pers, nil
}

func (*repoPerson) GetManyPersons(ctx context.Context, sex string) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	if sex == "G" {
		query := "SELECT idPerson, firstName, secondName, lastName, secondLastName, DPI, sex FROM VPO_Person " +
			" ORDER BY firstName ASC;"
		rows, err := db.QueryContext(ctx, query)
		if err == sql.ErrNoRows {
			return persons, lib.ErrNotFound
		}

		for rows.Next() {
			err := rows.Scan(&person.ID,
				&person.Firstname,
				&person.Secondname,
				&person.Lastname,
				&person.Secondlastname,
				&person.DPI,
				&person.Sexo,
			)

			if err != nil {
				return persons, err
			}

			persons = append(persons, person)
		}
	} else {
		query := "SELECT idPerson, firstName, secondName, lastName, secondLastName, DPI, sex FROM VPO_Person WHERE sex = ? " +
			" ORDER BY firstName ASC;"
		rows, err := db.QueryContext(ctx, query, sex)
		if err == sql.ErrNoRows {
			return persons, lib.ErrNotFound
		}

		for rows.Next() {
			err := rows.Scan(&person.ID,
				&person.Firstname,
				&person.Secondname,
				&person.Lastname,
				&person.Secondlastname,
				&person.DPI,
				&person.Sexo,
			)

			if err != nil {
				return persons, err
			}

			persons = append(persons, person)
		}

	}

	return persons, nil
}

func (*repoPerson) GetBaptizedPerson(ctx context.Context, limit int) ([]models.Baptism, error) {
	partida := models.Baptism{}
	partidas := []models.Baptism{}
	if limit == 0 {
		limit = 10
	}

	query := "SELECT " +
		"bap.idBaptism, " +
		"per.firstName, " +
		"per.secondName, " +
		"per.lastName, " +
		"per.secondLastname, " +
		"fat.firstName AS firstNameFather, " +
		"fat.secondName AS secondnameFather, " +
		"mot.firstName AS firstNameMother, " +
		"mot.secondName AS secondnameMother " +
		"FROM VPO_Baptism bap " +
		"INNER JOIN VPO_Person per ON per.idBaptism = bap.idBaptism " +
		"LEFT JOIN vpo_person fat ON fat.idPerson = per.idFather " +
		"LEFT JOIN VPO_Person mot ON mot.idPerson = per.idMother " +
		"LIMIT ?;"

	rows, err := db.QueryContext(ctx, query, limit)

	if err == sql.ErrNoRows {
		return partidas, lib.ErrNotFound
	}

	count := 1
	for rows.Next() {
		err := rows.Scan(
			&partida.IDBaptism,
			&partida.Firstname,
			&partida.Secondname,
			&partida.Lastname,
			&partida.Secondlastname,
			&partida.FirstnameFather,
			&partida.SecondnameFather,
			&partida.FirstnameMother,
			&partida.SecondnameMother,
		)

		if err != nil {
			return partidas, err
		}

		partida.Position = count
		count++
		partidas = append(partidas, partida)
	}

	return partidas, nil
}

func (*repoPerson) GetBaptizedPersonByFilter(ctx context.Context, filter string) ([]models.Baptism, error) {
	partida := models.Baptism{}
	partidas := []models.Baptism{}

	query := "SELECT " +
		"bap.idBaptism, " +
		"per.firstName, " +
		"per.secondName, " +
		"per.lastName, " +
		"per.secondLastname, " +
		"fat.firstName AS firstNameFather, " +
		"fat.secondName AS secondnameFather, " +
		"mot.firstName AS firstNameMother, " +
		"mot.secondName AS secondnameMother " +
		"FROM VPO_Baptism bap " +
		"INNER JOIN VPO_Person per ON per.idBaptism = bap.idBaptism " +
		"LEFT JOIN vpo_person fat ON fat.idPerson = per.idFather " +
		"LEFT JOIN VPO_Person mot ON mot.idPerson = per.idMother " +
		"WHERE  CONCAT(per.firstName, ' ', per.secondName, ' ', per.lastName, ' ', per.secondLastName) " +
		"LIKE ?;"

	rows, err := db.QueryContext(ctx, query, filter)

	if err == sql.ErrNoRows {
		return partidas, lib.ErrNotFound
	}

	count := 1
	for rows.Next() {
		err := rows.Scan(
			&partida.IDBaptism,
			&partida.Firstname,
			&partida.Secondname,
			&partida.Lastname,
			&partida.Secondlastname,
			&partida.FirstnameFather,
			&partida.SecondnameFather,
			&partida.FirstnameMother,
			&partida.SecondnameMother,
		)

		if err != nil {
			return partidas, err
		}

		partida.Position = count
		count++
		partidas = append(partidas, partida)
	}

	return partidas, nil
}

func (*repoPerson) GetBaptizedPartida(ctx context.Context, id string) (models.Baptism, error) {
	baptized := models.Baptism{}
	query := "SELECT * FROM batismforpdf WHERE idBaptism = ?;"

	row := db.QueryRowContext(ctx, query, id).Scan(
		&baptized.IDBaptism,
		&baptized.NumberBaptism,
		&baptized.Folio,
		&baptized.Book,
		&baptized.Borndate,
		&baptized.BaptismDate,
		&baptized.Firstname,
		&baptized.Secondname,
		&baptized.Lastname,
		&baptized.Secondlastname,
		&baptized.Sex,
		&baptized.FirstnameFather,
		&baptized.SecondnameFather,
		&baptized.LastnameFather,
		&baptized.SecondlastnameFather,
		&baptized.FirstnameMother,
		&baptized.SecondnameMother,
		&baptized.LastnameMother,
		&baptized.SecondlastnameMother,
		&baptized.FirstnameGodfather,
		&baptized.SecondnameGodfather,
		&baptized.LastnameGodfather,
		&baptized.SecondlastnameGodfather,
		&baptized.FirstnameGodmother,
		&baptized.SecondnameGodmother,
		&baptized.LastnameGodmother,
		&baptized.SecondlastnameGodmother,
		&baptized.FirstnameManager,
		&baptized.SecondnameManager,
		&baptized.LastnameManager,
		&baptized.SecondlastnameManager,
		&baptized.FirstnamePriest,
		&baptized.SecondnamePriest,
		&baptized.LastnamePriest,
		&baptized.SecondLastnamePriest,
	)

	if row == sql.ErrNoRows {
		return baptized, lib.ErrNotFound
	}

	return baptized, nil
}

func (*repoPerson) DeleteBaptizedPartida(ctx context.Context, id string) error {

	query := "CALL baptism_delete(?);"

	_, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (*repoPerson) GetManyPersonByFilter(ctx context.Context, limit int, filter string) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	var query string

	if filter == "GMF" {
		query = "SELECT idPerson, firstName, secondName, lastName, secondLastName " +
			" FROM VPO_Person " +
			"ORDER BY firstName ASC " +
			" limit ?;"

		rows, err := db.QueryContext(ctx, query, limit)
		if err == sql.ErrNoRows {
			return persons, lib.ErrNotFound
		}

		if err != nil {
			return persons, nil
		}

		var count int = 1
		for rows.Next() {

			if err := rows.Scan(
				&person.ID,
				&person.Firstname,
				&person.Secondname,
				&person.Lastname,
				&person.Secondlastname,
			); err != nil {
				return persons, err
			}
			person.Count = count
			persons = append(persons, person)
			count++
		}

	} else {
		query = "SELECT idPerson, firstName, secondName, lastName, secondLastName " +
			" FROM VPO_Person " +
			" WHERE  CONCAT(firstName, ' ', secondName, ' ', lastName, ' ', secondLastName) " +
			"LIKE ? " +
			"ORDER BY firstName ASC " +
			" limit ?;"

		rows, err := db.QueryContext(ctx, query, "%"+filter+"%", limit)
		if err == sql.ErrNoRows {
			return persons, lib.ErrNotFound
		}

		if err != nil {
			return persons, nil
		}

		var count int = 1
		for rows.Next() {

			if err := rows.Scan(
				&person.ID,
				&person.Firstname,
				&person.Secondname,
				&person.Lastname,
				&person.Secondlastname,
			); err != nil {
				return persons, err
			}

			person.Count = count
			persons = append(persons, person)
			count++
		}
	}

	return persons, nil
}

func (*repoPerson) UpdatePerson(ctx context.Context, person models.Person) error {
	querydb := "UPDATE VPO_Person SET " +
		"firstName = ?, " +
		"secondName = ?, " +
		"lastName = ?, " +
		"secondLastName = ?, " +
		"sex = ? " +
		"WHERE idPerson = ?; "

	_, err := db.QueryContext(
		ctx,
		querydb,
		person.Firstname,
		person.Secondname,
		person.Lastname,
		person.Secondlastname,
		person.Sexo,
		person.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
