package deps

import (
	"time"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/embeddeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/iodeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/requestdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/verbdeps"
)

// Deps is the dependency contract every adapter must satisfy. It is a
// struct of function fields, not an interface: an adapter fills every
// field with the behavior it provides, and the library calls those fields
// directly.
//
// Each field is one injectable behavior. The brain's own behavior is built
// from six of them — the clock every task stamps a record with, the pause the
// `watch` command sleeps between ticks, the formatted writer the command-line
// interface reports through, the argv parser it reads its flags with, the
// schema database every registry is persisted in, and the filesystem every
// task file is read from and every visualization is written to.
//
// The contract is wider than this one brain uses. Because the repository is
// a template meant to be forked, EmbedDeps and NewRequest are carried here as
// standing capabilities: a derived brain reads embedded assets or speaks HTTP
// without designing a contract for it first, and the standard adapter already
// fills both. The brain in sandbox/ never calls them, which is the point —
// they are offered, not required. Every field is documented in PublicApi.md
// whether the brain exercises it or not.
//
// VerbLib, KeepLib, EmbedDeps, IoLib and NewRequest show what changes when
// the dependency is another library built with this same pattern: the whole
// library arrives as one struct field, with no getter method and no bridging
// type around it. The sandbox never imports Verb, Keep, the `embed`
// machinery, os, or net/http — it declares a copy of each api in verbdeps,
// keepdeps, embeddeps, iodeps and requestdeps, and the adapter, which lives
// outside the sandbox, fills it.
type Deps struct {
	// Now returns the current time. It is what decides which month is the
	// open one, and what every record a task writes is stamped with.
	Now func() time.Time
	// Sleep pauses for the given duration. It is the only way the library
	// waits: the `watch` command runs a tick, sleeps for the interval it was
	// given, and runs another — so the loop never busy-waits and a test can
	// drive a hundred ticks without spending a hundred seconds.
	Sleep func(d time.Duration)
	// Printf writes one formatted message to the interface's output,
	// returning the number of bytes written and the write failure, if any.
	// It is the only way the library emits text: api.Lib.Sandboxmain — the
	// command-line interface — reports every result, every error, and its
	// usage screen through it, so the sandbox never touches a stream itself.
	Printf func(format string, a ...any) (n int, err error)
	// VerbLib is the embedded Verb argv-parser library, already initialized
	// by the adapter over the argument vector that adapter chose. Every
	// command word and every `--flag` of the interface is drained from it.
	VerbLib verbdeps.Lib
	// KeepLib is the embedded Keep schema-database library, already wired
	// by the adapter to the storage backend that adapter chose. The five
	// registries the brain keeps — accounts, categories, transactions,
	// recurrences and credit cards — live in it.
	KeepLib keepdeps.Lib
	// EmbedDeps is the embedded-asset library, already rooted by the adapter
	// at the asset directory that adapter chose. It is how the sandbox reads
	// the files shipped under /assets/ — the default Task.yaml and
	// Visualization.yaml the `start` command writes, templates, long-form
	// text, images — without importing the `embed` machinery itself.
	EmbedDeps embeddeps.Lib
	// IoLib is the filesystem library, reaching whatever root the adapter
	// gave it: reading and writing files, creating directories, testing
	// existence, and listing a tree one level deep or all the way down.
	//
	// It is what a tick is made of: Task.yaml and Visualization.yaml are read
	// through it, every rendered visualization is written through it, and
	// Error.md is created through it.
	IoLib iodeps.Lib
	// NewRequest opens an HTTP request against the given url, handing back
	// the request object the caller configures — method, headers, body — and
	// then sends with Fetch. It is a function field rather than a library
	// struct because a request is created per call, not injected once.
	//
	// The brain does not use it: nothing it does leaves the machine. A
	// derived brain that must speak HTTP — a task pulling a bank statement,
	// say — has the contract ready, filled by the standard adapter over
	// net/http.
	NewRequest func(url string) requestdeps.Request
}
