package gocode

import (
	helm "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func BitnamiFluxCD(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := helm.NewChart(ctx, "flux", helm.ChartArgs{
		Chart:     pulumi.String("flux"),
		Version:   pulumi.String("1.10.1"),
		Namespace: pulumi.String(nameSpaceName),
		FetchArgs: &helm.FetchArgs{
			Repo: pulumi.String("https://charts.bitnami.com/bitnami"),
		},
	})

	if err != nil {
		return err
	}

	return nil

}
