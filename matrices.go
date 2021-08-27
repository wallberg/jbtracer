package jbtracer

type Matrix struct {
	size  int
	cells []float32
}

// NewMatrix returns a newly initialized size x size matrix
func NewMatrix(size int) *Matrix {
	return &Matrix{
		size:  size,
		cells: make([]float32, size*size),
	}
}

// Set sets the value of the matrix cell at (x,y)
func (m *Matrix) Set(i, j int, value float32) {
	m.cells[i*m.size+j] = value
}

// Get returns the value of the matrix cell at (x,y)
func (m *Matrix) Get(i, j int) float32 {
	return m.cells[i*m.size+j]
}

// Equal determines if this matrix is the same as the provided matrix
func (a *Matrix) Equal(b *Matrix) bool {
	if a.size != b.size {
		return false
	}
	for i, value := range a.cells {
		if b.cells[i] != value {
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
			var value float32 = 0.0
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
