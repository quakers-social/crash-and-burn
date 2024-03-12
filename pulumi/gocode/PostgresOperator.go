package gocode

import (
	// "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	helm "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func PostgresOperator(ctx *pulumi.Context, nameSpaceName string) error {

	// cfg := config.New(ctx, "")
	// s3AccessSecret := cfg.RequireSecret("s3_access_secret")
	// mastodonSecretKeyBase := cfg.RequireSecret("mastodon_secret_key_base")
	// otpSecret := cfg.RequireSecret("mastodon_otp_secret")
	// mastodonPostgresPassword := cfg.RequireSecret("mastodon_postgres_password")
	// mastodonVapidPrivateKey := cfg.RequireSecret("mastodon_vapid_private_key")
	// mastodonVapidPublicKey := cfg.Require("mastodon_vapid_public_key")
	// mastodonAdminPassword := cfg.RequireSecret("mastodon_admin_password")

	chartValuse := pulumi.Map{
		"debug": pulumi.Bool(false),
		"controllerImages": pulumi.Map{
			"cluster": pulumi.String("registry.developers.crunchydata.com/crunchydata/postgres-operator:ubi8-5.5.1-0"),
		},
	}

	// Deploy Mastodon Helm chart
	_, err := helm.NewChart(ctx, "mastodon", helm.ChartArgs{
		Path: pulumi.String("./charts/postgres-operator-examples-main/helm/install"),
		// Namespace: pulumi.String(nameSpaceName),
		Values: chartValuse,
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
