package harvester

import (
	"context"

	volume "github.com/harvester/packer-plugin-harvester/builder/harvester/volume"
	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"

	"k8s.io/client-go/tools/clientcmd"

	harvclient "github.com/harvester/harvester/pkg/generated/clientset/versioned"
)

type stepPrepareVolumes struct {
}

func (s *stepPrepareVolumes) Run(ctx context.Context, state sdkmultistep.StateBag) sdkmultistep.StepAction {
	ui := state.Get("ui").(sdkpacker.Ui)
	config := state.Get("config").(*Config)

	ui.Say("Loading Kubeconfig from file")
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", config.HarvesterKubeconfig)
	if err != nil {
		return sdkmultistep.ActionHalt
	}

	ui.Say("Creating Harvester Client")
	client, err := harvclient.NewForConfig(kubeconfig)
	if err != nil {
		return sdkmultistep.ActionHalt
	}

	ui.Say("Preparing Volumes...")

	for _, volumeConfig := range config.Volumes {
		vpctx := volume.NewVolumePreparationContext(ctx, client, &ui, &state)
		volumeConfig.Prepare(vpctx)

		if volumeConfig.Source.Type == "download" {
			ui.Say(volumeConfig.Source.Image.URL)
		}
	}

	return sdkmultistep.ActionContinue
}

func (s *stepPrepareVolumes) Cleanup(state sdkmultistep.StateBag) {
	ui := state.Get("ui").(sdkpacker.Ui)
	ui.Say("Cleaning up Volumes...")
}
