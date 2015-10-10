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

## Examples / use cases

A (small) list of usecases how this library can be used:

* Determine the leader of a Mesos cluster
* Get a list of all completed Mesos tasks
* Get the stdout and stderr of a failed mesos task
* Get the statistics and push it to a different backend

Further more a few examples how the API can be used and the code looks like.

### TODO

TODO

## Version compatibility

This library was tested with Apache Mesos in version 0.22.1.
In theory this should work up to version 0.24.x (including).

In version 0.25.x they renamed various API endpoints (like state.json to /state).
See [Upgrading Mesos - Upgrading from 0.24.x to 0.25.x](http://mesos.apache.org/documentation/latest/upgrades/) for details.

If you are running a Mesos cluster >= 0.25.x and you can make this library working with it, please start a Pull Request or open an issue.
We are happy to get this support into.

## TODO-List

* Provide unit tests
* Add travis-ci
* Provide (executable) examples
* Add doc.go

## Contribution

* You have a question?
* Don`t know if a feature is supported?
* Want to implement a new feature, but don`t know how?
* You like the library and use it for your implementation / use case?
* You found a bug or incompatibility?
* Something is not working as expected?

Feel free to open a [new issue](https://github.com/andygrunwald/megos/issues/new).
I will be happy to answer them and try to help you.
It might be useful to add as much information as possible into the issue like Mesos version, example URL, (parts) of your code and the expected and current behaviour.

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
