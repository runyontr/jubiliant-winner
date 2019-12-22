# Chart Bump

The bash script in this folder will update the version of a Helm Chart.

## Usage

The application expects two arguments. The first is a path to a Helm chart (specifically, just a folder containing an `Chart.yaml` file), the second argument being
either `p`, `m`, or `M` depending on which part of the SemVer should be incremented. `p` specifies the `Patch` value should be incremented, `m` for `Minor` and `M` for `Major`.

## Example

```bash
$ ./bump.sh chart m
Upgrading Chart version from 1.2.3 to 1.3.0
$ ./bump.sh chart p
Upgrading Chart version from 1.3.0 to 1.3.1
$ ./bump.sh chart p
Upgrading Chart version from 1.3.1 to 1.3.2
$ ./bump.sh chart m
Upgrading Chart version from 1.3.2 to 1.4.0
$ ./bump.sh chart M
Upgrading Chart version from 1.4.0 to 2.0.0
```

## Future Work

- Error Checking. We don't handle any exceptions about when things aren't as expected
