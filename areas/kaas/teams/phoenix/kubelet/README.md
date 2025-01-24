## Kubelet

This document describes the SLO for `Kubelet` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

Kubelet is the primary "node agent" that runs on each node.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Kubelet    |    |    |
| PLEG Duration        | The proportion of sufficiently fast PLEG relist calls, as measured from Kubelet. | 99% success  |
| Started Containers        | The percentage of successful Container creation, measured from the Kubelet. | 99% success  |
| Stared Pods        | The percentage of successful Pod creation, measured from the Kubelet. | 99% success  |
| Storage    |    |    |
| Operations        | The percentage of successful storage operations per volume plugin, measured from Kubelet.<br>The status "fail-unknown" is considered as an error. | 99% success  |


## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
