# grpc-example
A simple gRCP example project

## Tests & benchmarks

Run all tests:
```bash
go test ./... 
```
Run all benchmarks
```bash
go test -bench=. ./... -check.b
```

Run client tests:
```bash
go test ./... -check.f ClientSuite
```