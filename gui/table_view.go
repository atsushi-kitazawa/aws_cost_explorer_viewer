package gui

import (
	"os"
	"strconv"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/model"
	"github.com/olekukonko/tablewriter"
)

func DisplayMetricSummaryTableView(ms []model.MetricSummary) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Users", "BlendedCost", "UnblendedCost", "UsageQuantity"})

	for _, m := range ms {
		table.Append([]string{
			m.User,
			strconv.FormatFloat(m.BlendedCost, 'f', 2, 64),
			strconv.FormatFloat(m.UnBlendedCost, 'f', 2, 64),
			strconv.FormatFloat(m.UsageQuantity, 'f', 2, 64),
		})
	}

	table.Render()
}
