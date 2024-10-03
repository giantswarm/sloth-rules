## K8s Scheduler

This document describes the SLO for `Kubernetes Scheduler` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The Scheduler is the component that assigns pods to nodes.
We want the assignment operation to be as fast as possible.

## SLIs and SLOs

|Category | SLI                                                                                                                                                                       |          SLO |
|---------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------:|
| Availability | The proportion of successful calls against the kubernetes internal API (code 200) against the amount of unsuccessful calls (code !=200). | 99% success  |
| Latency | The proportion of sufficiently fast pod-node assignments, as measured from scheduler pods's metrics.<br>“Sufficiently fast” is any assignment taking less than 5 minutes. | 90% success |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
