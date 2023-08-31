package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"githup.com/apibe/yifu/cmd/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var Logger *zap.SugaredLogger
var Writer *rotatelogs.RotateLogs

const logFilePath = "./"
const logFileName = "apibe"

func Init() {
	// 根据mode设置日志输出格式
	level := zapcore.InfoLevel
	switch config.C.Mod {
	case "debug":
		level = zapcore.DebugLevel
	case "test":
		level = zapcore.InfoLevel
	case "release":
		level = zapcore.ErrorLevel
	}
	// 日志输出相关配置
	fileName := path.Join(logFilePath, logFileName)
	Writer, _ = rotatelogs.New(
		fileName+".%Y%m%d.log",                  // 分割后的文件名称
		rotatelogs.WithLinkName(fileName),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(365*24*time.Hour), // 设置最大保存时间(365天)
		rotatelogs.WithRotationTime(time.Hour),  // 设置日志切割时间间隔(1小时)
	)
	writer := zapcore.AddSync(Writer)
	// debug 模式输出级别为debugLevel
	if level == zapcore.DebugLevel {
		writer = zapcore.AddSync(os.Stdout)
	}
	// 格式相关的配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间戳的格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别使用大写显示
	encoder := zapcore.NewConsoleEncoder(encoderConfig)     // 用json格式化的日志格式

	// 设置日志级别
	core := zapcore.NewCore(encoder, writer, level) // 将日志级别设置为 DEBUG
	logger := zap.New(core, zap.AddCaller(), zap.Fields())
	Logger = logger.Sugar()
}
