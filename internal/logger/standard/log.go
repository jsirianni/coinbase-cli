package standard

import (
    "fmt"
    "strings"
)

const (
    errorLVL = 0
    warnLVL  = 1
    infoLVL  = 2
    traceLVL = 3
)

type Standard struct {
    level int
}

func New(level string) Standard {
    switch strings.ToUpper(level) {
    case "ERROR":
        return Standard{errorLVL}
    case "WARN", "WARNING":
        return Standard{warnLVL}
    case "INFO":
        return Standard{infoLVL}
    case "TRACE":
        return Standard{traceLVL}
    default:
        fmt.Println(fmt.Sprintf("Invalid log level %s, defaulting to INFO", level))
        return Standard{infoLVL}
    }
}

func (s Standard) Error(msg error) {
    if s.level >= errorLVL {
        fmt.Println(msg.Error())
    }
}

func (s Standard) Warn(msg string) {
    if s.level >= warnLVL {
        fmt.Println(msg)
    }
}

func (s Standard) Info(msg string) {
    if s.level >= infoLVL {
        fmt.Println(msg)
    }
}

func (s Standard) Trace(msg string) {
    if s.level >= traceLVL {
        fmt.Println(msg)
    }
}
