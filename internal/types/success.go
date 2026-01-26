package types

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EntityRequest struct {
	Description string `json:"description"`
}
