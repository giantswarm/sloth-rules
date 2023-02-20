## Cloud Authentication

This document describes the SLO for `Container Network Interface` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The Container Network Interface is the component responsible for setting up networking on pods and enforce networking policies. Pod networking is a critical service for workloads to work properly. Cilium is used as the implementation in all Giant Swarm clusters.

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Endpoint Regeneration Latency | The proportion of sufficiently fast endpoint regeneration loops, as measured from cilium agent pods's metrics.<br>“Sufficiently fast” is any reconciliation taking less than 0.05 seconds. | 99% success  |
| Node Connectivity Latency | The proportion of sufficiently fast connectivity latency between cilium agents in other nodes, as measured from cilium agent pods's metrics.<br>“Sufficiently fast” is latency should be less than 0.001 seconds. | 99% success  |

## Clarifications and Caveats

- Endpoint regeneration: This metric measures how long it takes for an endpoint to regenerate. Cilium utilizes a portion of a cluster’s available resources to operate, and a high value could indicate that a cluster is churning more pods than Cilium can process.

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
