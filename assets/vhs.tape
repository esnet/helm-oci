Output assets/preview.gif

Require helm

Set Shell "zsh"
Set Theme "Dracula"
Set FontSize 24
Set Width 1400
Set Height 1100
Set TypingSpeed 30ms

Hide
# Uninstall any existing version so the visible install succeeds cleanly
Type "helm oci remove cert-manager 2>/dev/null; true"
Type "helm oci remove envoy-gw 2>/dev/null; true"
Type "helm plugin uninstall oci 2>/dev/null; true"
Enter
Sleep 4s
Type "export PROMPT='%F{magenta}❯%f '"
Enter
Type "clear"
Enter
Show

# 1 — Install the plugin
#
Type "helm plugin install https://github.com/esnet/helm-oci/releases/download/v1.0.1/oci-1.0.1.tgz"
Sleep 500ms
Enter
Sleep 4s

Type "clear"
Enter

# 2 — Browse available commands
#
Type "helm oci"
Sleep 500ms
Enter
Sleep 5s

Type "clear"
Enter

# 3 — Bookmark two OCI charts by name
#
Type "helm oci add envoy-gw oci://docker.io/envoyproxy/gateway-helm"
Sleep 500ms
Enter
Sleep 2s

Type "helm oci add cert-manager oci://quay.io/jetstack/cert-manager"
Sleep 500ms
Enter
Sleep 2s

# 4 — Verify bookmarks
#
Type "helm oci list"
Sleep 500ms
Enter
Sleep 3s

# 5 — Query available versions from the OCI registry
#
Type "helm oci versions envoy-gw"
Sleep 500ms
Enter
Sleep 4s

Sleep 4s
