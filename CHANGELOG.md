# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.28.3] - 2024-02-27

### Changed

- Lower `AtlasOperatorsReconciliationError`'s objective from 99 to 95.

## [0.28.2] - 2024-02-27

### Changed

- Add `cancel_if_prometheus_agent_down` label to unsilenced API server alerts

## [0.28.1] - 2024-02-21

### Changed

- Add `cancel_if_cluster_status_updating` labels to unsilenced API server alerts

## [0.28.0] - 2024-02-14

### Added

- Add alert on API server instance availability, `ApiServerInstanceUnavailable`.

## [0.27.3] - 2024-02-06

### Changed

- Check `kubernetes` service availability in `ApiServerAvailabilityErrorRateTooHigh` alert.

## [0.27.2] - 2024-01-29

### Fixed

- Fixed missing parenthesis in `NodeConnTrackAlmostExhausted`

## [0.27.1] - 2024-01-22

### Changed

- Removed silence from ApiServerAvailabilityErrorRateTooHigh alert.

## [0.27.0] - 2024-01-16

### Added

- Add Atlas kube-builder operator reconciliation SLOs. Target is defined at 99% for now.

## [0.26.0] - 2024-01-10

### Changed

- Silence `VPAComponentUnavailable`.

## [0.25.0] - 2023-12-19

### Added

- Add conntrack exhaustion SLO.

### Changed

- Exclude `karpenter-renewer` pods from Karpenter SLO

## [0.24.0] - 2023-12-12

### Added

- Uptime alert for Karpenter

## [0.23.0] - 2023-10-04

### Removed

- Removed unneeded node role in `KubeletPLEGDurationTooHigh`.

## [0.22.0] - 2023-09-12

### Changed

- Change ownership from Atlas to Turtles for vertical pod autoscaler related alerts.
- Promote DNSServiceUnavailable, CoreDNSLatencyTooHigh and LocalDNSCacheLatencyTooHigh to 24h.

## [0.21.0] - 2023-08-25

### Added

- New DNSServiceUnavailable SLO based on DNS queries.

### Changed

- Add option to disable alerts on template.
- Disable DNS alerts for external zones.

## [0.20.0] - 2023-08-21

### Changed

- Add `cluster_type` label in template.

## [0.19.1] - 2023-08-17

### Changed

- Made the `PrometheusUnavailable` SLO only evaluate MCs' prometheuses.

## [0.19.0] - 2023-08-17

### Changed

- Move `dns` alerts from kaas/phoenix to empowerment/cabbage.
- Upgrade `PrometheusUnavailable` SLO.

### Removed

- Removed `starboard-exporter` alert.

## [0.18.0] - 2023-08-03

### Changed

- Update vpa's SLO to `VpaComponentUnavailable`.

## [0.17.0] - 2023-07-21

### Added

- Add `CoreDNSExternalDomainsLatencyTooHigh` and `LocalDNSCacheExternalDomainsLatencyTooHigh` DNS SLOs.

### Changed

- Unsilence `DNSResolutionUnavailable` and make it only page during BH.

## [0.16.0] - 2023-07-13

### Changed

- Enable Phoenix's `CoreDNSLatencyTooHigh` and `LocalDNSCacheLatencyTooHigh` SLOs.

## [0.15.0] - 2023-07-13

### Changed

- Only consider `local` DNS queries for latency SLO.

## [0.14.0] - 2023-07-12

### Changed

- Raise success threshold for `CoreDNSLatencyTooHigh` to 0.032s.

## [0.13.0] - 2023-07-12

### Changed

- Add 'silence: true' label to all phoenix's SLOs to be able to selectively enable them.
- Only use `*.local` zone to measure the latency of local DNS cache.

## [0.12.1] - 2023-07-11

### Fixed

- Added quoting to `alertLabels` to avoid validation errors on non-string-like types.

## [0.12.0] - 2023-07-11

### Changed

- Add ops recipe to DNS latency SLOs.
- Make DNS latency SLOs page on BH only.

## [0.11.0] - 2023-07-10

### Changed

- Removed kaas-phoenix-cni-latency-intranode (cilium_node_connectivity_latency_seconds dropped)

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

[Unreleased]: https://github.com/giantswarm/sloth-rules/compare/v0.28.3...HEAD
[0.28.3]: https://github.com/giantswarm/sloth-rules/compare/v0.28.2...v0.28.3
[0.28.2]: https://github.com/giantswarm/sloth-rules/compare/v0.28.1...v0.28.2
[0.28.1]: https://github.com/giantswarm/sloth-rules/compare/v0.28.0...v0.28.1
[0.28.0]: https://github.com/giantswarm/sloth-rules/compare/v0.27.3...v0.28.0
[0.27.3]: https://github.com/giantswarm/sloth-rules/compare/v0.27.2...v0.27.3
[0.27.2]: https://github.com/giantswarm/sloth-rules/compare/v0.27.1...v0.27.2
[0.27.1]: https://github.com/giantswarm/sloth-rules/compare/v0.27.0...v0.27.1
[0.27.0]: https://github.com/giantswarm/sloth-rules/compare/v0.26.0...v0.27.0
[0.26.0]: https://github.com/giantswarm/sloth-rules/compare/v0.25.0...v0.26.0
[0.25.0]: https://github.com/giantswarm/sloth-rules/compare/v0.24.0...v0.25.0
[0.24.0]: https://github.com/giantswarm/sloth-rules/compare/v0.23.0...v0.24.0
[0.23.0]: https://github.com/giantswarm/sloth-rules/compare/v0.22.0...v0.23.0
[0.22.0]: https://github.com/giantswarm/sloth-rules/compare/v0.21.0...v0.22.0
[0.21.0]: https://github.com/giantswarm/sloth-rules/compare/v0.20.0...v0.21.0
[0.20.0]: https://github.com/giantswarm/sloth-rules/compare/v0.19.1...v0.20.0
[0.19.1]: https://github.com/giantswarm/sloth-rules/compare/v0.19.0...v0.19.1
[0.19.0]: https://github.com/giantswarm/sloth-rules/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/giantswarm/sloth-rules/compare/v0.17.0...v0.18.0
[0.17.0]: https://github.com/giantswarm/sloth-rules/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/giantswarm/sloth-rules/compare/v0.15.0...v0.16.0
[0.15.0]: https://github.com/giantswarm/sloth-rules/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/giantswarm/sloth-rules/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/giantswarm/sloth-rules/compare/v0.12.1...v0.13.0
[0.12.1]: https://github.com/giantswarm/sloth-rules/compare/v0.12.0...v0.12.1
[0.12.0]: https://github.com/giantswarm/sloth-rules/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/giantswarm/sloth-rules/compare/v0.10.0...v0.11.0
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
