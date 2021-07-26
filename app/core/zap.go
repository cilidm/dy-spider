package core

import (
	"fmt"
	"github.com/cilidm/dy-spider/app/constant"
	"os"
	"path"

	"github.com/cilidm/toolbox/file"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"time"
)

var (
	level zapcore.Level
	zType = constant.ZapLogType
)

func InitZap() (logger *zap.Logger) {
	file.IsNotExistMkDir(constant.ZapDirector)
	switch Conf.Zaplog.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(zType), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore(zType))
	}
	if Conf.Zaplog.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  Conf.Zaplog.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case Conf.Zaplog.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case Conf.Zaplog.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case Conf.Zaplog.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case Conf.Zaplog.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if Conf.Zaplog.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(zapType int) (core zapcore.Core) {
	var writer zapcore.WriteSyncer
	switch zapType {
	case 1:
		writer = GetWriteSyncerLumber()
	case 2:
		writer, err = GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	default:
		writer = GetWriteSyncerLumber()
	}
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(Conf.Zaplog.Prefix + " 2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(constant.ZapDirector, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if Conf.Zaplog.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func GetWriteSyncerLumber() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: path.Join(constant.ZapDirector, time.Now().Format("2006-01-02")+".log"), // 日志输出文件
		MaxSize:  5,                                                                       // 日志最大保存5M
		MaxAge:   30,                                                                      // 最多保留30个日志 和MaxBackups参数配置1个就可以
		Compress: false,                                                                   // 自导打 gzip包 默认false
		//MaxBackups: 5,                                                                       // 日志保留5个备份
	}
	return zapcore.AddSync(lumberJackLogger)
}
