# Install the CLI

## Description
Covers installing the `wraith` command-line interface globally, so it runs from any directory and survives a terminal restart. To create your first vault once installed, follow [StartABrain.md](/docs/Tutorials/StartABrain.md); to consume the same behavior as a Go library instead, follow [LibInitialization.md](/docs/Tutorials/LibInitialization.md).

### Rules
- You do **not** need Go installed to run `wraith`. You can download the pre-compiled binary for your operating system directly.
- The installation commands use `curl` to fetch the latest release from the repository.

---

## Workflow

### macOS

1. Copy and paste the command below corresponding to your Mac's processor into your terminal to download and install the binary:

   **Apple Silicon (M1/M2/M3 / arm64)**
   ```bash
   curl -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/macarm64.bin -o wraith && \
     chmod +x wraith && \
     sudo mv wraith /usr/local/bin/ && \
     wraith version
   ```

   **Intel (x86_64)**
   ```bash
   curl -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/mac86.bin -o wraith && \
     chmod +x wraith && \
     sudo mv wraith /usr/local/bin/ && \
     wraith version
   ```

2. Open a **new terminal** to make sure the binary is available everywhere.

### Linux

1. Copy and paste the command below corresponding to your architecture into your terminal to download and install the binary:

   **x86_64 (64-bit)**
   ```bash
   curl -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/linux86.out -o wraith && \
     chmod +x wraith && \
     sudo mv wraith /usr/local/bin/ && \
     wraith version
   ```

   **ARM64**
   ```bash
   curl -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/linuxarm64.out -o wraith && \
     chmod +x wraith && \
     sudo mv wraith /usr/local/bin/ && \
     wraith version
   ```

   **x86 (32-bit)**
   ```bash
   curl -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/linuxi32.out -o wraith && \
     chmod +x wraith && \
     sudo mv wraith /usr/local/bin/ && \
     wraith version
   ```

2. Open a **new terminal** to make sure the binary is available everywhere.

### Windows (PowerShell)

1. Open PowerShell and copy and paste the command below corresponding to your architecture. It will download the binary, install it to your user profile, and update your `PATH`:

   **x86_64 (64-bit)**
   ```powershell
   $dir = "$HOME\.local\bin"; `
     New-Item -ItemType Directory -Force -Path $dir | Out-Null; `
     curl.exe -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/windows86.exe -o "$dir\wraith.exe"; `
     if ($env:PATH -notlike "*$dir*") { `
       [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH', 'User') + ";$dir", 'User'); `
       $env:PATH += ";$dir"; `
       Write-Host "Added $dir to your PATH (restart the terminal for full effect)"; `
     }; `
     wraith.exe version
   ```

   **x86 (32-bit)**
   ```powershell
   $dir = "$HOME\.local\bin"; `
     New-Item -ItemType Directory -Force -Path $dir | Out-Null; `
     curl.exe -sL https://github.com/MateusMoutinhoOrg/Wraith/releases/latest/download/windowsi32.exe -o "$dir\wraith.exe"; `
     if ($env:PATH -notlike "*$dir*") { `
       [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH', 'User') + ";$dir", 'User'); `
       $env:PATH += ";$dir"; `
       Write-Host "Added $dir to your PATH (restart the terminal for full effect)"; `
     }; `
     wraith.exe version
   ```

2. Open a **new terminal** to make sure the binary is available everywhere.

### Verify after reboot

3. After restarting the machine (or opening a fresh terminal), confirm the binary is still found globally:
   ```bash
   # On macOS / Linux
   wraith version
   ```
   ```powershell
   # On Windows
   wraith.exe version
   ```

### Troubleshooting

4. If `curl` fails with a network error, check your internet connection and ensure GitHub is accessible.
5. If `wraith version` says `command not found` after installing, the directory you moved the binary to (e.g., `/usr/local/bin` ou `$HOME\.local\bin`) is not in your `PATH`. Add it to your shell profile manually.

### Install from a Clone (Requires Go)

1. Use this instead of the steps above when you are working on the project itself and want the binary built from your checkout:
   ```bash
   go build -o $(go env GOPATH)/bin/wraith ./cmd/main
   ```
2. Or skip installing entirely and run it straight from source:
   ```bash
   go run ./cmd/main tasks
   ```
