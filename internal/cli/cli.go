package cli

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// Flags defines each flag
type Flags struct {
	// Name of flag
	Name string
	// Help message of flag
	Message string
}

// Structure defines the CLI structure of tool
type Structure struct {
	// Name of Executable
	CliName string
	// Additional arguments associated with CLI in formatted manner
	Args []string
	// Heading of flag. It will be printed as it is
	FlagHeading string
	// Slice of Flags
	FlagData []Flags
	// Tab with to make output formatted
	TabWidth int
}

// PrintHelp prints out cli structure in a nice formatted way on the passed writer interface
func (c *Structure) PrintHelp(w io.Writer) {
	usage := fmt.Sprintf("Usage: %v  ", c.CliName)
	for _, v := range c.Args {
		usage += v + "  "
	}
	_, err := w.Write([]byte(usage))
	if err != nil {
		return
	}
	_, err = w.Write([]byte("\n" + c.FlagHeading + "\n"))
	if err != nil {
		return
	}
	for _, v := range c.FlagData {
		flagLen := utf8.RuneCountInString(v.Name)
		tTabLen := c.TabWidth - flagLen
		flagString := v.Name
		for i := 1; i <= tTabLen; i++ {
			flagString += " "
		}
		_, err2 := w.Write([]byte(fmt.Sprintf("%v%v\n", flagString, v.Message)))
		if err2 != nil {
			return
		}
	}
}
