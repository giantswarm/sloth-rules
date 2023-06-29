# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Update VpaComponentTooManyRestarts SLO alert to prevent it from paging too much.

## [0.10.0] - 2023-06-26

### Added

- Add missing external labels in aggregations.

## [0.9.0] - 2023-06-22

### Changed

- Add missing dashboard and opsrecipe to Prometheus Availability SLO alert.

## [0.8.4] - 2023-06-22

### Changed

- Reduce Prometheus Availability SLO target to 75%.

### Fixed

- Fix missing cluster id for Prometheus Availability SLO alert.

## [0.8.3] - 2023-06-19

### Fixed

- Rerelease with correct version.

## [0.8.1] - 2023-06-19

### Fixed

- Fix Prometheus Availability SLO target

## [0.8.2] - 2023-06-15

### Changed

- revert kaas-slo team for the SLOs investigations and adjustments

## [0.8.1] - 2023-06-15

### Changed

- create kaas-slo team for the SLOs investigations and adjustments

## [0.8.0] - 2023-06-14

### Added

- Added `ContainerdDiskUsage` slo.

## [0.7.1] - 2023-06-08

### Changed

- Rename ticket severity from slack to notify.

## [0.7.0] - 2023-06-08

### Added

- Added support for annotation in slo definition.

## [0.6.0] - 2023-06-06

### Added

- Added `PrometheusUnavailable` slo.

## [0.5.0] - 2023-06-02

### Changed

- Allow users to add specific labels to their SLO definition.

## [0.4.1] - 2023-06-02

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

[Unreleased]: https://github.com/giantswarm/sloth-rules/compare/v0.10.0...HEAD
[0.10.0]: https://github.com/giantswarm/sloth-rules/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/giantswarm/sloth-rules/compare/v0.8.4...v0.9.0
[0.8.4]: https://github.com/giantswarm/sloth-rules/compare/v0.8.3...v0.8.4
[0.8.3]: https://github.com/giantswarm/sloth-rules/compare/v0.8.1...v0.8.3
[0.8.1]: https://github.com/giantswarm/sloth-rules/compare/v0.8.2...v0.8.1
[0.8.2]: https://github.com/giantswarm/sloth-rules/compare/v0.8.1...v0.8.2
[0.8.1]: https://github.com/giantswarm/sloth-rules/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/giantswarm/sloth-rules/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/sloth-rules/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/sloth-rules/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/sloth-rules/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/sloth-rules/compare/v0.4.1...v0.5.0
[0.4.1]: https://github.com/giantswarm/sloth-rules/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/sloth-rules/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/sloth-rules/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/sloth-rules/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/sloth-rules/compare/v0.0.0...v0.1.0
