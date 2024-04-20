## Cloud APIs availability for core components

This document describes the SLO for `Cloud APIs availability for core components` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

`kube-controller-manager` or `kube-scheduler` might make requests to the cloud provider APIs. 
We want to measure the availability of those requests, as if they start to fail they might compromise the cluster functionality.

## SLIs and SLOs

|Category | SLI                                                                      |SLO |
|---------|:-------------------------------------------------------------------------|---:|
| Availability | The proportion of successful calls against the against the amount of unsuccessful calls. | 99% success  |

## Clarifications and Caveats

- This alert has been ported over from prometheus-rules and might need some changes or fine tuning.
