# SLO Documentation for Teleport Service

This document outlines the Service Level Objectives for the Teleport service managed by the Shield team.

| Status | Published |
| --- | :---: |
| Author | Shield Team |
| Date | 2024-03-20 |
| Reviewers | |
| Approvers | |
| Approval date | |
| Revisit date | |

## Teleport Service

The Teleport service provides secure access for servers, databases, and Kubernetes clusters.

| Category | SLI | SLO |
| --- | :--- | :--- |
| Availability | Rate of successful authentication requests | 95% |
| Latency | Response time for authentication requests | P95 < 300ms |

Note: Alerts will be triggered if the SLOs for availability or latency are not met.
