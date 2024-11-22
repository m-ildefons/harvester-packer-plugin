//go:generate packer-sdc mapstructure-to-hcl2 -type ImageSource

package volume

import (
	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"

	sdkmultistep "github.com/hashicorp/packer-plugin-sdk/multistep"
	// apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ImageSource struct {
	Namespace string `mapstructure:"namespace" required:"false"`
	URL       string `mapstructure:"url" required:"false"`
	Checksum  string `mapstructure:"checksum" required:"false"`
}

func (is *ImageSource) Prepare(vpctx *VolumePreparationContext) sdkmultistep.StepAction {
	image := &harvesterv1.VirtualMachineImage{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "packer-build-img-",
		},
		Spec: harvesterv1.VirtualMachineImageSpec{
			DisplayName: "packer-build-image-foo",
			SourceType:  "download",
			URL:         is.URL,
		},
	}

	if is.Namespace == "" {
		is.Namespace = "packer-build"
	}

	_, err := vpctx.Client.
		HarvesterhciV1beta1().
		VirtualMachineImages(is.Namespace).
		Create(vpctx.Context, image, metav1.CreateOptions{})
	if err != nil {
		vpctx.HaltOnError(err)
	}

	return sdkmultistep.ActionContinue
}
