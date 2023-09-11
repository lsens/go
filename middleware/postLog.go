package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lss/config"
	"lss/dao"
	"lss/model"
	"lss/utils"
	"strings"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		data, err := c.GetRawData()
		if err != nil {
			config.Log.Error("GetRawData error:", err.Error())
		}
		body := string(data)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点

		// 处理请求
		c.Next()
		responseBody := bodyLogWriter.body.String()
		// 日志格式
		if strings.Contains(c.Request.RequestURI, "/docs") || c.Request.RequestURI == "/" {
			return
		}

		var result any
		if responseBody != "" && responseBody[0:1] == "{" {
			err := json.Unmarshal([]byte(responseBody), &result)
			if err != nil {
				result = map[string]any{"status": -1, "msg": "解析异常:" + err.Error()}
			}
		}

		// 结束时间
		endTime := time.Now()
		// 日志格式
		var params, reqBody any
		if strings.Contains(c.ContentType(), "application/json") && body != "" {
			utils.FromJSON(body, &reqBody)
		}
		params = utils.GinParamMap(c)
		postLog := new(model.PostLog)
		postLog.Time = startTime.Format("2006-01-02 15:04:05")
		postLog.Uri = c.Request.RequestURI
		postLog.Method = c.Request.Method
		postLog.AppName = config.Cfg.App.Name
		postLog.ContentType = c.ContentType()
		postLog.RequestHeader = utils.GinHeaders(c)
		ip := c.GetHeader("X-Forward-For")
		if ip == "" {
			ip = c.GetHeader("X-Real-IP")
			if ip == "" {
				ip = c.ClientIP()
			}
		}
		postLog.ClientIP = ip
		postLog.RequestParam = params
		postLog.RequestBody = reqBody
		postLog.ResponseTime = endTime.Format("2006-01-02 15:04:05")
		postLog.ResponseMap = result
		postLog.TTL = int(endTime.UnixNano()/1e6 - startTime.UnixNano()/1e6)

		dao.Ld.PostLog(*postLog)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
