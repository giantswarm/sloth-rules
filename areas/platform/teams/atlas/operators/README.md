# SLO Documentation for Atlas operators

This document describes the SLO for `Atlas operators` on Giant Swarm's Management Clusters.

## Service Overview

Operators are developped using kube-builder so they expose similar metrics.
When they cannot reconcile over a period of time, we should be alerted

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Reconciliation | The reconciliations of the operators over time. | 95% success  |
