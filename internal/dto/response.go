package dto

type Response struct {
	ErrorCode int
	Data      *any   `json:"data"`
	Error     string `json:"error"`
}
