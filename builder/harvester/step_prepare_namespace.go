package harvester

import (
	"context"

	namespace "github.com/harvester/packer-plugin-harvester/builder/harvester/namespace"
	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type stepPrepareNamespace struct {
}

func (s *stepPrepareNamespace) Run(ctx context.Context, state sdkmultistep.StateBag) sdkmultistep.StepAction {
	ui := state.Get("ui").(sdkpacker.Ui)
	config := state.Get("config").(*Config)

	ui.Say("Loading Kubeconfig from file")
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", config.HarvesterKubeconfig)
	if err != nil {
		return sdkmultistep.ActionHalt
	}

	ui.Say("Creating Harvester Client")
	client, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return sdkmultistep.ActionHalt
	}

	ns := config.Namespace
	npctx := namespace.NewNamespacePreparationContext(ctx, client, &ui, &state)
	ns.Prepare(npctx)

	return sdkmultistep.ActionContinue
}

func (s *stepPrepareNamespace) Cleanup(state sdkmultistep.StateBag) {
}
