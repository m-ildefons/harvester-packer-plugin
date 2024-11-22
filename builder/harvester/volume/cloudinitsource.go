//go:generate packer-sdc mapstructure-to-hcl2 -type CloudInitSource

package volume

import (
	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
)

type CloudInitSource struct {
	MetaData    *string `mapstructure:"meta_data" required:"false"`
	UserData    *string `mapstructure:"user_data" required:"false"`
	NetworkData *string `mapstructure:"network_data" required:"false"`
}

func (cis *CloudInitSource) Prepare(vpctx *VolumePreparationContext) sdkmultistep.StepAction {
	return sdkmultistep.ActionContinue
}
