# kubectl-nodepool
kubectl plugin to inspect group of nodes

The plugin uses common node labels and taints to group nodes into node pools.

## Installation

```
make install
```

## Usage

```
> kubectl nodepool
NAME         ARCH   OS     TYPE              NODES
nodepool1    amd64  linux  Standard_D4s_v4       1
nodepool2    amd64  linux  Standard_D2as_v4      4
```