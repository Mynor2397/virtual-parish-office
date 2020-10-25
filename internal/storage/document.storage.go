package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*repoDocument) GetFolio(ctx context.Context, idBook int) ([]models.Folio, error) {
	folio := models.Folio{}
	folios := []models.Folio{}

	query := "SELECT idFolio, numberFolio, idBook FROM VPO_Folio WHERE idBook = ?;"
	rows, err := db.QueryContext(ctx, query, idBook)

	if err == sql.ErrNoRows {
		return folios, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&folio.IDFolio, &folio.NumberFolio, &folio.IDBook)

		if err != nil {
			return folios, err
		}

		folios = append(folios, folio)
	}

	return folios, nil
}

func (*repoDocument) Baptism(ctx context.Context, baptism models.Baptism) (string, error) {
	query := "call registbaptism(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := db.QueryContext(ctx,
		query,
		baptism.IDAddress,
		baptism.Address,
		baptism.IDBaptism,
		baptism.Folio,
		baptism.BaptismDate,
		baptism.IDPriest,
		baptism.ID,
		baptism.IDBaptized,
		baptism.Firstname,
		baptism.Secondname,
		baptism.Lastname,
		baptism.Secondlastname,
		baptism.Borndate,
		baptism.DPI,
		baptism.Sex,
		baptism.IDFather,
		baptism.IDMother,
		baptism.IDGodfather,
		baptism.IDGodMother,
		baptism.IDManager,
	)

	if err != nil {
		return "", err
	}

	return baptism.IDBaptism, nil
}

func (*repoDocument) GetBaptisms(ctx context.Context, filter string) ([]models.Baptism, error) {
	baptism := models.Baptism{}
	baptisms := []models.Baptism{}

	query := "select bap.idBaptism, per.firstName, per.secondName, per.lastName, per.secondLastname from VPO_Baptism bap " +
		" INNER JOIN VPO_Person per on per.idBaptism = bap.idBaptism;"
	rows, err := db.QueryContext(ctx, query)

	if err == sql.ErrNoRows {
		return baptisms, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&baptism.IDBaptism, &baptism.Firstname, &baptism.Secondname, &baptism.Lastname, &baptism.Secondlastname)
		if err != nil {
			return baptisms, err
		}

		baptisms = append(baptisms, baptism)
	}

	return baptisms, nil
}

func (*repoDocument) GetBook(ctx context.Context) ([]models.Book, error) {
	book := models.Book{}
	books := []models.Book{}

	query := "SELECT IDBook, NumberBook FROM VPO_baptismBook;"

	rows, err := db.QueryContext(ctx, query)

	if err == sql.ErrNoRows {
		return books, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&book.IDBook, &book.NumberBook)
		if err != nil {
			return books, err
		}

		books = append(books, book)
	}
	return books, nil
}

func (*repoDocument) CreateBook(ctx context.Context, book models.Book) error {
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer trans.Rollback()

	query := "INSERT INTO VPO_BaptismBook (idBook, numberBook, startDate, Commentary) VALUES(?,?,?,?)"
	_, err = db.QueryContext(ctx, query, book.IDBook, book.NumberBook, book.StartDate, book.Commentary)

	if err != nil {
		return err
	}

	for i := 1; i <= book.Folios; i++ {
		queryfolio := "INSERT INTO VPO_Folio(numberFolio, idBook) VALUES(?,?);"
		_, errfolio := db.QueryContext(ctx, queryfolio, i, book.IDBook)

		if err != nil {
			return errfolio
		}
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return errtrans
	}

	return nil
}

func (*repoDocument) GetLastBook(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) AS LAST FROM VPO_BaptismBook;"
	var last int
	err := db.QueryRowContext(ctx, query).Scan(&last)
	if err != nil {
		return 0, err
	}
	return last, nil
}

func (*repoDocument) GetAudit(ctx context.Context) ([]models.Audit, error) {
	audi := models.Audit{}
	audits := []models.Audit{}

	query := "SELECT u.userName, hb.dateEmitted, b.numberBaptism FROM vpo_user u "+
		     "INNER JOIN vpo_baptismhistory hb ON u.idUser = hb.idUser " +
			 " INNER JOIN vpo_baptism b ON b.idBaptism = hb.idBaptism;"

	data, err := db.QueryContext(ctx, query)
	if err == sql.ErrNoRows {
		return audits, lib.ErrNotFound
	}

	for data.Next() {
		err := data.Scan(&audi.UserName, &audi.DateEmited, &audi.IdBaptism)
		if err != nil {
			return audits, err
		}

		audits = append(audits, audi)
	}

	return audits, nil
}

func (*repoDocument) CreateAudi(ctx context.Context, audit models.Audit) error {
	query := "INSERT INTO VPO_BaptismHistory (idUser, idBaptism) VALUES (?,?)"

	_, err := db.QueryContext(ctx, query,
		audit.IdUser,
		audit.IdBaptism)

	if err != nil {
		return err
	}

	return nil
}
