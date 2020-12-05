# koron-go/aquestalk

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/aquestalk)](https://pkg.go.dev/github.com/koron-go/aquestalk)
[![GoDoc](https://godoc.org/github.com/koron-go/aquestalk?status.svg)](https://godoc.org/github.com/koron-go/aquestalk)
[![Actions/Go](https://github.com/koron-go/aquestalk/workflows/Go/badge.svg)](https://github.com/koron-go/aquestalk/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/aquestalk)](https://goreportcard.com/report/github.com/koron-go/aquestalk)

[Aquestalk][aq] for Go.

## Gettings started

Get `aquestalk` package.

```console
$ go get github.com/koron-go/aquestalk
$ cd ~/go/src/github.com/koron-go/aquestalk
```

Download **AquesTalk1** engine of your platform from [download section][dl].

Extract a zip file, and copy AquesTalk.dll (or so) to project dir or one of
dirs in PATH environment. Next example command copies a DLL of F1 voice for
windows/amd64.

```console
$ cp aqtk1-win-eva/x64/f1/AquesTalk.dll .
```

You'll hear "Konnichi wa Gopher" (こんにちはGopher) with next command.

```console
$ go run ./examples/01_hello/main.go
```

[aq]:https://www.a-quest.com/products/aquestalk_1.html
[dl]:https://www.a-quest.com/download.html

## Supported Platforms

* Windows (x86 and amd64)
