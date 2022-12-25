package service

import (
	"belajarlagi/model/web"
	"context"
)

type YayasanService interface {
	Create(ctx context.Context, request web.YayasanCreateRequest) web.YayasanResponse
	Update(ctx context.Context, request web.YayasanUpdateRequest) web.YayasanResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.YayasanResponse
	FindAll(ctx context.Context) []web.YayasanResponse
}
