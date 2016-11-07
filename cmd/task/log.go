package task

import (
	"github.com/lcaballero/griller/config"
	"log"
)

// Log outputs log writes when the --debug flag is present on the
// command line.
type Log struct {
	conf *config.Conf
}

// NewLog requires a Conf to conditionally output logging.
func NewLog(conf *config.Conf) *Log {
	return &Log{
		conf: conf,
	}
}

// Println simply outputs each argument with a common separator.
func (g *Log) Println(args ...interface{}) {
	if !g.conf.Debug {
		return
	}
	log.Println(args...)
}

// Printf interpolates the values into the given format string.
func (g *Log) Printf(format string, args ...interface{}) {
	if !g.conf.Debug {
		return
	}
	log.Printf(format, args...)
}
