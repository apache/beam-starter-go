# Apache Beam starter for Go

If you want to clone this repository to start your own project,
you can choose the license you prefer and feel free to delete anything related to the license you are dropping.
See also https://beam.apache.org/get-started/quickstart/go/ for additional instructions for this quickstart.

## Before you begin

Make sure you have a [Go](https://go.dev/) development environment ready.
If you don't, you can follow the instructions in the
[Download and install](https://go.dev/doc/install) page.

## Running the pipeline

Running your pipeline in Go is as easy as running the script file directly.

```sh
# You can run the script file directly.
go run main.go

# To run passing command line arguments.
go run main.go --input-text="🎉"

# To run the tests.
go test ./...
```

## Using other runners

To keep this template small, it only includes the [Direct Runner](https://beam.apache.org/documentation/runners/direct/).

For a comparison of what each runner currently supports, look at the [Beam Capability Matrix](https://beam.apache.org/documentation/runners/capability-matrix/).

To add a new runner, visit the runner's page for instructions on how to include it.

## Contributing

Thank you for your interest in contributing!
All contributions are welcome! 🎉🎊

Please refer to the [`CONTRIBUTING.md`](CONTRIBUTING.md) file for more information.

# License

This software is distributed under the terms of both the MIT license and the
Apache License (Version 2.0).

See [LICENSE](LICENSE) for details.
