## Admission Webhooks

This document describes the SLO for `Admission Webhooks` on Giant Swarm's Management Clusters.

## Service Overview

The admission webhooks are responsible for mutating and validating custom ressources.

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Latency | The proportion of sufficiently fast calls, as measured from the Kubernetes API server metrics.<br>“Sufficiently fast” is any reconciliation taking less than 0.5 seconds. | 99% success  |
