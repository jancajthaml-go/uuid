## universally unique identifier v4

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/uuid)](https://goreportcard.com/report/jancajthaml-go/uuid)

Algorithm generate 128bit UUID with cryptographically secure PRNG in RFC4122 format.

### Usage ###

```
import "github.com/jancajthaml-go/uuidv4"

uuid.Generate()
```

### Performance ###

- 48 B/op
- 1 allocs/op

verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Universally Unique IDentifier](https://en.wikipedia.org/wiki/Universally_unique_identifier)
