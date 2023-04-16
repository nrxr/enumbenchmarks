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
    * [Enumer vs Combined](#enumer-vs-combined)

<!-- vim-markdown-toc -->

## What are the implementations

I'm trying three implementations:

- The way enumer implements enums using a string with all the values and then a
  slice of integers as the indexes for getting data from that string.
- The most idiomatic way, mimicking how the HTTP library implements enums.
- The manual way I tend to write, with an optimization for the String method
  using a separate slice of strings for faster retrievel than from map lookups.
- The combined is the fastest from idiomatic with the fastest from manual. So
  string slice for String() and idiomatic switch for ExecutionStatusString
  function.

## Benchmarks

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/combined/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         0.5797 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         0.5248 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         0.5227 ns/op
BenchmarkExecutionStatusString_LastItem-8      	440918524	         2.727 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	587417516	         2.045 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	627972648	         1.912 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/combined/executionstatus	6.737s
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         0.9765 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         0.9774 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         0.9775 ns/op
BenchmarkExecutionStatusString_LastItem-8      	133323987	         9.020 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	132739348	         9.036 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	149557675	         8.093 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/enumer/executionstatus	10.619s
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/idiomatic/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         0.6260 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         0.6257 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         0.6251 ns/op
BenchmarkExecutionStatusString_LastItem-8      	439873963	         2.725 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	586818466	         2.046 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	627604122	         1.912 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/idiomatic/executionstatus	6.614s
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/manual/executionstatus
BenchmarkExecutionStatus_String_LastItem-8     	1000000000	         0.5239 ns/op
BenchmarkExecutionStatus_String_MiddleItem-8   	1000000000	         0.5230 ns/op
BenchmarkExecutionStatus_String_FirstItem-8    	1000000000	         0.5232 ns/op
BenchmarkExecutionStatusString_LastItem-8      	171471234	         6.886 ns/op
BenchmarkExecutionStatusString_MiddleItem-8    	198201745	         6.044 ns/op
BenchmarkExecutionStatusString_FirstItem-8     	224331176	         5.444 ns/op
PASS
ok  	github.com/nrxr/enumbenchmarks/manual/executionstatus	7.485s
```

### Enumer vs Manual

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
                                    │  enumer.txt  │              manual.txt              │
                                    │    sec/op    │    sec/op     vs base                │
ExecutionStatus_String_LastItem-8     0.9767n ± 0%   0.5234n ± 0%  -46.41% (p=0.000 n=20)
ExecutionStatus_String_MiddleItem-8   0.9767n ± 2%   0.5234n ± 1%  -46.41% (p=0.000 n=20)
ExecutionStatus_String_FirstItem-8    0.9778n ± 1%   0.5325n ± 8%  -45.54% (p=0.000 n=20)
ExecutionStatusString_LastItem-8       9.386n ± 0%    7.881n ± 1%  -16.04% (p=0.000 n=20)
ExecutionStatusString_MiddleItem-8     8.936n ± 0%    7.981n ± 1%  -10.68% (p=0.000 n=20)
ExecutionStatusString_FirstItem-8      8.158n ± 0%    5.444n ± 0%  -33.27% (p=0.000 n=20)
geomean                                2.934n         1.919n       -34.60%
```

### Enumer vs Idiomatic

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
                                    │  enumer.txt  │            idiomatic.txt             │
                                    │    sec/op    │    sec/op     vs base                │
ExecutionStatus_String_LastItem-8     0.9767n ± 0%   0.6259n ± 0%  -35.92% (p=0.000 n=20)
ExecutionStatus_String_MiddleItem-8   0.9767n ± 2%   0.6257n ± 0%  -35.94% (p=0.000 n=20)
ExecutionStatus_String_FirstItem-8    0.9778n ± 1%   0.6259n ± 0%  -35.99% (p=0.000 n=20)
ExecutionStatusString_LastItem-8       9.386n ± 0%    2.724n ± 0%  -70.98% (p=0.000 n=20)
ExecutionStatusString_MiddleItem-8     8.936n ± 0%    2.044n ± 0%  -77.12% (p=0.000 n=20)
ExecutionStatusString_FirstItem-8      8.158n ± 0%    1.912n ± 0%  -76.56% (p=0.000 n=20)
geomean                                2.934n         1.173n       -60.01%
```

### Enumer vs Combined

```console
goos: darwin
goarch: arm64
pkg: github.com/nrxr/enumbenchmarks/enumer/executionstatus
                                    │  enumer.txt  │             combined.txt             │
                                    │    sec/op    │    sec/op     vs base                │
ExecutionStatus_String_LastItem-8     0.9767n ± 0%   0.5771n ± 9%  -40.91% (p=0.000 n=20)
ExecutionStatus_String_MiddleItem-8   0.9767n ± 2%   0.5235n ± 0%  -46.41% (p=0.000 n=20)
ExecutionStatus_String_FirstItem-8    0.9778n ± 1%   0.5233n ± 2%  -46.48% (p=0.000 n=20)
ExecutionStatusString_LastItem-8       9.386n ± 0%    2.725n ± 0%  -70.97% (p=0.000 n=20)
ExecutionStatusString_MiddleItem-8     8.936n ± 0%    2.043n ± 2%  -77.14% (p=0.000 n=20)
ExecutionStatusString_FirstItem-8      8.158n ± 0%    1.915n ± 2%  -76.53% (p=0.000 n=20)
geomean                                2.934n         1.091n       -62.82%
```
