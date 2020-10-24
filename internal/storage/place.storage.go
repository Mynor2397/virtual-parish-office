package storage

import (
	"context"
	"database/sql"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*repoPerson) GetPlaces(ctx context.Context) ([]models.Place, error) {
	place := models.Place{}
	places := []models.Place{}

	// var datas []interface{}

	query := "SELECT idPlace, place, description FROM VPO_Place;"

	rows, err := db.QueryContext(ctx, query)

	if err == sql.ErrNoRows {
		return places, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&place.ID, &place.Name, &place.Description)
		if err != nil {
			return places, err
		}

		places = append(places, place)
	}

	return places, nil
}

func (*repoPerson) CreatePlace(ctx context.Context, place models.Place) (int, error) {
	var count int

	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return count, err
	}

	defer trans.Rollback()

	query := "INSERT INTO VPO_Place VALUES (?,?,?);"

	_, err = db.QueryContext(ctx, query, place.ID, place.Name, place.Description)
	if err != nil {
		return count, err
	}

	querycount := "SELECT COUNT(*) AS COUNT FROM VPO_Place;"
	res := db.QueryRowContext(ctx, querycount).Scan(&count)
	if res != nil {
		return count, err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return count, errtrans
	}

	return count, nil
}

func (*repoPerson) GetLastPlace(ctx context.Context) (int, error) {
	var count int
	query := "SELECT COUNT(*) AS COUNT FROM VPO_Place;"

	if err := db.QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
