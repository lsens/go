package model

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	//Page    *Page  `json:"page,omitempty"`
}

type Page struct {
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Count int `json:"count"`
}

func RError(code int, msg string) Res {
	return Res{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
func RSuccess(data interface{}) Res {
	return Res{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

type DataWithPage struct {
	List any `json:"list"`
	Page any `json:"page"`
}

func RSuccessWithPage(data interface{}, page *Page) Res {
	return Res{
		Code:    200,
		Message: "success",
		Data: map[string]interface{}{
			"list": data,
			"page": page,
		},
	}
}
