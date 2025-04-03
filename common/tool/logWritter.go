package tool

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

// LogWriter 日志处理器
type LogWriter struct {
	serverName string
	logger     *logx.RotateLogger
}

// Info 写入Info日志
func (l *LogWriter) Info(format string, v ...any) {
	l.writeLog("info", format, v...)
}

// 写日志
func (l *LogWriter) writeLog(tag, format string, v ...any) {
	// 获取调用者信息（文件路径 + 行号）
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		file = filepath.Base(file) // 仅保留文件名，避免路径过长
	}
	// 设置日志时间
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00") // logx 的时间格式
	logMsg := fmt.Sprintf(format, v...)
	// 设置日志内容
	formattedLog := fmt.Sprintf("%s\t %s \t%s\tcaller=%s:%d\n", timestamp, tag, logMsg, file, line)

	// 写入日志
	_, err := l.logger.Write([]byte(formattedLog)) // 写入日志
	if err != nil {
		logx.Error(fmt.Sprintf("【%s】无法写入日志: %v", l.serverName, err))
		return
	}
}

// 初始化日志
func iniLogWriter(serverName, logPath string) *LogWriter {
	// 确保日志目录存在
	_ = os.MkdirAll(logPath, os.ModePerm)

	// 设置日志文件
	logFilePath := path.Join(logPath, fmt.Sprintf("%s.log", serverName))

	// 定义日志轮转规则（按天分割）
	rule := logx.DefaultRotateRule(
		logFilePath, // 日志路径
		"-",         // 用 "-" 分隔历史日志
		7,           // 只保留最近 7 天的日志
		false,       // 不启用 gzip 压缩
	)

	// 创建 RotateLogger
	logger, err := logx.NewLogger(logFilePath, rule, false) // false: 不压缩
	if err != nil {
		logx.Error(fmt.Sprintf("【%s】无法创建日志文件: %v", serverName, err))
		return nil
	}

	return &LogWriter{
		serverName: serverName,
		logger:     logger,
	}
}

// logWriter
var logWriter *LogWriter

// GetLogWriter 获取日志记录器
func GetLogWriter(serverName, logPath string) *LogWriter {
	if logWriter == nil {
		logWriter = iniLogWriter(serverName, logPath)
	}
	return logWriter
}
