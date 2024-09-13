# glog - A Simple Logging Package for Go

`glog` is a custom logging package built on top of the popular [logrus](https://github.com/sirupsen/logrus) library. It allows for structured logging, supports multiple output targets (e.g., files and console), and includes the caller’s file name and line number in each log entry.

## Features

- **Multiple Output Destinations**: Logs can be written to both files and the console.
- **Structured Logging**: Add custom fields to log entries to enhance log structure.
- **Caller Info**: Logs include the function name, file, and line number of the log call.
- **Log Level Support**: Configure log levels to control the verbosity of logs.
- **Easy Integration**: Quickly integrate with your Go projects.

## Installation

Install the package by running:

```bash
go get github.com/EvvTim/glog
```

Import it into your project:

```go
import "github.com/EvvTim/glog"
```

## Usage

**Basic Logging**

To get started with logging, simply initialize the logger and start logging messages:

```go
package main

import (
    "github.com/EvvTim/glog"
)

func main() {
    // Get the logger instance
    log := logging.GetLogger()

    // Log messages
    log.Info("This is an info message")
    log.Warn("This is a warning message")
    log.Error("This is an error message")
}
```

**Log with Fields**

You can use structured logging by adding custom fields to your log entries:

```go
func main() {
    log := logging.GetLogger()

    log.WithField("user_id", 1234).Info("User login")
}
```

Or use the provided helper method:

```go
func main() {
    log := logging.GetLogger()

    // Log with additional fields
    logWithUser := log.GetLoggerWithField("user_id", 1234)
    logWithUser.Info("User login")
}
```

**Log Levels**

The package supports all log levels from logrus. You can configure the log level to control the verbosity of the logs.

```go
log.Trace("This is a trace message")
log.Debug("This is a debug message")
log.Info("This is an info message")
log.Warn("This is a warning message")
log.Error("This is an error message")
log.Fatal("This is a fatal message")
```

**Configuration**

You can customize the logger by modifying the following configurations in the init() function:

• **Log Formatter**: Customize the log format (e.g., plain text or JSON).
• **Log Level**: Set the minimum log level.
• **Output Writers**: Specify where the logs should be written (e.g., files, stdout).

**Changing the Log Level**

If you need to adjust the log level, modify the line in the init() function that sets the log level:
```go
l.SetLevel(logrus.TraceLevel) // Change to desired level: Trace, Debug, Info, Warn, Error, Fatal
```

**Adding More Output Destinations**

To write logs to multiple destinations, add more writers to the writerHook:
```go
l.AddHook(&writerHook{
    Writer:    []io.Writer{fileWriter, os.Stdout, anotherWriter}, // Add more writers
    LogLevels: logrus.AllLevels,
})
```

**Contributing**

Contributions are welcome! Feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/EvvTim/glog).

**License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
