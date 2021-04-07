## Functions

A function's *signature* refers to its type, composed of the parameter types and the result types (obeying sequence). A sequence of parameters or results can use the declaration shorthand if they are the same type. Go does not support default parameters and requires that all parameters are provided by a function call's arguments. All arguments will be passed by value to the function. Since we have reference types, this does not mean that local variables from the caller cannot not be mutated.

```go
func name(parameter-list) (result-list) {
	// body
}

// these functions have the same signature
func myFunc1(x int, y int) int { // .. }
func myFunc2(x, y int) int { // .. }
```

Function results can even be named. If they are named, then each name declares a local variable initialized to the type's zero value. We can complement this with a *bare return*, in which we do not specify the values to be returned with the `return` statement. Go knows to implicitly return the local variables of the function's named results.

```go
func main() {
	list := []int{1,2,3,4,5}
	fmt.Println(sumInts(list)) // "15"
}

// sum initialized to 0 before function executes
func sumInts(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return // example of a bare return
}
```

### Recursion
Recursion is typically limited by the function call stack since it is a fixed size in many languages. Go on the contrary has a variable-size stack that grows as needed. This means that we typically don't have to worry about stack overflows.

### Errors
Go functions handle errors by providing some indicitive result to the caller. This is typically the last parameter. Sticking to conventions, the last result should be a boolean `ok` if there was only one possible cause of failure. If there are multiple causes of failure, then the type should be of the `error` interface (interfaces will be covered in a later chapter).

Convention has us handling errors inside an `if` statement and the success handling code being minimally indented (i.e. not inside an `else`).

```go
if val, ok := someFunction(); !ok {
	// error handling code
}
// success handling code

if err := someOtherFunction(); err != nil {
	// error handling code
}
// success handling code
```

One error that often needs its own distinction is EOF when reading inputs. It is declared as `io.EOF` and may be used as follows:

```go
in := bufio.NewReader(os.Stdin)
for {
	r, _, err := in.ReadRune()
	if err == io.EOF {
		break // finished reading
	}
	if err != nil {
		// error handling code
	}
	// use r..
}
```

### Function Values
In Go, functions have a type and as such may be assigned to variables or passed between functions. Functions have a zero value of nil (which causes a panic if called). Function values may be compared with `nil` but are not themselves comparable and cannot be compared against other function values or used as keys in a map.

### Anonymous Functions
We also have *function literals* to denote a function within any expression. The value of said functions is called an *anonymous function* since it does not have a name.

```go
// not a great example but still shows how anonymous functions work
func performOperation(x, y int, operation func(a, b int) int) int {
	return operation(x, y)
}

func main() {
	// here we provide an anonymous function as an argument
	result := performOperation(6, 8, func(a, b) int {
		return a * b
	}))
	fmt.Println(result) // "48"
}
```

// TODO closures

### Variadic Functions

### Defer

### Panic & Recovery
