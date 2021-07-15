# TPoller Server
## Overview
The purpose of this application is to poll time-series data per time interval from any (telemetry) application running a gRPC server which implements the [the following gRPC service definition](/proto/tpoller.proto):

```proto
service TPoller {
    rpc PollTimeSeriesData (google.protobuf.Empty) returns (stream PolledTimeSeriesDatum) {}
}

message PolledLabel {
    string name = 1;
    string value = 2;
}

message PolledTimeSeriesDatum {
    string metric = 1;
    repeated PolledLabel labels = 2;
    double value = 3;
    google.protobuf.Timestamp timestamp = 4;
}
```

Your server which implemented the gRPC service definition is called a **data reader**. `tpoller-server` will then send your polled data over gRPC to a **fast time-series data storage server** called [tstorage-server](https://github.com/bartmika/tstorage-server).

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.16.3``
* ``tstorage-server`` running in background

## Installation
<!-- 1. Please visit the [sparkfunweathershield-arduino](https://github.com/bartmika/sparkfunweathershield-arduino) repository and setup the external device and connect it to your development machine.

2. Please visit the [serialreader-server](https://github.com/bartmika/serialreader-server) repository and setup that application on your device.

3. Please visit the [tstorage-server](https://github.com/bartmika/tstorage-server) repository and setup that application on your device. -->

Get our latest code.

```bash
go get -u github.com/bartmika/tpoller-server
```

## Usage
Run our application.

    go run main.go serve

If the server successfully starts you should see a message in your **termnal**:

    2021/07/10 15:40:36 Synching with local time...
    2021/07/10 15:41:00 Synchronized with local time.
    2021/07/10 15:41:00 TPoller is now running.

## Used by:
This server is confirmed to successfully poll from the following application(s):
* [tarduinoreader-server](https://github.com/bartmika/tarduinoreader-server)

If you'd like to add your app, please create a [pull request](https://github.com/bartmika/tpoller-server/pulls) with a link to your app.

## Documentation

### ``serve``
Run the gRPC server to allow other services to access the storage application

#### Fields

* `-a` or `--tstorage_addr` is for the port for this server to run on. Default value is 50051 if you don't use this option.
* `-b` or `--tstorage_port` is for the port for this server to run on. Default value is 50051 if you don't use this option.
* `-c` or `--datareader_addr` is for the location path to use to save the database files to. Default value is './tsdb'.
* `-d` or `--datareader_port` is for the location path to use to save the database files to. Default value is './tsdb'.


#### Example:

```bash
$GOBIN/tpoller-server serve -a=127.0.0.1 -b=50051 -c=127.0.0.1 -d=50052
```

#### Output:

```bash
2021/07/10 15:40:36 Synching with local time...
2021/07/10 15:41:00 Synchronized with local time.
2021/07/10 15:41:00 TPoller is now running.
```

### ``version``
Prints the current version of our application.

#### Example:

```bash
$GOBIN/tpoller-server version
```

#### Output:

```bash
tpoller-server v1.0
```

## License

This application is licensed under the **BSD 3-Clause License**. See [LICENSE](LICENSE) for more information.
