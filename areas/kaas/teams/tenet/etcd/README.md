## ETCD

This document describes the SLO for `ETCD Service` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The ETCD service is the persistence layer of the Kubernetes cluster.

## SLIs and SLOs

|Category | SLI                                                                                                                                                  |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Availability | The proportion of GRPC calls that have a successfull status code ("OK") as seen by etcd pods themselves.                                             | 99% success |
| Latency | The proportion of sufficiently quickly responded requests, as seen by etcd process itself. Any request is considered fast enough if it takes < 0.5s. | 99% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
