# [Go (Golang)](https://golang.org)


## Quick Start

To install: Download the .pkg [here](https://golang.org/dl/)

### To set up environment
```bash
$ mkdir golang
$ export GOPATH=[full_path_to_dir]/golang
```

###Recommended folder structure
**bin/** and **src/**<br>
**bin/** will containe the compiled files<br>
**src/** should have one package file per folder

Go files has file ending **.go**<br>
To compile the files:<br>
```bash
$ go install folder_path*
```
*folder_path is the path from src/, do not include src/ in the folder path


To run the compiled file:
```bash
$ bin/folder_path
```

## What is Golang

Go, also commonly referred to as golang, is a programming language developed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a statically typed language with syntax loosely derived from that of C, adding garbage collection, type safety, some structural typing capabilities, additional built-in types such as variable-length arrays & key-value maps, and a large standard library.
Android support was added in version 1.4, which has since been ported to also run on iOS. - [Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language))<br>

## Resources
###[Interactive tutorial](https://tour.golang.org/welcome/1)
###[Repl.it like editor](http://play.golang.org/)


## Cheat Sheet
### Go app skeleton:
```
packages main

import (
	"fmt"		// e.g. fmt.Println(), optional
	"math"	// e.g. math.Sqrt(), optional
	"time"	// e.g. time.Sleep(), optional
)

func main() {
	
}
```
### Variables:
variables can be declared with specified type
```
var i, j int = 1, 4
```
or implicitly with :=
```
i := 1
j := 4
```
Constants can not be declared with :=
```
const Pi = 3.14
```
**Basic types**
```
bool
string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr
rune is an alias for int32
uint is at least 32 bits but not an alias for uint32
float32, float64
complex64, complex128
```
Default values:
```
0 for numeric
false for bool
"" for string
```

### Functions:
functions defines its arguments types:<br>
```
func add(x int, y int) int { // the last int defines the return type
	return x+y
}
```

functions can return multiple results
```
a, b := split(sum)
```

functions can also return naked results
```
func split(sum int) (x, y int) {
	x = sum * .4
	y = sum - x
	return
}
```

### For loops:
```
for i := 0; i < 10; i++ {
	sum += i
}
```
shorthand:
```
for ; sum < 10 ; {
	sum += sum
}
```
which is basically a while loop:
```
for sum < 10 {
	sum += sum
}
```
### If statements:
```
if x < 0 {
	
} else {
	
}
```
if statments can execute code
```
if x:=3; x < y {
	
}
```
x is only accessible inside the scope of the if statement including subsequent else etc.

### Concurrency
Goroutine is a lightweight thread managed by Go runtime<br>
The below function will run say() as two different goroutines in the same address space
```
func main() {
	go say("string")
	say("other string")
}
```
**Channels**
Channels behave like an array that can only be pushed to and shift from. First in is first out. <br>
To make a channel:
```
ch := make(chan int)
```
To send to a channel:
```
ch <- some_value
```
To receive from channel:
```
variable := <- ch
```
**Switch**
 Evaluate from top to bottom and stops when a case is successful, or chooses the fallthrough/default case.

```
func main() {
	fmt.Println("When's Saturday?")
	
	//implicitly declares today as today's day
	today := time.Now().Weekday()
	
	//takes the time of the day "Saturday"
	switch time.Saturday {
	//cheks each case
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
		
	//if none of the cases are succesful it chooses the default case
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(time.Now().Weekday() + 2)
}
```




