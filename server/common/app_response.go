package common

type appRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewAppResponse(data, paging, filter interface{}) *appRes {
	return &appRes{data, paging, filter}
}

func NewSimpleAppResponse(data interface{}) *appRes {
	return &appRes{data, nil, nil}
}
