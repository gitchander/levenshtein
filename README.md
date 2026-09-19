# Levenshtein

An implementation of the Levenshtein distance for Go.

## Sources

- [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)

- [Go Programming Language](https://golang.org/)

## Installation

```bash
go get github.com/gitchander/levenshtein
```

Examples
--------

Basic string distance
---------------------

```go
package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = "exponential"
		b = "polynomial"
	)
	distance := lev.Strings(a, b)
	fmt.Printf("the levenshtein distance = %d\n", distance)
}
```

result:
```
the levenshtein distance = 6
```

Distance between word slices
----------------------------
```go
package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		line1 = "one two three four"
		line2 = "one two three"
	)
	var (
		a = strings.Fields(line1)
		b = strings.Fields(line2)
	)
	distance := lev.StringSlices(a, b)
	fmt.Printf("the levenshtein distance = %d\n", distance)
}
```

result:
```
the levenshtein distance = 1
```

Printing the distance matrix
----------------------------
```go
package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = []rune("sitting")
		b = []rune("kitten")
	)
	costs := lev.DefaultCosts
	fmt.Print(lev.PrintableMatrix(a, b, costs, ""))
}
```

result:
```
. . k i t t e n 
. 0 1 2 3 4 5 6 
s 1 1 2 3 4 5 6 
i 2 2 1 2 3 4 5 
t 3 3 2 1 2 3 4 
t 4 4 3 2 1 2 3 
i 5 5 4 3 2 2 3 
n 6 6 5 4 3 3 2 
g 7 7 6 5 4 4 3 
```

Using custom types via the Interface
------------------------------------
```go
package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

type Person struct {
	Name string
	Age  int
}

type PersonPair [2][]Person

var _ lev.Interface = PersonPair{}

func (p PersonPair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p PersonPair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

func main() {
	var (
		a = []Person{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}}
		b = []Person{{"one", 1}, {"two", 2}, {"three", 3}}
	)
	distance := lev.Distance(PersonPair{a, b})
	fmt.Printf("the levenshtein distance = %d\n", distance)
}
```

result:
```
the levenshtein distance = 1
```
