## Vertical-pod-autoscaler

This document describes the SLO for the `Vertical-pod-autoscaler` (VPA) on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

the VPA allows to monitor targeted objects' resources and adjust (raising or lowering) those if they're not adapted anymore to the object's needs.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| VPA    |    |    |
| Component Availability        | The percentage of availability for the VPA components | 95% success  |


## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
