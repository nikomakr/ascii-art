# Audit Checklist

This document mirrors the official audit questions and provides the exact commands
to run to prove each requirement is met.

---

## Functional

### Allowed packages

**Requirement:** Only standard Go packages are used.

**How to prove it:**

```bash
cat go.mod
```

The `go.mod` file contains no `require` block, meaning zero external dependencies.
Every import in the codebase (`fmt`, `os`, `strings`) is from the Go standard library.

```bash
grep -r "import" main.go art/art.go
```

---

### Try passing "hello" as an argument

```bash
go run . "hello" | cat -e
```

Expected:
```
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```

---

### Try passing "HELLO" as an argument

```bash
go run . "HELLO" | cat -e
```

Expected:
```
 _    _   ______   _        _         ____   $
| |  | | |  ____| | |      | |       / __ \  $
| |__| | | |__    | |      | |      | |  | | $
|  __  | |  __|   | |      | |      | |  | | $
| |  | | | |____  | |____  | |____  | |__| | $
|_|  |_| |______| |______| |______|  \____/  $
                                             $
                                             $
```

---

### Try passing "HeLlo HuMaN" as an argument

```bash
go run . "HeLlo HuMaN" | cat -e
```

Expected:
```
 _    _          _        _                 _    _           __  __           _   _  $
| |  | |        | |      | |               | |  | |         |  \/  |         | \ | | $
| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \  / |   __ _  |  \| | $
|  __  |  / _ \ | |      | |  / _ \        |  __  | | | | | | |\/| |  / _` | | . ` | $
| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\  | $
|_|  |_|  \___| |______| |_|  \___/        |_|  |_|  \__,_| |_|  |_|  \__,_| |_| \_| $
                                                                                     $
                                                                                     $
```

---

### Try passing "1Hello 2There" as an argument

```bash
go run . "1Hello 2There" | cat -e
```

Expected:
```
     _    _          _   _                         _______   _                           $
 _  | |  | |        | | | |                ____   |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                         $
                                                                                         $
```

---

### Try passing "Hello\nThere" as an argument

```bash
go run . "Hello\nThere" | cat -e
```

Expected:
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

---

### Try passing "Hello\n\nThere" as an argument

```bash
go run . "Hello\n\nThere" | cat -e
```

Expected:
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

---

### Try passing "{Hello & There #}" as an argument

```bash
go run . "{Hello & There #}" | cat -e
```

---

### Try passing 'hello There 1 to 2!' as an argument

```bash
go run . 'hello There 1 to 2!' | cat -e
```

---

### Try passing "MaD3IrA&LiSboN" as an argument

```bash
go run . "MaD3IrA&LiSboN" | cat -e
```

---

### Try passing "1a\"#FdwHywR&/()=" as an argument

```bash
go run . '1a"#FdwHywR&/()=' | cat -e
```

> Note: the outer single quotes let the shell pass the inner `"` as a literal character.

---

### Try passing "{|}~" as an argument

```bash
go run . "{|}~" | cat -e
```

---

### Try passing "[\]^_ 'a" as an argument

```bash
go run . "[\]^_ 'a" | cat -e
```

---

### Try passing "RGB" as an argument

```bash
go run . "RGB" | cat -e
```

---

### Try passing ":;<=>?@" as an argument

```bash
go run . ":;<=>?@" | cat -e
```

---

### Try passing '\!" #$%&'"'"'()*+,-./' as an argument

```bash
go run . '! "#$%&'"'"'()*+,-./' | cat -e
```

---

### Try passing "ABCDEFGHIJKLMNOPQRSTUVWXYZ" as an argument

```bash
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
```

---

### Try passing "abcdefghijklmnopqrstuvwxyz" as an argument

```bash
go run . "abcdefghijklmnopqrstuvwxyz" | cat -e
```

---

### Random string tests

These are chosen by the auditor at the time of review. Use any combination matching
the described criteria and verify the output is correctly rendered:

```bash
# At least 4 lowercase letters and 3 uppercase letters
go run . "helloWORLD" | cat -e

# At least 5 lowercase letters, a space and two numbers
go run . "hello world 42" | cat -e

# At least 1 uppercase letter and 3 special characters
go run . "Hello#@!" | cat -e
or (for zsh:)
go run . 'Hello#@!' | cat -e 

# At least 2 lowercase, 2 spaces, 1 number, 2 special characters, 3 uppercase
go run . "Hi There 1 AB#!" | cat -e
or (for zsh:)
go run . 'Hi There 1 AB#!' | cat -e
```

