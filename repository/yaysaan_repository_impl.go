package repository

import (
	"belajarlagi/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type YayasanRepositoryImpl struct {
}

func NewYayasanRepository() YayasanRepository {
	return &YayasanRepositoryImpl{}
}

func (repository *YayasanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan) domain.Yayasan {
	SQL := "insert into yayasan(nama,uname,pass) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, yayasan.Nama, yayasan.Uname, yayasan.Pass)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	id, err1 := result.LastInsertId()
	if err1 != nil {
		panic(err1)
	}
	yayasan.Id = int(id)
	fmt.Println("yayasan id = ", yayasan.Pass)
	return yayasan
}

func (repository *YayasanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan) domain.Yayasan {
	SQL := "update yayasan set nama=?,uname=?,pass=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, yayasan.Nama, yayasan.Uname, yayasan.Pass, yayasan.Id)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	return yayasan
}

func (repository *YayasanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, yayasan domain.Yayasan) {
	SQL := "delete from yayasan where id=?"
	_, err := tx.ExecContext(ctx, SQL, yayasan.Id)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}

func (repository *YayasanRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Yayasan {
	SQL := "select id, nama, pass, uname from yayasan"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	var yayasans []domain.Yayasan
	defer rows.Close()
	for rows.Next() {
		yayasan := domain.Yayasan{}
		rows.Scan(&yayasan.Id, &yayasan.Nama, &yayasan.Pass, &yayasan.Uname)
		yayasans = append(yayasans, yayasan)
	}
	return yayasans
}

func (repository *YayasanRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (domain.Yayasan, error) {
	SQL := "select id, nama, pass, uname from yayasan where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	yayasan := domain.Yayasan{}
	if rows.Next() {
		err = rows.Scan(&yayasan.Id, &yayasan.Nama, &yayasan.Uname, &yayasan.Pass)
		if err != nil {
			panic(err)
		}
		return yayasan, nil
	} else {
		return yayasan, errors.New("data tidak ditemukan")
	}
}
