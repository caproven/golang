# Chapter 2 - Program Structure

## Names
Identifiers for functions, variables, constants, types, statement labels, and packages all have the same requirement: they must begin with a letter or an underscore. This may be followed by any number of letters, digits, and underscores. Go programmers typically use camelCase for combining words. For acronyms, try to keep all letters rendered in the same case (ex. HTML).

There are a fair amount of predeclared names such as int, nil, false, etc. These are not reserved and may be redeclared (should you want to do so).

If a name begins with an uppercase letter, it is exported and may be accessed from outside the source package.

## Declarations
There are four major types of declarations: var, const, type, and func.

Each .go file begins with a package declaration, any import declarations, and a sequence of package-level declarations. Package-level declarations may be in any order (unlike C).

```go
package main

import "fmt"

const secret = 25.0

func divideBySecret(val float64) float64 {
	return val / secret
}

func main() {
	fmt.Printf("7 / secret = %f\n", divideBySecret(7))
}
```

## Variables
Variable declarations have the general form `var name type = expression`. Either the type of the expression may be omitted, but not both. If the type is omitted, the type is inferred from the given expression. If the expression is omitted, the variable will be given an initial value of the type's 'zero value'. This concept of zero values means that variables will never be uninitialized.

| Type | Zero Value |
| ---- | ---------- |
| number | 0 |
| bool | false |
| string | "" |
| interface | nil |
| reference | nil |

There is also short variable declaration which may be used for initializing local variables that takes the form `name := expression`. The type of name is determined by the type of expression. Short variable declaration doesn't always declare the variables on the left hand side, meaning that it also works for reassignment. A caveat here is that at least one new variable must be declared on the left hand side.

```go
var a, b, c int // 0, 0, 0
var x, y, z = false, "a string", 3.14
f,g := 5, 6.0
```

## Pointers
Go has support for pointers that function identically to pointers in C. Pointers are a reference to the memory address of a value rather than the value directly. Just like in C, pointers are denoted with `*` while a pointer can be taken for any value by using the address-of operator `&`. Go does NOT support pointer arithmetic, so it could not be used for things such as iterating over an array. This was an intentional decision that promotes safety.

```go
x := 5 // int
myPointer := &x
fmt.Println(*myPointer) // "5"
```

The `new` function may be used as a shorthand for creating a type instance and taking its address. `new(T)` will create an unnamed variable, initialize it to the zero value of type T, and return the address.

```go
// the following
intPtr1 := new(int) // type *int
// is the same as
var myInt int
intPtr2 := &myInt // type *int
fmt.Printf("%T %d %T %d\n", intPtr1, *intPtr1, intPtr2, *intPtr2) // "*int 0 *int 0"
```

## Assignment
Assignment in Go is about what you would expect from other languages. On top of the standard assignment operator `=`, we also have assignment operators for each of the arithmetic and bitwise operators (ex. `+=, *=, <<=`).

Additionally, we have tuple assignment which permits the assigning of several variables at once. This can be utilized for things such as easily swapping values.

```go
x, y, z := 1, 2, 3
fmt.Println(x, y, z) // "1 2 3"
// tuple assignment!
x, y, z = z, y, x
fmt.Println(x, y, z) // "3 2 1"
```

Here is an example of tuple assignment being put to use to calculate the fibonnaci series.

```go
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
```

While they are brought up later, functions may return more than one value. By combining this with tuple assignment, we can easily break down the returns from a function in the following way:

```go
// err will be nil if nothing went wrong
file, err = os.Open("foo.txt")
```

Many functions in the standard library will return this optional err value. On top of this, some operators in Go will return a bool `ok` to state if anything went wrong. The operators are:
```go
v, ok = m[key]  // map lookup
v, ok = x.(T)   // type assertion
v, ok = <-ch    // channel receive
```

## Type Declarations
A type declaration defines a new *named type* that has the same *underlying type* as an existing type. This can help differentiate between uses for the same underlying type (as in the example below) and shorten the names for complex structures such as structs which will be discussed later. Type declarations take the form of `type name underlying-type`.

```go
type Feet float64
type Meters float64

const (
	HeightF Feet = 5.93
	WingspanM Meters = 2.1
)

// cannot combine different types, causes compiler error
fmt.Println(HeightF + WingspanM)
```

As mentioned before, the above example has the benefit of differentiating between the measurements of feet and meters, even though they have the same underlying type of float64. If we had variables of both types, the compiler would prevent us from accidentally interacting the two types (ex. adding Feet and Meters). An explicit cast or conversion would be required, helping make sure the programmer intended for such an interaction between types.
