package logger

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Print(message string) {
	fmt.Println("Console:", message)
}

type SyslogLogger struct{}

func (s *SyslogLogger) WriteLog(text string) {
	fmt.Println("SYSLOG:", text)
}
