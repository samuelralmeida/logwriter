package logwriter

import "os"

type LogWriter struct {
	file *os.File
}

func NewLogWriter(filename string) (*LogWriter, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &LogWriter{file: file}, nil
}

func (l *LogWriter) Close() error {
	return l.file.Close()
}

func (l *LogWriter) Write(text string) error {
	_, err := l.file.WriteString(text)
	return err
}
