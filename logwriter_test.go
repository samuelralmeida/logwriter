package logwriter_test

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/samuelralmeida/logwriter"
)

func TestNewLogWriter(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := logwriter.NewLogWriter(filename)
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

	lw, err := logwriter.NewLogWriter(filename)
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

	lw, err := logwriter.NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}

	if err := lw.Close(); err != nil {
		t.Fatalf("Expected no error when closing the log file, but got: %v", err)
	}

	if err := lw.Write("This should fail"); err == nil {
		t.Fatalf("Expected an error when writing to the closed file, but did not get one")
	}
}

func TestLogWriter_ConcurrentWrite(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := logwriter.NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}
	defer lw.Close()

	var wg sync.WaitGroup
	messages := []string{"Hello, log!", "Another log message", "Yet another log message"}

	for _, msg := range messages {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			if err := lw.Write(m + "\n"); err != nil {
				t.Errorf("Expected no error when writing to the log file, but got: %v", err)
			}
		}(msg)
	}

	wg.Wait()

	contentByte, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error when reading the log file, but got: %v", err)
	}

	contentString := string(contentByte)
	for _, msg := range messages {
		if !strings.Contains(contentString, msg) {
			t.Fatalf("Expected the file content to contain '%s'", msg)
		}
	}
}

func TestLogWriter_Writeln(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := logwriter.NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}
	defer lw.Close()

	text := "Hello, log!"
	if err := lw.Writeln(text); err != nil {
		t.Fatalf("Expected no error when writing to the log file, but got: %v", err)
	}

	// Verify that the text was actually written to the file
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error when reading the log file, but got: %v", err)
	}
	if string(content) != fmt.Sprintln(text) {
		t.Fatalf("Expected the file content to be '%s', but got '%s'", text, string(content))
	}
}

func TestLogWriter_Writef(t *testing.T) {
	filename := "test.log"
	defer os.Remove(filename)

	lw, err := logwriter.NewLogWriter(filename)
	if err != nil {
		t.Fatalf("Expected no error when creating the log writer, but got: %v", err)
	}
	defer lw.Close()

	text := "Hello, %s %d!\n"
	if err := lw.Writef(text, "world", 2024); err != nil {
		t.Fatalf("Expected no error when writing to the log file, but got: %v", err)
	}

	// Verify that the text was actually written to the file
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error when reading the log file, but got: %v", err)
	}
	if string(content) != fmt.Sprintf(text, "world", 2024) {
		t.Fatalf("Expected the file content to be '%s', but got '%s'", text, string(content))
	}
}
