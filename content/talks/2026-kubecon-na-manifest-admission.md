---
title: "From Disk to Defense: Kubernetes Admission Policies the API Can't Touch"
date: 2026-11-10
venue: "KubeCon NA 2026"
recording: "https://youtu.be/gN7vWn5YRRg"
slides: "https://sessionize.com/download/cbalnewla~hzLjagxSNqCR29bv1NB1Qt.pdf~kubecon-na-2026-5793-slides-v10.pdf"
image: "https://img.youtube.com/vi/gN7vWn5YRRg/maxresdefault.jpg"
year: 2026
tags: ["kubernetes", "sig-auth", "security", "admission-control", "kubecon"]
---

Admission webhooks and ValidatingAdmissionPolicies can be deleted through the very API they protect, don't exist during bootstrap, and disappear if etcd goes down. With Benjamin Elder, this session shows Manifest-Based Admission Control ([KEP-5793](https://k8s.dev/resources/keps/5793)): webhooks and policies loaded from files on disk, active before the API server serves its first request, invisible to the REST API, and able to protect admission resources themselves. It includes a live demo of an attacker deleting critical policies and being stopped, and patterns for shipping tamper-proof configs across a fleet. ([Session page](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/schedule/?id=1245768))
