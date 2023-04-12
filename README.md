# inmemcache

A simple, thread-safe, in-memory cache implementation in Go.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Creating a Cache](#creating-a-cache)
  - [Setting a Value](#setting-a-value)
  - [Getting a Value](#getting-a-value)
  - [Deleting a Value](#deleting-a-value)
- [License](#license)

## Installation

To use inmemcache, simply import the package in your Go project:

\```bash
go get github.com/username/inmemcache
\```

Replace `username` with your GitHub username if you have forked this repository.

## Usage

### Creating a Cache

First, import the `inmemcache` package in your Go project:

\```go
import "github.com/username/inmemcache"
\```

To create a new cache, use the `NewCache()` function:

\```go
cache := inmemcache.NewCache()
\```

### Setting a Value

To set a value in the cache, use the `Set(key, value)` method:

```go
cache.Set("key", "value")
```

`key` must be a `string`, and `value` can be any `interface{}` type.

### Getting a Value

To retrieve a value from the cache, use the `Get(key)` method:

```go
value := cache.Get("key")
```

If the key is not found in the cache, the method will return `nil`.

### Deleting a Value

To delete a value from the cache, use the `Delete(key)` method:

```go
cache.Delete("key")
```

This will remove the key-value pair associated with the specified key from the cache.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
