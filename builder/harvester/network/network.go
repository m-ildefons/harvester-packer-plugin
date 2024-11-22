//go:generate packer-sdc mapstructure-to-hcl2 -type NetworkInterface

package network

type NetworkInterface struct {
	Model string `mapstructure:"model" required:"false"`
}
