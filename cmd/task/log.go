package task

import (
	"github.com/lcaballero/griller/config"
	"log"
)

type Log struct {
	conf *config.Conf
}

func NewLog(conf *config.Conf) *Log {
	return &Log{
		conf: conf,
	}
}

func (g *Log) Println(args ...interface{}) {
	if !g.conf.Debug {
		return
	}
	log.Println(args...)
}

func (g *Log) Printf(format string, args ...interface{}) {
	if !g.conf.Debug {
		return
	}
	log.Printf(format, args...)
}
