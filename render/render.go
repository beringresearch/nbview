package render

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	chroma "github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
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
	Outputs        []string `json:"outputs"`
}

// Render accepts a Notebook struct and returns a rendered string
func Render(notebook Notebook) string {

	var output string

	for _, cell := range notebook.Cells {
		var source string
		var outputs string

		lexer := lexers.Get("python")
		lexer = chroma.Coalesce(lexer)
		style := styles.Get("dracula")
		formatter := formatters.Get("terminal256")

		for _, s := range cell.Source {
			source += s
		}

		if cell.CellType == "code" {
			iterator, err := lexer.Tokenise(nil, source)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			buf := new(bytes.Buffer)
			err = formatter.Format(buf, style, iterator)
			source = buf.String()

			source = "\x1B[38;2;249;38;114m[" + strconv.Itoa(cell.ExecutionCount) + "]\x1B[0m " + source
		}

		source = indent.String(source, 4)

		for _, o := range cell.Outputs {
			outputs += o
		}

		output += source + "\n"
		output += "\t" + outputs + "\n"
	}

	output = wordwrap.String(output, 100)
	return output

}
