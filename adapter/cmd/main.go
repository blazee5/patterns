package main

import (
	"github.com/vladislavkutsenko/patterns/adapter/internal/adapter"
	"github.com/vladislavkutsenko/patterns/adapter/internal/logger"
)

func main() {
	consoleLogger := logger.ConsoleLogger{}
	syslogLogger := logger.SyslogLogger{}

	consoleAdapter := adapter.NewConsoleAdapter(&consoleLogger)
	syslogAdapter := adapter.NewSyslogAdapter(&syslogLogger)

	message := "This is a test message"

	consoleAdapter.Log(message)
	syslogAdapter.Log(message)
}
