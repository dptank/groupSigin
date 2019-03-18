package ginlog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
/**
*控制台打印日志信息
*/
func LogPrint(str ...interface{}) bool {
	zerolog.TimeFieldFormat = ""
	log.Print(str)
	return true
}