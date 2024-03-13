KUBERNETES SECRETS
==================

MISSION
-------



PULUME SECRETS
--------------

### Example

```go
package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Define Kubernetes Secret
		secret, err := v1.NewSecret(ctx, "my-secret", &v1.SecretArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Secret"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("my-secret"),
			},
			Data: pulumi.StringMap{
				"username": pulumi.String("admin"),
				"password": pulumi.String("secret"),
			},
		})
		if err != nil {
			return err
		}

		// Export the name of the Secret
		ctx.Export("secretName", secret.Metadata.Name())

		return nil
	})
}
```


See:
- [Pulumi project docs](pulumi.md)
- [Pulumi referenc](https://www.pulumi.com/registry/packages/kubernetes/api-docs/core/v1/secret/)


EXTERNAL SECRETS OPERATOR
-------------------------


See:
- [Docu](https://external-secrets.io/latest/)
- [Provider](https://external-secrets.io/latest/provider/aws-secrets-manager/)
- [Fake provider](https://external-secrets.io/latest/provider/fake/)