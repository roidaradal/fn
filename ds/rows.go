package ds

import (
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/str"
)

const rowGlue string = ","

type DataRows struct {
	Rows string
}

func (d DataRows) GetRows() [][]string {
	rows := str.CleanSplit(d.Rows, "\n")
	return fn.Map(rows, func(row string) []string {
		return str.CleanSplit(row, rowGlue)
	})
}
