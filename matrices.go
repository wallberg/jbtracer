package jbtracer

import (
	"fmt"
)

type Matrix struct {
	size  int
	cells []float64
}

// NewMatrix returns a newly initialized size x size matrix
func NewMatrix(size int) *Matrix {
	return &Matrix{
		size:  size,
		cells: make([]float64, size*size),
	}
}

// Set sets the value of the matrix cell at (x,y)
func (m *Matrix) Set(i, j int, value float64) {
	m.cells[i*m.size+j] = value
}

// Get returns the value of the matrix cell at (x,y)
func (m *Matrix) Get(i, j int) float64 {
	return m.cells[i*m.size+j]
}

// Equal determines if this matrix is the same as the provided matrix
func (a *Matrix) Equal(b *Matrix) bool {
	if a.size != b.size {
		return false
	}
	for i, value := range a.cells {
		if !EqualFloat64(b.cells[i], value) {
			return false
		}
	}
	return true
}

// Multiply multiplies this matrix with the provided matrix
func (a *Matrix) Multiply(b *Matrix) *Matrix {
	c := NewMatrix(a.size)

	// Iterate over the cells of matrix C
	for i := 0; i < c.size; i++ {
		for j := 0; j < c.size; j++ {
			var value float64 = 0.0
			// Iterate over a row of A and a column of C
			for k := 0; k < c.size; k++ {
				value += a.Get(i, k) * b.Get(k, j)
			}
			// Set the new cell value
			c.Set(i, j, value)
		}
	}
	return c
}

// MultiplyTuple multiplies this matrix with the provided tuple
// Assumes a 4x4 matrix, since Tuple is fixed with a size of 4
func (a *Matrix) MultiplyTuple(b *Tuple) *Tuple {
	c := &Tuple{}

	// Iterate over the rows of A
	for i := 0; i < a.size; i++ {
		value := a.Get(i, 0) * b.X
		value += a.Get(i, 1) * b.Y
		value += a.Get(i, 2) * b.Z
		value += a.Get(i, 3) * b.W
		switch i {
		case 0:
			c.X = value
		case 1:
			c.Y = value
		case 2:
			c.Z = value
		case 3:
			c.W = value
		}
	}

	return c
}

// IdentityMatrix returns an identity matrix of size 4x4
func IdentityMatrix() *Matrix {
	m := NewMatrix(4)
	for i := 0; i < 4; i++ {
		m.Set(i, i, 1.0)
	}
	return m
}

// Transpose returns the transpose of this matrix
func (a *Matrix) Transpose() *Matrix {
	b := NewMatrix(a.size)
	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			b.Set(j, i, a.Get(i, j))
		}
	}
	return b
}

// Determinant returns the determinant of this matrix
func (a *Matrix) Determinant() float64 {
	if a.size == 2 {
		return a.Get(0, 0)*a.Get(1, 1) - a.Get(0, 1)*a.Get(1, 0)
	}

	var d float64 = 0
	for j := 0; j < a.size; j++ {
		d += a.Get(0, j) * a.Cofactor(0, j)
	}
	return d
}

// Submatrix returns a copy of this matrix, but with row i
// column j removed
func (a *Matrix) Submatrix(i, j int) *Matrix {
	b := NewMatrix(a.size - 1)
	for ai := 0; ai < a.size; ai++ {
		for aj := 0; aj < a.size; aj++ {
			if ai != i && aj != j {
				bi := ai
				if bi > i {
					bi--
				}
				bj := aj
				if bj > j {
					bj--
				}
				b.Set(bi, bj, a.Get(ai, aj))
			}
		}
	}
	return b
}

// Minor returns the determinant of the submatrix at row i
// and column j
func (a *Matrix) Minor(i, j int) float64 {
	return a.Submatrix(i, j).Determinant()
}

// Cofactor returns the cofactor of this matrix
func (a *Matrix) Cofactor(i, j int) float64 {
	m := a.Minor(i, j)
	if (i+j)%2 == 1 {
		m *= -1
	}
	return m
}

// Inverse returns the inverse of this matrix, which is obtained by
// creating a matrix of cofactors, transponsing it, then dividing it
// by the determinant.
// Return an error if the determinant is 0, which means the matrix
// is not invertible
func (a *Matrix) Inverse() (*Matrix, error) {
	d := a.Determinant()
	if EqualFloat64(d, 0) {
		return nil, fmt.Errorf("matrix is not invertible")
	}

	b := NewMatrix(a.size)
	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			b.Set(j, i, a.Cofactor(i, j)/d)
		}
	}
	return b, nil
}
