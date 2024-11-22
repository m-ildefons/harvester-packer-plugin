//go:generate packer-sdc mapstructure-to-hcl2 -type Volume,VolumeSource

package volume

type VolumeSource struct {
	Type      string          `mapstructure:"type" required:"true"`
	CloudInit CloudInitSource `mapstructure:",squash"`
	Image     ImageSource     `mapstructure:",squash"`
}

type Volume struct {
	Source *VolumeSource `mapstructure:"source" required:"false"`
	Bus    string        `mapstructure:"bus" required:"false"`
}

func (v *Volume) Prepare(vpctx *VolumePreparationContext) {
	switch v.Source.Type {
	case "cloud-init":
		v.Source.CloudInit.Prepare(vpctx)
	case "download":
		v.Source.Image.Prepare(vpctx)
	default:
	}
}
