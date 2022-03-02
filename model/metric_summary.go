package model

import (
	"os"
	"strconv"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/setting"
	"github.com/olekukonko/tablewriter"
)

type MetricSummary struct {
	User          string
	BlendedCost   float64
	UnBlendedCost float64
	UsageQuantity float64
}

func DisplayMetricSummaryTableView(metrics setting.Metrics, ms []MetricSummary) {
	table := tablewriter.NewWriter(os.Stdout)

	headers := make([]string, 0)
	headers = append(headers, "USERS")
	for _, m := range metrics.M {
		headers = append(headers, m)
	}
	table.SetHeader(headers)

	for _, s := range ms {
		values := make([]string, 0)
		values = append(values, s.User)
		for _, m := range metrics.M {
			if m == "BLENDED_COST" {
				values = append(values, strconv.FormatFloat(s.BlendedCost, 'f', 2, 64))
			}
			if m == "UNBLENDED_COST" {
				values = append(values, strconv.FormatFloat(s.UnBlendedCost, 'f', 2, 64))
			}
			if m == "USAGE_QUANTITY" {
				values = append(values, strconv.FormatFloat(s.UsageQuantity, 'f', 2, 64))
			}
		}
		table.Append(values)
	}

	table.Render()
}
