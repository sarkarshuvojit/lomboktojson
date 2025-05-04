# LombokToJson

A lightweight utility that converts Lombok's default `toString()` output format to standard JSON format.

![Conversion Example](docs/assets/high-level-idea-clear.png)

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
go get github.com/yourusername/lombok-to-json
```

## Usage

```go
import "github.com/yourusername/lombok-to-json"

// Convert a Lombok toString string to JSON
lombokStr := "Customer(name=Raju,email=raju@gmail.com,age=15)"
jsonStr, err := lombokToJson.Convert(lombokStr)

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
