package gocode

import (
	// "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	core "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	meta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PulumiNamespace(ctx *pulumi.Context, nameSpaceName string) error {
	// _, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
	// 	File: "yaml/pulumi/*.yaml",
	// })

	_, err := core.NewNamespace(ctx, nameSpaceName, &core.NamespaceArgs{
		Metadata: &meta.ObjectMetaArgs{
			Name: pulumi.String(nameSpaceName),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
