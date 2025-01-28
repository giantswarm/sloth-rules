## Kubelet

This document describes the SLO for `Kubelet` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

Kubelet is the primary "node agent" that runs on each node.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Availability | The percentage of available kubelet in the cluster. | 99% success  |
| Node Pool Availability | The percentage of available kubelet in a node pool. | 99% success  |
| PLEG Duration        | The proportion of sufficiently fast PLEG relist calls, as measured from Kubelet. | 99% success  |
| Storage Operations        | The percentage of successful storage operations per volume plugin, measured from Kubelet.<br>The status "fail-unknown" is considered as an error. | 99% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
