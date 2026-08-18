//go:build ignore

// This file is an illustrative sample, not part of the build.
package main

import (
	"os"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

const (
	// vaultRoot is the folder the library reads and writes in: the one the
	// command was run from.
	vaultRoot = "."
	// databaseDirName is the folder inside it the records live in.
	databaseDirName = "data"
)

// main is the whole executable: wire, run, exit. No command is branched on
// here and nothing is printed here — that all lives inside the sandbox.
func main() {
	// 1. Build deps through the adapter (the opinionated layer).
	deps := wraithadapter.New(vaultRoot)

	// 2. Inject them into the pure library.
	l := wraithlib.New(deps, databaseDirName)

	// 3. Run the interface and exit with its return — the same os.Args[1:]
	//    the adapter wired the argv parser over.
	os.Exit(l.Sandboxmain(os.Args[1:]))
}
