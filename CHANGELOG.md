# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added `vpa-component-restarts` slo.

## [0.4.0] - 2023-05-10

### Changed

- Changed repo and chart name to `sloth-rules`
- Configured the repo according to GS repo's template
- added circleci jobs to push `sloth-rules` to all collections

## [0.3.0] - 2023-02-28

### Added

- Add `application.giantswarm.io/team` label to PrometheusRules.

## [0.2.0] - 2023-02-21

### Added

- Add etcd SLO.

### Changed

- Silence alerts when cluster is creating or deleting.
- Check if generated files are up to date.

### Fixed

- Fix typo in alert name.

## [0.1.0] - 2023-02-20

### Added

- Add API server SLO.
- Add CNI SLO.
- Add Controller-manager SLO.
- Add DNS SLO.
- Add Scheduler SLO.

[Unreleased]: https://github.com/giantswarm/sloth-rules/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/sloth-rules/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/sloth-rules/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/sloth-rules/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/sloth-rules/compare/v0.0.0...v0.1.0
