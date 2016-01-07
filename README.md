# Templater

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/templater/status.svg)](http://github.dronehippie.de/webhippie/templater)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

Templater is used by our docker containers to provide a functionality for clean
templating based on the golang `text/template` package. The variables are always
provided through environment variables. Before we integrated templater every
template has been handled by `envsubst` which doesn't provide any control
structures.


## Install

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html)

```bash
go get github.com/webhippie/templater
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2015 Thomas Boerger <http://www.webhippie.de>
```
