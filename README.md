# kubectl-nodepool
kubectl plugin to inspect group of nodes

The plugin uses common node labels to group nodes into nodepools.

Currently supports *AKS*, *GKE* and *EKS* clusters

## Installation

```
go install github.com/zmalik/kubectl-nodepool@latest
```

## Usage

```
> kubectl nodepool
NAME         ARCH   OS     TYPE              NODES
nodepool1    amd64  linux  Standard_D4s_v4       1
nodepool2    amd64  linux  Standard_D2as_v4      4
```
