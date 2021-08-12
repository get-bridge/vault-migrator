package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "vault-migrate",
	Short: "Facilitates migrating stuff in and out of Vault.",
}

func Run() error {
	return rootCmd.Execute()
}
