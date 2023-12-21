package api

type BaseRequest struct {
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
