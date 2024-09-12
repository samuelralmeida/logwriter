# logwriter

`logwriter` is a simple Go package for writing log messages to a file. It is safe for use with goroutines, allowing multiple goroutines to write to the log file concurrently without causing race conditions.

## Installation

To install the package, use the following command:

```bash
go get github.com/samuelralmeida/logwriter
```

## Usage

Here's a basic example of how to use the logwriter package:

```go

package main

import (
    "log"
    "github.com/samuelralmeida/logwriter"
)

func main() {
    lw, err := logwriter.NewLogWriter("logfile.log")
    if err != nil {
        log.Fatalf("Failed to create log writer: %v", err)
    }
    defer lw.Close()

    err = lw.Write("This is a log message.\n")
    if err != nil {
        log.Fatalf("Failed to write to log file: %v", err)
    }
}
```

## Goroutine Safety

The logwriter package is safe for use with goroutines. You can safely use the same logWriter instance across multiple goroutines. Here is an example demonstrating concurrent writes:

```go
package main

import (
    "log"
    "sync"
    "github.com/yourusername/logwriter"
)

func main() {
    lw, err := logwriter.NewLogWriter("logfile.log")
    if err != nil {
        log.Fatalf("Failed to create log writer: %v", err)
    }
    defer lw.Close()

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            err := lw.Write(fmt.Sprintf("Log message from goroutine %d\n", i))
            if err != nil {
                log.Printf("Failed to write log message: %v", err)
            }
        }(i)
    }
    wg.Wait()
}
```

## API

### NewLogWriter

```go

func NewLogWriter(filename string) (*logWriter, error)
```

Creates a new logWriter that writes to the specified file. If the file does not exist, it will be created. If the file exists, new log messages will be appended to the file.

#### Parameters:

    filename: The name of the log file.

#### Returns:

    *logWriter: A pointer to the created logWriter.
    error: An error if there was a problem creating or opening the file.

### Writing

```go
func (l *LogWriter) Write(a ...any) error
```

This method formats and writes the provided arguments as a log message. It converts the arguments into a single string using fmt.Sprint, then writes to the file.

**Parameters:**

`a ...any`: A variadic parameter that accepts any number of arguments.

**Returns:**

`error`: An error, if any occurred during the write operation.

```go
func (l *LogWriter) Writeln(a ...any) error
```

This method formats and writes the provided arguments as a log message with a newline appended. It converts the arguments into a single string using fmt.Sprintln, then writes to the file.

**Parameters:**

`a ...any`: A variadic parameter that accepts any number of arguments.

**Returns:**

`error`: An error, if any occurred during the write operation.

```go
func (l *LogWriter) Writef(format string, a ...any) error
```

This method formats and writes a log message according to the specified format string. It uses fmt.Sprintf to format the message, then writes to the file.

**Parameters:**

`format string`: A format string as per fmt.Sprintf syntax.

`a ...any`: Additional arguments to be formatted according to the format string.

**Returns:**

`error`: An error, if any occurred during the write operation.

```go
func (l *LogWriter) WriteAsJson(msg string, fields map[string]any) error
```

This method writes a log message in JSON format. It includes the provided message and additional fields in the JSON object. If the fields map is nil, it initializes it to an empty map. The JSON is then written to the file.

**Parameters:**

`msg string`: The log message to include in the JSON object.

`fields map[string]any`: A map of additional fields to include in the JSON object. If nil, it defaults to an empty map.

**Returns:**

`error`: An error, if any occurred during the JSON marshaling or write operation.


### Close

```go
func (l *logWriter) Close() error
```
Closes the log file.

#### Returns:

    error: An error if there was a problem closing the file.

## Contributing

If you would like to contribute to this project, please open an issue or submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
Contact
