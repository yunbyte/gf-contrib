# drivers

Powerful database drivers for package gdb.

# Installation

Let's take `pgsql` for example.

```shell
go get -u github.com/yunbyte/gf-contrib/drivers/pgsql/v2
```

Choose and import the driver to your project:

```go
import _ "github.com/yunbyte/gf-contrib/drivers/pgsql/v2"
```

Commonly imported at top of `main.go`:

```go
package main

import (
	_ "github.com/yunbyte/gf-contrib/drivers/pgsql/v2"

	// Other imported packages.
)

func main() {
	// Main logics.
}
```
