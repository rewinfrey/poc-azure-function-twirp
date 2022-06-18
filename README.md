# Azure Function and Protobuf using Twirp PoC

This is a demo app as a PoC to work with [Azure Functions][AzureFunctions] and [protobufs][protobufs] via [Twirp][Twirp].

### Setup

This PoC uses a Makefile to build the executables.  The binaries are written to the `bin/` directory.

```shell
cd poc-azure-function-twirp
make
```

[Azure Core Tools][AzureCoreTools] are also required.

### Run

Run the Azure Function runtime and event listener (this implicitly runs the `bin/server` binary):

```shell
cd poc-azure-function-twirp
func start [--verbose]
```

In a new terminal session run the binary that issues Twirp requests:

```shell
cd poc-azure-function-twirp
bin/runner
```

[AzureCoreTools]: https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local
[AzureFunctions]: https://docs.microsoft.com/en-us/azure/azure-functions/functions-overview
[Twirp]: https://twitchtv.github.io/twirp/
[protobufs]: https://developers.google.com/protocol-buffers



