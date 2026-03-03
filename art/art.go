package art

import (
	"fmt"
	"os"
	"strings"
)

const (
	// Each character in a banner file occupies 8 lines of art
	// plus 1 blank separator line = 9 lines per character block.
	charHeight    = 8
	linesPerChar  = 9 // charHeight + 1 blank separator
	firstASCII    = 32 // ASCII code of space ' ', the first character in banner files
)

// banner holds the parsed ASCII art for all 95 printable characters.
// Index 0 = space (ASCII 32), index 94 = '~' (ASCII 126).
type banner [95][charHeight]string // just called "banner" for simplicity, not a map or slice

// Render takes the raw input string and a banner name, and returns the full ASCII art output as a single string ready to be printed.
// The input string may contain literal "\n" sequences (backslash + n), which are treated as line separators — not real newlines.

// An empty input string returns an empty string (no output at all).
func Render(input, bannerName string) (string, error) {
	// Empty string produces zero output
	if input == "" {
		return "", nil
	}

	b, err := loadBanner(bannerName) // Load the specified banner file and parse it into a banner struct. If the banner name is invalid or the file cannot be read, return an error.
	if err != nil {
		return "", err
	}

	// Split the input on literal "\n" sequences. This allows us to handle multi-line input where lines are separated by "\n" in the input string. Each segment between "\n" will be rendered as a separate line of ASCII art.
	lines := strings.Split(input, `\n`)

	var sb strings.Builder // Use a strings.Builder to efficiently build the final output string. This allows us to concatenate the rendered lines of ASCII art without creating intermediate strings, which can be costly in terms of memory and performance.

	for i, line := range lines {
		if line == "" {
			// An empty segment from a "\n" split produces a single blank line,
			// EXCEPT for the very last empty segment when input ends with "\n".
			// Example: "\n" splits into ["", ""] — we print one blank line total.
			// Example: "Hello\n" splits into ["Hello", ""] — we print Hello's 8 lines then one blank line.
			// The last empty string in the slice is always printed as a blank line to match the expected audit output.
			if i < len(lines)-1 || len(lines) == 2 && lines[0] == "" {
				sb.WriteByte('\n')
			} else if i == len(lines)-1 && lines[0] != "" {
				// Trailing \n after content: write the blank line
				sb.WriteByte('\n')
			}
			continue
		}

		// Validate all characters in the line are printable ASCII (32–126).
		for _, ch := range line {
			if ch < firstASCII || ch > 126 {
				return "", fmt.Errorf("invalid character %q: only printable ASCII (32–126) supported", ch)
			}
		}

		// Render all 8 rows of this line by concatenating each character's row.
		for row := 0; row < charHeight; row++ {
			for _, ch := range line {
				idx := int(ch) - firstASCII
				sb.WriteString(b[idx][row])
			}
			sb.WriteByte('\n')
		}
	}

	return sb.String(), nil
}

// loadBanner reads and parses the named banner file into a banner struct.
// The banner file must be located in the "banners/" directory relative to where the program is run.
func loadBanner(name string) (banner, error) {
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !validBanners[name] {
		return banner{}, fmt.Errorf("unknown banner %q: choose standard, shadow, or thinkertoy", name)
	}

	path := "banners/" + name + ".txt"
	data, err := os.ReadFile(path)
	if err != nil {
		return banner{}, fmt.Errorf("could not read banner file %q: %w", path, err)
	}

	return parseBanner(data)
}

// parseBanner parses the raw bytes of a banner file into a banner struct.
// Banner files use \r\n or \n line endings — we normalise to \n.
// The file starts with a blank line, then each character block is:
//   8 lines of art followed by 1 blank line.
func parseBanner(data []byte) (banner, error) {
	// Normalise Windows line endings.
	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	// Split into individual lines.
	lines := strings.Split(content, "\n")

	// The banner file begins with a blank line (line 0), so the first character (space, ASCII 32) starts at line 1.
	// Character at index i starts at line: 1 + i*linesPerChar
	var b banner

	for i := 0; i < 95; i++ {
		startLine := 1 + i*linesPerChar

		// Guard against malformed/short banner files.
		if startLine+charHeight > len(lines) {
			return banner{}, fmt.Errorf(
				"banner file too short: expected at least %d lines, got %d",
				1+95*linesPerChar, len(lines),
			)
		}

		for row := 0; row < charHeight; row++ {
			b[i][row] = lines[startLine+row]
		}
	}

	return b, nil
}