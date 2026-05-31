# helm-oci

A Helm plugin that adds local bookmarking for OCI-based Helm charts.

![helm-oci demo](assets/preview.gif)

## Table of Contents

- [Install](#install)
- [Problem](#problem)
- [Solution](#solution)
- [Usage](#usage)
  - [Managing Bookmarks](#managing-bookmarks)
  - [Inspecting Charts](#inspecting-charts)
  - [Installing and Managing Releases](#installing-and-managing-releases)
- [Bookmark Storage](#bookmark-storage)
- [Development](#development)
- [License](LICENSE.md)

## Install

```sh
helm plugin install https://github.com/esnet/helm-oci/releases/download/v1.0.1/oci-1.0.1.tgz
```

## Problem

Helm's traditional repository system lets you add a repo once and reference charts by short name:

```sh
helm repo add jetstack https://charts.jetstack.io
helm install jetstack/cert-manager cert-manager -n cert-manager --create-namespace
```

OCI-based charts have no equivalent. Every install, upgrade, or inspect requires the full OCI URL:

```sh
helm install envoy-gateway oci://docker.io/envoyproxy/gateway-helm --version 1.7.0
helm show values oci://docker.io/envoyproxy/gateway-helm --version 1.8.0
helm upgrade envoy-gateway oci://docker.io/envoyproxy/gateway-helm --version 1.8.0
```

When managing multiple OCI charts across clusters, remembering and retyping these URLs becomes impractical and cumbersome.

## Solution

`helm-oci` is a Helm plugin which solves this UX problem and lets you bookmark OCI chart URLs and reference them by name.

```sh
❯ helm oci
Manage local bookmarks for OCI-based Helm charts.

Add OCI chart URLs once, then reference them by name for install,
upgrade, pull, show, values, template, and version listing.

Usage:
  oci [command]

Available Commands:
  add         Bookmark an OCI chart reference
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install a bookmarked OCI chart
  list        List all bookmarked OCI chart references
  pull        Pull a bookmarked OCI chart
  remove      Remove a bookmarked OCI chart reference
  show        Show chart metadata for a bookmarked OCI chart
  template    Render templates for a bookmarked OCI chart
  upgrade     Upgrade a release using a bookmarked OCI chart
  values      Show values for a bookmarked OCI chart
  versions    List available versions for a bookmarked OCI chart
```

## Usage

### Managing Bookmarks

```bash
# Add a bookmark
helm oci add <name> <oci-url>
helm oci add envoy-gw oci://docker.io/envoyproxy/gateway-helm
helm oci add cert-manager oci://quay.io/jetstack/cert-manager

# List all bookmarks
helm oci list

# Remove a bookmark
helm oci remove <name>
```

### Inspecting Charts

```bash
# List available versions (tags) from the OCI registry
helm oci versions <name>

# Show chart metadata
helm oci show <name> [--version <version>]

# Show default values
helm oci values <name> [--version <version>]
```

### Installing and Managing Releases

```bash
# Install a chart
helm oci install <name> <release> [--version <version>] [helm flags...]
helm oci install envoy-gw my-gateway --version 1.7.0 --namespace envoy --create-namespace

# Upgrade a release
helm oci upgrade <name> <release> [--version <version>] [helm flags...]
helm oci upgrade envoy-gw my-gateway --version 1.8.0

# Render templates locally
helm oci template <name> <release> [--version <version>] [helm flags...]
helm oci template envoy-gw my-gateway --version 1.8.0 --set foo=bar

# Pull chart archive to local directory
helm oci pull <name> [--version <version>]
```

All flags after the bookmark name and release name are passed through directly to the underlying `helm` command. Any flag that `helm install`, `helm upgrade`, etc. accept will work — `--set`, `--values`, `--namespace`, `--create-namespace`, `--wait`, and so on.

## Bookmark Storage

Bookmarks are stored in `$HELM_DATA_HOME/oci-bookmarks.yaml`. The default location is:

- **macOS**: `~/Library/helm/oci-bookmarks.yaml`
- **Linux**: `~/.local/share/helm/oci-bookmarks.yaml`

The file is plain YAML:

```yaml
bookmarks:
  - name: envoy-gw
    url: oci://docker.io/envoyproxy/gateway-helm
  - name: cert-manager
    url: oci://quay.io/jetstack/cert-manager
```

## Development

Requires Go 1.22+.

```bash
make build      # Build binary to bin/helm-oci
make test       # Run all tests with race detector
make install    # Build and install into Helm plugins directory
make uninstall  # Remove from Helm plugins directory
make dist       # Cross-compile release archives
```
