# LombokToJson

A lightweight utility that converts Lombok's default `toString()` output format to standard JSON format.

![Conversion Example](docs/assets/high-level-idea-clear.png)

## Installation

### CLI Tool

Install the command-line interface tool:

```bash
go install github.com/sarkarshuvojit/lomboktojson@latest
```
## Usage

### As a Go Package

```go
import "github.com/sarkarshuvojit/lomboktojson"

// Convert a Lombok toString string to JSON
lombokStr := "Customer(name=Raju,email=raju@gmail.com,age=15)"
jsonStr, err := lombokToJson.Convert(lombokStr)

// Result: {"name":"Raju","email":"raju@gmail.com","age":15}
```

### CLI Usage

The CLI tool offers three ways to provide input:

#### 1. Interactive Mode

Simply run the tool without any arguments and paste or type your Lombok output. Press Ctrl+D when finished.

```bash
lomboktojson > $(dirseq)-customer-data.json
# Paste your Lombok output and press Ctrl+D
# Creates: 1-customer-data.json with the JSON result
```

#### 2. File Input

Use the `-i` flag to specify an input file:

```bash
lomboktojson -i lombok-output.log > $(dirseq)-parsed-output.json
# Creates: 1-parsed-output.json with the JSON result
```

#### 3. Pipe Input

Pipe content directly into the tool:

```bash
cat lombok-output.log | lomboktojson > $(dirseq)-customer-records.json
# Creates: 1-customer-records.json with the JSON result
```

#### Example Workflow

```bash
# Extract logs containing Customer objects
grep "Customer(" application.log > lombok-logs.txt

# Convert to JSON and save to sequentially named file
cat lombok-logs.txt | lomboktojson > $(dirseq)-converted-customers.json
# Creates: 1-converted-customers.json

# Process another batch
grep "Order(" application.log | lomboktojson > $(dirseq)-converted-orders.json
# Creates: 2-converted-orders.json
```

## Problem

When using Lombok's `@ToString` or `@Data` annotations in Java classes, log output is generated in Lombok's specific format:

```
Customer(name=Raju,email=raju@gmail.com,age=15)
```

This format isn't easily parseable by standard tools that expect JSON:

```json
{ "name": "Raju", "email": "raju@gmail.com", "age": 15 }
```

LombokToJson bridges this gap by providing a simple conversion utility.

## Features

- Convert Lombok's toString format to standard JSON
- Simple API for integration into existing projects
- No external dependencies
- High performance parsing
- Supports nested objects and collections

## Installation

```bash
go get github.com/sarkarshuvojit/lombok-to-json
```

## Usage

```go
import "github.com/sarkarshuvojit/lombok-to-json"

// Convert a Lombok toString string to JSON
lombokStr := "Customer(name=Raju,email=raju@gmail.com,age=15)"
jsonStr, err := lombokToJson.GetJson(lombokStr)

// Result: {"name":"Raju","email":"raju@gmail.com","age":15}
```

## Testing

### Using Go's built-in test tool

```bash
# Run tests
go test

# Run tests with verbose output
go test -v ./...
```

### Development with live reload

Prerequisites: Install [air](https://github.com/air-verse/air)

```bash
# Run tests automatically on file changes
air
```

### Using Makefile

```bash
# Run tests
make test

# Run tests with verbose output
make testv
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
