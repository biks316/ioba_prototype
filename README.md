# IOBA Prototype - Intent-Oriented Behavior Automation

A Go-based prototype for defining and executing high-level intents through YAML configuration. This project demonstrates a proof-of-concept system where complex behaviors are defined as declarative intents with reusable actions.

## Overview

IOBA (Intent-Oriented Behavior Automation) allows you to:
- Define high-level goals as **Intents**
- Specify **Actions** (logging, command execution) within each intent
- Execute intents by name through a simple CLI interface
- Manage multi-step workflows in a declarative YAML format

## Project Structure

```
.
├── main.go          # Core application logic
├── intent.yml       # Intent definitions and configuration
├── go.mod           # Go module definition
└── README.md        # This file
```

## Getting Started

### Prerequisites
- Go 1.24.3 or later
- gopkg.in/yaml.v3

### Running the Prototype

```bash
go run main.go
```

This will:
1. Load all intents from `intent.yml`
2. Execute each intent sequentially
3. Display logs for each action (logging and command execution)

## Configuration (intent.yml)

Define your intents in `intent.yml`. Each intent contains:

```yaml
intents:
  - name: IntentName
    description: "What this intent does"
    actions:
      - type: log
        message: "Log message"
      - type: execute
        command: "bash command to run"
```

### Action Types

- **log**: Print a message to stdout
- **execute**: Run a bash command and output its result

## Example

The included `intent.yml` defines two intents:

- **CreateUser**: Logs creation and executes an echo command
- **DeleteUser**: Logs deletion and executes an echo command

Run the prototype to see these intents in action.

## Architecture

### Core Components

1. **Intent**: High-level goal with a name, description, and actions
2. **Action**: Individual operation (log or execute)
3. **IntentSpec**: Container holding all intents
4. **Executor**: Runs actions and intents in sequence

## Future Enhancements

- CLI argument support to execute specific intents
- Conditional action execution
- Error handling and rollback mechanisms
- Action dependencies and sequencing
- Output capture and transformation
