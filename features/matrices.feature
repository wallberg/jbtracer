Feature: Matrices

Scenario: Constructing and inspecting a 4x4 matrix
  Given the following 4x4 matrix M:
    |  1   |  2   |  3   |  4   |
    |  5.5 |  6.5 |  7.5 |  8.5 |
    |  9   | 10   | 11   | 12   |
    | 13.5 | 14.5 | 15.5 | 16.5 |
  Then M[0,0] = 1
    And M[0,3] = 4
    And M[1,0] = 5.5
    And M[1,2] = 7.5
    And M[2,2] = 11
    And M[3,0] = 13.5
    And M[3,2] = 15.5

Scenario: A 2x2 matrix ought to be representable
  Given the following 2x2 matrix M:
    | -3 |  5 |
    |  1 | -2 |
  Then M[0,0] = -3
    And M[0,1] = 5
    And M[1,0] = 1
    And M[1,1] = -2

Scenario: A 3x3 matrix ought to be representable
  Given the following 3x3 matrix M:
    | -3 |  5 |  0 |
    |  1 | -2 | -7 |
    |  0 |  1 |  1 |
  Then M[0,0] = -3
    And M[1,1] = -2
    And M[2,2] = 1

Scenario: Matrix equality with identical matrices
  Given the following matrix A:
      | 1 | 2 | 3 | 4 |
      | 5 | 6 | 7 | 8 |
      | 9 | 8 | 7 | 6 |
      | 5 | 4 | 3 | 2 |
    And the following matrix B:
      | 1 | 2 | 3 | 4 |
      | 5 | 6 | 7 | 8 |
      | 9 | 8 | 7 | 6 |
      | 5 | 4 | 3 | 2 |
  Then A = matrix B

Scenario: Matrix equality with different matrices
  Given the following matrix A:
      | 1 | 2 | 3 | 4 |
      | 5 | 6 | 7 | 8 |
      | 9 | 8 | 7 | 6 |
      | 5 | 4 | 3 | 2 |
    And the following matrix B:
      | 2 | 3 | 4 | 5 |
      | 6 | 7 | 8 | 9 |
      | 8 | 7 | 6 | 5 |
      | 4 | 3 | 2 | 1 |
  Then A != matrix B

Scenario: Multiplying two matrices
  Given the following matrix A:
      | 1 | 2 | 3 | 4 |
      | 5 | 6 | 7 | 8 |
      | 9 | 8 | 7 | 6 |
      | 5 | 4 | 3 | 2 |
    And the following matrix B:
      | -2 | 1 | 2 |  3 |
      |  3 | 2 | 1 | -1 |
      |  4 | 3 | 6 |  5 |
      |  1 | 2 | 7 |  8 |
    And the following matrix C:
      | 20|  22 |  50 |  48 |
      | 44|  54 | 114 | 108 |
      | 40|  58 | 110 | 102 |
      | 16|  26 |  46 |  42 |
  Then A * B = matrix C

Scenario: A matrix multiplied by a tuple
  Given the following matrix A:
      | 1 | 2 | 3 | 4 |
      | 2 | 4 | 4 | 2 |
      | 8 | 6 | 4 | 1 |
      | 0 | 0 | 0 | 1 |
    And b ← tuple(1, 2, 3, 1)
    And c ← tuple(18, 24, 33, 1)
  Then A * b = tuple c

Scenario: Multiplying a matrix by the identity matrix
  Given the following matrix A:
    | 0 | 1 |  2 |  4 |
    | 1 | 2 |  4 |  8 |
    | 2 | 4 |  8 | 16 |
    | 4 | 8 | 16 | 32 |
  Then A * identity_matrix = matrix A

Scenario: Multiplying the identity matrix by a tuple
  Given a ← tuple(1, 2, 3, 4)
  Then identity_matrix * a = tuple a

Scenario: Transposing a matrix
  Given the following matrix A:
      | 0 | 9 | 3 | 0 |
      | 9 | 8 | 0 | 8 |
      | 1 | 8 | 5 | 3 |
      | 0 | 0 | 5 | 8 |
    And the following matrix B:
      | 0 | 9 | 1 | 0 |
      | 9 | 8 | 8 | 0 |
      | 3 | 0 | 5 | 5 |
      | 0 | 8 | 3 | 8 |
    And C ← transpose(A)
  Then B = matrix C

Scenario: Transposing the identity matrix
  Given A ← transpose(identity_matrix)
  Then A = matrix identity_matrix

Scenario: Calculating the determinant of a 2x2 matrix
  Given the following matrix A:
      |  1 | 5 |
      | -3 | 2 |
    And b ← determinant(A)
    And c ← scalar(17)
  Then b = scalar c

Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix
  Given the following matrix A:
      |  1 | 5 |  0 |
      | -3 | 2 |  7 |
      |  0 | 6 | -3 |
    And the following matrix B:
      | -3 | 2 |
      |  0 | 6 |
    And C ← submatrix(A, 0, 2)
  Then B = matrix C

Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix
  Given the following matrix A:
      | -6 |  1 |  1 |  6 |
      | -8 |  5 |  8 |  6 |
      | -1 |  0 |  8 |  2 |
      | -7 |  1 | -1 |  1 |
    And the following matrix B:
      | -6 |  1 | 6 |
      | -8 |  8 | 6 |
      | -7 | -1 | 1 |
    And C ← submatrix(A, 2, 1)
  Then B = matrix C

Scenario: Calculating a minor of a 3x3 matrix
  Given the following matrix A:
      |  3 |  5 |  0 |
      |  2 | -1 | -7 |
      |  6 | -1 |  5 |
    And B ← submatrix(A, 1, 0)
    And b ← determinant(B)
    And c ← minor(A, 1, 0)
  Then b = scalar 25
      And c = scalar 25