---

## Basic

### Does the project run quickly and effectively?

**How to prove it:**

"The banner file is read once per invocation"

Every time you run go run . "hello", the programme starts, reads the banner file exactly one time using os.ReadFile(), then uses that data for the entire render. It never re-reads the file mid-execution. Compare that to a bad implementation that opens the file once per character — for "hello" that would be 5 file reads instead of 1.

"Character lookup is O(1) direct array index access"

O(1) means constant time — it takes the same amount of time to find any character regardless of how many characters exist. This is because of the formula:

```
goidx := int('h') - 32  // → 72
b[72][row]            // → jump directly to slot 72
```

Go jumps straight to that memory slot. It does not loop through the array searching for 'h'. Compare that to a bad implementation using a loop — to find 'z' it would check 'a', then 'b', then 'c'... all the way to 'z'. That would be O(n).

"No loops over the banner data at render time"

The only loop that happens over banner data is during the initial parse — loading the file into the array once at startup. After that, rendering never touches the raw file data again. It only reads from the already-parsed array using direct index access.

"Output is built with a single strings.Builder and printed once"

strings.Builder accumulates all the output in memory efficiently, then fmt.Print(result) writes everything to the terminal in one operation. A bad implementation might call fmt.Print inside the inner loop — once per character row — causing many small writes to the terminal instead of one large efficient one.

The time command just measures this in practice:
```
bashtime go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" > /dev/null
```
You should see a very small real time, confirming the programme finishes near-instantly even for long inputs.

---

### Does the code obey good practices?

**How to prove it:**

```bash
# Verify code formatting — should produce no output if correctly formatted
gofmt -l .

# Run the Go vet static analyser — should produce no warnings
go vet ./...
```

Key practices followed in the codebase:
- `main.go` handles only CLI concerns — no business logic
- `art/art.go` contains all rendering logic, fully separated
- All exported functions have doc comments
- Errors are returned, not swallowed
- Error messages are written to `stderr`, output to `stdout`
- Constants are named and documented rather than using magic numbers

---

### Is there a test file for this code?

**How to prove it:**

```bash
ls art/art_test.go
```

---

### Are the tests checking each possible case?

**How to prove it:**

```bash
go test ./... -v
```

The test file covers:
- Empty string input
- Single `\n` input
- Trailing `\n`
- Double `\n` between words
- Lowercase letters
- Uppercase letters
- Mixed case
- Numbers
- Spaces between words
- Special characters from every ASCII range
- All three banners
- Invalid banner name error handling
- Invalid character error handling

---

### Is the output well-structured? Are the characters displayed correctly in line?

**How to prove it:**

```bash
go run . "Hello There" | cat -e
```

Each line ends with `$` (shown by `cat -e`), confirming no missing or extra
newlines. Characters are horizontally aligned because each banner row is
concatenated in order — row 0 of every character forms line 1, row 1 forms
line 2, and so on, across all 8 rows.

---

## Social

### Did you learn anything from this project?

Topics encountered and understood through this project:

- Go module system (`go mod init`, import paths)
- Fixed-size arrays vs maps as lookup structures and their trade-offs
- ASCII character encoding and the printable range (32–126)
- File I/O with `os.ReadFile`
- String manipulation with `strings.Split`, `strings.Builder`, `strings.ReplaceAll`
- Separating CLI concerns from business logic
- Writing table-driven unit tests in Go
- The difference between a real newline and a literal `\n` escape sequence

---

### Can it be open-sourced / be used for other sources?

Yes. The project uses no external dependencies and no proprietary assets.
The rendering logic in `art/art.go` is fully self-contained and could be imported as a package by any other Go project using the `Render()` function directly.

---

### Would you recommend/nominate this programme as an example for the rest of the school?

The codebase demonstrates a clear separation of concerns, meaningful naming,
documented functions, comprehensive tests, and correct handling of all edge
cases identified in the audit. It follows Go idioms throughout and contains
no unnecessary complexity.