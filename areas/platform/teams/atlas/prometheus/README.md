## Prometheus

This document describes the SLO for `Prometheus` on both Giant Swarm's Management Clusters.

## Service Overview

Prometheus is the component that monitors all our resources.
When Prometheus is unavailable, we are not alerted

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Availability | The availability of the prometheus container over time. | 99% success  |