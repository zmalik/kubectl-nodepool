package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/zmalik/kubectl-nodepool/pkg/printer"
)

func RunNodepool(cmd *cobra.Command, args []string) {
	config, err := genericclioptions.NewConfigFlags(true).ToRESTConfig()
	if err != nil {
		fmt.Print("Error: could not load kubeconfig", err)
	}
	// new corev1 client using config
	client := corev1.NewForConfigOrDie(config)
	nodes, err := client.Nodes().List(cmd.Context(), metav1.ListOptions{})
	if err != nil {
		fmt.Print("Error: could not list nodes", err)
	}
	printer.Print(nodes)
}
