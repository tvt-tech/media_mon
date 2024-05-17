package utils

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.InfoLevel)
	// Logger.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors:          false,
	// 	FullTimestamp:          false,
	// 	TimestampFormat:        "15:04:05",
	// 	DisableLevelTruncation: true,
	// 	DisableQuote:           true,

	// })
	Logger.SetFormatter(&customFormatter{})
}

// customFormatter is a custom log formatter
type customFormatter struct{}

// Format formats the log entry
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := entry.Level.String()
	message := entry.Message

	// Format the log entry as "LEVEL: message"
	formatted := fmt.Sprintf("%s: %s\n", strings.ToUpper(level), message)
	return []byte(formatted), nil
}
