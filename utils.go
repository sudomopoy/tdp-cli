package main

import (
	"fmt"
	"os"

	"github.com/lensesio/tableprinter"
)

const tab = "\t"
const row = "|\t\t\t\t  |"
const row_line = "--------------------"

func showTable(arr [][]string) {
	if len(arr) == 0 {
		fmt.Println(cli_NOTHING_TO_DO)
	}
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.RowLine = true
	printer.Print(arr)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
