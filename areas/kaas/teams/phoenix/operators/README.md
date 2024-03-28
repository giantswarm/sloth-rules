## Operators

This document describes the SLO for `Operators` on Giant Swarm's Management Clusters.

## Service Overview

Operators which run on Giant Swarm's Management Clusters ensure workload cluster creation, upgrade and deletion.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Management Cluster Operators |    |    |
| aws-operator                | The percentage of successful reconciled operations by app version. | 99% success  |
| irsa-operator               | The percentage of successful reconciled operations.   | 99% success  |
| cluster-operator | The percentage of successful storage operations by pod (app version) | 99% success  |


## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
