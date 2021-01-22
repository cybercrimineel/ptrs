# ptrs

This package is based on [go-pointer](https://github.com/mattn/go-pointer/).
Using the allocation implementation of the C runtime to create references to Go
data structures can be expensive. The realization that these references can be
arbitrary as long as the Go runtime can keep track, made it possible for me to
create this alternative.

This package is designed to be more memory efficient and its API is meant to
feel more familliar to people that know C. The algorithm used by this package
has a constant time complexity.
