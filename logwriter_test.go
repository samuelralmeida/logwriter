package logwriter

import (
	"os"
	"testing"
)

func TestNewLogWriter(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}
	if lw == nil {
		t.Fatalf("Expected the log writer not to be nil")
	}
	lw.Close()
}

func TestLogWriter_Write(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}
	defer lw.Close()

	text := "Hello, log!"
	if err := lw.Write(text); err != nil {
		t.Fatalf("Expected no error when writing to the log file, but got: %v", err)
	}

	// Verify that the text was actually written to the file
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error when reading the log file, but got: %v", err)
	}
	if string(content) != text {
		t.Fatalf("Expected the file content to be '%s', but got '%s'", text, string(content))
	}
}

func TestLogWriter_Close(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}

	if err := lw.Close(); err != nil {
		t.Fatalf("Expected no error when closing the log file, but got: %v", err)
	}

	// Try writing to the file after closing to ensure it is really closed
	if err := lw.Write("This should fail"); err == nil {
		t.Fatalf("Expected an error when writing to the closed file, but did not get one")
	}
}
