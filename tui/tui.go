package tui

import (
	"fmt"
	"log"

	"github.com/gbowne1/Diskscape/core/diskstats"
	"github.com/rivo/tview"
)

func Run() error {
	app := tview.NewApplication()

	// Get disk statistics for the root directory (or target directory from config)
	stats, err := diskstats.GetDiskStats("/")
	if err != nil {
		log.Printf("Error retrieving disk stats: %v\n", err)

		// Display error message in TUI
		table := tview.NewTable().
			SetBorders(true).
			SetTitle("Error").
			SetTitleAlign(tview.AlignCenter)
		table.SetCell(0, 0, tview.NewTableCell(fmt.Sprintf("Failed to retrieve disk stats: %v", err)))

		if err := app.SetRoot(table, true).Run(); err != nil {
			return err
		}
		return nil
	}

	// Display disk statistics in a table
	table := tview.NewTable().
		SetBorders(true).
		SetTitle("Disk Usage Statistics").
		SetTitleAlign(tview.AlignCenter)

	table.SetCell(0, 0, tview.NewTableCell("Total Space").SetTextAlign(tview.AlignLeft))
	table.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%d KB", stats.TotalSpace/1024)))

	table.SetCell(1, 0, tview.NewTableCell("Free Space").SetTextAlign(tview.AlignLeft))
	table.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%d KB", stats.FreeSpace/1024)))

	table.SetCell(2, 0, tview.NewTableCell("Used Percentage").SetTextAlign(tview.AlignLeft))
	table.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%.2f%%", stats.UsedPercentage)))

	// Set the table as the root and run the TUI application
	if err := app.SetRoot(table, true).Run(); err != nil {
		return err
	}
	return nil
}
