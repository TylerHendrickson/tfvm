package tfvm

import (
	"fmt"
	"os"
)

func RunTerraformCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			fmt.Printf("No terraform version configured. Place .tfvmrc or .terraform-version in current or parent dir.\n")
			os.Exit(1)
		}

		return err
	}

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(config.version)
	if err != nil {
		return nil
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if !installed {
		err = inventory.InstallTerraform(tfRelease)
		if err != nil {
			return nil
		}
	}

	tf, err := inventory.GetTerraform(tfRelease)
	if err != nil {
		return err
	}

	exitCode, err := tf.Run(args...)
	if err != nil {
		return fmt.Errorf("running tf failed: %v exitCode=%d", err, exitCode)
	}

	os.Exit(exitCode)

	return nil
}
