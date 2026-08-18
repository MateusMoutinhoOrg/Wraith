package commands

import (
	"time"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Watch runs a tick every interval, until the process is interrupted. It is
// how the vault is normally used: leave it running, edit `Task.yaml` in your
// editor, save, and the pages redraw themselves.
//
// A failing tick does not stop the loop. The failure is printed and written
// to Error.md, the task file is disarmed by the tick itself, and the next
// interval comes around — so a typo costs you one tick, not the session.
func Watch(l *api.Lib, quiet bool) int {
	value, err := l.Deps.VerbLib.GetStringOption([]string{config.TimeFlag}, 0)
	if err != nil || value == "" {
		return UsageError(l, config.MissingInterval)
	}
	interval, parseErr := time.ParseDuration(value)
	if parseErr != nil || interval <= 0 {
		return UsageError(l, config.InvalidInterval, value)
	}
	if !quiet {
		l.Deps.Printf(config.Watching, interval.String())
	}
	for {
		if err := l.PerformFullTick(); err != nil {
			l.Deps.Printf(config.Failed, err.Error())
		}
		l.Deps.Sleep(interval)
	}
}
