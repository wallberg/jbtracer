package jbtracer

import (
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
)

func matrix(m1name string, table *godog.Table) error {
	rows := len(table.Rows)
	cols := len(table.Rows[0].Cells)
	if rows != cols {
		return fmt.Errorf("Matrix currently only supports square dimensions")
	}

	m := NewMatrix(rows)
	for i, row := range table.Rows {
		for j, col := range row.Cells {
			if f, err := strconv.ParseFloat(col.Value, 32); err != nil {
				return err
			} else {
				m.Set(i, j, (float32)(f))
			}
		}
	}

	matrices[m1name] = m

	return nil
}

func equalsMatrixCell(m1name string, i, j int, expected float32) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	got := m1.Get(i, j)
	if got != expected {
		return fmt.Errorf("Expected %s[%d,%d] = %v; got %v", m1name, i, j, expected, got)
	}
	return nil
}

func equalsMatrix(m1name, op, m2name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	if m2, ok = matrices[m2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m2name)
	}

	var expected bool
	if op == "=" {
		expected = true
	} else {
		expected = false
	}
	got := m1.Equal(m2)

	if got != expected {
		return fmt.Errorf("Expected %s = %s is %t; got %t", m1name, m2name, expected, got)
	}
	return nil
}
