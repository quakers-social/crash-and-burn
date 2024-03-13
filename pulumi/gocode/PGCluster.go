package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func PGCluster(ctx *pulumi.Context, nameSpaceName string) error {
	s3AccessSecret := cfg.RequireSecret("s3_access_secret")

	transformations := []yaml.Transformation{
		// Make every service private to the cluster, i.e., turn all services into ClusterIP
		// instead of LoadBalancer.
		func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
			if state["kind"] == "PostgresCluster" {
				spec := state["spec"].(map[string]interface{})
				// spec["type"] = "ClusterIP"
				backups := spec["backups"].(map[string]interface{})
				pgbackrest := backups["pgbackrest"].(map[string]interface{})
				repos := pgbackrest["repos"].([]interface{})

				repo01 := make(map[string]interface{})
				repo01["name"] = "repo1"

				s3 := make(map[string]interface{})
				s3["bucket"] = "bucket"
				s3["endpoint"] = "endpoint"
				s3["region"] = "region"

				repo01["s3"] = s3
				repos[0] = repo01
			}
		},
	}

	_, err := yaml.NewConfigFile(ctx, "pg-cluster", &yaml.ConfigFileArgs{
		File:            "yaml/pg-cluster/*.yaml",
		Transformations: transformations,
	})
	if err != nil {
		return err
	}
	return nil

}
