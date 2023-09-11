package config

type App struct {
	Name   string `json:"name" yaml:"name"`
	Listen string `json:"listen" yaml:"listen"`
}

type Logger struct {
	Path     string `json:"path" yaml:"path"`
	Level    string `json:"level" yaml:"level"`
	PostLog  string `json:"postLog" yaml:"postLog"`
	ThirdLog string `json:"thirdLog" yaml:"thirdLog"`
}

type Api struct {
	VmList     string `json:"vmList" yaml:"vmList"`
	NicList    string `json:"nicList" yaml:"nicList"`
	MetricData string `json:"metricData" yaml:"metricData"`
	SendMsg    string `json:"sendMsg" yaml:"sendMsg"`
}

type PaoPaoUser struct {
	GroupId  string `json:"groupId" yaml:"groupId"`
	ToCustId string `json:"toCustId" yaml:"toCustId"`
}

type Mysql struct {
	Conn string `json:"conn" yaml:"conn"`
}

type Mongo struct {
	Conn string `json:"conn" yaml:"conn"`
	Db   string `json:"db" yaml:"db"`
}

type runtimeParam struct {
	RootDir string `json:"-" yaml:"-"` // 此软件运行后的工作目录
}

type Configs struct {
	App          App          `json:"app" yaml:"app"`
	Logger       Logger       `json:"logger" yaml:"logger"`
	Api          Api          `json:"api" yaml:"api"`
	PaoPaoUser   PaoPaoUser   `json:"paoPaoUser" yaml:"paoPaoUser"`
	Mysql        Mysql        `json:"mysql" yaml:"mysql"`
	Mongo        Mongo        `json:"mongo" yaml:"mongo"`
	Spec         string       `json:"spec" yaml:"spec"`
	RuntimeParam runtimeParam `json:"-" yaml:"-"`
}

// Cfg 全局的Config配置，解析dns.yaml的结果
var Cfg *Configs

var configFileName = "lss.yml"
