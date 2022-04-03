package render

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	chroma "github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/crypto/ssh/terminal"
)

// Notebook contains Jupyter Notebook definitions
type Notebook struct {
	Cells []Cell `json:"cells"`
}

// Cell is the basic building block of a Notebook structure
type Cell struct {
	CellType       string   `json:"cell_type"`
	ExecutionCount int      `json:"execution_count"`
	Source         []string `json:"source"`
	Outputs        []Output `json:"outputs"`
}

type Output struct {
	Name       string   `json:"name"`
	OutputType string   `json:"output_type"`
	Text       []string `json:"text"`
}

// Render accepts a Notebook struct and returns a rendered string
func Render(notebook Notebook) string {
	terminalWidth, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var output string
	line := strings.Repeat("─", terminalWidth-10)

	for _, cell := range notebook.Cells {
		var source string
		var outputs string

		cellNumber := "\x1B[38;2;249;38;114m[" + strconv.Itoa(cell.ExecutionCount) + "]\x1B[0m "
		lexer := lexers.Get("python")
		lexer = chroma.Coalesce(lexer)
		style := styles.Get("dracula")
		formatter := formatters.Get("terminal256")

		for i, s := range cell.Source {

			if cell.CellType == "markdown" {
				if strings.HasPrefix(s, "#") {
					s = "\n\u001b[34m" + s + "\n" + line + "\x1B[0m "
				}

			}

			if cell.CellType == "code" {

				iterator, err := lexer.Tokenise(nil, s)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				buf := new(bytes.Buffer)
				err = formatter.Format(buf, style, iterator)
				s = buf.String()
			}

			if i > -1 {
				s += "\t"
			}

			source += s
		}

		if cell.CellType == "code" {
			source = cellNumber + source
		}

		for _, o := range cell.Outputs {
			for _, t := range o.Text {
				outputs += t
			}
		}

		output += source + "\n"
		output += "\t" + outputs + "\n"
	}

	output = indent.String(wordwrap.String(output, terminalWidth-50), 4)
	return output

}
