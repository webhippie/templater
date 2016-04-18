# Templater

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/templater/status.svg)](http://github.dronehippie.de/webhippie/templater)
[![Coverage Status](http://coverage.dronehippie.de/badges/webhippie/templater/coverage.svg)](http://coverage.dronehippie.de/webhippie/templater)
[![Go Doc](https://godoc.org/github.com/webhippie/templater?status.svg)](http://godoc.org/github.com/webhippie/templater)
[![Go Report](http://goreportcard.com/badge/webhippie/templater)](http://goreportcard.com/report/webhippie/templater)

Templater is used by our docker containers to provide a functionality for clean
templating based on the golang `text/template` package. The variables are always
provided through environment variables. Before we integrated templater every
template has been handled by `envsubst` which doesn't provide any control
structures.


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/templater). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).
If you are missing an architecture just write us on our nice
[Gitter](https://gitter.im/webhippie/general) chat. Take a look at the help
output, per default we enabled an auto updater for the binary to avoid bugs
related to old versions. If you find a security issue please contact
thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/webhippie/templater
cd $GOPATH/src/github.com/webhippie/templater
make deps build

bin/templater -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2015-2016 Thomas Boerger <http://www.webhippie.de>
```
