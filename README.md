<div align="center">

# WebUI Go

<div>
  <!-- <a href="https://github.com/webui-dev/go-webui/actions?query=branch%3Amain">
    <img
      alt="Build Status"
      src="https://img.shields.io/github/actions/workflow/status/webui-dev/go-webui/ci.yml?branch=main&style=for-the-badge&logo=V&labelColor=414868&logoColor=C0CAF5"
    >
  </a> -->
  <a href="https://github.com/webui-dev/go-webui/pulse">
    <img
      alt="Last Commit"
      src="https://img.shields.io/github/last-commit/webui-dev/go-webui?style=for-the-badge&logo=github&logoColor=C0CAF5&labelColor=414868"
    />
  </a>
  <a href="https://github.com/webui-dev/go-webui/releases/latest">
    <img
      alt="Go-WebUI Release Version"
      src="https://img.shields.io/github/v/tag/webui-dev/go-webui?style=for-the-badge&logo=webtrees&logoColor=C0CAF5&labelColor=414868&color=7664C6&label=release"
    >
  </a>
  <a href="https://github.com/webui-dev/go-webui/blob/main/LICENSE">
    <img
      alt="License"
      src="https://img.shields.io/github/license/webui-dev/go-webui?style=for-the-badge&amp&logo=opensourcehardware&label=License&logoColor=C0CAF5&labelColor=414868&color=8c73cc"
    >
  </a>
</div>

<br>

![Screenshot](https://github.com/webui-dev/webui/assets/34311583/57992ef1-4f7f-4d60-8045-7b07df4088c6)

> WebUI is not a web-server solution or a framework, but it allows you to use any web browser as a GUI, with your preferred language in the backend and HTML5 in the frontend. All in a lightweight portable lib.

</div>

## Features

- Fully Independent (No need for any third-party runtimes)
- Lightweight ~200 Kb & Small memory footprint
- Fast binary communication protocol between WebUI and the browser (*Instead of JSON*)
- Multi-platform & Multi-Browser
- Using private profile for safety
- Original library written in Pure C

## Usage

### Installation

```sh
go get github.com/webui-dev/go-webui@latest
```

### Minimal Example

```go
package main

import "github.com/webui-dev/go-webui"

func main() {
	w := webui.NewWindow()
	w.Show("<html>Hello World</html>")
	webui.Wait()
}
```

Find more examples in the [`examples/`](https://github.com/webui-dev/go-webui/tree/main/examples) directory.

## Documentation

> **Note**
> We are currently writing documentation.

- [Online Documentation](https://webui.me/docs/#/go)

## UI & The Web Technologies

[Borislav Stanimirov](https://ibob.bg/) discusses using HTML5 in the web browser as GUI at the [C++ Conference 2019 (*YouTube*)](https://www.youtube.com/watch?v=bbbcZd4cuxg).

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

|  | WebView | Qt | WebUI |
| ------ | ------ | ------ | ------ |
| Runtime Dependencies on Windows | *WebView2* | *QtCore, QtGui, QtWidgets* | ***A Web Browser*** |
| Runtime Dependencies on Linux | *GTK3, WebKitGTK* | *QtCore, QtGui, QtWidgets* | ***A Web Browser*** |
| Runtime Dependencies on macOS | *Cocoa, WebKit* | *QtCore, QtGui, QtWidgets* | ***A Web Browser*** |

## Wrappers

| Language | Status | Link |
| ------ | ------ | ------ |
| Go | ✔️ | [Go-WebUI](https://github.com/webui-dev/go-webui) |
| Nim | ✔️ | [Nim-WebUI](https://github.com/webui-dev/nim-webui) |
| Pascal | ✔️ | [Pascal-WebUI](https://github.com/webui-dev/pascal-webui) |
| Python | ✔️ | [Python-WebUI](https://github.com/webui-dev/python-webui) |
| Rust | *not complete* | [Rust-WebUI](https://github.com/webui-dev/rust-webui) |
| TypeScript / JavaScript | ✔️ | [Deno-WebUI](https://github.com/webui-dev/deno-webui) |
| V | ✔️ | [V-WebUI](https://github.com/webui-dev/v-webui) |
| Zig | *not complete* | [Zig-WebUI](https://github.com/webui-dev/zig-webui) |

## Supported Web Browsers

| Browser | Windows | macOS | Linux |
| ------ | ------ | ------ | ------ |
| Mozilla Firefox | ✔️ | ✔️ | ✔️ |
| Google Chrome | ✔️ | ✔️ | ✔️ |
| Microsoft Edge | ✔️ | ✔️ | ✔️ |
| Chromium | ✔️ | ✔️ | ✔️ |
| Yandex | ✔️ | ✔️ | ✔️ |
| Brave | ✔️ | ✔️ | ✔️ |
| Vivaldi | ✔️ | ✔️ | ✔️ |
| Epic | ✔️ | ✔️ | *not available* |
| Apple Safari | *not available* | *coming soon* | *not available* |
| Opera | *coming soon* | *coming soon* | *coming soon* |

### License

> Licensed under the MIT License.

### Stargazers

[![Stargazers repo roster for @webui-dev/go-webui](https://reporoster.com/stars/webui-dev/go-webui)](https://github.com/webui-dev/webui/stargazers)
