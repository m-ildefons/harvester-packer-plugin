//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package harvester

import (
	"fmt"

	sdkcommon "github.com/hashicorp/packer-plugin-sdk/common"
	sdkcommunicator "github.com/hashicorp/packer-plugin-sdk/communicator"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
	sdktemplateconfig "github.com/hashicorp/packer-plugin-sdk/template/config"
	sdktemplateinterpolate "github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"github.com/harvester/packer-plugin-harvester/builder/harvester/namespace"
	"github.com/harvester/packer-plugin-harvester/builder/harvester/network"
	"github.com/harvester/packer-plugin-harvester/builder/harvester/volume"
)

const (
	DefaultCPUCount   = 1
	DefaultMemorySize = "512Mi"
)

type Config struct {
	sdkcommon.PackerConfig `mapstructure:",squash"`

	// Path to a kubeconfig file to access the Harvester API
	HarvesterKubeconfig string `mapstructure:"kubeconfig" required:"true"`

	// Namespace to do the build in
	Namespace namespace.Namespace `mapstructure:"namespace" required:"false"`

	Communicator sdkcommunicator.Config `mapstructure:"communicator"`

	// VM Configuration
	// These properties can be used to set up specific hardware configuration of
	// the VM for the installation. E.g. attach additional disks with to supply
	// cloud-init data, answer files or drivers.
	//
	// Number of CPUs for VM during install
	CPUCount int `mapstructure:"cpu" required:"false"`
	// Amount of memory during install
	MemorySize string `mapstructure:"memory" required:"false"`
	// Volumes
	Volumes []volume.Volume `mapstructure:"volume" required:"false"`
	// Network interfaces
	NetworkInterfaces []network.NetworkInterface `mapstructure:"network_interface" required:"false"`

	ctx sdktemplateinterpolate.Context
}

func (c *Config) Prepare(meta ...interface{}) ([]string, error) {
	err := sdktemplateconfig.Decode(c,
		&sdktemplateconfig.DecodeOpts{
			PluginType:         "harvester",
			Interpolate:        true,
			InterpolateContext: &c.ctx,
			InterpolateFilter:  &sdktemplateinterpolate.RenderFilter{},
		},
		meta...,
	)
	if err != nil {
		return nil, err
	}

	errors := &sdkpacker.MultiError{}
	warnings := make([]string, 0)

	errors = sdkpacker.MultiErrorAppend(errors, c.Communicator.Prepare(&c.ctx)...)

	if c.CPUCount < 1 {
		c.CPUCount = 1
		warnings = append(warnings,
			fmt.Sprintf("no CPU count given, using defalt of %v", DefaultCPUCount),
		)
	}

	if c.MemorySize == "" {
		c.MemorySize = "512Mi"
		warnings = append(warnings,
			fmt.Sprintf("no memory size given, using default of %v", DefaultMemorySize),
		)
	}

	if c.Namespace == "" {
		c.Namespace = "packer-build"
		warnings = append(warnings,
			fmt.Sprintf("no namespace given, using default: \"packer-build\""),
		)
	}

	if len(errors.Errors) > 0 {
		return warnings, errors
	}

	return warnings, nil
}
