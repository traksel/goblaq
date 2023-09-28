package daemon

import (
	"time"

	"github.com/traksel/goblaq/internal/status"
)

func Run(mode string) {
	if mode == "daemon" {
		for {
			var s status.Status
			s.WriteStatus()
			time.Sleep(60 * time.Second)
		}
	}
}
