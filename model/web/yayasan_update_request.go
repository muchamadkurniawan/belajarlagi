package web

type YayasanUpdateRequest struct {
	Id    int    `validate:"required"`
	Name  string `validate:"required, max=200,min=1" json:"name""`
	Uname string `validate:"required, max=200,min=1" json:"uname"`
	Pass  string `validate:"required, max=200,min=1" json:"pass"`
}
