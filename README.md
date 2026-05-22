````md
# greeting

Simple Go package for Indonesian greetings based on current time.

## Installation

```bash
go get github.com/johnylie/gogreeting
````

## Usage

```go
package main

import (
	"fmt"

	"github.com/johnylie/gogreeting"
)

func main() {
	fmt.Println(greeting.Selamat("NAME"))
}
```

## Example Output

```text
Selamat pagi, NAME
Selamat siang, NAME
Selamat sore, NAME
Selamat malam, NAME
```

## License

MIT

```
```
