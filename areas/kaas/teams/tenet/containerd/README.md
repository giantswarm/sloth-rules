## Containerd

This document describes the SLO for `Containerd` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

Containerd is a high-level container runtime that provides a platform for managing container images, runtime, and storage.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Containerd |    |    |
| Disk Usage | We start counting errors at less than 10% free space, proportionally inverse to the percentage of remaining free space. | 99% |


## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
