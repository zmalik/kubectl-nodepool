package printer

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	corev1 "k8s.io/api/core/v1"
	"os"
)

const no_name = "TBD"

type NodePool struct {
	Arch         string
	Os           string
	InstanceType string
	Name         string
	Nodes        []corev1.Node
}

func Print(nodes *corev1.NodeList) {

	nodepool := make(map[string]*NodePool)
	for _, node := range nodes.Items {
		hash := getMapToken(node)
		if nodepoolItem, ok := nodepool[hash]; !ok {
			nodepool[hash] = &NodePool{}
			nodepool[hash].Arch = node.Labels["beta.kubernetes.io/arch"]
			nodepool[hash].Os = node.Labels["kubernetes.io/os"]
			nodepool[hash].InstanceType = node.Labels["beta.kubernetes.io/instance-type"]
			nodepool[hash].Name = getNodepoolName(node.Labels)
			nodepool[hash].Nodes = append(nodepool[hash].Nodes, node)
		} else {
			nodepoolItem.Nodes = append(nodepoolItem.Nodes, node)
		}

	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Name", "Arch", "OS", "Type", "Nodes"})
	for _, nodepoolItem := range nodepool {
		name := nodepoolItem.Name
		if name == no_name {
			name = findCommonPrefix(nodepoolItem.Nodes)
		}
		t.AppendRow(table.Row{name, nodepoolItem.Arch, nodepoolItem.Os, nodepoolItem.InstanceType, len(nodepoolItem.Nodes)})
	}
	t.SetStyle(table.StyleColoredDark)
	t.Render()
}

func findCommonPrefix(nodes []corev1.Node) string {
	// find common longest prefix of node names
	var prefix string
	for _, node := range nodes {
		if prefix == "" {
			prefix = node.Name
		} else {
			prefix = commonPrefix(prefix, node.Name)
		}
	}
	return prefix
}

func commonPrefix(prefix string, name string) string {
	for i := 0; i < len(prefix); i++ {
		if prefix[i] != name[i] {
			return prefix[:i]
		}
	}
	return prefix
}

func getMapToken(node corev1.Node) string {
	arch := node.Labels["beta.kubernetes.io/arch"]
	os := node.Labels["kubernetes.io/os"]
	instanceType := node.Labels["beta.kubernetes.io/instance-type"]
	name := getNodepoolName(node.Labels)
	taints := ""
	if name == no_name {
		// use taints if name is not set
		taints = formatTaintsAsString(node.Spec.Taints)
	}

	return fmt.Sprintf("%s-%s-%s-%s-%s", arch, os, instanceType, name, taints)
}

func formatTaintsAsString(taints []corev1.Taint) string {
	var taintsString string
	for _, taint := range taints {
		taintsString += taint.Key + taint.Value + string(taint.Effect)
	}
	return taintsString
}

func getNodepoolName(labels map[string]string) string {
	//TODO: add support for other providers
	if name, ok := labels["kubernetes.azure.com/agentpool"]; ok {
		return name
	}
	if name, ok := labels["cloud.google.com/gke-nodepool"]; ok {
		return name
	}
	if name, ok := labels["eks.amazonaws.com/nodegroup"]; ok {
		return name
	}
	return no_name
}
