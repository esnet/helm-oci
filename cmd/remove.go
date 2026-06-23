package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "remove <name>",
		Short:             "Remove a bookmarked OCI chart reference",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeBookmarkNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if err := getStore().Remove(name); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Removed %q\n", name)
			return nil
		},
	}
}
