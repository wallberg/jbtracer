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
