package dtos

type ErrorPageDto struct {
	Header  HeaderDto
	Message string `json:"message"`
	Code    int    `json:"code"`
	Details string `json:"details"`
}
