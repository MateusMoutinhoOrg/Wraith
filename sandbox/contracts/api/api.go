package api

import (
	"github.com/MateusMoutinhoOrg/Agnos-Cli/sandbox/contracts/deps"
)

type Lib struct {
	Deps deps.Deps

	Sandboxmain func(args []string) int
}
