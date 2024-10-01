## K8s Controller manager

This document describes the SLO for `Kubernetes Controller Manager` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The Controller manager is the component that reconciles the most important kubernetes resources.
When the reconciliation speed is too slow, the cluster might become slow responding to changes.

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Availability | The proportion of successful calls against the kubernetes internal API (code 200) against the amount of unsuccessful calls (code !=200). | 99% success  |
| Latency | The proportion of sufficiently fast reconciliation loops, as measured from controller managers pods's metrics.<br>“Sufficiently fast” is any reconciliation taking less than 10 seconds. | 90% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
