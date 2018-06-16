# Validation Benchmark

Spec: AMD Ryzen 7 2700x stock clock, RAM DDR4 3000 on Windows 10 Pro

vgo test -bench Simple -benchmem
```
goos: windows
goarch: amd64
pkg: github.com/CJ-Jackson/scratchpad/benchmark/validation
BenchmarkValidationSimple-16                    20000000               114 ns/op              16 B/op          1 allocs/op
BenchmarkGoValidatorSimple-16                     500000              3071 ns/op             352 B/op         14 allocs/op
BenchmarkSchemaSimple-16                         2000000               971 ns/op             216 B/op         10 allocs/op
BenchmarkValidationSimpleWithError-16            2000000               652 ns/op             600 B/op         11 allocs/op
BenchmarkGoValidatorSimpleWithError-16            500000              3331 ns/op             720 B/op         15 allocs/op
BenchmarkSchemaSimpleWithError-16                1000000              1087 ns/op             264 B/op         12 allocs/op
PASS
ok      github.com/CJ-Jackson/scratchpad/benchmark/validation   11.821s
```

vgo test -bench Complex -benchmem
```
goos: windows
goarch: amd64
pkg: github.com/CJ-Jackson/scratchpad/benchmark/validation
BenchmarkValidationComplex-16                    1000000              1247 ns/op              80 B/op          6 allocs/op
BenchmarkGoValidatorComplex-16                    100000             14520 ns/op            1745 B/op         61 allocs/op
BenchmarkSchemaComplex-16                         500000              3665 ns/op             656 B/op         36 allocs/op
BenchmarkValidationComplexWithError-16            500000              2696 ns/op            2184 B/op         40 allocs/op
BenchmarkGoValidatorComplexWithError-16           100000             14630 ns/op            1897 B/op         64 allocs/op
BenchmarkSchemaComplexWithError-16                500000              3183 ns/op             656 B/op         33 allocs/op
PASS
ok      github.com/CJ-Jackson/scratchpad/benchmark/validation   9.497s
```