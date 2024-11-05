#!/usr/bin/env bash

# Ensure the current directory is part of a go module.
if [ ! -f "go.mod" ]; then
	echo "error: failed to find go.mod file in current directory."
	echo "       To set up the go-webui module, use it in a module directory."
	exit 1
fi

go mod tidy
go get github.com/webui-dev/go-webui/v2@main
go get github.com/webui-dev/webui@main > /dev/null 2>&1

go_path="${GOPATH:-$(go env GOPATH)}"

go_webui_full_version=$(grep "github.com/webui-dev/go-webui/v2" go.sum | awk '{print $2}' | head -n 1)
webui_full_version=$(grep "github.com/webui-dev/webui" go.sum | awk '{print $2}' | head -n 1)

go_webui_path="$go_path/pkg/mod/github.com/webui-dev/go-webui/v2@$go_webui_full_version"
webui_path="$go_path/pkg/mod/github.com/webui-dev/webui@$webui_full_version"

# Validate paths.
if [ ! -d "$go_webui_path" ]; then
	echo "Failed to find go-webui in '$go_webui_path'"
	has_error=true
fi
if [ ! -d "$webui_path" ]; then
	echo "Failed to find webui in '$webui_path'"
	has_error=true
fi
if [ "$has_error" = true ]; then
	exit 1
fi

# Exit if link already exists in the directory of the used go-webui version.
if [ -d "$go_webui_path/webui" ]; then
	exit 0
fi

# Store original permissions.
# Not strictly necessary, yet we ensure end without changes to the original permissions.
og_perms=$(stat -c "%a" "$go_webui_path")
chmod +w "$go_webui_path"

# Linking allows using WebUI C even in cases of multiple go-webui versions without creating bloat.
ln -s "$webui_path" "$go_webui_path/webui"

# Restore original permissions.
chmod "$og_perms" "$go_webui_path"
go mod tidy
