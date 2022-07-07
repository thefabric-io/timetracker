# Time tracker

Time tracker is a lightweight utility to track the duration of a scoped process.

## Usage

```go
package main

import (
	"github.com/thefabric-io/timetracker"
	"time"
)

func main() {
	defer timetracker.Log(time.Now(), "process name")

	time.Sleep(2 * time.Second)
}
```
