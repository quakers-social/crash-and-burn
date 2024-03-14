package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	core "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	meta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func PGCluster(ctx *pulumi.Context, nameSpaceName string) error {
	cfg := config.New(ctx, "")
	postgresS3BackupSecret := cfg.RequireSecret("postgres_s3_backup_secret")

	_, err := core.NewSecret(ctx, "postgres-s3-backup-secret", &core.SecretArgs{
		Metadata: &meta.ObjectMetaArgs{
			Name:      pulumi.String("postgres-s3-backup-secret"),
			Namespace: pulumi.String(nameSpaceName),
		},
		Type: pulumi.String("Opaque"),
		Data: pulumi.StringMap{
			"s3.conf": postgresS3BackupSecret,
		},
	})
	if err != nil {
		return err
	}

	// transformations := []yaml.Transformation{
	// 	// Make every service private to the cluster, i.e., turn all services into ClusterIP
	// 	// instead of LoadBalancer.
	// 	func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
	// 		if state["kind"] == "PostgresCluster" {
	// 			spec := state["spec"].(map[string]interface{})
	// 			// spec["type"] = "ClusterIP"
	// 			backups := spec["backups"].(map[string]interface{})
	// 			pgbackrest := backups["pgbackrest"].(map[string]interface{})
	// 			repos := pgbackrest["repos"].([]interface{})

	// 			repo01 := make(map[string]interface{})
	// 			repo01["name"] = "repo1"

	// 			s3 := make(map[string]interface{})
	// 			s3["bucket"] = "bucket"
	// 			s3["endpoint"] = "endpoint"
	// 			s3["region"] = "region"

	// 			repo01["s3"] = s3
	// 			repos[0] = repo01
	// 		}
	// 	},
	// }

	_, err2 := yaml.NewConfigFile(ctx, "pg-cluster", &yaml.ConfigFileArgs{
		File: "yaml/pg-cluster/*.yaml",
		// Transformations: transformations,
	})
	if err2 != nil {
		return err2
	}
	return nil

}
