package model

type Result struct {
	Status  int         `json:"status" bson:"status"`
	Message string      `json:"message" bson:"msg"`
	Data    interface{} `json:"data" bson:"data"`
}

func Error(code int, msg string) Result {
	return Result{
		Status:  code,
		Message: msg,
		Data:    nil,
	}
}

func Success(data interface{}) Result {
	return Result{
		Status:  200,
		Message: "success",
		Data:    data,
	}
}
