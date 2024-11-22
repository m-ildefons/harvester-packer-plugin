//go:generate packer-sdc mapstructure-to-hcl2 -type Namespace

package namespace

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace string

func (n *Namespace) Prepare(npctx *NamespacePreparationContext) {
	client := npctx.Client

	namespace, err := client.CoreV1().Namespaces().Get(context.TODO(), string(*n), metav1.GetOptions{})
	if err == nil {
		return
	} else if err != nil && !apierrors.IsNotFound(err) {
		npctx.HaltOnError(err)
	}

	namespace = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: string(*n),
		},
	}
	_, err = client.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		npctx.HaltOnError(err)
	}
}
