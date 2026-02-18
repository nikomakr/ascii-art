# ASCII Art

A command line programme written in Go that receives a string as an argument and outputs it as a graphical representation using ASCII characters.

## Description

ASCII Art renders any input string using pre-defined banner styles. Each character is represented as an 8-line tall ASCII drawing, with characters placed side by side to form words and sentences. The programme handles letters, numbers, spaces, special characters and newline sequences.

## Project Structure

```
ascii-art/
├── go.mod
├── main.go
├── art/
│   ├── art.go
│   └── art_test.go
└── banners/
```

- **main.go** — Entry point. Handles CLI argument parsing and delegates rendering to the art package
- **art/art.go** — Core rendering engine. Loads banner files, parses character maps and builds the ASCII output
- **art/art_test.go** — Unit tests covering all character categories and edge cases
- **banners/** — Plain text files containing the ASCII template for each banner style

## Usage

```bash
go run . "<string>"
go run . "<string>" <banner>
```

The banner argument is optional and defaults to `standard`.

### Examples

```bash
go run . "Hello"
go run . "Hello There"
go run . "Hello\nThere"
go run . "hello" shadow
go run . "hello" thinkertoy
```

### Newline Support

Use the literal `\n` sequence within your string to split output across multiple lines:

```bash
go run . "Hello\nThere"     # two separate rendered blocks
go run . "Hello\n\nThere"   # two blocks with a blank line between
```

## Banner Styles


## Supported Characters

All 95 printable ASCII characters are supported, covering:

- Uppercase letters A to Z
- Lowercase letters a to z
- Digits 0 to 9
- Space
- Special characters including `!`, `"`, `#`, `$`, `%`, `&`, `'`, `(`, `)`, `*`, `+`, `,`, `-`, `.`, `/`, `:`, `;`, `<`, `=`, `>`, `?`, `@`, `[`, `\`, `]`, `^`, `_`, `` ` ``, `{`, `|`, `}`, `~`

## Running Tests

```bash
go test ./...
```

## Requirements

- Go 1.22 or higher
- No external dependencies — standard library only

## Allowed Packages

Only Go standard library packages are used in this project, in compliance with the project specification.