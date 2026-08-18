// Package commands holds one file per command of the interface. Each one
// reads what it needs off the injected Verb parser, calls one field of
// api.Lib, reports through the injected Printf, and returns the process exit
// code.
//
// No command does any work of its own. Everything a command can make happen
// is a function field on api.Lib, which is why the same behavior is reachable
// from Go code with no command line in sight.
package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// UsageError reports a command line that could not be understood, prints the
// usage screen after it, and returns the usage exit code.
func UsageError(l *api.Lib, format string, values ...any) int {
	l.Deps.Printf(format, values...)
	l.Deps.Printf("%s", config.Usages)
	return api.ExitUsage
}

// Failure reports an error the command could not recover from.
func Failure(l *api.Lib, failure error) int {
	l.Deps.Printf(config.Failed, failure.Error())
	return api.ExitError
}

// Version prints the interface version.
func Version(l *api.Lib) int {
	l.Deps.Printf("%s\n", config.Version)
	return api.ExitOk
}
