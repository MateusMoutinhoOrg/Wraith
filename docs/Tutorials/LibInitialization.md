# Library Initialization

## Description
Covers installing the library and initializing it with the standard adapter in a new program. Running actions from code afterwards is [RunApiSample.md](/docs/Tutorials/RunApiSample.md), and driving the same behavior from a terminal is [RunTasks.md](/docs/Tutorials/RunTasks.md). For other ways to build the dependencies, see [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md).

---

## Workflow

1. Install the lib:

```bash
go get github.com/MateusMoutinhoOrg/Wraith@latest
```

2. Create a file called `main.go`:

```go
package main

// 1. Import the standard adapter and the lib
import (
	"fmt"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
)

func main() {
	// 2. Create deps via an adapter (the "opinionated" part), rooted at the
	//    folder the vault lives in.
	deps := wraithadapter.New("my-brain")

	// 3. Inject deps into the pure library, naming the folder inside the
	//    vault the registries are persisted in.
	l := wraithlib.New(deps, "data")

	// 4. Run a task. The entries map is the same one a Task.yaml decodes to.
	if err := l.PerformTask("AddAccount", map[string]any{
		"account": "Bank",
	}); err != nil {
		panic(err)
	}

	// 5. Render a visualization. It hands back the files it produced and
	//    writes nothing itself.
	renders, err := l.PerformVisualization("DashBoard", map[string]any{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(renders[0].Content))
}
```

3. Run the code:

```bash
go run main.go
```

The second argument to `wraithlib.New` is required: a library that did not know where its data lived could not answer a question about it. Passing `""` takes the default, `data`.
