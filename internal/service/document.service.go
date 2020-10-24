package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*documentService) GetFolio(ctx context.Context, idBook int) ([]models.Folio, error) {
	return myDocumentStorage.GetFolio(ctx, idBook)
}

func (*documentService) Baptism(ctx context.Context, baptism models.Baptism) (string, error) {
	baptism.IDBaptized = uuid.New().String()
	baptism.IDBaptism = uuid.New().String()
	baptism.IDAddress = uuid.New().String()
	return myDocumentStorage.Baptism(ctx, baptism)
}

func (*documentService) GetBaptisms(ctx context.Context, filter string) ([]models.Baptism, error) {
	return myDocumentStorage.GetBaptisms(ctx, filter)
}

func (*documentService) GetBook(ctx context.Context) ([]models.Book, error) {
	return myDocumentStorage.GetBook(ctx)
}

func (*documentService) CreateBook(ctx context.Context, book models.Book) error {
	last, err := myDocumentStorage.GetLastBook(ctx)

	if err != nil {
		return err
	}

	book.IDBook = last + 1
	return myDocumentStorage.CreateBook(ctx, book)
}

func (*documentService) GetCountBook(ctx context.Context) (int, error) {
	return myDocumentStorage.GetLastBook(ctx)
}

func (*documentService) GetAudit(ctx context.Context) ([]models.Audit, error) {
	return myDocumentStorage.GetAudit(ctx)
}

func (*documentService) CreateAudi(ctx context.Context, idUser string, audit models.Audit) error {
	audit.IdUser = idUser

	return myDocumentStorage.CreateAudi(ctx, audit)
}
