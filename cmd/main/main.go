package main

import (
	"os"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

const (
	// vaultRoot is the directory the brain reads and writes in: the one the
	// command was run from. A vault is a folder, so the brain you drive is
	// the folder you are standing in — which is what lets one binary serve
	// as many vaults as you have directories.
	vaultRoot = "."
	// databaseDirName is the folder inside the vault the registries are
	// persisted in. The `--database` flag overrides it for one invocation.
	databaseDirName = "data"
)

// main is the whole executable: it wires an adapter into the library, hands
// the command line to api.Lib.Sandboxmain — the interface itself, which lives
// inside the sandbox — and exits with the code it returns. Choosing where the
// vault lives is an OS-bound decision, so it is made here, outside the
// sandbox, and never inside it.
func main() {
	// 1. Build deps via the standard adapter: a real clock, a real sleep,
	//    standard output, the Verb parser over os.Args[1:], a Keep database
	//    on disk, the filesystem rooted at the vault, and the defaults
	//    compiled into this binary.
	deps := wraithadapter.New(vaultRoot)

	// 2. Inject them into the pure library, telling it which folder of the
	//    vault its registries live in.
	l := wraithlib.New(deps, databaseDirName)

	// 3. Run the interface and exit with its return — the same os.Args[1:]
	//    the adapter wired the parser over.
	os.Exit(l.Sandboxmain(os.Args[1:]))
}
