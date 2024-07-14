package logwriter

import (
	"os"
	"sync"
)

type LogWriter struct {
	file *os.File
	mu   sync.Mutex
}

func NewLogWriter(filename string) (*LogWriter, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &LogWriter{file: file}, nil
}

func (l *LogWriter) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.file.Close()
}

func (l *LogWriter) Write(text string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.file.WriteString(text)
	return err
}
