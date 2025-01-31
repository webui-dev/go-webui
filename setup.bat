@echo off

rem 1) Ensure the current directory has go.mod
if not exist "go.mod" (
    echo Error: failed to find go.mod file in current directory.
    echo        To set up the go-webui module, use this script in a module directory.
    exit /b 1
)

rem 2) Run go commands
go mod tidy
go get github.com/webui-dev/go-webui/v2@main
go get github.com/webui-dev/webui@main >NUL 2>&1

rem 3) Retrieve GOPATH (use environment variable if defined, otherwise go env)
if defined GOPATH (
    set "go_path=%GOPATH%"
) else (
    for /f "delims=" %%I in ('go env GOPATH') do (
        set "go_path=%%I"
    )
)

rem 4) Parse the first matching version lines from go.sum
rem    For go-webui/v2:
set "go_webui_full_version="
for /f "tokens=2" %%I in ('type go.sum ^| findstr /i "github.com/webui-dev/go-webui/v2"') do (
    if not defined go_webui_full_version (
        set "go_webui_full_version=%%I"
    )
)

rem For webui:
set "webui_full_version="
for /f "tokens=2" %%I in ('type go.sum ^| findstr /i "github.com/webui-dev/webui"') do (
    if not defined webui_full_version (
        set "webui_full_version=%%I"
    )
)

rem 5) Construct paths based on the parsed versions
set "go_webui_path=%go_path%\pkg\mod\github.com\webui-dev\go-webui\v2@%go_webui_full_version%"
set "webui_path=%go_path%\pkg\mod\github.com\webui-dev\webui@%webui_full_version%"

rem 6) Validate that these paths actually exist
set "has_error=false"
if not exist "%go_webui_path%\" (
    echo Failed to find go-webui in "%go_webui_path%"
    set "has_error=true"
)
if not exist "%webui_path%\" (
    echo Failed to find webui in "%webui_path%"
    set "has_error=true"
)

if "%has_error%"=="true" (
    exit /b 1
)

rem 7) If a "webui" directory or link already exists in the go_webui_path, do nothing
if exist "%go_webui_path%\webui\" (
    exit /b 0
)

rem 8) Create a directory symlink for the webui folder
rem    (Requires Administrator privileges or Developer Mode)
echo Creating symlink: %go_webui_path%\webui  ->  %webui_path%
mklink /D "%go_webui_path%\webui" "%webui_path%"
if errorlevel 1 (
    echo mklink failed. Run as Administrator if needed.
    exit /b 1
)

rem 9) Final tidy
go mod tidy

exit /b 0
