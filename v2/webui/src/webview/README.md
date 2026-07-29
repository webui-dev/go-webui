## 1. `WebView2.h`

1. Go to https://www.nuget.org/packages/microsoft.web.webview2.
2. Download the `.nupkg` file.
3. Rename it to `.zip`.
4. Extract the archive.
5. Copy the file `microsoft.web.webview2.x.x.xxxx.xx\build\native\include\WebView2.h`.
6. Paste `WebView2.h` into this folder (`webui\src\webview`).

## 2. `EventToken.h`

1. Go to https://learn.microsoft.com/en-us/windows/apps/windows-sdk/downloads.
2. Download and install the Windows SDK.
3. Copy `C:\Program Files (x86)\Windows Kits\10\Include\10.x.xxxx.x\winrt\EventToken.h`.
4. Paste `EventToken.h` into this folder (`webui\src\webview`).

## 3. Prevent warnings

Add the following snippet at the top of `WebView2.h`:

```cpp
// Disable all warnings (WebUI)
#ifdef _MSC_VER
#pragma warning(push, 0)
#pragma warning(disable: 4996)
#elif defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#endif
```
