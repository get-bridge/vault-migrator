package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"vault-migrate/lib"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export {root} {prefix}",
	Short: "Exports a secret path to JSON.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := vault.NewClient(nil)
		if err != nil {
			return err
		}

		m := &lib.Migrator{
			Client: c,
			Root:   args[0],
		}

		prefix := args[1]
		if !strings.HasPrefix(prefix, "/") {
			prefix = fmt.Sprintf("/%s", prefix)
		}
		if !strings.HasSuffix(prefix, "/") {
			prefix = fmt.Sprintf("%s/", prefix)
		}

		data, err := m.ReadData(prefix)
		if err != nil {
			return err
		}

		return json.NewEncoder(cmd.OutOrStdout()).Encode(data)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
