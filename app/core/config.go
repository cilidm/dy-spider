package core

import (
	"sync"
)

var (
	once sync.Once
	Conf = &conf{}
)

type conf struct {
	DB     DBConf
	Redis  RedisConf
	App    SettingConf
	Zaplog ZapLogConf
}

type DBConf struct {
	DBType string
	DBUser string
	DBPwd  string
	DBHost string
	DBName string
}

type RedisConf struct {
	RedisAddr string
	RedisPWD  string
	RedisDB   int
}

type SettingConf struct {
	HttpPort     int `json:"http-port"`
	ReadTimeout  int `json:"read-timeout"`
	WriteTimeout int `json:"write-timeout"`
	RunMode      string
	PageSize     int
	JwtSecret    string
	UploadTmpDir string
	ImgSavePath  string
	ImgUrlPath   string
}

type ZapLogConf struct {
	Level         string `json:"level" yaml:"level"`
	Format        string ` json:"format" yaml:"format"`
	Prefix        string ` json:"prefix" yaml:"prefix"`
	Director      string ` json:"director"  yaml:"director"`
	LinkName      string ` json:"linkName" yaml:"link-name"`
	ShowLine      bool   ` json:"showLine" yaml:"showLine"`
	EncodeLevel   string ` json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `json:"logInConsole" yaml:"log-in-console"`
}

func InitConfig() {
	once.Do(func() {
		Conf.App.JwtSecret = "JwtSecret"
		Conf.App.UploadTmpDir = "static/upload"
		Conf.App.ImgSavePath = "static/upload"
		Conf.App.ImgUrlPath = "runtime/upload/images"
		Conf.App.HttpPort = 8009
		Conf.App.ReadTimeout = 60
		Conf.App.WriteTimeout = 60
		Conf.Zaplog.Level = "info"
		Conf.Zaplog.Format = "console"
		Conf.Zaplog.Prefix = "[base-spider]"
		Conf.Zaplog.Director = "runtime/log"
		Conf.Zaplog.LinkName = "latest_log"
		Conf.Zaplog.ShowLine = true
		Conf.Zaplog.EncodeLevel = "LowercaseColorLevelEncoder"
		Conf.Zaplog.StacktraceKey = "stacktrace"
		Conf.Zaplog.LogInConsole = true
	})
}
