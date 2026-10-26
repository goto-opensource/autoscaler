# GoTo fork: OCI patches and internal images

This fork keeps version-specific OCI recovery backports for GoTo. Upstream Kubernetes documentation follows below.

**For GoTo image builds and releases, use the [k8s-image-build guide](https://github.com/goto-dois/k8s-image-build/blob/main/README.md).** That internal repository owns the Jenkins recipe, source pins, build parameters and Artifactory publication. The 1.35 recipe and expanded guide are being added in [image-build PR #5](https://github.com/goto-dois/k8s-image-build/pull/5).

- Patched source lives on version-specific branches such as [`cluster-autoscaler-goto-1.35.2-pr-9703`](https://github.com/goto-opensource/autoscaler/tree/cluster-autoscaler-goto-1.35.2-pr-9703), based on an upstream release plus the selected backport.
- This fork's Actions run source checks and include Helm-chart release workflows. They do not publish GoTo's internal container images.
- Internal image builds are opt-in. Selected PR builds publish development images; the documented main-branch release flow publishes release images. Opening or merging a source PR here is not the publication trigger.
- Review a backport against its release baseline. Do not merge a versioned branch into this fork's `master` merely to obtain an internal image. Publishing an image does not deploy it.

Maintain the backport for 1.35 patch updates, repinning and retesting each new source version through `k8s-image-build`. The fix is merged into upstream's 1.36 release branch in [#10109](https://github.com/kubernetes/autoscaler/pull/10109); retire our backport when adopting a qualified 1.36 release containing it.

---

# Kubernetes Autoscaler

[![Release Charts](https://github.com/kubernetes/autoscaler/actions/workflows/chart-release.yaml/badge.svg)](https://github.com/kubernetes/autoscaler/actions/workflows/chart-release.yaml) [![CA Tests](https://github.com/kubernetes/autoscaler/actions/workflows/ca-test.yaml/badge.svg)](https://github.com/kubernetes/autoscaler/actions/workflows/ca-test.yaml) [![VPA Tests](https://github.com/kubernetes/autoscaler/actions/workflows/vpa-test.yaml/badge.svg)](https://github.com/kubernetes/autoscaler/actions/workflows/vpa-test.yaml) [![GoDoc Widget]][GoDoc]

This repository contains autoscaling-related components for Kubernetes.

## What's inside

[Cluster Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler) - a component that automatically adjusts the size of a Kubernetes
Cluster so that all pods have a place to run and there are no unneeded nodes. Supports several public cloud providers. Version 1.0 (GA) was released with kubernetes 1.8.

[Cluster Autoscaler Helm Chart](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/charts) - Supported Helm chart for Cluster Autoscaler.

[Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) - a set of components that automatically adjust the
amount of CPU and memory requested by pods running in the Kubernetes Cluster. Current state - beta.

[Vertical Pod Autoscaler Helm Chart](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler/charts) - Supported Helm chart for Vertical Pod Autoscaler.

[Addon Resizer](https://github.com/kubernetes/autoscaler/tree/master/addon-resizer) - a simplified version of vertical pod autoscaler that modifies
resource requests of a deployment based on the number of nodes in the Kubernetes Cluster. Current state - beta.

## Contact Info

Interested in autoscaling? Want to talk? Have questions, concerns or great ideas?

Please join us on #sig-autoscaling at https://kubernetes.slack.com/, or join one
of our weekly meetings.  See [the Kubernetes Community Repo](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md) for more information.

## Getting the Code

Fork the repository in the cloud:
1. Visit https://github.com/kubernetes/autoscaler
1. Click Fork button (top right) to establish a cloud-based fork.

The code must be checked out as a subdirectory of `k8s.io`, and not `github.com`.

```shell
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
# Replace "$YOUR_GITHUB_USERNAME" below with your github username
git clone https://github.com/$YOUR_GITHUB_USERNAME/autoscaler.git
cd autoscaler
```

Please refer to Kubernetes [Github workflow guide] for more details.

[GoDoc]: https://godoc.org/k8s.io/autoscaler
[GoDoc Widget]: https://godoc.org/k8s.io/autoscaler?status.svg
[Github workflow guide]: https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md
