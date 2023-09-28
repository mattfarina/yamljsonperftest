# Testing YAML and JSON parsing Via Kubernetes YAML Library

It's possible to parse both JSON and YAML files with the Kubernets YAML parser. Is there a difference in performance? That's what this repository is designed to test.

This test is meant to support Helm and the placement of JSON in index.yaml files for performance.

## &tldr Results

Three things were tested:

1. Parsing a large YAML file into a Go struct instance
2. Parsing a large JSON file into a Go struct instance using the YAML parser. This is important to assess backwards compatibility for older versions of Helm that will use the YAML parser.
3. Parsing a large JSON file into a Go struct instance using JSON parsing, which is what newer Helm would do with JSON.

A snapshot of results:
```shell
goos: linux
goarch: amd64
pkg: github.com/mattfarina/yamljsonperftest
cpu: 12th Gen Intel(R) Core(TM) i7-1280P
BenchmarkYaml-20                       1        18928784545 ns/op       8416900424 B/op 151667973 allocs/op
BenchmarkJsonThroughYaml-20            1        17394763008 ns/op       7843400360 B/op 137570777 allocs/op
BenchmarkJson-20                       1        2420351573 ns/op        444798936 B/op  10001927 allocs/op
PAS
```

It appears that JSON parsing through the YAML parser is a slight bit faster than YAML parsing. JSON parsing of the JSON is significantly faster. For Helm, older versions would continue to function with approximately their existing performance while newer versions would be faster with JSON.

## Setup

In order to run the tests you need to generate the test data. That can be done with Go installed by running the command:

```shell
go run ./cmd/gen
```

The test data is not included due to its size.

## Running Tests

To run the tests use the following command:

```shell
go test -bench=. -benchmem
```
