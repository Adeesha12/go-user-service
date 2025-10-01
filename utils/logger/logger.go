package logger

import "log"

var (
	Info  = log.New(log.Writer(), "\033[32mINFO:\033[0m ", log.Ldate|log.Ltime)
	Warn  = log.New(log.Writer(), "\033[33mWARN:\033[0m ", log.Ldate|log.Ltime)
	Error = log.New(log.Writer(), "\033[31mERROR:\033[0m ", log.Ldate|log.Ltime)
)
