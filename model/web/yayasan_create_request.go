package web

type YayasanCreateRequest struct {
	Nama  string `validate:"required, min = 1, max = 100:" json:"nama"`
	Uname string `validate:"required, min = 1, max = 100:" json:"nama"`
	Pass  string `validate:"required, min = 1, max = 100:" json:"nama"`
}
