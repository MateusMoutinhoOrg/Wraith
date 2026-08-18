# `lib.New`

**Type:** Function

## Signature

```go
func New(d deps.Deps, databasePath string) api.Lib
```

## Description

Injects a filled dependency contract into the library and returns the [`api.Lib`](/docs/References/PublicApi/api.Lib.md) entry point. It stores the deps and the database path on the struct, reads the task and visualization registries once, then runs the factories in `sandbox/lib/publicfunctions/` over it, each filling one function field with a closure that reads those deps at call time. This is the only wiring point: `sandbox` never imports an adapter, so the caller chooses which implementation to pass. The package is named `lib` and lives at `sandbox/`, so importers alias it: `wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"` — matching the `wraithadapter` / `wraithlib` / `wraithtypes` alias convention used by every consumer of this module.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `d` | [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | A dependency contract with every field filled, usually built by an adapter. |
| `databasePath` | `string` | The folder inside the vault the registries are persisted in. Required — a library that did not know where its data lived could not answer a question about it. Passing `""` takes the default, `data`. |

## Returns

| Type | Description |
| :--- | :--- |
| [`api.Lib`](/docs/References/PublicApi/api.Lib.md) | A ready-to-use library instance carrying the injected deps. |

## Examples

### Basic Initialization

```go
package main

import (
	"log"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	// 1. Build the deps through an adapter
	deps := wraithadapter.New("my-brain")

	// 2. Inject them into the library
	l := wraithlib.New(deps, "data")

	// The library instance 'l' is now ready for use.
	log.Println("tasks this brain carries:", len(l.Tasks))
}
```
