package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

// UserCreate es el servicio de conexion al storage de crear usuario
func (*userService) Create(ctx context.Context, user *models.User) (string, error) {
	user.ID = uuid.New().String()

	return myUserstorage.Create(ctx, user)
}

// UserLogin es el servicio de conexion al storage de login de usuario
func (*userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	return myUserstorage.Login(ctx, user)
}

func (*userService) Update(ctx context.Context, id, rol string) error {
	return myUserstorage.Update(ctx, id, rol)
}

func (*userService) GetManyUsers(ctx context.Context) ([]models.User, error) {
	return myUserstorage.GetManyUsers(ctx)
}

func (*userService) Roles(ctx context.Context)([]models.Rol, error){
	return myUserstorage.Roles(ctx)
}