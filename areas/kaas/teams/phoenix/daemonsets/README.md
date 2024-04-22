## Daemonsets

This document describes the SLO for `DaemonSets` on Giant Swarm's clusters.

## Service Overview

Some core platform components are scheduled as `DaemonSets` in the `giantswarm` and `kube-system` namespaces. These `DaemonSets` include `kube-proxy`, cloud controllers, CNI components and exporters.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Daemonsets |    |    |
| Availability | The percentage of daemonsets pods scheduled and available on nodes. | 99% success  |

## Clarifications and Caveats

- The alert was mutuated from the one in prometheus-rules, the SLO will need to be tested in the real environments and adjusted to a reasonable value.
