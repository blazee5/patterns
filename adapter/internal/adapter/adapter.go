package adapter

import "github.com/vladislavkutsenko/patterns/adapter/internal/logger"

type ConsoleAdapter struct {
	consoleLogger *logger.ConsoleLogger
}

func NewConsoleAdapter(consoleLogger *logger.ConsoleLogger) *ConsoleAdapter {
	return &ConsoleAdapter{
		consoleLogger: consoleLogger,
	}
}

func (c *ConsoleAdapter) Log(message string) {
	c.consoleLogger.Print(message)
}

type SyslogAdapter struct {
	syslogLogger *logger.SyslogLogger
}

func NewSyslogAdapter(syslogLogger *logger.SyslogLogger) *SyslogAdapter {
	return &SyslogAdapter{
		syslogLogger: syslogLogger,
	}
}

func (s *SyslogAdapter) Log(message string) {
	s.syslogLogger.WriteLog(message)
}
