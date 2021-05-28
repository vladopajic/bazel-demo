# bazel-demo

Bazel [installation guide](https://docs.bazel.build/versions/4.1.0/install.html)

Run target:
``` 
bazel run //01-hello:01-hello
```

Run tests
```
// run all tests inside workspace
bazel test //... 

// run all test from pacage
bazel test //01-hello/...
```
