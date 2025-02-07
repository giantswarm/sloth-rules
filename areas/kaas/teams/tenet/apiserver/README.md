## K8s API server

This document describes the SLO for `Kubernetes API server` on both Giant Swarm's Management Clusters and Workload Clusters.

## Service Overview

The API server service is the entry point for all users (as in humans) and all services (as in software) running inside or outside the cluster.

## SLIs and SLOs

|Category |SLI |SLO |
|---------|:---|---:|
| API    |    |    |
| Availability        | The percentage of successful calls to the API, measured from the API server pods's metrics.<br>Any HTTP status < 500 and != 429 is considered successful. If the API server pods are down (and thus no request metrics are available), the error rate is set to 100%. | 99% success  |
| Availability        | The availability-instances alert is based on the `up{app="kubernetes"}` metric and is necessary to alert when the API server is down: the `apiserver_request_total` metric is exposed by the API server itself and won't be available while it is down, and thus the error rate computed by the other availability alert won't be precise (due to aggregation functions over time masking the gap in availability).  | 95% success  |
| Latency        | The proportion of sufficiently fast requests, as measured from API server pods's metrics.<br>“Sufficiently fast” is different based on the resource type and the operation involved, as follows:<br><ul><li>Latency of mutating API calls for single objects for every (resource, verb) pair (excluding virtual and aggregated resources and Custom Resource Definitions), measured as 99th percentile over last 5 minutes: 1s</li><li>Latency of non-streaming read-only API calls for every (resource, scope) pair (excluding virtual and aggregated resources and Custom Resource Definitions), measured as 99th percentile over last 5 minutes: <= 1.25s if `scope=resource`, <= 30s otherwise (if `scope=namespace` or `scope=cluster`)</li></ul><br>See [upstream SLO definition](https://github.com/kubernetes/community/blob/master/sig-scalability/slos/api_call_latency.md) for more details.   | 99% success  |

## Clarifications and Caveats

- The SLO was set arbitrarily for now, will need to be tested in the real environments and adjusted to a reasonable value.
