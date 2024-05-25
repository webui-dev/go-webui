# Download helper for WebUI wrapper users to simplify the setup with the latest
# WebUI-C versions - Go Prototype.
#
# Source: https://github.com/webui-dev/go-webui
# License: MIT
#
# Currently the downloader works for tagged release versions.
# Usage via web: `sh -c "$(curl -fsSL https://raw.githubusercontent.com/webui-dev/go-webui/main/setup.sh)"`
# Local execution e.g., `sh $GOPATH/pkg/mod/github.com/webui-dev/go-webui/v2@v2.4.0/setup.sh` would require
# less logic but the idea is to eventually dynamically determine the latest version to also support versions
# like `@latest` or commit SHAs.

# Store the current location to restore it at the end of the script.
$current_location = Get-Location

$module = "github.com/webui-dev/go-webui/v2"
$webui_version="v2.4.1" # TODO: fetch latest version automatically and allow to set version via flag
$release_base_url = "https://github.com/webui-dev/webui/releases/"

# Determine the release archive for the used platform and architecture.
# For this Windows script this is currently only x64.
$platform = [System.Environment]::OSVersion.Platform
$architecture = [System.Environment]::Is64BitOperatingSystem
switch -wildcard ($platform)
{
	"Win32NT"
	{
		switch -wildcard ($architecture)
		{
			"True"
			{
				$archive = "webui-windows-gcc-x64.zip"
			}
			default
			{
				Write-Host "The setup script currently does not support $arch architectures on Windows."
				exit 1
			}
		}
	}
	default
	{
		Write-Host "The setup script currently does not support $platform."
		exit 1
	}
}

# Parse CLI arguments.
# Defaults
$output = "webui"
for ($i = 0; $i -lt $args.Length; $i++)
{
	switch -wildcard ($args[$i])
	{
		'--output'
		{
			$output = $args[$i + 1]
			$i++
			break
		}
		'--nightly'
		{
			$nightly = $true
			break
		}
		'--local'
		{
			$local = $true
			break
		}
		'--help'
		{
			Write-Host "Usage: setup.ps1 [flags]"
			Write-Host ""
			Write-Host "Flags:"
			Write-Host "  -o, --output: Specify the output directory"
			Write-Host "  --nightly: Download the latest nightly release"
			Write-Host "  --local: Save the output into the current directory"
			Write-Host "  -h, --help: Display this help message"
			exit 0
		}
		default
		{
			Write-Host "Unknown option: $($args[$i])"
			exit 1
		}
	}
}

if ($local -eq $true)
{
	Set-Location v2
	# TODO: add path verification for local setup
} else
{
	# Verify GOPATH.
	if ([string]::IsNullOrEmpty($Env:GOPATH))
	{
		Write-Host "Warning: GOPATH is not set."
		$go_path = "$Env:USERPROFILE\go"
		Write-Host "Trying to use $go_path instead."
	} else
	{
		$go_path = $Env:GOPATH
	}

	# Verify that module package is installed.
	$module_path = Join-Path $go_path "pkg\mod\$module@$webui_version"
	if (-not (Test-Path $module_path -PathType Container))
	{
		Write-Host "Error: '$module_path' does not exist in GOPATH."
		Write-Host "Make sure to run 'go get $module@$webui_version' first."
		exit 1
	}

	Set-Location $module_path
}

$archive_dir = $archive.Replace(".zip", "")

# Clean old library files in case they exist.
Remove-Item -Path $archive -ErrorAction SilentlyContinue
Remove-Item -Path $archive_dir -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item -Path $output -Recurse -Force -ErrorAction SilentlyContinue

# Download and extract the archive.
Write-Host "Downloading..."
if ($nightly -eq $true)
{
	$version="nightly"
	$url = "${release_base_url}download/nightly/$archive"
} else
{
	$version=$webui_version
	$url = "${release_base_url}latest/download/$archive"
}
Invoke-WebRequest -Uri $url -OutFile $archive

# Move the extracted files to the output directory.
Write-Host "Extracting..."
Expand-Archive -LiteralPath $archive
Move-Item -Path $archive_dir\$archive_dir -Destination $output

# Clean downloaded files and residues.
Remove-Item -Path $archive -Force
Remove-Item -Path $archive_dir -Recurse -Force

Write-Host "Done."
Set-Location $current_location
