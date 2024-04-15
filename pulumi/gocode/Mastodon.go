package gocode

import (
	// "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func Mastodon(ctx *pulumi.Context, nameSpaceName string) error {

	cfg := config.New(ctx, "")
	s3AccessSecret := cfg.RequireSecret("s3_access_secret")
	mastodonSecretKeyBase := cfg.RequireSecret("mastodon_secret_key_base")
	otpSecret := cfg.RequireSecret("mastodon_otp_secret")

	chartValuse := pulumi.Map{
		"mastodon": pulumi.Map{
			"local_domain": pulumi.String("mastodon.the-independent-friend.de"),
			"s3": pulumi.Map{
				"enabled":       pulumi.String("true"),
				"access_key":    pulumi.String("21D554754B5AB6CF0F635"),
				"access_secret": s3AccessSecret,
				"bucket":        pulumi.String("mastodon001"),
				"endpoint":      pulumi.String("us-west-1.storage.impossibleapi.net"),
				"region":        pulumi.String("us-west-1"),
			},
			"secrets": pulumi.Map{
				"secret_key_base": mastodonSecretKeyBase,
				"otp_secret":      otpSecret,
			},
		},
		"ingress": pulumi.Map{
			"enabled": pulumi.String("false"),
		},
		"postgresql": pulumi.Map{
			"enabled": pulumi.String("true"),
		},
	}

	// Deploy Mastodon Helm chart
	_, err := helmv3.NewChart(ctx, "mastodon", &helmv3.ChartArgs{
		Chart:     pulumi.String("mastodon"),
		Version:   pulumi.String("5.0.0"),
		Namespace: nameSpaceName,
		FetchArgs: &helmv3.FetchArgs{
			Repo: pulumi.String("https://charts.example.com/"),
		},
		Values: chartValuse,
	}, pulumi.Provider("k8s"))
	if err != nil {
		return err
	}

	return nil

}