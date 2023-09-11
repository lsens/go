package dao

import (
	"errors"
	"fmt"
	"lss/config"
	"lss/model"
	"time"
)

var Ld = &logDao{}

type logDao struct{}

func (log logDao) ApiLog(api, requestTime string, requestBody, response interface{}, resErr error) error {
	conn := config.DBObj.GetMongoConn()
	if conn == nil {
		fmt.Println("获取数据库连接错误")
		return errors.New("获取数据库连接错误")
	}
	obj := model.ApiLog{
		Api:          api,
		RequestBody:  requestBody,
		RequestTime:  requestTime,
		Response:     response,
		ResponseTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if resErr != nil {
		obj.Err = resErr.Error()
	}
	err := conn.C(config.Cfg.Logger.ThirdLog).Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

func (log logDao) PostLog(obj model.PostLog) error {
	conn := config.DBObj.GetMongoConn()
	if conn == nil {
		fmt.Println("获取数据库连接错误")
		return errors.New("获取数据库连接错误")
	}
	err := conn.C(config.Cfg.Logger.PostLog).Insert(&obj)
	if err != nil {
		return err
	}
	return nil
}
