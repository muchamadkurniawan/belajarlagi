package service

import (
	"belajarlagi/app"
	"belajarlagi/model/web"
	"belajarlagi/repository"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestNewYayasanService_FindAllFindAll(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepo := repository.NewYayasanRepository()
	YayasanService := NewYayasanService(yayasanRepo, db, validate)
	all := YayasanService.FindAll(context.Background())
	fmt.Println(all)
	for _, yayasan := range all {
		fmt.Println("id yayasan : ", yayasan.Id)
		fmt.Println("nama yayasan : ", yayasan.Nama)
		fmt.Println("user name yayasan : ", yayasan.Uname)
		fmt.Println("password yayasan : ", yayasan.Pass)
		fmt.Println("-----------------------")
	}
}
func TestYayasanServiceImpl_FindById(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepo := repository.NewYayasanRepository()
	YayasanService := NewYayasanService(yayasanRepo, db, validate)
	var id int = 3
	yayasan := YayasanService.FindById(context.Background(), id)
	fmt.Println("------------------------")
	fmt.Println("id yayasan : ", yayasan.Id)
	fmt.Println("nama yayasan : ", yayasan.Nama)
	fmt.Println("user name yayasan : ", yayasan.Uname)
	fmt.Println("password yayasan : ", yayasan.Pass)
}

func TestYayasanServiceImpl_Update(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepo := repository.NewYayasanRepository()
	YayasanService := NewYayasanService(yayasanRepo, db, validate)
	yayasan := web.YayasanUpdateRequest{
		Id:    2,
		Name:  "uswah4",
		Uname: "uname_uswah4",
		Pass:  "pass_uswah4",
	}
	YayasanService.Update(context.Background(), yayasan)
}

func TestYayasanServiceImpl_Delete(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepo := repository.NewYayasanRepository()
	YayasanService := NewYayasanService(yayasanRepo, db, validate)
	id := 3
	YayasanService.Delete(context.Background(), id)
}
