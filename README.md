# Go JSON Parser

A lightweight JSON parser written in Go that supports JSON formatting and selective querying using dot notation. This parser implements a custom lexer and parser to handle JSON data structures.

## Features

- Custom lexer and parser implementation
- JSON string formatting with customizable indentation
- Query JSON data using dot notation path selectors
- Support for all JSON data types:
  - Objects
  - Arrays
  - Strings
  - Numbers
  - Booleans
  - Null

## Installation

```bash
git clone https://github.com/yourusername/go-json.git
cd go-json
go build
```

## Usage

The parser supports two main operations: `format` and `select`.

### Formatting JSON

Format JSON with custom indentation:

```bash
# Format JSON from a file
go-json format -f input.json -t "    "

# Format JSON from stdin
echo '{"name": "John", "age": 30}' | go-json format
```

### Querying JSON

Select specific values using dot notation:

```bash
# Query from a file
go-json select -q "user.address.street" -f input.json

# Query from stdin
echo '{"user": {"address": {"street": "123 Main St"}}}' | go-json select -q "user.address.street"
```

#### Query Syntax

- Use dot notation to access object properties: `user.name`
- Use bracket notation with indices to access array elements: `users.[0].name`
- Combine both for complex queries: `users.[0].addresses.[1].street`

### Examples

Given this JSON:

```json
{
    "users": [
        {
            "name": "John",
            "age": 30,
            "addresses": [
                {
                    "type": "home",
                    "street": "123 Main St"
                },
                {
                    "type": "work",
                    "street": "456 Market St"
                }
            ]
        }
    ]
}
```

Query examples:

```bash
# Get the first user's name
go-json select -q "users.[0].name" -f user.json
# Output: "John"

# Get the first user's work address
go-json select -q "users.[0].addresses.[1]" -f user.json
# Output: {
#     "type": "work",
#     "street": "456 Market St"
# }
```

## Implementation Details

The parser is implemented in several components:

- `lexer`: Tokenizes JSON input into a stream of tokens
- `parser`: Constructs an AST (Abstract Syntax Tree) from tokens
- `ast`: Defines the node types for the syntax tree
- `operations`: Implements formatting and query operations
