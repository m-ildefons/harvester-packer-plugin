package namespace

import (
	"context"

	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"

	"k8s.io/client-go/kubernetes"
)

type NamespacePreparationContext struct {
	Client  *kubernetes.Clientset
	Ui      *sdkpacker.Ui
	State   *sdkmultistep.StateBag
	Context context.Context
}

func NewNamespacePreparationContext(ctx context.Context, client *kubernetes.Clientset, ui *sdkpacker.Ui, state *sdkmultistep.StateBag) *NamespacePreparationContext {
	return &NamespacePreparationContext{
		Client:  client,
		Ui:      ui,
		State:   state,
		Context: ctx,
	}
}

func (npctx *NamespacePreparationContext) HaltOnError(err error) sdkmultistep.StepAction {
	(*npctx.State).Put("error", err)
	(*npctx.Ui).Error(err.Error())
	return sdkmultistep.ActionHalt
}
