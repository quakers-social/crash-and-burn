package gocode

import (
	// "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	helm "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func BitnamiMastodon(ctx *pulumi.Context, nameSpaceName string) error {

	cfg := config.New(ctx, "")
	s3AccessSecret := cfg.RequireSecret("s3_access_secret")
	mastodonSecretKeyBase := cfg.RequireSecret("mastodon_secret_key_base")
	otpSecret := cfg.RequireSecret("mastodon_otp_secret")
	mastodonPostgresPassword := cfg.RequireSecret("mastodon_postgres_password")
	mastodonVapidPrivateKey := cfg.RequireSecret("mastodon_vapid_private_key")
	mastodonVapidPublicKey := cfg.Require("mastodon_vapid_public_key")
	mastodonAdminPassword := cfg.RequireSecret("mastodon_admin_password")

	chartValuse := pulumi.Map{
		"adminUser":       pulumi.String("olaf"),
		"adminEmail":      pulumi.String("briefkasten@olaf-radicke.de"),
		"adminPassword":   mastodonAdminPassword,
		"localDomain":     pulumi.String("mastodon.the-independent-friend.de"),
		"webDomain":       pulumi.String("mastodon.the-independent-friend.de"),
		"secretKeyBase":   mastodonSecretKeyBase,
		"otpSecret":       otpSecret,
		"vapidPrivateKey": mastodonVapidPrivateKey,
		"vapidPublicKey":  pulumi.String(mastodonVapidPublicKey),
		"externalS3": pulumi.Map{
			"existingSecretAccessKeyIDKey": pulumi.String("21D554754B5AB6CF0F635"),
			"existingSecretKeySecretKey":   s3AccessSecret,
			"bucket":                       pulumi.String("mastodon001"),
			"host":                         pulumi.String("us-west-1.storage.impossibleapi.net"),
			"region":                       pulumi.String("us-west-1"),
		},
		"postgresql": pulumi.Map{
			"enabled": pulumi.Bool(false),
			"auth": pulumi.Map{
				"password": mastodonPostgresPassword,
			},
		},
	}

	// Deploy Mastodon Helm chart
	_, err := helm.NewChart(ctx, "mastodon", helm.ChartArgs{
		Path:      pulumi.String("./charts/charts-main/bitnami/mastodon"),
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
