# `deps.Deps`

**Type:** Struct

## Definition

```go
type Deps struct {
	Now        func() time.Time
	Sleep      func(d time.Duration)
	Printf     func(format string, a ...any) (n int, err error)
	VerbLib    verbdeps.Lib
	KeepLib    keepdeps.Lib
	EmbedDeps  embeddeps.Lib
	IoLib      iodeps.Lib
	NewRequest func(url string) requestdeps.Request
}
```

## Description

The dependency contract every adapter must fill. A filled `Deps` is built by an adapter — see [`standard.New`](/docs/References/PublicApi/standard.New.md) — and passed to [`lib.New`](/docs/References/PublicApi/lib.New.md).

Six fields are what the brain is built from: `Now` is the clock (injecting it fixes what a record is stamped with and which month is the open one, which is what makes a brain testable), `Sleep` is the pause `watch` waits with between two ticks (injecting it lets a test drive a hundred ticks without spending a hundred seconds), `Printf` is the writer the command-line interface reports through, `VerbLib` is how that interface reads its command line, `KeepLib` is the schema database every registry is persisted in, and `IoLib` is how a tick reads `Task.yaml` and writes every rendered page.

`EmbedDeps` is what `wraith start` copies the default `Task.yaml` and `Visualization.yaml` out of. `NewRequest` is a **standing capability**: the brain never calls it, and the contract carries it anyway because this repository is a template meant to be forked — a derived brain whose task pulls a bank statement finds the contract already declared and already filled. It is offered, not required.

`Printf` is the library's only way of emitting text: `Sandboxmain` prints every result, every error, and its usage screen through it, so the sandbox never touches a stream itself and the interface can be run against a buffer as easily as against a terminal. What it prints is written in [`sandbox/config`](/docs/References/Structure.md#sandboxconfig) as compile-time constants, so changing the interface's wording is editing that package and every reference stays under the compiler's eye.

`VerbLib`, `KeepLib`, `EmbedDeps` and `IoLib` are the exceptions to "every field is a function": the dependency is itself a library built with this pattern, so it arrives as one plain struct field — [`verbdeps.Lib`](/docs/References/PublicApi/verbdeps.Lib.md), [`keepdeps.Lib`](/docs/References/PublicApi/keepdeps.Lib.md), [`embeddeps.Lib`](/docs/References/PublicApi/embeddeps.Lib.md), [`iodeps.Lib`](/docs/References/PublicApi/iodeps.Lib.md) — with no getter around it. `NewRequest` stays a function because a [`requestdeps.Request`](/docs/References/PublicApi/requestdeps.Request.md) is created per call rather than injected once. The sandbox never imports the embedded Verb or Keep libraries, Go's `embed` machinery, `os`, or `net/http`; it declares a copy of each api under `sandbox/contracts/deps/`, and the adapter, which lives outside the sandbox, initializes the real library and assigns its fields onto that copy.

Because it is a struct and not an interface, a value returned by an adapter can be patched field by field before injection, and a custom contract needs no type declaration at all. The trade-off: the compiler cannot detect a field you forgot to fill — it stays nil and panics on first call. See [StructContracts.md](/docs/References/StructContracts.md) and [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md).

## Fields

| Field | Description |
| :--- | :--- |
| `Now func() time.Time` | Returns the current time — what a record is stamped with, and what decides which month is the open one. |
| `Sleep func(d time.Duration)` | Pauses for a duration. The only way the library waits: `watch` sleeps between two ticks. |
| `Printf func(format string, a ...any) (n int, err error)` | Writes one formatted message to the interface's output — the only way the library emits text. |
| `VerbLib verbdeps.Lib` | The embedded Verb argv parser, already initialized by the adapter over the argument vector that adapter chose. |
| `KeepLib keepdeps.Lib` | The embedded Keep schema database every registry is stored in, already wired by the adapter to the storage backend that adapter chose. |
| `EmbedDeps embeddeps.Lib` | The embedded assets under [`/assets/`](/docs/References/Structure.md#assets), including the defaults `wraith start` writes, already rooted by the adapter at the asset directory that adapter chose. |
| `IoLib iodeps.Lib` | The filesystem: how a tick reads `Task.yaml`, writes every rendered page, and reports a failure in `Error.md`. |
| `NewRequest func(url string) requestdeps.Request` | Opens an HTTP request bound to `url`, handed back for the caller to configure and send. A standing capability: nothing this brain does leaves the machine. |

## Examples

```go
package main

import (
	"fmt"
	"time"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	// Start from an adapter — it is what fills KeepLib, the one dependency
	// the brain cannot work without — then patch the single behavior this
	// program wants to control.
	d := wraithadapter.New("my-brain")

	frozen := time.Unix(0, 0)
	d.Now = func() time.Time { return frozen }

	l := wraithlib.New(d, "data")

	l.PerformTask("AddAccount", map[string]any{"account": "Bank"})

	// Every page renders against the frozen clock, so "today" is 1-jan-1970
	// however long the test takes.
	renders, _ := l.PerformVisualization("DashBoard", map[string]any{})
	fmt.Println(string(renders[0].Content))
}
```
