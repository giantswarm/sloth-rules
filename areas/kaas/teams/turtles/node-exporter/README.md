## Node Exporter

This document describes the SLO for the `Node Exporter` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The node exporter allows to monitor node metrics. 

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| Node Exporter |    |    |
| Conntrack Entries | The percentage of used connections compared to maximum number of connections on a node | 85% success  |


## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
