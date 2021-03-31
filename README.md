# redis-proto-vs-json

Benchmark of storing `proto` and `json` on redis 🥊

> `docker-compose up`
> 
> `go test ./... -bench=.`

```
pkg: github.com/benka-me/redis-proto-vs-json/src/bch
BenchmarkProto-16            733           1807597 ns/op
--- BENCH: BenchmarkProto-16
    bench_test.go:60: proto len:  46
    bench_test.go:60: proto len:  4600
    bench_test.go:60: proto len:  33718
BenchmarkJson-16             698           1817650 ns/op
--- BENCH: BenchmarkJson-16
    bench_test.go:76: json len:  130
    bench_test.go:76: json len:  13000
    bench_test.go:76: json len:  90740
PASS
```


