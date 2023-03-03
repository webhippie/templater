# Templater

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/templater?sort=semver)](https://github.com/webhippie/templater) [![Build Status](https://github.com/webhippie/templater/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/templater/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/templater.svg)](https://pkg.go.dev/github.com/webhippie/templater) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/templater)](https://goreportcard.com/report/github.com/webhippie/templater) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/d95dc8cbd6a14ee78b3d52a6a0104304)](https://www.codacy.com/gh/webhippie/templater/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/templater&amp;utm_campaign=Badge_Grade)

Templater is used by our docker containers to provide a functionality for clean
templating based on the Golang `text/template` package. The variables are always
provided through environment variables. Before we integrated templater every
template has been handled by `envsubst` which doesn't provide any control
structures while Golang templates got this builtin.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/templater.git
cd templater

make generate build

./bin/templater -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/templater/releases
[dockerhub]: https://hub.docker.com/r/webhippie/templater/tags/
[quay]: https://quay.io/repository/webhippie/templater?tab=tags
[docs]: https://webhippie.github.io/templater/#getting-started
[golang]: http://golang.org/doc/install.html
