# megos

[![GoDoc](https://godoc.org/github.com/andygrunwald/megos?status.svg)](https://godoc.org/github.com/andygrunwald/megos)
[![Build Status](https://travis-ci.org/andygrunwald/megos.svg)](https://travis-ci.org/andygrunwald/megos)

[Go(lang)](https://golang.org/) client library for accessing [Apache Mesos](http://mesos.apache.org/) cluster.

## Features

* Determine the Mesos leader
* Get the current state of every mesos node (master or slave)
* Retrieve stdout and stderr of tasks
* Covered with unit tests

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

### Determine the leader node

```go
node1, _ := url.Parse("http://192.168.1.120:5050/")
node2, _ := url.Parse("http://192.168.1.122:5050/")

mesos := megos.NewClient([]*url.URL{node1, node2})
leader, err := mesos.DetermineLeader()
if err != nil {
	panic(err)
}

fmt.Println(leader)
// Output:
// 	master@192.168.1.122:5050
```

### Get the version of Mesos

```
node1, _ := url.Parse("http://192.168.1.120:5050/")
node2, _ := url.Parse("http://192.168.1.122:5050/")

mesos := megos.NewClient([]*url.URL{node1, node2})
state, err := mesos.GetStateFromCluster()
if err != nil {
	panic(err)
}

fmt.Printf("Mesos v%s", state.Version)
// Output:
// Mesos 0.22.1
```

### Get stdout and stderr of a task

Get stdout and stderr from a task of the [chronos](https://github.com/mesos/chronos) framework. Error checks are dropped for simplicity.

```go
node1, _ := url.Parse("http://192.168.1.120:5050/")
node2, _ := url.Parse("http://192.168.1.122:5050/")
mesos := megos.NewClient([]*url.URL{node1, node2})

frameworkPrefix := "chronos"
taskID := "ct:1444578480000:0:example-chronos-task:"

mesos.DetermineLeader()
state, _ := mesos.GetStateFromLeader()

framework, _ := mesos.GetFrameworkByPrefix(state.Frameworks, frameworkPrefix)
task, _ := mesos.GetTaskByID(framework.CompletedTasks, taskID)
slave, _ := mesos.GetSlaveByID(state.Slaves, task.SlaveID)

pid := mesos.ParsePidInformation(slave.PID)
slaveState, _ := mesos.GetStateFromPid(pid)

framework, _ = mesos.GetFrameworkByPrefix(slaveState.Frameworks, frameworkPrefix)
executor, _ := mesos.GetExecutorByID(framework.CompletedExecutors, taskID)

stdOut, _ := mesos.GetStdOutOfTask(pid, executor.Directory)
stdErr, _ := mesos.GetStdErrOfTask(pid, executor.Directory)

fmt.Println(stdOut)
fmt.Println("================")
fmt.Println(stdErr)
// Output:
// Registered executor on 192.168.1.123
// Starting task ct:1444578480000:0:example-chronos-task:
// sh -c 'MY COMMAND'
// Forked command at 10629
// ...
// ================
// I1011 17:48:00.390614 10602 exec.cpp:132] Version: 0.22.1
// I1011 17:48:00.532158 10618 exec.cpp:206] Executor registered on slave 20150603-103119-2046951690-5050-24382-S1
```

## Version compatibility

This library was tested with Apache Mesos in version 0.22.1.
In theory this should work up to version 0.24.x (including).

In version 0.25.x they renamed various API endpoints (like state.json to /state).
See [Upgrading Mesos - Upgrading from 0.24.x to 0.25.x](http://mesos.apache.org/documentation/latest/upgrades/) for details.

If you are running a Mesos cluster >= 0.25.x and you can make this library working with it, please start a Pull Request or open an issue.
We are happy to get this support into.

## TODO-List

* Add doc.go

## Other/Similar projects

* [boldfield/go-mesos](https://github.com/boldfield/go-mesos): A client for discovering information about a Mesos exposed via HTTP API
* [antonlindstrom/mesos_stats](https://github.com/antonlindstrom/mesos_stats): Statistics definition for Mesos /monitor/statistics.json
* [Clever/marathon-stats](https://github.com/Clever/marathon-stats): A simple container which queries marathon and mesos for stats about their current state, and logs these stats to stderr
* [bolcom/mesos_metrics](https://github.com/bolcom/mesos_metrics): Go definitions for the Mesos `{master}:5050/metrics/snapshot` and `{slave}:5051/metrics/snapshot` endpoints

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
