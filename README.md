# Nuc Node LED Controller

A few models of Intel NUCs have a dope [software-controllable LED](https://github.com/milesp20/intel_nuc_led) on their front panel. This is a Kubernetes controller that makes that LED reflect the status of the Kubernetes node running on that NUC.

![image](https://user-images.githubusercontent.com/47/43054490-724eea12-8df7-11e8-955c-d074b4817f93.png)

## Default LED Settings

|Node Status   |LED Ring       |
|--------------|---------------|
|Ready         |Solid green    |
|Unschedulable |Pulsing yellow |
|Not Ready     |Pulsing red    |
|Error         |Blinking red   |

These can be be tweaked to your liking by editing [the `nuc-node-led-controller` configmap](./config/kubernetes/default/configmaps/nuc-node-led-controller.yaml).

## Installation

The Kubernetes manifests in [`config/kubernetes/default`](./config/kubernetes/default) will create a [ClusterRole](./config/kubernetes/default/clusterroles/nuc-node-led-controller.yaml) that allows this controller to get/list/watch Nodes and a [DaemonSet](./config/kubernetes/default/daemonsets/nuc-node-led-controller.yaml) that runs a copy of [`main.go`](./main.go) on each node.

* Install https://github.com/milesp20/intel_nuc_led on each of your NUCs
* Install [Skaffold](https://github.com/GoogleContainerTools/skaffold)

```
go get github.com/urcomputeringpal/nuc-node-led-controller
cd $GOPATH/github.com/urcomputeringpal/nuc-node-led-controller
skaffold deploy
```
