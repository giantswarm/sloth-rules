## Cloud Authentication for Pods

This document describes the SLO for `Cloud Authentication for Pods` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The Cloud Authentication for Pods is responsible for providing credentials to Pods in order to be able to manage resources in a cloud provider.

IAM Roles for Service Accounts (IRSA) is used on AWS.

## SLIs and SLOs

|Category | SLI                                                                                                                                                                                      |SLO |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---:|
| Availability | The proportion of successful calls against the mutating webhooks (code 200) against the amount of unsuccessful calls (code !=200). | 99% success  |
| Latency | The proportion of sufficiently fast mutation calls, as measured from the webhook pods's metrics.<br>“Sufficiently fast” is any reconciliation taking less than 0.25 seconds. | 99% success  |


## Clarifications and Caveats

- Metrics are lacking in the webhook itself.

- There is very little traffic in an empty cluster and SLO can be skewed pretty fast.

- Even if a role does not exist, the webhook return a 200 http code.
