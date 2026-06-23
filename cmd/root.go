package cmd

import (
	"os"
	"path/filepath"
	"runtime"

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
		switch runtime.GOOS {
		case "darwin":
			dataHome = filepath.Join(home, "Library", "helm")
		case "windows":
			dataHome = filepath.Join(os.Getenv("APPDATA"), "helm")
		default:
			if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
				dataHome = filepath.Join(xdg, "helm")
			} else {
				dataHome = filepath.Join(home, ".local", "share", "helm")
			}
		}
	}
	return filepath.Join(dataHome, "oci-bookmarks.yaml")
}

// completeBookmarkNames provides tab-completion for the <name> argument
// on any command that operates on an existing bookmark.
func completeBookmarkNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	items, err := getStore().List()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	var names []string
	for _, b := range items {
		names = append(names, b.Name)
	}
	return names, cobra.ShellCompDirectiveNoFileComp
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
