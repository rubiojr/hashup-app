# HashUp App

Search your indexed files.

Graphical user interface for [HashUp](https://github.com/rubiojr/hashup).

![screenshot](/docs/screenshot.png)

## Installation

### macOS

* Requires Go 1.24 and Xcode command line tools

```
script/package
```

Generates an .app bundle in `dist/HashUp.app`

### Linux

* Requires Go 1.24 and build tools

Ubuntu/Debian:

```
sudo apt install libgtk-3-dev libwebkit2gtk-4.1-dev build-essential
```

Fedora:

```
dnf install gtk3-devel webkit2gtk4.1-devel make automake gcc pkg-config
```

```
go install github.com/rubiojr/hashup-app@latest
```
