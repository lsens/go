package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

// MarshalJSON implements json.Marshaler.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	//entity your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	var ts int64
	ts, err = strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	theTime := time.Unix(ts, 0)
	*t = Timestamp(theTime)
	return nil
}

func (t Timestamp) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *Timestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Timestamp(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Timestamp) GetTime() time.Time {
	return time.Time(t)
}

// GetUnixTimeSql 获取unix时间戳sql
func GetUnixTimeSql(unixTime int64) string {
	return fmt.Sprintf("FROM_UNIXTIME(%d)", unixTime)
}

func FromJSON(j string, o interface{}) *interface{} {
	err := json.Unmarshal([]byte(j), &o)
	if err != nil {
		return nil
	} else {
		return &o
	}
}

func ToJSON(o interface{}) string {
	j, err := json.Marshal(o)
	if err != nil {
		return "{}"
	} else {
		js := string(j)
		js = strings.Replace(js, "\\u003c", "<", -1)
		js = strings.Replace(js, "\\u003e", ">", -1)
		js = strings.Replace(js, "\\u0026", "&", -1)
		return js
	}
}

func GinParamMap(c *gin.Context) map[string]string {
	params := make(map[string]string)
	if c.Request.Method == "GET" {
		for k, v := range c.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else if c.Request.Method == "POST" {
		if strings.Contains(c.ContentType(), "x-www-form-urlencoded") {
			c.Request.ParseForm()
			for k, v := range c.Request.PostForm {
				params[k] = v[0]
			}
			for k, v := range c.Request.URL.Query() {
				params[k] = v[0]
			}
		} else if strings.Contains(c.ContentType(), "multipart/form-data") {
			c.Request.ParseMultipartForm(100 * 1024 * 1024)
			for k, v := range c.Request.MultipartForm.Value {
				params[k] = v[0]
			}
			for k, v := range c.Request.URL.Query() {
				params[k] = v[0]
			}
		}
	}
	return params
}

func GinHeaders(c *gin.Context) map[string]string {
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}
	return headers
}

func CopyStruct(dest interface{}, src interface{}, tag string) error {

	destValue := reflect.ValueOf(dest).Elem()
	destKey := reflect.TypeOf(dest).Elem()
	srcValue := reflect.ValueOf(src)
	srcKey := reflect.TypeOf(src)

	if reflect.TypeOf(dest).Elem().Kind() != reflect.Struct || reflect.TypeOf(src).Kind() != reflect.Struct {
		fmt.Println("格式错误", reflect.TypeOf(dest).Elem().Kind(), reflect.TypeOf(src).Kind())
		return errors.New("格式错误")
	}

	srcKeyMap := make(map[string]int)

	for i1 := 0; i1 < srcKey.NumField(); i1++ {
		if srcKey.Field(i1).Tag.Get(tag) != "" {
			srcKeyMap[srcKey.Field(i1).Tag.Get(tag)] = i1
		} else {
			srcKeyMap[srcKey.Field(i1).Name] = i1
		}
	}

	num := 0
	if destValue.NumField() > srcValue.NumField() {
		num = destValue.NumField()
	} else {
		num = srcValue.NumField()
	}

	for i := 0; i < num; i++ {
		if i < destValue.NumField() {
			destValueF := destValue.Field(i)
			destKeyF := destKey.Field(i)

			descKeyName := ""
			if destKeyF.Tag.Get(tag) != "" {
				descKeyName = destKeyF.Tag.Get(tag)
			} else {
				descKeyName = destKeyF.Name
			}

			ite, ok := srcKeyMap[descKeyName]
			if ok {
				srcValueF := srcValue.Field(ite)
				if destValueF.Kind() == srcValueF.Kind() {
					destValueF.Set(srcValueF)
				}
			}
		}
	}
	fmt.Println("源：", src, "目标：", dest)
	return nil
}
