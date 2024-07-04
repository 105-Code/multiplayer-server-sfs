package logger

import (
	"fmt"
	"html/template"
	"os"

	"github.com/105-Code/multiplayer-server-sfs/pkg/config"
)

var m = map[string]string{
	"reset":       "\033[0m",
	"cyan":        "\033[36m",
	"boldgray":    "\033[0;37m",
	"blue":        "\033[0;34m",
	"lightblue":   "\033[1;34m",
	"yellow":      "\033[0;33m",
	"lightyellow": "\033[1;33m",
	"green":       "\033[0;32m",
	"lightgreen":  "\033[1;32m",
	"red":         "\033[0;31m",
}

func WithColors(msg string, args ...interface{}) {
	//agregar aqui opción si se desea hacer quiet los comandos
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	t := template.Must(template.New("").Parse(msg))
	t.Execute(os.Stdout, m)
}

func Error(msg string, args ...interface{}) {
	errorMsg := "{{.red}}[ERROR] {{.reset}}" + msg + "\n"
	if len(args) > 0 {
		errorMsg = fmt.Sprintf(msg, args...)
	}
	t := template.Must(template.New("").Parse(errorMsg))
	t.Execute(os.Stdout, m)
}

func Warn(msg string, args ...interface{}) {
	warnMsg := "{{.yellow}}[WARNING] {{.reset}}" + msg + "\n"
	WithColors(warnMsg, args...)
}

func Info(msg string, args ...interface{}) {
	infoMsg := "[INFO] " + msg + "\n"
	WithColors(infoMsg, args...)
}

func Debug(msg string, args ...interface{}) {
	if !config.AppConfig.Debug {
		return
	}
	debugMsg := "{{.boldgray}}[DEBUG] {{.reset}}" + msg + "\n"
	WithColors(debugMsg, args...)
}
