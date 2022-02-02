package utils

import "github.com/rl404/fairy/log"

var l log.Logger

// InitLog to init global logger.
func InitLog(t log.LogType, lvl log.LogLevel, json, color bool) (err error) {
	l, err = log.New(log.Config{
		Type:       t,
		Level:      lvl,
		JsonFormat: json,
		Color:      color,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetLogger to get logger.
func GetLogger() log.Logger {
	if l == nil {
		l, _ = log.New(log.Config{
			Type:       log.Zerolog,
			Level:      log.TraceLevel,
			JsonFormat: false,
			Color:      true,
		})
	}
	return l
}

// Log to log with custom field.
func Log(field map[string]interface{}) {
	GetLogger().Log(field)
}

// Info to log info.
func Info(str string, args ...interface{}) {
	GetLogger().Info(str, args...)
}

// Error to print error.
func Error(str string, args ...interface{}) {
	GetLogger().Error(str, args...)
}

// Fatal to log fatal.
func Fatal(str string, args ...interface{}) {
	GetLogger().Fatal(str, args...)
}
