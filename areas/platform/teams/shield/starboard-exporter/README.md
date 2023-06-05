## K8s API server

This document describes the SLO for `Starboard exporter` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

Starboard exporter generates metrics out of different vulnerability reports.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| StarboardExporterErrorRate        | The error rate per cycle of reconciliation of the underlying controller | 99% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
