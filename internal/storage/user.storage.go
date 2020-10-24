package storage

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/Mynor2397/virtual-parish-office/internal/helper"
	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*repoUser) Create(ctx context.Context, user *models.User) (string, error) {
	user.Username = strings.TrimSpace(user.Username)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	query := "INSERT INTO VPO_User (idUser, userName, password, uuidRol) values (?, ?, ?, ?);"
	_, err = db.QueryContext(ctx, query, user.ID, user.Username, string(hashedPassword), user.Rol)

	if err != nil {
		log.Println(err)
		return "", lib.ErrDuplicateUser
	}
	return user.ID, nil

}

func (*repoUser) Login(ctx context.Context, user *models.User) (models.User, error) {
	var response models.User
	var passwordClient string

	query := "SELECT u.idUser, u.userName, u.password, r.typeRol FROM VPO_User u "
	query += "INNER JOIN VPO_userRole r ON u.uuidRol = r.uuidRol "
	query += "WHERE binary userName = ?;"

	row := db.QueryRowContext(ctx, query, user.Username).Scan(&user.ID, &user.Username, &passwordClient, &user.Rol)

	if row == sql.ErrNoRows {
		return response, lib.ErrUserNotFound
	}

	if row != nil {
		return response, row
	}

	hashedPasswordDatabase := []byte(passwordClient)
	valuePassword := bcrypt.CompareHashAndPassword(hashedPasswordDatabase, []byte(user.Password))
	if valuePassword != nil {
		return response, lib.ErrUserNotFound
	}

	user.Password = ""
	response.Username = user.Username
	response.Rol = user.Rol

	token := helper.GenerateJWT(user)
	response.Token = token

	return response, nil
}

func (*repoUser) Update(ctx context.Context, id, rol string) error {
	query := "UPDATE VPO_User SET uuidRol=? WHERE idUser = ?;"

	_, err := db.QueryContext(ctx, query, rol, id)

	if err != nil {
		return err
	}

	return nil
}
func (*repoUser) GetManyUsers(ctx context.Context) ([]models.User, error) {
	user := models.User{}
	users := []models.User{}

	query := "SELECT u.idUser, u.userName, r.typeRol FROM VPO_User u " +
		 	 "INNER JOIN VPO_UserRole r ON u.uuidRol = r.uuidRol;"
	rows, err := db.QueryContext(ctx, query)
	if err == sql.ErrNoRows {
		return users, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Rol)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
func (*repoUser) Roles(ctx context.Context) ([]models.Rol, error) {
	rol := models.Rol{}
	rols := []models.Rol{}

	query := "SELECT uuidRol, typeRol FROM VPO_UserRole;"
	rows, err := db.QueryContext(ctx, query)

	if err == sql.ErrNoRows {
		return rols, lib.ErrNotFound
	}

	for rows.Next(){
		err := rows.Scan(&rol.IDRol, &rol.TypeRol)
		if err != nil{
			return rols, err
		}
		rols = append(rols, rol)
	}

	return rols, err
}
