package repository

import (
	"belajarlagi/app"
	"belajarlagi/model/domain"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestGetAll(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := NewYayasanRepository()
	var yayasans []domain.Yayasan
	yayasans = repo.GetAll(context.Background(), tx)
	for _, yayasan := range yayasans {
		fmt.Println("id yayasan : ", yayasan.Id)
		fmt.Println("nama yayasan : ", yayasan.Nama)
		fmt.Println("user name yayasan : ", yayasan.Uname)
		fmt.Println("password yayasan : ", yayasan.Pass)
		fmt.Println("-----------------------")
	}
}

func TestGetById(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := NewYayasanRepository()
	yayasan := domain.Yayasan{}
	id := 90
	yayasan, err1 := repo.GetById(context.Background(), tx, id)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("------------------------")
	fmt.Println("id yayasan : ", yayasan.Id)
	fmt.Println("nama yayasan : ", yayasan.Nama)
	fmt.Println("user name yayasan : ", yayasan.Uname)
	fmt.Println("password yayasan : ", yayasan.Pass)
}

func TestSaveYayasan(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var yayasan = domain.Yayasan{
		Nama:  "uswah2",
		Uname: "U_uswah2",
		Pass:  "passUswah2",
	}
	repo := NewYayasanRepository()
	save := repo.Save(context.Background(), tx, yayasan)
	tx.Commit()
	fmt.Println(save.Id, " ", save.Nama)
}

func TestUpdateYayasan(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := NewYayasanRepository()
	yayasan := domain.Yayasan{
		Id:    1,
		Nama:  "update islam",
		Uname: "update uname islam",
		Pass:  "pass islma",
	}
	repo.Update(context.Background(), tx, yayasan)
	tx.Commit()
}

func TestDeleteYayasan(t *testing.T) {
	db := app.NewDB()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	repo := NewYayasanRepository()
	yayasan := domain.Yayasan{
		Id: 4,
	}
	repo.Delete(context.Background(), tx, yayasan)
	tx.Commit()
}
