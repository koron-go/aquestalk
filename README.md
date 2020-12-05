# koron-go/aquestalk

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/aquestalk)](https://pkg.go.dev/github.com/koron-go/aquestalk)
[![GoDoc](https://godoc.org/github.com/koron-go/aquestalk?status.svg)](https://godoc.org/github.com/koron-go/aquestalk)
[![Actions/Go](https://github.com/koron-go/aquestalk/workflows/Go/badge.svg)](https://github.com/koron-go/aquestalk/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/aquestalk)](https://goreportcard.com/report/github.com/koron-go/aquestalk)

[Aquestalk][aq] for Go.

## Gettings started

To get `aquestalk` package:

```console
$ go get github.com/koron-go/aquestalk
$ cd ~/go/src/github.com/koron-go/aquestalk
```

Download **AquesTalk1** engine of your platform from [download section][dl].

Extract a zip file, and copy AquesTalk.dll (or so) to project dir or one of
dirs in PATH environment. Next example command copies a DLL of F1 voice for
windows/amd64.

```console
# for Windows
$ cp aqtk1-win-eva/x64/f1/AquesTalk.dll .

# for Linux
$ cp aqtk1-lnx-eva/lib64/f1/libAquesTalk.so .
```

You'll hear "Konnichi wa Gopher" (こんにちはGopher) with next command.

```console
$ go run ./examples/01_hello/main.go
```

### for Linux

You may need libasound2-dev to install, before `go run`.

```console
$ sudo apt install -y libasound2-dev
```

### for macOS

You should use **AquesTalk10** (not AquesTalk1).

Download "AquesTalk10 Mac  1.1.0 (2017/11/01)".

Extract a zip file, and copy AquesTalk.framework to /Library/Frameworks

```console
$ cp -Rp AquesTalk.framework /Library/Frameworks/
```

## Supported Platforms

* Windows (x86 and amd64)
* Linux (x86 and amd64)
* macOS (confirmed on M1 macmini)

## IMPORTANT: License of AquesTalk

This package is a just a wrapper of AquesTalk engine, doesn't include
AquesTalk.

Using AquesTalk engine with your programs, you must obtain AquesTalk's license.

This package is developped by using evaluation version of AquesTalk.

[aq]:https://www.a-quest.com/products/aquestalk_1.html
[dl]:https://www.a-quest.com/download.html
