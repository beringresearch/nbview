package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/beringresearch/nbview/render"
)

const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type model struct {
	name     string
	content  string
	ready    bool
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render(m.name)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func loadNotebookFromUrl(url string) (render.Notebook, error) {
	var notebook render.Notebook
	var byteValue []byte

	response, err := http.Get(url)
	if err != nil {
		return notebook, err
	}

	if response.StatusCode == http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		byteValue = buf.Bytes()
	}

	if len(byteValue) == 0 {
		err = errors.New("unable to read from URL")
	} else {
		json.Unmarshal(byteValue, &notebook)
	}

	return notebook, err
}

func loadNotebookFromFile(fileName string) (render.Notebook, error) {
	var notebook render.Notebook
	jsonFile, err := os.Open(fileName)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &notebook)

	return notebook, err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("view Jupyter Notebooks in your terminal\n\n Usage:\n  nbview FILENAME")
		os.Exit(0)
	}
	var notebook render.Notebook
	fileName := os.Args[1]

	if _, err := os.Stat(fileName); err == nil {
		notebook, err = loadNotebookFromFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		notebook, err = loadNotebookFromUrl(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	content := render.Render(notebook)

	p := tea.NewProgram(
		model{content: string(content), name: fileName},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
