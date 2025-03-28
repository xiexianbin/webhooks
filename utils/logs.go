package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

// consoleLogs for dev env
var consoleLogs *logs.BeeLogger

// fileLogs for production env
var fileLogs *logs.BeeLogger

// run mode
var runmode string

// init logs
func InitLogs() {
	level := beego.AppConfig.String("logs::level")
	logFile := beego.AppConfig.String("logs::file")

	// console logs
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async()

	// file logs
	fileLogs = logs.NewLogger(10000)
	fileLogs.SetLogger(logs.AdapterMultiFile, `{
		"filename": "logs/"`+logFile+`,
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
		"level":`+level+`,
		"daily": true,
		"maxdays": 10}`)
	fileLogs.Async()

	// run mode
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

func LogEmergency(v interface{}) {
	log("emergency", v)
}

func LogAlert(v interface{}) {
	log("alert", v)
}
func LogCritical(v interface{}) {
	log("critical", v)
}

func LogError(v interface{}) {
	log("error", v)
}

func LogWarning(v interface{}) {
	log("warning", v)
}

func LogNotice(v interface{}) {
	log("notice", v)
}
func LogInfo(v interface{}) {
	log("info", v)
}

func LogDebug(v interface{}) {
	log("debug", v)
}

func LogTrace(v interface{}) {
	log("trace", v)
}

// Logs
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, v)
		case "alert":
			fileLogs.Alert(format, v)
		case "critical":
			fileLogs.Critical(format, v)
		case "error":
			fileLogs.Error(format, v)
		case "warning":
			fileLogs.Warning(format, v)
		case "notice":
			fileLogs.Notice(format, v)
		case "info":
			fileLogs.Info(format, v)
		case "debug":
			fileLogs.Debug(format, v)
		case "trace":
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
		}
	}

	switch level {
	case "emergency":
		consoleLogs.Emergency(format, v)
	case "alert":
		consoleLogs.Alert(format, v)
	case "critical":
		consoleLogs.Critical(format, v)
	case "error":
		consoleLogs.Error(format, v)
	case "warning":
		consoleLogs.Warning(format, v)
	case "notice":
		consoleLogs.Notice(format, v)
	case "info":
		consoleLogs.Info(format, v)
	case "debug":
		consoleLogs.Debug(format, v)
	case "trace":
		consoleLogs.Trace(format, v)
	default:
		consoleLogs.Debug(format, v)
	}
}
