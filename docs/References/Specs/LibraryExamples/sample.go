//go:build ignore

// This file is an illustrative sample, not part of the build.
package main

import (
	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	// 1. Build deps through an adapter (the opinionated layer).
	deps := wraithadapter.New("my-brain")

	// 2. Inject deps into the pure library.
	l := wraithlib.New(deps, "data")

	// 3. Exercise the library — it never knows which adapter is behind it.
	//    Its functions are struct fields, called like any method.
	if err := l.PerformTask("ExampleTask", map[string]any{"example": "value"}); err != nil {
		panic(err)
	}
}
