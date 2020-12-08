# Go Overview

## Resources

- Much of this overview is taken from [A Tour of Go](https://tour.golang.org/welcome/1)
- [Go Time Podcast](https://changelog.com/gotime)

## Go as a language

- Originally developed by Google to address critisisms of other languages. The designers were primarily motivated by their shared dislike of C++.
- Statically typed
- Similar to C (but memory safe, with garbage collection and structural typing)
- Great for concurrency
- v1.0 was publically released in March 2012

## Why I think you should learn (and use) Go

- Readability is high (i.e. it's easy to reason what a program is meant to do)
- Cross platform - Write code once, compile to your target and run.
- It's easy to learn and use, but is still powerful
  - If you have some programming experience you could learn the syntax in about a day and start being productive in it right after.
  - Like any language, you'll become more productive as you learn the standard library and other popular packages
- The community is awesome and is growing every day
- Makes concurrency easy (easier) - but be careful not to go crazy with this at first
  - Concurrency is also easier to reason about (IMO) than that of dotnet's async-await
- Write your test files along side you package files (no testing framework needed). Test package is included in the standard library.

## What is Go good for?

Go is often called the "Devops" language, but that's a myth (though there is a ton of devops tooling written in Go). It's good for just about anything you want to do.

- Web Backends
  - Helpful Packages:
    - [buffalo](https://gobuffalo.io/en/)
    - [gorilla](https://github.com/gorilla)
    - [gin](https://github.com/gin-gonic/gin)
- Cli Tools
  - Docker, K8s, and many more are written in Go
  - Helpful Packages
    - [cobra](https://github.com/spf13/cobra)
- MicroControllers
  - [tiny go](https://tinygo.org/) - Go for small spaces
- Web Assembly
  - [Getting started with WASM using Go](https://github.com/golang/go/wiki/WebAssembly)
- "Serverless"
  - AWS Lamda, GCP Functions, others
  - Some Links
    - [AWS Lamda](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
    - [Google Cloud Platform Functions](https://cloud.google.com/functions/docs/concepts/go-runtime)
- Anywhere that you need a scripting language
  - Go can be compiled to target nearly any operating system, so it's perfect for scripting.
- Just about anything you can think of.

## What does a Go program look like?

- Go programs are made up of packages
- Programs start running in main()
- Builds produce a single binary

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, SDG!")
}
```

### Helpful tips

- Packages named after folders (typically)
  - Files in the same folder are all a part of the same package (unless it is named [package name]\_test)
- Downloading packages
  - `go get <package name>`
- Running a go program
  - `go run <package name>.go`
  - executable programs start in main, so the file could be server.go, main.go, foobar.go, etc. so long as it has the main function
  - typically main.go
- Building binary
  - `go build .`
- Running a binary
  - `.\<binary name>`
- Use modules for version control
  - `go mod init <name of your package>`
  - [Using Go Modules](https://blog.golang.org/using-go-modules)
- Errors, not exceptions

  - Go doesn't have exceptions, it has errors, which are explicitly returned.
  - IMO, this is better than try-catch-finally, because it forces you to acknowledge and handle or explicitly ignore the error where as languages that allow for exceptions can obfuscate the exceptions that are thrown if they are not documented.

    ```go
    package data

    import (
        "fmt"
        "github.com/pkg/errors"
    )

    // For custom errors, can use errors.New()
    // Good idea to export these so that they can be checked for

    // ErrSomeCustom is a custom error
    var ErrSomeCustom = errors.New("Some Custom Error")

    func getPersonById(id string) *person, error {
        person, err := db.GetPerson(id)
        if err != nil {
            return nil, errors.Wrap(err, fmt.Sprintf("Could not get person with id: %s from database", id))
        }
        return person, nil
    }
    ```

- Go doesn't have generics [yet](https://blog.golang.org/generics-next-step). Because of this an empty interface is often used in a function signature so that it can take in any type (every type meets the requirements of any empty interface)
  ```go
  func Print(a interface{}) (n int, err error){}
  func Println(a ...interface{}) (n int, err error){}
  ```
  - Hopefully we'll have generics soon.

## Useful Packages

- Database
  - [Migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4@v4.14.1?utm_source=gopls)
    - Useful for handling database migrations
- Web
  - [gqlgen](https://github.com/99designs/gqlgen)
    - Generate boilerplate code for GraphQL based on your schema
- Environment configuration
  - [godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.3.0?utm_source=gopls)
    - Useful for pulling in .env files
  - [envconfig](https://pkg.go.dev/github.com/kelseyhightower/envconfig@v1.4.0?utm_source=gopls)
    - Useful for decoding environment variables
- Testing
  - [dockertest](https://pkg.go.dev/github.com/ory/dockertest@v3.3.5+incompatible?utm_source=gopls)
    - Useful for running docker containers during test
    - I use this to spin up a postgres container, then run my migrations on it when testing my database access. Testing against the real thing is alway preferable to testing against some in memory version of it (to me anyways).
  - [require](https://pkg.go.dev/github.com/stretchr/testify@v1.6.1/require?utm_source=gopls)
    - Useful for assertions. There's also a package called assert that's exactly the same, except that require will stop your tests as soon as you have a failure.
  - [mock](https://pkg.go.dev/github.com/stretchr/testify@v1.6.1/mock?utm_source=gopls)
    - Useful for mocking.

# The Language

Note, this section has a lot of sudo code for brevity

## Syntax

- Go only has 25 keywords
  - C# has 110
  - JavaScript has 64
- Go has 45 operators and punctuation

### package, import, func

```go
// define package that file belongs to
package main

// import packages that source file depends on
import (
    "fmt"
    "math/rand"
)

// function declaration
func main() {
    fmt.Println("The random number is", rand.Intn(10))
}
```

### var, const

```go
// variable declaration
var y = 1
// or
x := 2

// constant declaration
const z = 3
```

### pointers

```go
package main

import "fmt"

func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	// The & operator generates a pointer to its operand.
	// The * operator denotes the pointer's underlying value.
	// Go has no pointer arithmetic.

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer - prints 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i - prints 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j - prints 73
}
```

### struct

```go
package main

import "fmt"

// Struct declaration
// A struct is a collection of fields.
type pet struct {
    name string
    age  int
}

func newPet(name string, age int) *pet {
    p := person{name: name, age: age}
    return &p
}

func main() {
    // How to create new structs:
    fmt.Println(pet{"Max", 5}) //{Max 5}
    fmt.Println(pet{name: "Buddy", age: 10}) //{Buddy 10}
    // Omited fields will be zero valued
    fmt.Println(pet{name: "Shadow"}) //{Shadow 0}
    // An & prefix yields a pointer to the struct.
    fmt.Println(&pet{name: "Honey", age: 14}) //&{Honey 14}

    // access struct fields with a dot
    fluffy := pet{name: "Fluffy", age: 7}
    fmt.Println(fluffy.name) //Fluffy

    // You can also use dots with struct pointers - the pointers are automatically dereferenced.
    fluffyPointer := &fluffy
    fmt.Println(fluffyPointer.age)

    // Structs are mutable
    fluffyPointer.age = 8
    fmt.Println(fluffyPointer.age) //8
}
```

### interfaces

```go
package messenger

// Note: Capitalized structs, interfaces, consts, vars and functions are exported for use in other packages. Exported fields should be commented.

// Mesage type
type Message struct {
    from    string
    to      string
    content string
}

// Interface declaration
// An interface type is defined as a set of method signatures.
// MessageService sends messages
type MessageService interface {
    SendMessage(msg message) error
}

// Example interface implementation:
// To implement an interface, implement all the methods in the interface
// Notice how the implementation is implicit, not explicit
type messageService struct {
    client *SES // AWS Simple Email Service
}

// Note: This is sudo code
func (ms *messageService) SendMessage(msg) error {
    err := ms.client.SendMessage(msg)
    return err
}
```

## Arrays, Slices, Maps

### Arrays

- The type [n]T is an array of n values of type T.

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1]) //Hello World
fmt.Println(a) //[Hello World]

primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes) //[2 3 5 7 11 13]
```

### Slices

- For a more in depth overview of slices, check out slices in [A Tour of Go](https://tour.golang.org/moretypes/7)
- The type []T is a slice with elements of type T.
- Slices are references to the underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.
- Form slice like this: a[low : high]

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]
fmt.Println(s) //[3 5 7]
```

```go
nums := []int{}
// append to a slice
nums = append(nums, 1)
nums = append(nums, 2, 3)
fmt.Println(nums) //[1 2 3]
```

- Slices have length and capacity
- Zero value of a slice is nil

### Maps

- Maps are similar to dictionaries
- They map keys to values
- Zero value of a map is nil
- The make function returns a map of the given type, initialized and ready for use.

```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// Using map literal
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// insert or update an element
m[key] = elem

// retrieve an element
elem = m[key]

// delete an element
delete(m, key)

// check that an element exists in the map
// If key is in m, ok is true. If not, ok is false.
elem, ok = m[key]
```

## Control Flow

### For loop

```go
// A familiar for loop (init, exit condition, post)
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
fmt.Println(sum) //45

// init and post statements are optional
// (i.e. for is go's while)
sum = 1
for sum < 1000 {
    sum += sum
}
fmt.Println(sum) //1024

// Loop forever by omiting the exit condition
for {
}

// Range
// Loop over a slice
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// i - index
// v - copy of the element in the array
for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
}
// Output:
// 2**0 = 1
// ... omitted
// 2**7 = 128

// If you only want the index:
i := range pow

// If you only want the value:
_, v := range pow
```

### If statment

```go
if x < 0 {
    // Do something
}

if x < 0 {
    // Do x < 0 stuff
} else {
    // Do x >= 0 stuff
}

if x < 0 {
    // Do x < 0 stuff
} else if x == 0 {
    // Do x == 0 stuff
} else {
    // Do x > 0 stuff
}
```

### Switch

```go
// Switch does top down evaluation, stopping when case succeeds
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
}
```

### defer

- A defer statement defers the execution of a function until the surrounding function returns.

```go
func main() {
    defer fmt.Println("world")
	fmt.Println("hello")
}

// Output:
// hello
// world
```

- defers can be stacked

```go
func main() {
    fmt.Println("counting")
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("done")
}

// Output:
// counting
// done
// 2
// 1
// 0
```

### Panic and Recover

- Panic is used to create a run time error
- Handle run time panics with the built-in recover function

```go
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// Output:
// PANIC
```

## Functions (in depth)

### Simple Function

```go
func main(){
    fmt.Println("Hello World")
}
```

### Function Arguments, returns

```go
func add(x, y int) int {
    return x + y
}
```

### Named Returns

Don't use this unless it's dead simple. It worsens readability

```go
func f2() (r int) {
  r = 1
  return
}
```

### Return multiple values

```go
func f() (int, int) {
  return 5, 6
}
```

### Variadic functions

Used to pass an indeterminante number of args

```go
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

// Note: This is how fmt.Println is implemented:
func Println(a ...interface{}) (n int, err error)
```

### Methods on a type

Note: Methods are functions

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

### Method using a pointer reciever

There are two reasons to use a pointer reciever

1. The first is so that the method can modify the value that its receiver points to.
1. The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs()) // Before scaling: &{X:3 Y:4}, Abs: 5
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs()) //After scaling: &{X:15 Y:20}, Abs: 25
}
```

### Closure

```go
func main() {
  add := func(x, y int) int {
    return x + y
  }
  fmt.Println(add(1,1)) //2
}
```

```go
func main() {
	x := 0
	increment := func() {
		x++
	}
	increment()
	increment()
	fmt.Println(x) //2
}
```

### Functions as return types

```go
func makeHelloFunc() func() {
    return func() {
        fmt.Println("Hello")
    }
}

func main() {
	x := makeHelloFunc()
	x()
}

// Output:
// Hello
```

### Recursion

Functions can call themselves

```go
func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
```

## Go Routines

- A goroutine is a lightweight thread managed by the Go runtime.
  ```go
  go f(x, y, z) // starts a new go routine
  ```
  - Evaluation of f, x, y, z happen in the current goroutine and the execution of f happens in the new goroutine.

### Go Keyword

- Any function can become a go routine just by using the `go` keyword
- You can kind of think of `go` as Go's async. We'll look at some examples.

### Channels

- Channels are a typed conduit through which you can send and receive values with the channel operator, <-. They are useful for syncing and allowing different go routines to communicate.
- By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

```go
ch := make(chan int) // Make a channel of type int
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

### Buffered Channels

- Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
- Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

```go
ch := make(chan int, 100)
```

### Range and Close

- Senders can close (`close(c)`)a channel to indicate that no more values will be sent
- Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:

```go
v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.
```

- The loop for i := range c receives values from the channel repeatedly until it is closed.

```go
for i := range c {
    fmt.Println(i)
}
```

### Select

- The select statement lets a goroutine wait on multiple communication operations.
- The default case in a select is run if no other case is ready.

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
            return
        default:
            // Runs if no other case is ready
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### Directional Channels

- Chanels can be send only and recieve only
- `chan<- T` send only
- `<-chan T` recieve only
- If you can't remember which is which, look at the arrow.
  - Pointing into `chan` - send
  - Pointing away from `chan` - recieve
- Useful for when you want to be sure that a function can only read or write (not both) to a channel

```go
func SendOnly(pings chan<- string) {
// pings is a send only chan
}

func RecieveOnly(pongs <-chan string) {
// pongs is a recieve only chan
}
```
