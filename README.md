
Run all tests & benchmarks:

```bash
go test -v -bench=. ./...
```

```bash
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel run //:gazelle_test
```
