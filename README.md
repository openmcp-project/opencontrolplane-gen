# opencontrolplane-gen

`opencontrolplane-gen` is a code transformation tool to use together with `go generate`.

## Commands

`opencontrolplane-gen` defines a `Command` interface to transform go source files line by line.

`Commands` are executed by a `Runner` which processes go source files in single pass fashion.

This requires `Commands` to be stateful in regards of their activation state and arguments.

## Runner

A `Runner` takes an abitrary number of `Commands` to transform individual go source files.

When nesting `Commands` in source files, be aware of the order that `Commands` are defined in the `Runner`.

## Usage

`opencontrolplane-gen` is expected to be used with `go generate`, see the following example source code file:

```go
//go:generate opencontrolplane-gen
package testdata

// opencontrolplane-gen:replace TestReconciler=RECONCILER_NAME
type TestReconciler struct {
	namespace string
	//opencontrolplane-gen:if OPTIONAL_FIELDS=include
	conditionalName string
	// test nested command
	//opencontrolplane-gen:replace Replace=FIELD_NAME
	conditionalReplace int
	//opencontrolplane-gen:fi
}

// opencontrolplane-gen:replace TestReconciler=RECONCILER_NAME
func (r *TestReconciler) Reconciler() {
	//opencontrolplane-gen:replace name=FIELD_NAME namespace=FIELD_NAMESPACE
	var name, namespace string
	//opencontrolplane-gen:replace name=FIELD_NAME
	_ = name
	//opencontrolplane-gen:replace namespace=FIELD_NAMESPACE
	_ = namespace
}
```

`opencontrolplane-gen` allows in memory transformation via the `DRY_RUN` environment variable:

```shell
DRY_RUN=true go generate ./...
```

Pass values by defining the environment variables used in the command directives:

```shell
DRY_RUN=true RECONCILER_NAME=MyReconciler OPTIONAL_FIELDS=include go generate ./...
```

Inspect debug output:

```shell
DRY_RUN=true DEBUG=true OPTIONAL_FIELDS=include go generate ./...
```

Save changes to disk:

```shell
go generate ./...
```

## Execution Options

Environment Variables:

- DRY_RUN (bool): 1 or true result in the in-memory source code transformation being printed to stdout.
- DEBUG (bool): 1 or true result in debug output being printed to stdout.

## Support, Feedback, Contributing

This project is open to feature requests/suggestions, bug reports etc. via [GitHub issues](https://github.com/openmcp-project/platform-service-template/issues). Contribution and feedback are encouraged and always welcome. For more information about how to contribute, the project structure, as well as additional contribution information, see our [Contribution Guidelines](https://github.com/openmcp-project/.github/blob/main/CONTRIBUTING.md).

## Security / Disclosure

If you find any bug that may be a security problem, please follow our instructions at [in our security policy](https://github.com/openmcp-project/platform-service-template/security/policy) on how to report it. Please do not create GitHub issues for security-related doubts or problems.

## Code of Conduct

We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for everyone. By participating in this project, you agree to abide by its [Code of Conduct](https://github.com/openmcp-project/.github/blob/main/CODE_OF_CONDUCT.md) at all times.

## Licensing

Copyright OpenControlPlane contributors. Please see our [LICENSE](LICENSE) for copyright and license information. Detailed information including third-party components and their licensing/copyright information is available [via the REUSE tool](https://api.reuse.software/info/github.com/openmcp-project/platform-service-template).

---

<p align="center">
  <a href="https://apeirora.eu/content/projects/">
    <img alt="BMWK-EU funding logo" src="https://apeirora.eu/assets/img/BMWK-EU.png" width="300"/>
  </a>
</p>

<p align="center">
  OpenControlPlane is part of <a href="https://apeirora.eu/content/projects/">ApeiroRA</a>, an EU Important Project of Common European Interest (IPCEI-CIS).
</p>

<p align="center">
  Copyright Linux Foundation Europe. For web site terms of use, trademark policy and other project policies please see <a href="https://linuxfoundation.eu/en/policies">https://linuxfoundation.eu/en/policies</a>.
</p>
