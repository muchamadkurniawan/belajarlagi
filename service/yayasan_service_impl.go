package service

import (
	"belajarlagi/helper"
	"belajarlagi/model/web"
	"belajarlagi/repository"
	"context"
	"database/sql"
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
	//TODO implement me
	panic("implement me")
}

func (service *YayasanServiceImpl) Update(ctx context.Context, request web.YayasanUpdateRequest) web.YayasanResponse {
	//TODO implement me
	panic("implement me")
}

func (service *YayasanServiceImpl) Delete(ctx context.Context, id int) {
	//TODO implement me
	panic("implement me")
}

func (service *YayasanServiceImpl) FindById(ctx context.Context, id int) web.YayasanResponse {
	//TODO implement me
	panic("implement me")
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
