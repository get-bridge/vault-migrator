package cmd

import (
	"encoding/json"
	"vault-migrate/lib"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import {root}",
	Short: "Imports a vault-migrate export into Vault",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := vault.NewClient(nil)
		if err != nil {
			return err
		}

		m := lib.Migrator{
			Client: c,
			Root:   args[0],
		}

		data := map[string]map[string]interface{}{}
		if err := json.NewDecoder(cmd.InOrStdin()).Decode(&data); err != nil {
			return err
		}

		return m.WriteData(data)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
