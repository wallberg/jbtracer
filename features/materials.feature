Feature: Materials

Background:
  Given m ← material()
    And position ← point(0, 0, 0)

Scenario: The default material
  Given m ← material()
  Then m.color = color(1, 1, 1)
    And m.ambient = 0.1
    And m.diffuse = 0.9
    And m.specular = 0.9
    And m.shininess = 200.0

Scenario: Reflectivity for the default material
  Given m ← material()
  Then m.reflective = 0.0

# Scenario: Transparency and Refractive Index for the default material
#   Given m ← material()
#   Then m.transparency = 0.0
#     And m.refractive_index = 1.0

Scenario: Lighting with the eye between the light and the surface
  Given eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(1.9, 1.9, 1.9)

Scenario: Lighting with the eye between light and surface, eye offset 45°
  Given eyev ← vector(0, 0.70711, -0.70711)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(1.0, 1.0, 1.0)

Scenario: Lighting with eye opposite surface, light offset 45°
  Given eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 10, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(0.7364, 0.7364, 0.7364)

Scenario: Lighting with eye in the path of the reflection vector
  Given eyev ← vector(0, -0.7071067812, -0.7071067812)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 10, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(1.6364, 1.6364, 1.6364)

Scenario: Lighting with the light behind the surface
  Given eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, 10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(0.1, 0.1, 0.1)

Scenario: Lighting with the surface in shadow
  Given eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
    And in_shadow ← true
  When result ← lighting(m, light, position, eyev, normalv, in_shadow)
  Then result = color(0.1, 0.1, 0.1)

Scenario: Lighting with a pattern applied
  Given black ← color(0, 0, 0)
    And white ← color(1, 1, 1)
    And pattern ← stripe_pattern(white, black)
    And m.pattern ← pattern
    And m.ambient ← 1
    And m.diffuse ← 0
    And m.specular ← 0
    And eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
  When p1 ← point(0.9, 0, 0)
    And p2 ← point(1.1, 0, 0)
    And c1 ← lighting(m, light, p1, eyev, normalv, in_shadow)
    And c2 ← lighting(m, light, p2, eyev, normalv, in_shadow)
  Then c1 = color(1, 1, 1)
    And c2 = color(0, 0, 0)
