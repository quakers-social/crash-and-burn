package main

import (
	// appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	// corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	// metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// "quakers-social/gocode"
)

var CrashAndBurn = func(ctx *pulumi.Context) error {
	// nameSpaceName := "pulumi-apps"

	// err := gocode.PulumiNamespace(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = gocode.LocalStorage(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = gocode.DebugContainer(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	pulumi.Run(CrashAndBurn)
}
