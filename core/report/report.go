package report

import (
	"os"
	"text/template"

	"github.com/gbowne1/Diskscape/diskstats"
	"github.com/gbowne1/Diskscape/filesystem"
)

const reportTemplate = `
# Disk Usage Report

## Total Space: {{.TotalSpace}} KB
## Used Space: {{.UsedSpace}} KB
## Free Space: {{.FreeSpace}} KB
## Used Percentage: {{.UsedPercentage}}%

## Top 10 Largest Files/Directories:
{{range .TopEntries}}
{{.Path}} - Size: {{.Size}} KB
{{end}}

## Detailed Breakdown:
{{range .Breakdown}}
{{.Name}}: {{.Size}} KB ({{.Percentage}}%)
{{end}}
`

type Report struct {
	TotalSpace     int64
	UsedSpace      int64
	FreeSpace      int64
	UsedPercentage float64
	TopEntries     []filesystem.FilesystemEntry
	Breakdown      map[string]struct {
		Size       int64
		Percentage float64
	}
}

func GenerateReport(stats diskstats.DiskStats, entries []filesystem.FilesystemEntry, breakdown map[string]struct {
	Size       int64
	Percentage float64
}) *Report {
	return &Report{
		TotalSpace:     int64(stats.TotalSpace),
		UsedSpace:      int64(stats.TotalSpace - stats.FreeSpace),
		FreeSpace:      int64(stats.FreeSpace),
		UsedPercentage: stats.UsedPercentage,
		TopEntries:     entries[:10],
		Breakdown:      breakdown,
	}
}

func PrintReport(report *Report) {
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		panic(err)
	}

	data := struct {
		TotalSpace     int64
		UsedSpace      int64
		FreeSpace      int64
		UsedPercentage float64
		TopEntries     []filesystem.FilesystemEntry
		Breakdown      map[string]struct {
			Size       int64
			Percentage float64
		}
	}{
		TotalSpace:     report.TotalSpace,
		UsedSpace:      report.UsedSpace,
		FreeSpace:      report.FreeSpace,
		UsedPercentage: report.UsedPercentage,
		TopEntries:     report.TopEntries,
		Breakdown:      report.Breakdown,
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
