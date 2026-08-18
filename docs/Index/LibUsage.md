# Library Usage

## Description
Index of the documentation for developers consuming Wraith as a Go library: wiring an adapter into the sandbox, running tasks and rendering visualizations from code, and looking up the public API. Driving the same behavior from a terminal is indexed by [Brain-Usage.md](/docs/Index/Brain-Usage.md); building a brain of your own is indexed by [Brain-Config.md](/docs/Index/Brain-Config.md).

The library is always built the same way: an adapter produces a `deps.Deps`, `lib.New` injects it into the closed sandbox along with the folder the data lives in, and the returned `api.Lib` carries every behavior — the command line included, as one field like any other.

---

## Tutorials

- [LibInitialization.md](/docs/Tutorials/LibInitialization.md)
  - **description:** Install the lib, create deps via an adapter, and run a first program
- [RunApiSample.md](/docs/Tutorials/RunApiSample.md)
  - **description:** Run one of the shipped Go examples from the source tree
  - [Run API Examples](/docs/Tutorials/RunApiSample.md#run-api-examples)
- [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md)
  - **description:** How the library receives its dependencies, and how to map your own adapters
  - [Find Dependencies Functions you can use](/docs/Tutorials/HandleDependencies.md#find-dependencies-functions-you-can-use)
  - [Add New Dependencie](/docs/Tutorials/HandleDependencies.md#add-new-dependencie)
  - [Overwrinting a adapter function](/docs/Tutorials/HandleDependencies.md#overwrinting-a-adapter-function)
  - [Creating a adapter in repo](/docs/Tutorials/HandleDependencies.md#creating-a-adapter-in-repo)
  - [Creating a adapter in your project](/docs/Tutorials/HandleDependencies.md#creating-a-adapter-in-your-project)

---

## References

- [PublicApi.md](/docs/References/PublicApi.md)
  - **description:** Index of every public-facing entry of the library, grouped by its role
  - [Entry Points](/docs/References/PublicApi.md#entry-points)
  - [The State Machine](/docs/References/PublicApi.md#the-state-machine)
  - [Declarations](/docs/References/PublicApi.md#declarations)
  - [Dependency Contracts](/docs/References/PublicApi.md#dependency-contracts)
- [Adapters.md](/docs/References/Adapters.md)
  - **description:** Every shipped adapter you can inject, and when to use each one
  - [Available Adapters](/docs/References/Adapters.md#available-adapters)
  - [Embedded Libraries](/docs/References/Adapters.md#embedded-libraries)
- [ApiSamplesList.md](/docs/References/ApiSamplesList.md)
  - **description:** Every Go example shipped in `examples/libraryExamples/`
  - [Examples](/docs/References/ApiSamplesList.md#examples)
