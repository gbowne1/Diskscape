package tui

import (
	"fmt"

	"github.com/gbowne1/Diskscape/diskstats"
	"github.com/rivo/tview"
)

func Run() error {
	app := tview.NewApplication()

	stats, err := diskstats.GetDiskStats("/")
	if err != nil {
		return err
	}

	table := tview.NewTable().
		SetBorders(true).
		SetTitle("Disk Usage Statistics").
		SetTitleAlign(tview.AlignLeft)

	table.SetCell(0, 0, tview.NewTableCell("Total Space").SetTextAlign(tview.AlignLeft))
	table.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%d KB", stats.TotalSpace)))

	table.SetCell(1, 0, tview.NewTableCell("Free Space").SetTextAlign(tview.AlignLeft))
	table.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%d KB", stats.FreeSpace)))

	table.SetCell(2, 0, tview.NewTableCell("Used Percentage").SetTextAlign(tview.AlignLeft))
	table.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%.2f%%", stats.UsedPercentage)))

	if err := app.SetRoot(table, true).Run(); err != nil {
		return err
	}
	return nil
}
