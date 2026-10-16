# SOAP Go (enuan fork)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/enuan/soap-go)](https://pkg.go.dev/github.com/enuan/soap-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/enuan/soap-go)](https://goreportcard.com/report/github.com/enuan/soap-go)
[![CI](https://github.com/enuan/soap-go/actions/workflows/release.yaml/badge.svg)](https://github.com/enuan/soap-go/actions/workflows/release.yaml)

A Go SDK and CLI tool for SOAP web services. Hard fork of
[`way-platform/soap-go`](https://github.com/way-platform/soap-go) maintained in
the `enuan` org.

Why a hard fork: upstream does not process external contributions (the fixes
below have been open upstream for months), while enuan services need correct
code generation for real-world .NET WSDLs. See
[`AGENTS.md`](./AGENTS.md) for maintenance and upstream-sync policy.

## Differences from upstream (base: v0.29.2)

Codegen fixes (enuan):

- `complexContent/extension` inheritance: base fields are merged into derived
  types, so session tokens (`BaseRequest.SessionId`) and result state
  (`ResultBase.State`/`Success`) are no longer silently dropped.
- Self-referential simple types (e.g. a local `duration` restricted from
  `xs:duration`, as shipped in the Microsoft serialization schema of .NET
  WSDLs) no longer cause a stack overflow during generation.
- One client method per portType operation: multiple SOAP 1.1 bindings
  referencing the same portType no longer produce duplicated methods.

Community fixes picked from [`tnymlr/soap-go`](https://github.com/tnymlr/soap-go)
(authorship preserved in git history):

- XSD enum identifier sanitisation and stable enum output order.
- Element/type name collision handling (`Wrapper` suffix).
- Namespace-scoped type naming (`--namespace-prefixes`) and namespace-aware
  element reference resolution.
- `xs:include` / `xs:import` support and cross-namespace lookups.
- `soap:header` support in the generated client.
- `xs:all` decomposition and raw XML heuristic fixes.
- `xs:dateTime` without embedded timezone (`soap.XSDDateTime`).
- Booleans encoded as `0`/`1` on the wire.

Packaging:

- Single Go module: the CLI is part of the root module, so remote
  installation works (upstream's `cmd/soap` module had a `replace` directive
  that broke `go install ...@latest`).

## Features

- Support for SOAP 1.1, WSDL 1.1, and XSD 1.0
- Code generation from WSDL files
- Documentation generation

## CLI tool

The `soap` CLI tool can generate code, documentation, and call SOAP APIs on the fly.

```bash
$ soap

  Multi-tool for SOAP APIs

  USAGE

    soap [command] [--flags]

  CODE GENERATION

    gen [--flags]         Generate code for a SOAP API

  DOCUMENTATION

    doc [--flags]         Display documentation for a SOAP API

  NETWORK OPERATIONS

    call [--flags]        Call a SOAP action

  UTILS

    completion [command]  Generate the autocompletion script for the specified shell
    help [command]        Help about any command

  FLAGS

    -h --help             Help for soap
    -v --version          Version for soap
```

### Installing

```bash
go install github.com/enuan/soap-go/cmd/soap@latest
```

Prebuilt binaries for Linux, Windows, and Mac are available from the [Releases](https://github.com/enuan/soap-go/releases).

## License

MIT, see [LICENSE](./LICENSE). Original work © Way Data Technologies Oy;
fork modifications © enuan.
