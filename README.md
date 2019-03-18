# Kubeception

**Warning: Proof of Concept**

Kubernetes in Kubernetes! This project makes it possible to run a Kubernetes cluster in Kubernetes. It is based off
some prior experimentation I did in [kubeception_legacy](https://github.com/plombardi89/kubeception_legacy).

While this tool can be used to run a full Kubernetes cluster inside of Kubernetes the rationale for the work is more
nuanced. I wanted to make it easy to run a "hybrid" Kubernetes setup with the master control plane in an existing
Kubernetes cluster and workers provisioned out in the cloud and then joined to the master.
