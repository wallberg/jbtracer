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
    And full_quarter ← rotation_y(1.57079)
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

