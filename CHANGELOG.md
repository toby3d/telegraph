# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.1] - 2020-01-14
### Added
- This `CHANGELOG.md` file

### Changed
- Renamed `SUPPORT.md` to `CONTRIBUTING.md`
- Updated patrons list in `CONTRIBUTING.md`
- Updated patreon link
- Updated current year in `LICENSE.md`

### Removed
- `go:generate` comments

## [1.2.0] - 2020-01-14
### Added
- `.github/workflows` folder with migration actions for GitHub mirror
- Internal methods structures for JSON marshling in requests

### Changed
- GitLab CI configuration for caches and modules support
- `Makefile` for go modules support
- Format of the code
- Removed last linter warnings
- Added actual package import path into `doc.go`
- Use `time.Time` structure instead `[]int` in `GetViews` method (yeah, this violates the semantic versioning because backward compatibility breaks, but this is the last time, I promise)
- Updated go modules in `go.mod`/`go.sum`
- Force type initialization in constants section

### Fixed
- Use JSON body instead query strings for better cyrillic support #6
- Removed `charset` parameter in `Content-Type` header, because sometimes Telegraph answer BadRequest status

## [1.1.0] - 2019-07-24
### Added
- [EditorConfig](https://editorconfig.org/)
- GitLab CI configuration
- [PreCommit](https://pre-commit.com/) hooks configuration
- `Makefile` with useful snippets
- `SUPPORT.md` list with all contributors, helpers and patrons
- `go.mod` and `go.sum` for go modules support
- Individual methods tests files
- Added `types.go` with all available structures and types

### Changed
- `LICENSE.md` current year
- Actual package URLs in `README.md`
- Format `ContentFormat` code
- Moved all types from methods files into `types.go` file
- Used copy of structures in all methods instead links, because this arguments is required
- Renamed `request` method into `makeRequest`
- Added comments for false-positives linters warnings
- Renamed `request.go` to `telegraph.go`

### Removed
- Travis CI configuration due to migration from GitHub to GitLab native CI
- `CONTRIBUTORS.md` in favor of a single `SUPPORT.md` list
- `PATRONS.md` in favor of a single `SUPPORT.md` list
- Badges collection from `README.md`
- `invalid_test.go`/`valid_test.go` due individual methods tests files

## [1.0.0] - 2018-01-09
### Added
- Travis CI configuration
- Code of conduct info
- Contributors list
- License information
- Patrons list
- ReadMe
- All available methods
- Coverage tests

[Unreleased]: https://gitlab.com/toby3d/telegraph/compare/v1.2.1...develop
[1.2.1]: https://gitlab.com/toby3d/telegraph/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitlab.com/toby3d/telegraph/compare/v1.1.0...v1.2.0
[1.1.0]: https://gitlab.com/toby3d/telegraph/compare/v1.0.0...v1.1.0
[1.0.0]: https://gitlab.com/toby3d/telegraph/tree/v1.0.0