package api

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
