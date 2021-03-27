package gui

import (
	"log"
	"strconv"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/model"
	"github.com/rivo/tview"
)

var (
	App *tview.Application = tview.NewApplication()
	Flex  *CostFlex  = &CostFlex{}
	Table *CostTable = &CostTable{}
)

type CostFlex struct {
	*tview.Flex
}

type CostTable struct {
	*tview.Table
}

func InitTview() {
	Table = &CostTable{tview.NewTable().
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectable(true, false)}
	Table.SetBorder(true).SetTitle("")

	Flex = &CostFlex{tview.NewFlex().
		AddItem(Table, 0, 1, true),
	}
	_ = Flex
}

func InitTable() {
	Table.SetCell(0, 0, &tview.TableCell{
		Text:            "Users",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           0,
		BackgroundColor: 0,
	})
	Table.SetCell(1, 0, &tview.TableCell{
		Text:            "BlendedCost",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           0,
		BackgroundColor: 0,
	})
	Table.SetCell(2, 0, &tview.TableCell{
		Text:            "UnblendedCost",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           0,
		BackgroundColor: 0,
	})
	Table.SetCell(3, 0, &tview.TableCell{
		Text:            "UsageQuantity",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           0,
		BackgroundColor: 0,
	})
}

func DisplayMetricSummary(ms []model.MetricSummary) {
	for i, m := range ms {
		Table.SetCell(0, i+1, &tview.TableCell{
			Text:            m.User,
			NotSelectable:   false,
			Align:           tview.AlignLeft,
			Color:           0,
			BackgroundColor: 0,
		})
		Table.SetCell(1, i+1, &tview.TableCell{
			Text:            strconv.FormatFloat(m.BlendedCost, 'f', 2, 64),
			NotSelectable:   false,
			Align:           tview.AlignLeft,
			Color:           0,
			BackgroundColor: 0,
		})
		Table.SetCell(2, i+1, &tview.TableCell{
			Text:            strconv.FormatFloat(m.UnBlendedCost, 'f', 2, 64),
			NotSelectable:   false,
			Align:           tview.AlignLeft,
			Color:           0,
			BackgroundColor: 0,
		})
		Table.SetCell(3, i+1, &tview.TableCell{
			Text:            strconv.FormatFloat(m.UsageQuantity, 'f', 2, 64),
			NotSelectable:   false,
			Align:           tview.AlignLeft,
			Color:           0,
			BackgroundColor: 0,
		})
	}

	if err := App.SetRoot(Flex, true).EnableMouse(true).Run(); err != nil {
	    log.Fatal(err)
	}
}
