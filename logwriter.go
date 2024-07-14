package logwriter

import "os"

type logWriter struct {
	file *os.File
}

func NewLogWriter(filename string) (*logWriter, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &logWriter{file: file}, nil
}

func (l *logWriter) Close() error {
	return l.file.Close()
}

func (l *logWriter) Write(text string) error {
	_, err := l.file.WriteString(text)
	return err
}
