

## Setup

```bash
mkdir hello
cd hello

#enable dependency tracking
go mod init github.com/hello
```

## Simple code
### Writing very basic code
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Note: a package is a way to group functions, and it's made up of all the files in the same directory


### Run the code
```bash
go run
# or
go run main.go
```

## Calling external code

Visit https://pkg.go.dev and check for new modules and versions. For now, import rsc.io/quote package 

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

Adding new module requirements and sums, Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module

```bash
go mod tidy
```

Then,
```bash
go run
# or
go run main.go
```

Adding an specific dependency is simple as:

```bash
# Add all dependencies for a package in te module
go get . 

# or add an specific dependency
go get example.com/newmodule
```
