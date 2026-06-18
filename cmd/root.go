package cmd

import (
	"os"
	"path/filepath"

	"github.com/esnet/helm-oci/pkg/bookmark"
	"github.com/spf13/cobra"
)

var bookmarkStorePath string

func setBookmarkPath(path string) {
	bookmarkStorePath = path
}

func getStore() *bookmark.Store {
	return bookmark.NewStore(bookmarkStorePath)
}

func defaultBookmarkPath() string {
	dataHome := os.Getenv("HELM_DATA_HOME")
	if dataHome == "" {
		home, _ := os.UserHomeDir()
		dataHome = filepath.Join(home, ".local", "share", "helm")
	}
	return filepath.Join(dataHome, "oci-bookmarks.yaml")
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oci",
		Short: "Bookmark and manage OCI Helm chart references",
		Long: `Manage local bookmarks for OCI-based Helm charts.

Add OCI chart URLs once, then reference them by name for install,
upgrade, pull, show, values, template, and version listing.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if bookmarkStorePath == "" {
				bookmarkStorePath = defaultBookmarkPath()
			}
		},
	}

	cmd.AddCommand(
		newAddCmd(),
		newRemoveCmd(),
		newListCmd(),
		newVersionsCmd(),
		newValuesCmd(),
		newShowCmd(),
		newInstallCmd(),
		newUpgradeCmd(),
		newPullCmd(),
		newTemplateCmd(),
	)

	return cmd
}
