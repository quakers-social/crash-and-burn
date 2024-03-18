package gocode

import (
	helm "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PostgresOperator(ctx *pulumi.Context, nameSpaceName string) error {

	chartValuse := pulumi.Map{
		"debug": pulumi.Bool(false),
		"controllerImages": pulumi.Map{
			"cluster": pulumi.String("registry.developers.crunchydata.com/crunchydata/postgres-operator:ubi8-5.5.1-0"),
		},
	}

	// Deploy Mastodon Helm chart
	_, err := helm.NewChart(ctx, "mastodon", helm.ChartArgs{
		Path:      pulumi.String("./charts/postgres-operator-examples-main/helm/install"),
		Namespace: pulumi.String(nameSpaceName),
		Values:    chartValuse,
	})
	// _, err := helm.NewChart(ctx, "mastodon", helm.ChartArgs{
	// 	Chart:     pulumi.String("mastodon"),
	// 	Version:   pulumi.String("5.0.0"),
	// 	Namespace: pulumi.String(nameSpaceName),
	// 	FetchArgs: &helmv3.FetchArgs{
	// 		Repo: pulumi.String("https://charts.example.com/"),
	// 	},
	// 	Values: chartValuse,
	// })

	if err != nil {
		return err
	}

	return nil

}
