package cli

// The command-line interface of the brain, written entirely inside the
// sandbox. It reads the command line through the injected Verb argv parser
// (deps.Deps.VerbLib), takes every word it displays from sandbox/config, and
// writes every line through the injected formatted writer (deps.Deps.Printf),
// so the whole program stays free of OS-bound and third-party imports — the
// process only hands it an argument vector and exits with the code it
// returns.
//
// No display text is written here: the usage screen, the version, and each
// message below are constants in sandbox/config, and this package addresses
// them by name. Changing what the interface says is editing that package —
// which keeps every message under the compiler's eye, so a renamed constant
// is a build failure rather than a blank line at runtime.
//
// Like sandbox/lib/vault, this package is neither an object package nor the
// entry point: it declares no types and no factories, and is called by
// SandboxmainFactory in sandbox/lib/publicfunctions.

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/cli/commands"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Run is the body of api.Lib.Sandboxmain: it dispatches one command line and
// returns the process exit code. args is only read to detect an empty command
// line; the arguments themselves are drained through the injected Verb
// parser, which the adapter wired over the same vector.
//
// The order below is deliberate. The flags that answer on their own come
// first, then the three path flags — every command takes them, and reading
// them here means no command has to — and only then the command word, which
// by that point is the first thing left that nothing has claimed.
func Run(l *api.Lib, args []string) int {
	if len(args) == 0 {
		l.Deps.Printf("%s", config.Usages)
		return api.ExitUsage
	}

	verb := l.Deps.VerbLib
	if verb.IsPresent(config.HelpFlags) {
		l.Deps.Printf("%s", config.Usages)
		return api.ExitOk
	}
	if verb.IsPresent(config.VersionFlags) {
		return commands.Version(l)
	}
	quiet := verb.IsPresent(config.QuietFlags)
	commands.Paths(l)

	command, err := verb.GetNextStringArg()
	if err != nil {
		return commands.UsageError(l, config.NoCommand)
	}

	switch command {
	case "help":
		l.Deps.Printf("%s", config.Usages)
		return api.ExitOk
	case "version":
		return commands.Version(l)
	case "start":
		return commands.Start(l, quiet)
	case "tick":
		return commands.Tick(l, quiet)
	case "watch":
		return commands.Watch(l, quiet)
	case "run":
		return commands.Run(l, quiet)
	case "render":
		return commands.Render(l, quiet)
	case "tasks":
		return commands.Tasks(l)
	case "visualizations":
		return commands.Visualizations(l)
	}
	return commands.UsageError(l, config.UnknownCommand, command)
}
