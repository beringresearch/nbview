package render

import (
	"strconv"
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
			source += "[" + strconv.Itoa(cell.ExecutionCount) + "] "
		}
		for _, s := range cell.Source {
			source += s
		}

		for _, o := range cell.Outputs {
			outputs += o
		}

		output += source + "\n"
		output += "\t" + outputs + "\n"
	}

	return output

}
