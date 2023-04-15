# Benchmarks of Enum implementations

Benchmarks for different enum implementations in Go.

The base is [the k6's execution status type][base-impl] implemented via the
generator enumer.

[base-impl]: https://github.com/grafana/k6/blob/1785ec07fd0e63bfc5323ca48e894a47cb02ac79/lib/execution_status_gen.go#L10-L19


<!-- vim-markdown-toc GFM -->

* [What are the implementations](#what-are-the-implementations)
* [Benchmarks](#benchmarks)
    * [Enumer vs Manual](#enumer-vs-manual)
    * [Enumer vs Idiomatic](#enumer-vs-idiomatic)

<!-- vim-markdown-toc -->

## What are the implementations

I'm trying three implementations:

- The way enumer implements enums using a string with all the values and then a
  slice of integers as the indexes for getting data from that string.
- The most idiomatic way, mimicking how the HTTP library implements enums.
- The manual way I tend to write, with an optimization for the String method
  using a separate slice of strings for faster retrievel than from map lookups.

## Benchmarks

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	697159990	         1.599 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	750704174	         1.599 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	749813324	         1.600 ns/op
BenchmarkExecutionStatusString_LastItem-8      	80869797	        15.02 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	75749607	        15.95 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	100000000	        13.18 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/enumer/executionstatus	8.039s
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/idiomatic/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         1.024 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         1.025 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         1.025 ns/op
BenchmarkExecutionStatusString_LastItem-8      	266512113	         4.503 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	358854044	         3.342 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	383289942	         3.127 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/idiomatic/executionstatus	8.336s
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/manual/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         0.8559 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         0.8574 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         0.9450 ns/op
BenchmarkExecutionStatusString_LastItem-8      	93189710	        12.81 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	93819939	        12.51 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	137130495	         8.751 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/manual/executionstatus	7.644s
```

### Enumer vs Manual

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
                                    │  enumer.txt   │               manual.txt               │
                                    │    sec/op     │    sec/op      vs base                 │
ExecutionStatus_String_LastItem-8     1.6010n ± ∞ ¹   0.8562n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatus_String_MiddleItem-8   1.5990n ± ∞ ¹   0.8563n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatus_String_FirstItem-8    1.6000n ± ∞ ¹   0.8566n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_LastItem-8       15.76n ± ∞ ¹    10.76n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_MiddleItem-8     13.39n ± ∞ ¹    10.25n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_FirstItem-8     13.140n ± ∞ ¹    8.748n ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                                4.741n          2.909n        -38.64%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
```

### Enumer vs Idiomatic

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
                                    │  enumer.txt   │             idiomatic.txt             │
                                    │    sec/op     │    sec/op     vs base                 │
ExecutionStatus_String_LastItem-8      1.601n ± ∞ ¹   1.025n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatus_String_MiddleItem-8    1.599n ± ∞ ¹   1.025n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatus_String_FirstItem-8     1.600n ± ∞ ¹   1.025n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_LastItem-8      15.760n ± ∞ ¹   4.501n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_MiddleItem-8    13.390n ± ∞ ¹   3.342n ± ∞ ¹        ~ (p=1.000 n=1) ²
ExecutionStatusString_FirstItem-8     13.140n ± ∞ ¹   3.127n ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                                4.741n         1.924n        -59.43%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
```
