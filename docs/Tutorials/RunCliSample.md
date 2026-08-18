# Run CLI Examples

## Description
How to run the CLI examples provided in the repository to understand the tool's behavior from the terminal.

---

## Run CLI Examples

CLI examples are shell scripts that demonstrate how to use the `wraith` command-line tool. They build the binary dynamically and run in an isolated scratch space, in a temporary vault of their own, so they never touch a brain of yours.

### Workflow

1. Browse the `/examples/cliExamples/` directory for a script that matches the workflow you want to learn (e.g., `StartAVault.sh`).
2. Run the script from the project root:
   ```bash
   bash ./examples/cliExamples/StartAVault.sh
   ```
3. Read the output. The script will output comments indicating what it's doing before executing the CLI commands.
4. To run all CLI examples sequentially:
   ```bash
   for script in ./examples/cliExamples/*.sh; do bash "$script"; done
   ```
