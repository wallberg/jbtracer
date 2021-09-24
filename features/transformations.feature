Feature: Matrix Transformations

Scenario: Multiplying by a translation matrix
  Given transform ← translation(5, -3, 2)
    And p ← point(-3, 4, 5)
    And e ← point(2, 1, 7)
   Then transform * p = tuple e

Scenario: Multiplying by the inverse of a translation matrix
  Given transform ← translation(5, -3, 2)
    And inv ← inverse(transform)
    And p ← point(-3, 4, 5)
    And e ← point(-8, 7, 3)
   Then inv * p = tuple e

Scenario: Translation does not affect vectors
  Given transform ← translation(5, -3, 2)
    And v ← vector(-3, 4, 5)
   Then transform * v = tuple v

Scenario: A scaling matrix applied to a point
  Given transform ← scaling(2, 3, 4)
    And p ← point(-4, 6, 8)
    And e ← point(-8, 18, 32)
   Then transform * p = tuple e
    And e is a point

Scenario: A scaling matrix applied to a vector
  Given transform ← scaling(2, 3, 4)
    And v ← vector(-4, 6, 8)
    And e ← vector(-8, 18, 32)
   Then transform * v = tuple e
    And e is a vector

Scenario: Multiplying by the inverse of a scaling matrix
  Given transform ← scaling(2, 3, 4)
    And inv ← inverse(transform)
    And v ← vector(-4, 6, 8)
    And e ← vector(-2, 2, 2)
   Then inv * v = tuple e
    And e is a vector

Scenario: Reflection is scaling by a negative value
  Given transform ← scaling(-1, 1, 1)
    And p ← point(2, 3, 4)
    And e ← point(-2, 3, 4)
  Then transform * p = tuple e
    And e is a point

Scenario: Rotating a point around the x axis
  Given p ← point(0, 1, 0)
                                # π / 4
    And half_quarter ← rotation_x(0.78539)
                                # π / 2
    And full_quarter ← rotation_x(1.57079)
                      # √2/2     √2/2
    And e1 ← point(0, 0.707106, 0.707106)
    And e2 ← point(0, 0, 1)
                      # √2/2     √2/2
    And e3 ← point(0, 0.707106, 0.707106)
    And e4 ← point(0, 0, 1)
  Then half_quarter * p = tuple e1
    And e1 is a point
    And full_quarter * p = tuple e2
    And e2 is a point
    And half_quarter * p = tuple e3
    And e3 is a point
    And full_quarter * p = tuple e4
    And e4 is a point

Scenario: The inverse of an x-rotation rotates in the opposite direction
  Given p ← point(0, 1, 0)
                                # π / 4
    And half_quarter ← rotation_x(0.78539)
    And inv ← inverse(half_quarter)
                             # √2/2    -√2/2
            And e ← point(0, 0.707106, -0.707106)
  Then inv * p = tuple e
    And e is a point

Scenario: Rotating a point around the y axis
  Given p ← point(0, 0, 1)
                                # π / 4
    And half_quarter ← rotation_y(0.78539)
                                # π / 2
    And full_quarter ← rotation_y(1.570796)
                              # √2/2         √2/2
    And e1 ← point(0.707106, 0, 0.707106)
    And e2 ← point(1, 0, 0)
  Then half_quarter * p = tuple e1
    And e1 is a point
    And full_quarter * p = tuple e2
    And e2 is a point

Scenario: Rotating a point around the z axis
  Given p ← point(0, 1, 0)
                                # π / 4
    And half_quarter ← rotation_z(0.78539)
                                # π / 2
    And full_quarter ← rotation_z(1.57079)
                             # √2/2    -√2/2
    And e1 ← point(-0.707106, 0.707106, 0)
    And e2 ← point(-1, 0, 0)
  Then half_quarter * p = tuple e1
    And e1 is a point
    And full_quarter * p = tuple e2
    And e2 is a point

Scenario: A shearing transformation moves x in proportion to y
  Given transform ← shearing(1, 0, 0, 0, 0, 0)
    And p ← point(2, 3, 4)
    And e ← point(5, 3, 4)
  Then transform * p = tuple e

Scenario: A shearing transformation moves x in proportion to z
  Given transform ← shearing(0, 1, 0, 0, 0, 0)
    And p ← point(2, 3, 4)
    And e ← point(6, 3, 4)
  Then transform * p = tuple e

Scenario: A shearing transformation moves y in proportion to x
  Given transform ← shearing(0, 0, 1, 0, 0, 0)
    And p ← point(2, 3, 4)
    And e ← point(2, 5, 4)
  Then transform * p = tuple e

Scenario: A shearing transformation moves y in proportion to z
  Given transform ← shearing(0, 0, 0, 1, 0, 0)
    And p ← point(2, 3, 4)
    And e ← point(2, 7, 4)
  Then transform * p = tuple e

Scenario: A shearing transformation moves z in proportion to x
  Given transform ← shearing(0, 0, 0, 0, 1, 0)
    And p ← point(2, 3, 4)
    And e ← point(2, 3, 6)
  Then transform * p = tuple e

Scenario: A shearing transformation moves z in proportion to y
  Given transform ← shearing(0, 0, 0, 0, 0, 1)
    And p ← point(2, 3, 4)
    And e ← point(2, 3, 7)
  Then transform * p = tuple e

Scenario: Individual transformations are applied in sequence
  Given p ← point(1, 0, 1)
                     # π / 2
    And A ← rotation_x(1.570796)
    And B ← scaling(5, 5, 5)
    And C ← translation(10, 5, 7)
  # apply rotation first
  When p2 ← A * tuple p
    And e2 ← point(1, -1, 0)
  Then p2 = tuple e2
  # then apply scaling
  When p3 ← B * tuple p2
    And e3 ← point(5, -5, 0)
  Then p3 = tuple e3
  # then apply translation
  When p4 ← C * tuple p3
    And e4 ← point(15, 0, 7)
  Then p4 = tuple e4

Scenario: Chained transformations must be applied in reverse order
  Given p ← point(1, 0, 1)
                     # π / 2
    And A ← rotation_x(1.570796)
    And B ← scaling(5, 5, 5)
    And C ← translation(10, 5, 7)
  When T1 ← C * B
    And T2 ← T1 * A
    And e ← point(15, 0, 7)
  Then T2 * p = tuple e

Scenario: The transformation matrix for the default orientation
  Given from ← point(0, 0, 0)
    And to ← point(0, 0, -1)
    And up ← vector(0, 1, 0)
  When t ← view_transform(from, to, up)
  Then t = matrix identity_matrix

Scenario: A view transformation matrix looking in positive z direction
  Given from ← point(0, 0, 0)
    And to ← point(0, 0, 1)
    And up ← vector(0, 1, 0)
  When t ← view_transform(from, to, up)
    And e ← scaling(-1, 1, -1)
  Then t = matrix e

Scenario: The view transformation moves the world
  Given from ← point(0, 0, 8)
    And to ← point(0, 0, 0)
    And up ← vector(0, 1, 0)
  When t ← view_transform(from, to, up)
    And e ← translation(0, 0, -8)
  Then t = matrix e

Scenario: An arbitrary view transformation
  Given from ← point(1, 3, 2)
    And to ← point(4, -2, 8)
    And up ← vector(1, 1, 0)
  When t ← view_transform(from, to, up)
    And the following 4x4 matrix e:
      | -0.50709 | 0.50709 |  0.67612 | -2.36643 |
      |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
      | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
      |  0.00000 | 0.00000 |  0.00000 |  1.00000 |
  Then t = matrix e
