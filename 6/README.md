# Chart Bump

The Golang application in this folder will update the version of a Helm Chart.

## Usage

The application expects two arguments. The first is a path to a Helm chart (specifically, just a folder containing an `Chart.yaml` file), the second argument being
either `p`, `m`, or `M` depending on which part of the SemVer should be incremented. `p` specifies the `Patch` value should be incremented, `m` for `Minor` and `M` for `Major`.

## Examples

```bash
$ go run main.go chart m
Upgrading Chart version from 1.2.3 to 1.3.0
$ go run main.go chart p
Upgrading Chart version from 1.3.0 to 1.3.1
$ go run main.go chart p
Upgrading Chart version from 1.3.1 to 1.3.2
$ go run main.go chart m
Upgrading Chart version from 1.3.2 to 1.4.0
$ go run main.go chart M
Upgrading Chart version from 1.4.0 to 2.0.0
```

## Future Work

- Backup existing chart metadata
- Prevent unused fields from being written. When running the Go application, (e.g. updating the Patch), new fields get inserted into the file:

```yaml
name: litecoind
home: ""
sources: []
version: 1.2.4
description: Sample Chart file for litecoind
keywords: []
maintainers:
  - name: runyontr
    email: tom@runyon.dev
    url: ""
    xxx_nounkeyedliteral: {}
    xxx_unrecognized: []
    xxx_sizecache: 0
engine: ""
icon: ""
apiversion: ""
condition: ""
tags: ""
appversion: ""
deprecated: false
tillerversion: ""
annotations: {}
kubeversion: ""
xxx_nounkeyedliteral: {}
xxx_unrecognized: []
xxx_sizecache: 0
```
