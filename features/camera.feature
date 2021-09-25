Feature: Camera

Scenario: Constructing a camera
                             # π/2
  Given c ← camera(160, 120, 1.570796)
  Then c.hsize = 160
    And c.vsize = 120
                        # π/2
    And c.field_of_view = 1.570796
    And c.transform = identity_matrix

Scenario: The pixel size for a horizontal canvas
                             # π/2
  Given c ← camera(200, 125, 1.5707963268)
  Then c.pixel_size = 0.01

Scenario: The pixel size for a vertical canvas
                             # π/2
  Given c ← camera(125, 200, 1.5707963268)
  Then c.pixel_size = 0.01

Scenario: Constructing a ray through the center of the canvas
                             # π/2
  Given c ← camera(201, 101, 1.5707963268)
  When r ← ray_for_pixel(c, 100, 50)
  Then r.origin = point(0, 0, 0)
    And r.direction = vector(0, 0, -1)

Scenario: Constructing a ray through a corner of the canvas
                             # π/2
  Given c ← camera(201, 101, 1.5707963268)
  When r ← ray_for_pixel(c, 0, 0)
  Then r.origin = point(0, 0, 0)
    And r.direction = vector(0.66519, 0.33259, -0.66851)

Scenario: Constructing a ray when the camera is transformed
                             # π/2
  Given c ← camera(201, 101, 1.5707963268)
                         # π/4
  When roty ← rotation_y(0.7853981634)
    And translation ← translation(0, -2, 5)
    And transform ← roty * translation
    And c.transform ← transform
    And r ← ray_for_pixel(c, 100, 50)
  Then r.origin = point(0, 2, -5)
                             # √2/2           -√2/2
    And r.direction = vector(0.7071067812, 0, -0.7071067812)

# Scenario: Rendering a world with a camera
#   Given w ← default_world()
#     And c ← camera(11, 11, 1.5707963268)
#     And from ← point(0, 0, -5)
#     And to ← point(0, 0, 0)
#     And up ← vector(0, 1, 0)
#     And c.transform ← view_transform(from, to, up)
#   When image ← render(c, w)
#   Then pixel_at(image, 5, 5) = color(0.38066, 0.47583, 0.2855)
