package logger

import (
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	WarnLevel
	ErrorLevel
	CriticalLevel
)

var levelNames = []string{"INFO", "WARN", "ERROR", "CRITICAL"}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
}

type Logger struct {
	log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Log(level LogLevel, message string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     levelNames[level],
		Message:   message,
	}

	l.log.Printf("[%s] %s: %s\n", entry.Timestamp.Format(time.RFC3339), entry.Level, entry.Message)

	streamer, err := streamer.NewLogStreamer()
	if err != nil {
		log.Fatalf("Failed to create log streamer: %v", err)
	}

	streamer.SendLog(entry.Level, entry.Message, nil)
}

func (l *Logger) Info(message string) {
	l.Log(InfoLevel, message)
}

func (l *Logger) Warn(message string) {
	l.Log(WarnLevel, message)
}

func (l *Logger) Error(message string) {
	l.Log(ErrorLevel, message)
}

func (l *Logger) Critical(message string) {
	l.Log(CriticalLevel, message)
}
