# striprtf

> This package is based on [IntelligenceX/fileconversion](https://github.com/IntelligenceX/fileconversion/blob/1b64e2d06acedecada44aff7a3942bef5b811409/RTF%202%20Text.go) and [J45k4/rtf-go](https://github.com/J45k4/rtf-go).

`striprtf` is a Go package designed to extract plain text or HTML content from RTF (Rich Text Format) documents, removing all formatting information.

## Installation

To install the package, use the `go get` command:

```bash
$ go get github.com/attilabuti/striprtf@latest
```

## Usage

Here's a basic example of how to use the `striprtf` package:

### Extracting text from RTF

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/attilabuti/striprtf"
)

func main() {
    file, err := os.Open("document.rtf")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    r, err := striprtf.ExtractText(file)
    if err != nil {
        panic(err)
    }

    text, err := io.ReadAll(r)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(text))
}
```

### Extracting HTML from RTF

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/attilabuti/striprtf"
)

func main() {
    file, err := os.Open("document.rtf")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    r, err := striprtf.ExtractHtml(file)
    if err != nil {
        panic(err)
    }

    html, err := io.ReadAll(r)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(html))
}
```

## Issues

Submit the [issues](https://github.com/attilabuti/striprtf/issues) if you find any bug or have any suggestion.

## Contribution

Fork the [repo](https://github.com/attilabuti/striprtf) and submit pull requests.

## License

This extension is licensed under the [MIT License](https://github.com/attilabuti/striprtf/blob/main/LICENSE).