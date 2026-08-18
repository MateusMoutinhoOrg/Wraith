# Wraith

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Wraith.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Wraith)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Wraith)](https://github.com/MateusMoutinhoOrg/Wraith/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](go.mod)
[![License](https://img.shields.io/badge/license-Unlicense-green)](LICENSE)

A **second brain** you drive with two files — a template for building one, shipped with a financial brain already in it.

---

## Overview

Wraith is a small state machine over a folder. You write an action into `Task.yaml`; it applies that action to your data and re-renders every dashboard you declared in `Visualization.yaml`. Nothing else happens, and nothing else needs to.

```yaml
# Task.yaml — one action, waiting to be armed
name: AddTransaction
account: Bank
category: Food
amount: -32.90
date: 2026-08-18
apply: true
```

```bash
wraith tick     # apply it, then redraw every page
```

Two ideas carry the whole project:

- **Tasks** are what can happen. One file each, under [`/sandbox/Tasks/Tasks/`](/sandbox/Tasks/Tasks/), declared in one array.
- **Visualizations** are what you get to see. One file each, under [`/sandbox/Visualization/Visualization/`](/sandbox/Visualization/Visualization/), declared in one array.

Both are meant to be replaced. The financial brain in this repository — accounts, categories, transactions, recurrences, credit cards — is a **worked example of the shape**, not the point of it. Fork the repo, swap the tasks for yours, and you have a brain for something else entirely. That path is [Brain-Config](/docs/Index/Brain-Config.md).

### Where the code lives

The core lives in **`/sandbox/`**: a **closed sandbox** that reaches nothing outside itself. Everything it can do arrives through an injected `Deps`.

```
adapters/  ──▶  sandbox/  ◀──  cmd/, examples/libraryExamples/
(reaches the OS)  (closed)     (wire the two together)
```

The command line is `api.Lib.Sandboxmain` — one field of the library like any other. The installed binary in **`/cmd/main/`** holds no command, no flag, and no output of its own: it wires an adapter into the library and calls that one field.

- **`/sandbox/`**: the closed library taking a `Deps` and returning an `api.Lib`.
- **`/sandbox/Tasks/`**: every action, one per file, plus the switcher that runs one by name.
- **`/sandbox/Visualization/`**: every renderer, one per file, plus the switcher that renders one by name.
- **`/adapters/`**: concrete implementations of the `Deps` contract.
- **`/cmd/`** & **`/examples/`**: where an adapter and the library are wired together.
- **`/assets/`**: files compiled into the binary and reached only through the injected `Deps` — including the default `Task.yaml` and `Visualization.yaml` that `wraith start` writes.

See [SandboxIsolation.md](/docs/References/SandboxIsolation.md) and [StructContracts.md](/docs/References/StructContracts.md) for the full mechanic.

---

## Doc Index

Documentation is split into four themes, one index page each under `docs/Index/`, listing that theme's **Tutorials** — step-by-step workflows — and its **References** — explanations and lookups. Start from the theme index matching what you want to do.

| Theme | Description |
| --- | --- |
| [Brain Usage](/docs/Index/Brain-Usage.md) | For people driving a brain: installing the binary, running tasks, choosing what gets rendered. |
| [Brain Config](/docs/Index/Brain-Config.md) | For people building their own brain: forking this one, adding tasks, adding visualizations. |
| [Library Usage](/docs/Index/LibUsage.md) | For Go callers: wiring an adapter, running tasks and rendering from code. |
| [Development](/docs/Index/Development.md) | For contributors: the rules, the mechanics, the per-goal workflows, and the specifications. |

New here? [Brain Usage → InstallCli.md](/docs/Tutorials/InstallCli.md) installs the binary; [StartABrain.md](/docs/Tutorials/StartABrain.md) turns an empty folder into a working vault.

---

## License

This project is licensed under the [Unlicense](./LICENSE).
