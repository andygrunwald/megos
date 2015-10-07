# megos

[![GoDoc](https://godoc.org/github.com/andygrunwald/megos?status.svg)](https://godoc.org/github.com/andygrunwald/megos)

[Go(lang)](https://golang.org/) client library for accessing [Apache Mesos](http://mesos.apache.org/) cluster.

## Installation

It is go gettable

    $ go get github.com/andygrunwald/megos

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/andygrunwald/megos
    $ go test -v

## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/megos) for a detailed API description.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/megos#pkg-examples).

### TODO

TODO

## Version compatibility

This library was tested with Apache Mesos in version 0.22.1.
In theory this should work up to version 0.24.x (including).

In version 0.25.x they renamed various API endpoints (like state.json to /state).
See [Upgrading Mesos - Upgrading from 0.24.x to 0.25.x](http://mesos.apache.org/documentation/latest/upgrades/) for details.

If you are running a Mesos cluster >= 0.25.x and you can make this library working with it, please start a Pull Request or open an issue.
We are happy to get this support into.

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
