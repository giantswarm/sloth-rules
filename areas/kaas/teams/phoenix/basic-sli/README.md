## Basic SLIs

This document describes the SLO for `Basic SLIs` on Giant Swarm's clusters.

## Service Overview

Some platform components (e.g., kyverno, cert-manager) are deployed as StatefulSets, Deployments or Daemonsets, with a label `giantswarm.io/monitoring_basic_sli: "true"`. We monitor the availability of those components.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Availability Daemonsets | The percentage of daemonsets pods scheduled and available on nodes. | 99% success  |
| Availability Deployments | The percentage of deployments pods scheduled and available on nodes. | 99% success  |
| Availability Statefulsets | The percentage of statefulsets pods scheduled and available on nodes. | 99% success  |

## Clarifications and Caveats

- The alert was mutuated from the one in prometheus-rules, the SLO will need to be tested in the real environments and adjusted to a reasonable value.
