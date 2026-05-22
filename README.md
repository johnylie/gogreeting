````md
# gogreeting

Simple and customizable Go package for time-based greetings.

`gogreeting` automatically returns greeting messages based on the current local time.

## Features

- 🇺🇸 English support
- 🇮🇩 Indonesian support
- Custom greeting translations
- Personalized greetings with names

---

## Installation

```bash
go get github.com/johnylie/gogreeting
````

---

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/johnylie/gogreeting"
)

func main() {
	message := gogreeting.Greeting(
		"NAME",
		gogreeting.ID,
		nil,
	)

	fmt.Println(message)
}
```

Example output:

```text
Selamat pagi, NAME
```

---

## Supported Languages

| Language   | Constant        |
| ---------- | --------------- |
| English    | `gogreeting.EN` |
| Indonesian | `gogreeting.ID` |

---

## English Example

```go
message := gogreeting.Greeting(
	"NAME",
	gogreeting.EN,
	nil,
)

fmt.Println(message)
```

Possible output:

```text
Good evening, NAME
```

---

## Indonesian Example

```go
message := gogreeting.Greeting(
	"NAME",
	gogreeting.ID,
	nil,
)

fmt.Println(message)
```

Possible output:

```text
Selamat malam, NAME
```

---

## Custom Translation Example

```go
custom := &gogreeting.Translation{
	Morning: "Halo pagi",
}

message := gogreeting.Greeting(
	"NAME",
	gogreeting.ID,
	custom,
)

fmt.Println(message)
```

Possible output:

```text
Halo pagi, NAME
```

---

## Time Rules

| Time Range    | Greeting  |
| ------------- | --------- |
| 05:00 - 10:59 | Morning   |
| 11:00 - 14:59 | Afternoon |
| 15:00 - 17:59 | Evening   |
| 18:00 - 04:59 | Night     |

---

## API

### Greeting

```go
func Greeting(
	name string,
	lang Language,
	custom *Translation,
) string
```

Returns greeting text based on the current local time.

### Parameters

| Parameter | Type           | Description                 |
| --------- | -------------- | --------------------------- |
| `name`    | `string`       | Name of the person          |
| `lang`    | `Language`     | Greeting language           |
| `custom`  | `*Translation` | Optional custom translation |

---

## Translation Structure

```go
type Translation struct {
	Morning   string
	Afternoon string
	Evening   string
	Night     string
}
```

---

## License

MIT

---

## Author

👤 Johny Lie

* GitHub: https://github.com/johnylie
* Package: https://pkg.go.dev/github.com/johnylie/gogreeting

---

## 🌱 Show your support

Please ⭐️ this repository if this project helped you!

[![](https://img.shields.io/static/v1?label=Sponsor\&message=%E2%9D%A4\&logo=GitHub\&color=%23fe8e86)](https://github.com/sponsors/johnylie)
