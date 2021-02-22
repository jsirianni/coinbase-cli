package logger

type Logger interface {
    Error(msg error)
    Warn(msg string)
    Info(msg string)
    Trace(msg string)
}
