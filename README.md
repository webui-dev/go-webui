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

## Documentation

- [Online Documentation](https://webui.me/docs/#/go_api)

### License

> Licensed under the MIT License.
