package volume

import (
	"context"

	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"

	harvclient "github.com/harvester/harvester/pkg/generated/clientset/versioned"
)

type VolumePreparationContext struct {
	Client  *harvclient.Clientset
	Ui      *sdkpacker.Ui
	State   *sdkmultistep.StateBag
	Context context.Context
}

func NewVolumePreparationContext(ctx context.Context, client *harvclient.Clientset, ui *sdkpacker.Ui, state *sdkmultistep.StateBag) *VolumePreparationContext {
	return &VolumePreparationContext{
		Client:  client,
		Ui:      ui,
		State:   state,
		Context: ctx,
	}
}

func (vpctx *VolumePreparationContext) HaltOnError(err error) sdkmultistep.StepAction {
	(*vpctx.State).Put("error", err)
	(*vpctx.Ui).Error(err.Error())
	return sdkmultistep.ActionHalt
}
