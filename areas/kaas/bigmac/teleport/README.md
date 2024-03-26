# SLO Documentation for Teleport Service

This document outlines the Service Level Objectives for the Teleport service managed by the Bigmac team.

| Status | Published |
| --- | :---: |
| Author | Bigmac Team |
| Date | 2024-03-20 |
| Reviewers | |
| Approvers | |
| Approval date | |
| Revisit date | |

## Service Overview

The Teleport service provides secure access to servers, databases, and Kubernetes clusters based on the principle of least privilege.

## SLIs and SLOs

| Category | SLI | SLO |
| --- | :--- | :--- |
| Availability | Rate of successful authentication requests | 95.0% |
| Latency | Response time for authentication requests | P95 < 300ms |

## Rationale

Maintaining high availability and low latency for the Teleport service is critical to ensure that users have consistent and reliable access.

## Error Budget

The error budget will be calculated based on the SLOs set for availability and latency.

## Clarifications and Caveats

Any deviations from the defined SLOs will trigger an investigation to identify root causes and implement corrective actions.