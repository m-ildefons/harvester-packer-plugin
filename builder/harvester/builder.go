package harvester

import (
	"context"

	hcl "github.com/hashicorp/hcl/v2/hcldec"
	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	sdkcommonsteps "github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
)

const BuilderId = "harvester.harvester"

type Builder struct {
	config Config
	runner sdkmultistep.Runner
}

func (b *Builder) ConfigSpec() hcl.ObjectSpec {
	return b.config.FlatMapstructure().HCL2Spec()
}

func (b *Builder) Prepare(meta ...interface{}) ([]string, []string, error) {
	warnings, errors := b.config.Prepare(meta...)
	return nil, warnings, errors
}

func (b *Builder) Run(ctx context.Context, ui sdkpacker.Ui, hook sdkpacker.Hook) (sdkpacker.Artifact, error) {
	var (
		artifact *Artifact = nil
		steps    []sdkmultistep.Step
		state    *sdkmultistep.BasicStateBag
	)

	state = new(sdkmultistep.BasicStateBag)
	state.Put("config", &b.config)
	state.Put("debug", b.config.PackerDebug)
	state.Put("hook", hook)
	state.Put("ui", ui)

	steps = append(steps,
		&stepPrepareNamespace{},
		&stepPrepareVolumes{},
	)

	b.runner = sdkcommonsteps.NewRunnerWithPauseFn(steps, b.config.PackerConfig, ui, state)
	b.runner.Run(ctx, state)

	return artifact, nil
}
