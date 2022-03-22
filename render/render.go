package render

import (
	"strconv"

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

		var textColour string

		if cell.CellType == "code" {
			source += "\x1B[38;2;249;38;114m[" + strconv.Itoa(cell.ExecutionCount) + "]\x1B[0m "
			textColour = "\033[97m"
		} else {
			textColour = "\033[36m"
		}
		for _, s := range cell.Source {
			source += textColour + s + "\033[0m"
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
