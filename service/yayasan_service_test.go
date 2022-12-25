package service

import (
	"belajarlagi/app"
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
}
