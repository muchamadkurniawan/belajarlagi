package repository

import (
	"belajarlagi/model/domain"
	"context"
	"database/sql"
)

type YayasanRepository interface {
	Save(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan) domain.Yayasan
	Update(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan) domain.Yayasan
	Delete(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Yayasan
	GetById(ctx context.Context, tx *sql.Tx, id int) (domain.Yayasan, error)
}
