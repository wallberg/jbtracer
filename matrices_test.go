package jbtracer

import (
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
)

func tableToMatrix(table *godog.Table) (*Matrix, error) {
	rows := len(table.Rows)
	cols := len(table.Rows[0].Cells)
	if rows != cols {
		return nil, fmt.Errorf("Matrix currently only supports square dimensions")
	}

	m := NewMatrix(rows)
	for i, row := range table.Rows {
		for j, col := range row.Cells {
			if f, err := strconv.ParseFloat(col.Value, 32); err != nil {
				return nil, err
			} else {
				m.Set(i, j, (float32)(f))
			}
		}
	}

	return m, nil
}

func matrix(m1name string, table *godog.Table) error {

	if m, err := tableToMatrix(table); err != nil {
		return err
	} else {
		matrices[m1name] = m
	}

	return nil
}

func matrixCellEqual(m1name string, i, j int, expected float32) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	got := m1.Get(i, j)
	if got != expected {
		return fmt.Errorf("Expected %s[%d,%d] = %v; got %v", m1name, i, j, expected, got)
	}
	return nil
}

func matrixEqual(m1name, op, m2name string) error {
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
		return fmt.Errorf("Expected %s %s %v; got %v", m1name, op, m1, m2)
	}
	return nil
}

func matrixMultiply(m1name, m2name, m3name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	if m2, ok = matrices[m2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m2name)
	}
	if m3, ok = matrices[m3name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m3name)
	}

	expected := m3
	got := m1.Multiply(m2)
	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s * %s = %v; got %v", m1name, m2name, expected, got)
	}
	return nil
}

func matrixMultiplyTuple(m1name, t1name, t2name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}
	if t1, ok = tuples[t1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t1name)
	}
	if t2, ok = tuples[t2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", t2name)
	}

	expected := t2
	got := m1.MultiplyTuple(t1)

	if !got.Equal(expected) {
		return fmt.Errorf("Expected %s * %s = %v; got %v", m1name, t1name, expected, got)
	}
	return nil
}

func matrixTranspose(m1name, m2name string) error {
	if m2, ok = matrices[m2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m2name)
	}

	matrices[m1name] = m2.Transpose()
	return nil
}

func matrixDeterminant(s1name, m1name string) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}

	scalars[s1name] = m1.Determinant()
	return nil
}

func scalar(s1name string, value float32) error {
	scalars[s1name] = value
	return nil
}

func scalarEqual(s1name, s2name string) error {
	if s1, ok = scalars[s1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", s1name)
	}
	if s2, ok = scalars[s2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", s2name)
	}

	expected := true
	got := s1 == s2

	if got != expected {
		return fmt.Errorf("Expected %s = %s is %t; got %t", s1name, s2name, expected, got)
	}
	return nil
}

func submatrix(m1name, m2name string, i, j int) error {
	if m2, ok = matrices[m2name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m2name)
	}

	matrices[m1name] = m2.Submatrix(i, j)
	return nil
}

func matrixMinor(s1name, m1name string, i, j int) error {
	if m1, ok = matrices[m1name]; !ok {
		return fmt.Errorf("Unknown symbol %s", m1name)
	}

	scalars[s1name] = m1.Minor(i, j)
	return nil
}
