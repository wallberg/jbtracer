Feature: Spheres

Scenario: A ray intersects a sphere at two points
  Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = 4.0
    And xs[1] = 6.0

Scenario: A ray intersects a sphere at a tangent
  Given r ← ray(point(0, 1, -5), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = 5.0
    And xs[1] = 5.0

Scenario: A ray misses a sphere
  Given r ← ray(point(0, 2, -5), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 0

Scenario: A ray originates inside a sphere
  Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = -1.0
    And xs[1] = 1.0

Scenario: A sphere is behind a ray
  Given r ← ray(point(0, 0, 5), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0] = -6.0
    And xs[1] = -4.0

Scenario: Intersect sets the object on the intersection
  Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
    And s ← sphere()
  When xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0].object = s
    And xs[1].object = s

Scenario: A sphere's default transformation
  Given s ← sphere()
  Then s.transform = identity_matrix

Scenario: Changing a sphere's transformation
  Given s ← sphere()
    And t ← translation(2, 3, 4)
  When set_transform(s, t)
  Then s.transform = t

Scenario: Intersecting a scaled sphere with a ray
  Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
    And s ← sphere()
    And t ← scaling(2, 2, 2)
  When set_transform(s, t)
    And xs ← intersect(s, r)
  Then xs.count = 2
    And xs[0].t = 3
    And xs[1].t = 7

Scenario: Intersecting a translated sphere with a ray
  Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
    And s ← sphere()
    And t ← translation(5, 0, 0)
  When set_transform(s, t)
    And xs ← intersect(s, r)
  Then xs.count = 0

Scenario: The normal on a sphere at a point on the x axis
  Given s ← sphere()
  When n ← normal_at(s, point(1, 0, 0))
  Then n = vector(1, 0, 0)

Scenario: The normal on a sphere at a point on the y axis
  Given s ← sphere()
  When n ← normal_at(s, point(0, 1, 0))
  Then n = vector(0, 1, 0)

Scenario: The normal on a sphere at a point on the z axis
  Given s ← sphere()
  When n ← normal_at(s, point(0, 0, 1))
  Then n = vector(0, 0, 1)

Scenario: The normal on a sphere at a nonaxial point
  Given s ← sphere()
                            # √3/3 = 0.57735
  When n ← normal_at(s, point(0.57735, 0.57735, 0.57735))
  Then n = vector(0.57735, 0.57735, 0.57735)

Scenario: The normal is a normalized vector
  Given s ← sphere()
  When n ← normal_at(s, point(0.57735, 0.57735, 0.57735))
  Then n = normalize(n)

Scenario: Computing the normal on a translated sphere
  Given s ← sphere()
    And m ← translation(0, 1, 0)
    And set_transform(s, m)
  When n ← normal_at(s, point(0, 1.70711, -0.70711))
  Then n = vector(0, 0.70711, -0.70711)

Scenario: Computing the normal on a transformed sphere
  Given s ← sphere()
    And scaling ← scaling(1, 0.5, 1)
                            # π/5
    And rotation ← rotation_z(0.628318)
    And m ← scaling * rotation
    And set_transform(s, m)
                               # √2/2,    -√2/2
  When n ← normal_at(s, point(0, 0.707106, -0.707106))
  Then n = vector(0, 0.97014, -0.24254)

Scenario: A sphere has a default material
  Given s ← sphere()
  When m1 ← s.material
    And m2 ← material()
  Then m1 = material m2

Scenario: A sphere may be assigned a material
  Given s ← sphere()
    And m ← material()
    And m.ambient ← 1
  When s.material ← m
    And m2 ← s.material
  Then m = material m2

Scenario: A helper for producing a sphere with a glassy material
  Given s ← glass_sphere()
    And m ← s.material
  Then s.transform = identity_matrix
    And m.transparency = 1.0
    And m.refractive_index = 1.5 

