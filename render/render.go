package render

import (
	"strconv"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
)

type Notebook struct {
	Cells []Cell `json:"cells"`
}

type Cell struct {
	CellType       string   `json:"cell_type"`
	ExecutionCount int      `json:"execution_count"`
	Source         []string `json:"source"`
	Outputs        []string `json:"outputs"`
}

func Render(notebook Notebook) string {

	var output string

	for _, cell := range notebook.Cells {
		var source string
		var outputs string

		if cell.CellType == "code" {
			source += "\x1B[38;2;249;38;114m[" + strconv.Itoa(cell.ExecutionCount) + "]\x1B[0m "
		}
		for _, s := range cell.Source {
			source += s
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
