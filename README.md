# Templater

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/templater/status.svg)](http://github.dronehippie.de/webhippie/templater)
[![Stories in Ready](https://badge.waffle.io/webhippie/templater.svg?label=ready&title=Ready)](http://waffle.io/webhippie/templater)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d95dc8cbd6a14ee78b3d52a6a0104304)](https://www.codacy.com/app/webhippie/templater?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/templater&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/templater?status.svg)](http://godoc.org/github.com/webhippie/templater)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/templater)](http://goreportcard.com/report/github.com/webhippie/templater)
[![](https://images.microbadger.com/badges/image/tboerger/templater.svg)](http://microbadger.com/images/tboerger/templater "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Templater is used by our docker containers to provide a functionality for clean templating based on the golang `text/template` package. The variables are always provided through environment variables. Before we integrated templater every template has been handled by `envsubst` which doesn't provide any control structures.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/misc/templater). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/webhippie/templater
cd $GOPATH/src/github.com/webhippie/templater
make clean retool sync generate build

./templater -h
```


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
