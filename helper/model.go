package helper

import (
	"belajarlagi/model/domain"
	"belajarlagi/model/web"
)

func ToYayasanResponses(yayasans []domain.Yayasan) []web.YayasanResponse {
	var yayasanResponses []web.YayasanResponse
	for _, yayasan := range yayasans {
		yayasanResponses = append(yayasanResponses, ToYayasanResponse(yayasan))
	}
	return yayasanResponses
}

func ToYayasanResponse(yayasan domain.Yayasan) web.YayasanResponse {
	return web.YayasanResponse{
		Id:    yayasan.Id,
		Nama:  yayasan.Nama,
		Uname: yayasan.Uname,
		Pass:  yayasan.Pass,
	}
}
