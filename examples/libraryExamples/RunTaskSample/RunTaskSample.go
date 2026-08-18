package main

// RunTaskSample — running a task from Go, with no command line in sight.
//
// The interface you drive from a terminal is one field of the library like
// any other, so everything it can do is reachable from code. This example
// wires the standard adapter into the sandbox and calls api.Lib.PerformTask
// directly, which is exactly what `wraith run` does one layer up.
//
// Run it with:
//   go run ./examples/libraryExamples/RunTaskSample/RunTaskSample.go

import (
	"fmt"
	"os"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	// A vault is a folder, so this example works in one of its own and
	// removes it on the way out.
	vault, err := os.MkdirTemp("", "wraith-sample")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(vault)

	// 1. Build deps through the adapter, rooted at the vault.
	deps := wraithadapter.New(vault)

	// 2. Inject them, naming the folder the registries live in.
	l := wraithlib.New(deps, "data")

	// 3. Run tasks. The entries map is the same one a Task.yaml decodes to,
	//    so anything the file can say, this call can say too.
	if err := l.PerformTask("AddAccount", map[string]any{
		"account": "Bank",
		"opening": 3000,
	}); err != nil {
		panic(err)
	}
	if err := l.PerformTask("AddCategory", map[string]any{
		"category":    "Food",
		"description": "Groceries and eating out",
		"revenues":    false,
		"expenses":    true,
	}); err != nil {
		panic(err)
	}
	if err := l.PerformTask("AddTransaction", map[string]any{
		"account":     "Bank",
		"category":    "Food",
		"amount":      -32.90,
		"date":        "2026-08-18",
		"description": "Market",
	}); err != nil {
		panic(err)
	}

	// 4. Every task the binary carries is on the library, so a program can
	//    ask what it is allowed to do rather than hard-coding a list.
	for _, task := range l.Tasks {
		fmt.Printf("%-20s %s\n", task.Name, task.Description)
	}
}
