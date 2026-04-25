---
title: "Kubernetes v1.36: Admission Policies That Can't Be Deleted"
date: 2026-05-04
source: "Kubernetes.io"
link: "https://kubernetes.io/blog/2026/05/04/kubernetes-v1-36-manifest-based-admission-control/"
year: 2026
---

Kubernetes v1.36 introduces an alpha feature that addresses the challenge of enforcing security policies across cluster bootstrap: manifest-based admission control. It lets you define admission webhooks and CEL-based policies as files on disk, loaded by the API server at startup, before it serves any requests, ensuring they can't be deleted or bypassed.
