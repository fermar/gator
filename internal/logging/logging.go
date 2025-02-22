package logging

import (
	"io"
	"log"
	"os"
)

type Lgr struct {
	Logger  *log.Logger
	Enabled bool
}

var Lg = Lgr{Logger: log.New(io.Discard, "GatorLog:", log.LstdFlags), Enabled: false}

//	func NewLogger(status bool) error {
//		// Lg.Logger = log.New(io.Discard, "GatorLog:", log.LstdFlags)
//		Lg.Enabled = status
//		if Lg.Enabled {
//			Lg.EnLog()
//		} else {
//			Lg.DisLog()
//		}
//		return nil
//	}
func (lg *Lgr) EnLog() {
	lg.Enabled = true
	lg.Logger.SetOutput(os.Stdout)
	lg.Logger.Println("Log habilitado")
}

func (lg *Lgr) DisLog() {
	lg.Enabled = false
	lg.Logger.SetOutput(io.Discard)
	lg.Logger.Println("Log deshabilitado")
}
