# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

A Go implementation of the ray tracer from *The Ray Tracer Challenge* by Jamis Buck, built chapter-by-chapter following the book's test-driven approach. The root package `github.com/wallberg/jbtracer` contains the ray tracer engine; `cmd/chapterN/` contains standalone `main` programs (one per book chapter) that render a PPM image using the engine as it existed at that point in the book.

## Commands

A `Taskfile.yml` wraps the commands below — run `task --list` to see them (`task build`, `task test`, `task test:feature FEATURE=spheres`, `task test:tag TAG=@wip`, `task vet`, `task render CHAPTER=chapter11`).

- Build the engine: `go build .`
- Run the full test suite (unit tests + Cucumber/godog scenarios): `go test .`
- Run a single godog feature file: `go test . -args features/spheres.feature` (paths are positional args, passed after `-args`, not a `--godog.paths` flag)
- Run godog scenarios by tag: `go test . -args --godog.tags=@wip`
- Run only Go unit tests matching a name: `go test . -run TestSphere` (only exercises Go `Test*` functions, not godog scenarios, since godog scenarios all run under `TestMain`)
- Render a chapter's scene: `go run ./cmd/chapter11` (outputs a PPM image to stdout — redirect to a file, e.g. `go run ./cmd/chapter11 > out.ppm`)

Do not use `go build ./...` or `go test ./...` to validate changes to the core engine: many `cmd/chapterN/` programs are frozen snapshots written against the engine's API *at that point in the book* and will not compile against the current API (e.g. `cmd/chapter4`, `cmd/chapter6`, `cmd/chapter7/snowman` currently fail to build). This is expected — those programs are historical artifacts, not maintained call sites. Only `cmd/chapter10`, `cmd/chapter11`, and the root package are expected to build cleanly against the current API; check `go build ./cmd/chapterN` for the specific chapter you're touching.

## Architecture

### Feature files: `features/` vs `code/`

- `features/*.feature` is the **live** Cucumber source used by `go test .` (referenced by relative path from the module root in `jbtracer_test.go`).
- `code/` (gitignored, not tracked) holds the original zip of book source downloaded from the publisher (feature files, pseudocode, sample `.obj` files) — it's a local reference copy, not part of the build. `code/features/*.feature` is often more complete than `features/*.feature` for later chapters (e.g. cubes, cylinders, groups, CSG, triangles, obj_file) that haven't been ported into `features/` yet.
- When implementing a new chapter, the workflow is: copy/adapt the relevant scenarios from `code/features/*.feature` into `features/*.feature`, then add matching step definitions and step regexes in `jbtracer_test.go`.

### Step definitions live in one file

All godog step regexes and their Go implementations are registered in `InitializeScenario` in `jbtracer_test.go`, grouped by topic (tuples, matrices, rays, spheres, materials, world, camera, patterns, etc.) with package-level `var` scenario state shared across steps. When adding a feature, extend this file rather than creating a parallel step-registration mechanism.

### Shape / Pattern dispatch pattern

`Shape` (shapes.go) and `Pattern` (patterns.go) are interfaces implemented by concrete types (`Sphere`, `Plane`, ...). Each interface has package-level wrapper functions — `Intersections(s Shape, r *Ray)`, `NormalAt(s Shape, worldPoint *Tuple)`, `PatternAt(pattern Pattern, object Shape, worldPoint *Tuple)` — that handle the common world-space/object-space transform math (inverting the shape's transform, converting points/normals between spaces) before delegating to the concrete type's own `Intersections`/`NormalAt`/`PatternAt` method for the shape-specific geometry. Concrete implementations work entirely in object space; always call the package-level wrapper, not the interface method directly, when operating in world space.

### Rendering pipeline

`World` (world.go) holds `Light` and `Objects []Shape`. `Camera.Render` (camera.go) casts one `Ray` per pixel via `RayForPixel`, resolves color via `World.ColorAt` → `Intersections` → `Hit` → `PreparedComputations` (intersections.go) → `ShadeHit`, which combines `Material.Lighting` (Phong model, materials.go) with recursive `ReflectedColor`/`RefractedColor` bounded by a depth counter (`DefaultReflectedDepth`). Output goes through `Canvas.NewPPM` (canvas.go) to produce a PPM image written to stdout by the `cmd/chapterN` programs.

### Scene description YAML (`code/cover.yml`)

Not currently consumed by any Go code — it's the book's YAML scene-description format (`add`/`define`/`extend`, material/transform references) for potential future implementation of a scene-file loader, kept as reference alongside `code/`.
