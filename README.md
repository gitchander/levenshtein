# levenshtein
An implementation of the Levenshtein distance for Golang

Sources
-------

[Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)

[Golang](https://golang.org/)

Install
-------

```
go get github.com/gitchander/levenshtein
```

Examples
--------

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
	fmt.Printf("the levenshtein distance between %q and %q = %d\n", a, b, distance)
}
```

result:
```
the levenshtein distance between "exponential" and "polynomial" = 6
```

Example the distance by words:
------------------------------
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
	distance := lev.Distance(lev.StringSlices{a, b})
	fmt.Printf("the levenshtein distance between %q and %q = %d\n", a, b, distance)
}
```

result:
```
the levenshtein distance between ["one" "two" "three" "four"] and ["one" "two" "three"] = 1
```

Example with print matrix:
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
	fmt.Print(lev.PrintableMatrix(costs, a, b, ""))
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

Example the distance by interfaces:
------------------------------
```go
package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = []Person{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}}
		b = []Person{{"one", 1}, {"two", 2}, {"three", 3}}
	)
	distance := lev.Distance(PersonSlices{a, b})
	fmt.Printf("the levenshtein distance = %d\n", distance)
}

type Person struct {
	Name string
	Age  int
}

type PersonSlices [2][]Person

var _ lev.Interface = PersonSlices{}

func (p PersonSlices) Len(k int) int {
	return len(p[k])
}

func (p PersonSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}
```

result:
```
the levenshtein distance = 1
```