Scenario: Calculating a cofactor of a 3x3 matrix
  Given the following matrix A:
      |  3 |  5 |  0 |
      |  2 | -1 | -7 |
      |  6 | -1 |  5 |
    And min00 ← minor(A, 0, 0)
    And cof00 ← cofactor(A, 0, 0)
    And min10 ← minor(A, 1, 0)
    And cof10 ← cofactor(A, 1, 0)
  Then min00 = scalar -12
    And cof00 = scalar -12
    And min10 = scalar 25
    And cof10 = scalar -25

Scenario: Calculating the determinant of a 3x3 matrix
  Given the following matrix A:
    |  1 |  2 |  6 |
    | -5 |  8 | -4 |
    |  2 |  6 |  4 |
    And cof00 ← cofactor(A, 0, 0)
    And cof01 ← cofactor(A, 0, 1)
    And cof02 ← cofactor(A, 0, 2)
    And det ← determinant(A)
  Then cof00 = scalar 56
    And cof01 = scalar 12
    And cof02 = scalar -46
    And det = scalar -196

Scenario: Calculating the determinant of a 4x4 matrix
  Given the following 4x4 matrix A:
    | -2 | -8 |  3 |  5 |
    | -3 |  1 |  7 |  3 |
    |  1 |  2 | -9 |  6 |
    | -6 |  7 |  7 | -9 |
    And cof00 ← cofactor(A, 0, 0)
    And cof01 ← cofactor(A, 0, 1)
    And cof02 ← cofactor(A, 0, 2)
    And cof03 ← cofactor(A, 0, 3)
    And det ← determinant(A)
  Then cof00 = scalar 690
    And cof01 = scalar 447
    And cof02 = scalar 210
    And cof03 = scalar 51
    And det = scalar -4071

Scenario: Testing an invertible matrix for invertibility
  Given the following 4x4 matrix A:
    |  6 |  4 |  4 |  4 |
    |  5 |  5 |  7 |  6 |
    |  4 | -9 |  3 | -7 |
    |  9 |  1 |  7 | -6 |
  And det ← determinant(A)
  Then det = scalar -2120
    And A is invertible

Scenario: Testing a noninvertible matrix for invertibility
  Given the following 4x4 matrix A:
    | -4 |  2 | -2 | -3 |
    |  9 |  6 |  2 |  6 |
    |  0 | -5 |  1 | -5 |
    |  0 |  0 |  0 |  0 |
  And det ← determinant(A)
  Then det = scalar 0
    And A is not invertible

Scenario: Calculating the inverse of a matrix
  Given the following 4x4 matrix A:
      | -5 |  2 |  6 | -8 |
      |  1 | -5 |  1 |  8 |
      |  7 |  7 | -6 | -7 |
      |  1 | -3 |  7 |  4 |
    And B ← inverse(A)
    And the following 4x4 matrix C:
      |  0.21805 |  0.45113 |  0.24060 | -0.04511 |
      | -0.80827 | -1.45677 | -0.44361 |  0.52068 |
      | -0.07895 | -0.22368 | -0.05263 |  0.19737 |
      | -0.52256 | -0.81391 | -0.30075 |  0.30639 |
    And det ← determinant(A)
    And cof23 ← cofactor(A, 2, 3)
    And cof32 ← cofactor(A, 3, 2)
  Then det = scalar 532
    And cof23 = scalar -160
    And cof32 = scalar 105
               # 105/532
    And B[2,3] = 0.19736
               # -160/532
    And B[3,2] = -0.30075

Scenario: Calculating the inverse of another matrix
  Given the following 4x4 matrix A:
      |  8 | -5 |  9 |  2 |
      |  7 |  5 |  6 |  1 |
      | -6 |  0 |  9 |  6 |
      | -3 |  0 | -9 | -4 |
    And the following 4x4 matrix B:
      | -0.15385 | -0.15385 | -0.28205 | -0.53846 |
      | -0.07692 |  0.12308 |  0.02564 |  0.03077 |
      |  0.35897 |  0.35897 |  0.43590 |  0.92308 |
      | -0.69231 | -0.69231 | -0.76923 | -1.92308 |
    And Ainv ← inverse(A)
  Then Ainv = matrix B

Scenario: Calculating the inverse of a third matrix
  Given the following 4x4 matrix A:
      |  9 |  3 |  0 |  9 |
      | -5 | -2 | -6 | -3 |
      | -4 |  9 |  6 |  4 |
      | -7 |  6 |  6 |  2 |
    And the following 4x4 matrix B:
      | -0.04074 | -0.07778 |  0.14444 | -0.22222 |
      | -0.07778 |  0.03333 |  0.36667 | -0.33333 |
      | -0.02901 | -0.14630 | -0.10926 |  0.12963 |
      |  0.17778 |  0.06667 | -0.26667 |  0.33333 |
    And Ainv ← inverse(A)
  Then Ainv = matrix B

Scenario: Multiplying a product by its inverse
  Given the following 4x4 matrix A:
      |  3 | -9 |  7 |  3 |
      |  3 | -8 |  2 | -9 |
      | -4 |  4 |  4 |  1 |
      | -6 |  5 | -1 |  1 |
    And the following 4x4 matrix B:
      |  8 |  2 |  2 |  2 |
      |  3 | -1 |  7 |  0 |
      |  7 |  0 |  5 |  4 |
      |  6 | -2 |  0 |  5 |
    And AB ← A * B
    And Binv ← inverse(B)
  Then AB * Binv = matrix A
