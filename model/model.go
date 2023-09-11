package model

// User 用户
type User struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"colum:name"`
}

func (User) TableName() string {
	return "user"
}

type ApiLog struct {
	Api          string      `json:"api" bson:"api"`
	RequestTime  string      `json:"requestTime" bson:"requestTime"`
	RequestBody  interface{} `json:"requestBody" bson:"requestBody"`
	ResponseTime string      `json:"responseTime" bson:"responseTime"`
	Response     interface{} `json:"response" bson:"response"`
	Err          string      `json:"err" bson:"err"`
}

type PostLog struct {
	Time          string            `json:"time" bson:"time"`
	ResponseTime  string            `json:"responseTime" bson:"responseTime"`
	TTL           int               `json:"ttl" bson:"ttl"`
	AppName       string            `json:"appName" bson:"appName"`
	Method        string            `json:"method" bson:"method"`
	ContentType   string            `json:"contentType" bson:"contentType"`
	Uri           string            `json:"uri" bson:"uri"`
	ClientIP      string            `json:"clientIP" bson:"clientIP"`
	RequestHeader map[string]string `json:"requestHeader" bson:"requestHeader"`
	RequestParam  any               `json:"requestParam" bson:"requestParam"`
	RequestBody   any               `json:"requestBody" bson:"requestBody"`
	ResponseStr   string            `json:"responseStr" bson:"responseStr"`
	ResponseMap   any               `json:"responseMap" bson:"responseMap"`
}
