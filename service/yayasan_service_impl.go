package service

import (
	"belajarlagi/helper"
	"belajarlagi/model/domain"
	"belajarlagi/model/web"
	"belajarlagi/repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

type YayasanServiceImpl struct {
	YayasanRepository repository.YayasanRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewYayasanService(yayasanRepository repository.YayasanRepository, DB *sql.DB, validate *validator.Validate) YayasanService {
	return &YayasanServiceImpl{
		YayasanRepository: yayasanRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *YayasanServiceImpl) Create(ctx context.Context, request web.YayasanCreateRequest) web.YayasanResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	yayasan := domain.Yayasan{
		Nama:  request.Nama,
		Uname: request.Uname,
		Pass:  request.Pass,
	}
	service.YayasanRepository.Save(ctx, tx, yayasan)
	if err != nil {
		return web.YayasanResponse{}
	}
	return helper.ToYayasanResponse(yayasan)
}

func (service *YayasanServiceImpl) Update(ctx context.Context, request web.YayasanUpdateRequest) web.YayasanResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	_, err1 := service.YayasanRepository.GetById(ctx, tx, request.Id)
	if err1 != nil {
		fmt.Println(err1)
	}
	yayasan := domain.Yayasan{
		Nama:  request.Name,
		Uname: request.Uname,
		Pass:  request.Pass,
		Id:    request.Id,
	}
	service.YayasanRepository.Update(ctx, tx, yayasan)
	return helper.ToYayasanResponse(yayasan)
}

func (service *YayasanServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	byId, err1 := service.YayasanRepository.GetById(ctx, tx, id)
	if err1 != nil {
		fmt.Println(err1)
	}
	service.YayasanRepository.Delete(ctx, tx, byId)
}

func (service *YayasanServiceImpl) FindById(ctx context.Context, id int) web.YayasanResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	yayasan, err := service.YayasanRepository.GetById(ctx, tx, id)
	if err != nil {
		return web.YayasanResponse{}
	}
	return helper.ToYayasanResponse(domain.Yayasan(yayasan))
}

func (service *YayasanServiceImpl) FindAll(ctx context.Context) []web.YayasanResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	getAll := service.YayasanRepository.GetAll(ctx, tx)
	return helper.ToYayasanResponses(getAll)
}
