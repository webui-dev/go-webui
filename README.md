<div align="center">

![Logo](https://raw.githubusercontent.com/webui-dev/webui-logo/main/webui_go.png)

# Go-WebUI v2.5.1

#### [Features](#features) · [Installation](#installation) · [Usage](#usage) · [Documentation](#documentation) · [WebUI](https://github.com/webui-dev/webui)

[build-status]: https://img.shields.io/github/actions/workflow/status/webui-dev/go-webui/ci.yml?branch=main&style=for-the-badge&logo=go&labelColor=414868&logoColor=C0CAF5
[last-commit]: https://img.shields.io/github/last-commit/webui-dev/go-webui?style=for-the-badge&logo=github&logoColor=C0CAF5&labelColor=414868
[release-version]: https://img.shields.io/github/v/tag/webui-dev/go-webui?style=for-the-badge&logo=webtrees&logoColor=C0CAF5&labelColor=414868&color=7664C6
[license]: https://img.shields.io/github/license/webui-dev/go-webui?style=for-the-badge&logo=opensourcehardware&label=License&logoColor=C0CAF5&labelColor=414868&color=8c73cc

[![][build-status]](https://github.com/webui-dev/go-webui/actions?query=branch%3Amain)
[![][last-commit]](https://github.com/webui-dev/go-webui/pulse)
[![][release-version]](https://github.com/webui-dev/go-webui/releases/latest)
[![][license]](https://github.com/webui-dev/go-webui/blob/main/LICENSE)

> Use any web browser or WebView as GUI.\
> With Go in the backend and modern web technologies in the frontend.

![Screenshot](https://raw.githubusercontent.com/webui-dev/webui-logo/main/screenshot.png)

</div>

## Features

- Parent library written in pure C
- Portable (*Needs only a web browser or a WebView at runtime*)
- Lightweight (*Few Kb library*) & Small memory footprint
- Fast binary communication protocol
- Multi-platform & Multi-Browser
- Using private profile for safety
- Cross-platform WebView

## Installation

> [!NOTE]
> Until the next stable release it is recommended to use go-webui's latest development version.

- ### As Go module

  The easiest way to setup go-webui as a Go module is to use the `setup.sh` or `setup.bat` script.

  It will run `go get` to retrieve the go-webui module and bootstrap the version of the WebUI C library that it is using.

  - Windows

  ```sh
  sh -c "$(curl -fsSL https://raw.githubusercontent.com/webui-dev/go-webui/main/setup.bat)"
  ```

  - Linux / macOS

  ```sh
  sh -c "$(curl -fsSL https://raw.githubusercontent.com/webui-dev/go-webui/main/setup.sh)"
  ```

- ### As submodule

  The instructions below set up go-webui in a `modules` subdirectory of a go project.

  ```
  go-project
  ├── modules
  │   └── go-webui
  ├── ...
  └── go.mod
  ```

  Add and init the submodule

  ```sh
  git submodule add https://github.com/webui-dev/go-webui.git modules/go-webui
  ```

  ```sh
  git submodule update --init --filter=blob:none --recursive
  ```

  `replace` the path accordingly in the `g.mod` file.

  ```
  require github.com/webui-dev/go-webui/v2 v2.5.1

  replace github.com/webui-dev/go-webui/v2 v2.5.1 => ./modules/go-webui
  ```

- ### As git clone - for development and contribution purposes

  The command below retrieves go-webui as a lightweight, filtered clone.

  ```sh
  git clone --recursive --shallow-submodules --filter=blob:none --also-filter-submodules \
    https://github.com/webui-dev/go-webui.git
  ```

## Usage

### Example

```html
<!-- index.html -->
<!doctype html>
<html>
   <head>
      <script src="webui.js"></script>
   </head>
   <body>
      <button onclick="test();">Test Go-WebUI</button>
      <script>
         async function test() {
            // Call a Go function.
            const result = await myGoFunction('Hello From JavaScript');
            alert(result); // "Hello From Go"
         }
      </script>
   </body>
</html>
```

```go
// main.go
package main

import (
	"fmt"

	ui "github.com/webui-dev/go-webui/v2"
)

func myGoFunction(e ui.Event) string {
   // Get first argument
	name, _ := ui.GetArg[string](e)
	fmt.Printf("JavaScript sent: %s\n", name) // Hello From JavaScript
   // Return a response to JavaScript
	response := fmt.Sprintf("Hello From Go")
	return response
}

func main() {
	// Create a window.
	w := ui.NewWindow()
	// Bind a Go function.
	ui.Bind(w, "myGoFunction", myGoFunction)
	// Show frontend.
	w.Show("index.html")
	// Wait until all windows get closed.
	ui.Wait()
}
```

Find more examples in the [`examples/`](https://github.com/webui-dev/go-webui/tree/main/examples) directory.

## Documentation

### Enable TLS/SSL

Enable WebUI's security layer by adding the `webui_tls` build tag.

```sh
go run -tags webui_tls <path>
```

### Debugging

To use WebUI's debug build, add the `webui_log` build tag. E.g.:

```sh
go run -tags webui_log minimal.go
```

- [Online Documentation](https://webui.me/docs/#/go) (WIP)

## UI & The Web Technologies

[Borislav Stanimirov](https://ibob.bg/) discusses using HTML5 in the web browser as GUI at the [C++ Conference 2019 (_YouTube_)](https://www.youtube.com/watch?v=bbbcZd4cuxg).

<!-- <div align="center">
  <a href="https://www.youtube.com/watch?v=bbbcZd4cuxg"><img src="https://img.youtube.com/vi/bbbcZd4cuxg/0.jpg" alt="Embrace Modern Technology: Using HTML 5 for GUI in C++ - Borislav Stanimirov - CppCon 2019"></a>
</div> -->

<div align="center">

![CPPCon](https://github.com/webui-dev/webui/assets/34311583/4e830caa-4ca0-44ff-825f-7cd6d94083c8)

</div>

Web application UI design is not just about how a product looks but how it works. Using web technologies in your UI makes your product modern and professional, And a well-designed web application will help you make a solid first impression on potential customers. Great web application design also assists you in nurturing leads and increasing conversions. In addition, it makes navigating and using your web app easier for your users.

### Why Use Web Browsers?

Today's web browsers have everything a modern UI needs. Web browsers are very sophisticated and optimized. Therefore, using it as a GUI will be an excellent choice. While old legacy GUI lib is complex and outdated, a WebView-based app is still an option. However, a WebView needs a huge SDK to build and many dependencies to run, and it can only provide some features like a real web browser. That is why WebUI uses real web browsers to give you full features of comprehensive web technologies while keeping your software lightweight and portable.

### How Does it Work?

<div align="center">

![Diagram](https://github.com/ttytm/webui/assets/34311583/dbde3573-3161-421e-925c-392a39f45ab3)

</div>

Think of WebUI like a WebView controller, but instead of embedding the WebView controller in your program, which makes the final program big in size, and non-portable as it needs the WebView runtimes. Instead, by using WebUI, you use a tiny static/dynamic library to run any installed web browser and use it as GUI, which makes your program small, fast, and portable. **All it needs is a web browser**.

### Runtime Dependencies Comparison

|                                 | WebView           | Qt                         | WebUI               |
| ------------------------------- | ----------------- | -------------------------- | ------------------- |
| Runtime Dependencies on Windows | _WebView2_        | _QtCore, QtGui, QtWidgets_ | **_A Web Browser_** |
| Runtime Dependencies on Linux   | _GTK3, WebKitGTK_ | _QtCore, QtGui, QtWidgets_ | **_A Web Browser_** |
| Runtime Dependencies on macOS   | _Cocoa, WebKit_   | _QtCore, QtGui, QtWidgets_ | **_A Web Browser_** |

## Wrappers

| Language        | v2.4.x API | v2.5.x API | Link                                                    |
| --------------- | --- | -------------- | ---------------------------------------------------------  |
| Python          | ✔️ | _not complete_ | [Python-WebUI](https://github.com/webui-dev/python-webui)  |
| Go              | ✔️ | _not complete_ | [Go-WebUI](https://github.com/webui-dev/go-webui)          |
| Zig             | ✔️ |  _not complete_ | [Zig-WebUI](https://github.com/webui-dev/zig-webui)        |
| Nim             | ✔️ |  _not complete_ | [Nim-WebUI](https://github.com/webui-dev/nim-webui)        |
| V               | ✔️ |  _not complete_ | [V-WebUI](https://github.com/webui-dev/v-webui)            |
| Rust            | _not complete_ |  _not complete_ | [Rust-WebUI](https://github.com/webui-dev/rust-webui)      |
| TS / JS (Deno)  | ✔️ |  _not complete_ | [Deno-WebUI](https://github.com/webui-dev/deno-webui)      |
| TS / JS (Bun)   | _not complete_ |  _not complete_ | [Bun-WebUI](https://github.com/webui-dev/bun-webui)        |
| Swift           | _not complete_ |  _not complete_ | [Swift-WebUI](https://github.com/webui-dev/swift-webui)    |
| Odin            | _not complete_ |  _not complete_ | [Odin-WebUI](https://github.com/webui-dev/odin-webui)      |
| Pascal          | _not complete_ |  _not complete_ | [Pascal-WebUI](https://github.com/webui-dev/pascal-webui)  |
| Purebasic       | _not complete_ |  _not complete_ | [Purebasic-WebUI](https://github.com/webui-dev/purebasic-webui)|
| - |  |  |
| Common Lisp     | _not complete_ |  _not complete_ | [cl-webui](https://github.com/garlic0x1/cl-webui)          |
| Delphi          | _not complete_ |  _not complete_ | [WebUI4Delphi](https://github.com/salvadordf/WebUI4Delphi) |
| C#              | _not complete_ |  _not complete_ | [WebUI4CSharp](https://github.com/salvadordf/WebUI4CSharp) |
| WebUI.NET       | _not complete_ |  _not complete_ | [WebUI.NET](https://github.com/Juff-Ma/WebUI.NET)          |
| QuickJS         | _not complete_ |  _not complete_ | [QuickUI](https://github.com/xland/QuickUI)                |
| PHP             | _not complete_ |  _not complete_ | [PHPWebUiComposer](https://github.com/KingBes/php-webui-composer) |

## Supported Web Browsers

| Browser         | Windows         | macOS         | Linux           |
| --------------- | --------------- | ------------- | --------------- |
| Mozilla Firefox | ✔️              | ✔️            | ✔️              |
| Google Chrome   | ✔️              | ✔️            | ✔️              |
| Microsoft Edge  | ✔️              | ✔️            | ✔️              |
| Chromium        | ✔️              | ✔️            | ✔️              |
| Yandex          | ✔️              | ✔️            | ✔️              |
| Brave           | ✔️              | ✔️            | ✔️              |
| Vivaldi         | ✔️              | ✔️            | ✔️              |
| Epic            | ✔️              | ✔️            | _not available_ |
| Apple Safari    | _not available_ | _coming soon_ | _not available_ |
| Opera           | _coming soon_   | _coming soon_ | _coming soon_   |

## Supported WebView

| WebView         | Status         |
| --------------- | --------------- |
| Windows WebView2 | ✔️ |
| Linux GTK WebView   | ✔️ |
| macOS WKWebView  | ✔️ |

### License

> Licensed under the MIT License.

### Stargazers

[![Stargazers repo roster for @webui-dev/go-webui](https://reporoster.com/stars/webui-dev/go-webui)](https://github.com/webui-dev/webui/stargazers)
