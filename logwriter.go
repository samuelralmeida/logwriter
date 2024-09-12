package logwriter

import (
	"encoding/json"
	"fmt"
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

func (l *LogWriter) write(text string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.file.WriteString(text)
	return err
}

func (l *LogWriter) Write(a ...any) error {
	text := fmt.Sprint(a...)
	return l.write(text)
}

func (l *LogWriter) Writeln(a ...any) error {
	text := fmt.Sprintln(a...)
	return l.write(text)
}

func (l *LogWriter) Writef(format string, a ...any) error {
	text := fmt.Sprintf(format, a...)
	return l.write(text)
}

func (l *LogWriter) WriteAsJson(msg string, fields map[string]any) error {
	if fields == nil {
		fields = map[string]any{}
	}

	fields["message"] = msg

	json, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return l.write(string(json))
}
