## universally unique identifier v4

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/uuid)](https://goreportcard.com/report/jancajthaml-go/uuid)

Algorithm generate 128bit UUID with cryptographically secure PRNG in format 8-4-4-4-12 lowercase hex digits.

### Usage ###

```
import "github.com/jancajthaml-go/uuid"

uuid.Generate()
```

### Performance ###

- 48 B/op
- 1 allocs/op

verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Universally Unique IDentifier](https://en.wikipedia.org/wiki/Universally_unique_identifier)
