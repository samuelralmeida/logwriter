# logwriter

logwriter is a simple Go package for writing log messages to a file.

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

### Write

```go

func (l *logWriter) Write(text string) error
```

Writes a log message to the file.

#### Parameters:

    text: The log message to be written.

#### Returns:

    error: An error if there was a problem writing to the file.

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
