## DNS Service

This document describes the SLO for `DNS Service` on both Giant Swarm's Management Clusters and Workload Clusters. CoreDNS and Node Local DNS are used as the implementation in all Giant Swarm clusters.

## Service Overview

The DNS service is responsible for resolving internat domains (services) as well as any external domain. Failing to resolve DNS will cause issues across all cluster workloads.

## SLIs and SLOs

| Category                             | SLI                                                                                                                                                                                                                                                      |SLO |
|--------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| DNS Availability                     | The proportion of sufficiently resolved DNS queries, as measured from BlackBox Exporter pods's metrics.                                                                                                                                                  | 99% success |
| CoreDNS Latency                      | The proportion of sufficiently fast DNS resolution for `local` zones, as measured from CoreDNS pods's metrics.<br>“Sufficiently fast” is any resolution taking less than 0.032 seconds.                                                                  | 99% success  |
| Local Node DNS Cache Latency         | The proportion of sufficiently fast DNS resolution for `local` zones, as measured from NodeLocalDNSCache pods's metrics.<br>“Sufficiently fast” is any resolution taking less than 0.008 seconds.                                                        | 99% success  |
| CoreDNS Latency for external domains | The proportion of sufficiently fast DNS resolution for `non local` zones, as measured from CoreDNS pods's metrics.<br>“Sufficiently fast” is any resolution taking less than 0.128 seconds on AWS and less that 0.512 seconds everywhere else.           | 99% success  |
| Local Node DNS Cache Latency for external domains | The proportion of sufficiently fast DNS resolution for `non local` zones, as measured from NodeLocalDNSCache pods's metrics.<br>“Sufficiently fast” is any resolution taking less than 0.128 seconds on AWS and less that 0.512 seconds everywhere else. | 99% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
