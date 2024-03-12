MASTODON
========


MISSION
-------

Installation of Mastodon on Kubernets.


CODE SNIPET
-----------

```golang
package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Initialize Kubernetes Provider
		_, err := kubernetes.NewProvider(ctx, "k8s", nil)
		if err != nil {
			return err
		}

		// Deploy Mastodon Helm chart
		mastodonChart, err := helmv3.NewChart(ctx, "mastodon", &helmv3.ChartArgs{
			Chart:    pulumi.String("mastodon"),
			Version:  pulumi.String("5.0.0"),
			FetchArgs: &helmv3.FetchArgs{
				Repo: pulumi.String("https://charts.example.com/"),
			},
			Values: pulumi.Map{
				"hostname": pulumi.String("social.example.com"),
				"email":    pulumi.String("admin@example.com"),
				// ... other required values
			},
		}, pulumi.Provider("k8s"))
		if err != nil {
			return err
		}

		// Export the status of the deployment
		ctx.Export("mastodonStatus", mastodonChart.GetResourceStatus("v1/Deployment", "mastodon", ""))
		return nil
	})
}
```

CREATE VAPID KEYs
-----------------

```bash
openssl ecparam -name prime256v1 -genkey -noout -out vapid_private_key.pem
openssl ec -in vapid_private_key.pem -pubout -out vapid_public_key.pem
echo -n VAPID_PRIVATE_KEY=;cat vapid_private_key.pem | sed -e "1 d" -e "$ d" | tr -d "\n"; echo
echo -n VAPID_PUBLIC_KEY=;cat vapid_public_key.pem | sed -e "1 d" -e "$ d" | tr -d "\n"; echo

```


EXTERNAL LINKS
--------------

- [Offc. Mastodon Helm Chart](https://github.com/mastodon/chart)
- [Bitnami package for Mastodon](https://github.com/bitnami/charts/tree/main/bitnami/mastodon/)
- [Mastodon on DigitalOcean Kubernetes](https://docs.digitalocean.com/developer-center/mastodon-on-digitalocean-kubernetes/)
- [Running Mastodon server on Kubernetes](https://softwaremill.com/running-mastodon-server-on-kubernetes/)
-