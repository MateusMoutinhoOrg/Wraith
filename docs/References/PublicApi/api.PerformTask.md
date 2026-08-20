# `api.Lib.PerformTask`

**Type:** Field

## Signature

```go
PerformTask func(taskName string, entries map[string]any) error
```

## Description

Runs one [task](/docs/References/PublicApi/api.Task.md) by name against the library's database. It is the whole of "do something": a tick calls it with what it read out of `Task.yaml`, and `wraith run` calls it with what it read off the command line.

The `entries` map is the same one a `Task.yaml` decodes to, so anything the file can say this call can say too. It is validated against the fields the task declares before the task is reached: a missing required field, an unknown field, or a field of the wrong type is reported and nothing is written.

It writes no file and renders nothing. Re-rendering is a separate step — a caller that wants both asks for both, or calls [`PerformFullTick`](/docs/References/PublicApi/api.PerformFullTick.md).

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `taskName` | `string` | The task to run. Must be one of `Lib.Tasks`. |
| `entries` | `map[string]any` | The task's fields, keyed by the name they carry in `Task.yaml`. |

## Returns

| Type | Description |
| :--- | :--- |
| `error` | An unknown task, an invalid field, or a failure the task reported. `nil` on success. |

## Examples

```go
package main

import (
	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	l := wraithlib.New(wraithadapter.New("my-brain"), "data")

	if err := l.PerformTask("AddAccount", map[string]any{
		"account": "Bank",
	}); err != nil {
		panic(err)
	}
}
```
