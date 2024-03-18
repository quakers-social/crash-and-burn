package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"quakers-social/gocode"
)

var CrashAndBurn = func(ctx *pulumi.Context) error {
	nameSpaceName := "quakers-social"

	err := gocode.PulumiNamespace(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.LocalStorage(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.DebugContainer(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.PostgresOperator(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.PGCluster(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.BitnamiFluxCD(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	// err = gocode.BitnamiMastodon(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	pulumi.Run(CrashAndBurn)
}
