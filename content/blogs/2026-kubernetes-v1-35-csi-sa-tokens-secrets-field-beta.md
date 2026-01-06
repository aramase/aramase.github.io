---
title: "Kubernetes v1.35: A Better Way to Pass Service Account Tokens to CSI Drivers"
date: 2026-01-07
source: "Kubernetes.io"
link: "https://kubernetes.io/blog/2026/01/07/kubernetes-v1-35-csi-sa-tokens-secrets-field-beta/"
year: 2026
---

Kubernetes v1.35 introduces a beta solution to address service account token handling in CSI drivers: CSI Driver Opt-in for Service Account Tokens via Secrets Field. This allows CSI drivers to receive service account tokens through the `secrets` field in `NodePublishVolumeRequest`, which is the appropriate place for sensitive data in the CSI specification.
