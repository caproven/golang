# Chapter 3 - Basic Data Types

Go has four categories of data types: *basic*, *aggregate*, *reference*, and *interface*.

| Category | Example Types |
| -------- | ------------- |
| Basic | numbers, string, bool |
| Aggregate | arrays, struct |
| Reference | pointers, slices, maps, functions, channels |
| Interface | ??? |

Below we will go over the basic data types. It is worth noting that all basic type values are comparable (==, !=). Even though all basic types are comparable, all binary arithmetic and logic operators (excluding shifts) require operands of the same type. This means that explicit casts must be used to prevent type mismatches.

Type conversions to an arbitrary type *T* may be done with the operation `T(x)`. This conversion may result in changed values or lost precision.

```go
var myInt int64 = 5
var myFloat float64 = 5.0

x := myInt + myFloat // this causes a type mismatch
y := float64(myInt) + myFloat // this works
```

## Integers
There are both signed and unsigned variants for integer types. For these, we have the *int* and *uint* types. These will default to either 32 or 64 bits, based on the compiler's decision.

| Signed | Unsigned |
| ------ | -------- |
| int8 | uint8 (byte) |
| int16 | uint16 |
| int32 (rune) | uint32 |
| int64 | uint64 |

## Floating-Point
There are only two sizes of floating-point numbers in Go: *float32* and *float64*. While you have the option of using float32, float64 is preferred for its much better precision.

A floating-point value can represent an invalid number, *NaN*. This value can be accessed by `math.NaN()` and can be checked for with `math.isNaN()`. the `math.isNaN()` function should be used to test values since the actual *NaN* value always returns false when used in comparisons with itself.

```go
// BAD - this will always return false
func test(val float64) bool {
	return val == math.NaN()
}
```

## Complex Numbers
There are two sizes of complex numbers, *complex64* and *complex128*. These are composed of *float32* and *float64* components respectively, for both the real and imaginary parts of the represented complex numbers. These components can be extracted with the `real()` and `imag()` built-in functions.

```go
var a complex128 = complex(5, 10) // 5+10i
b := 8+9i

fmt.Println(real(a)) // "5"
fmt.Println(imag(b)) // "9"
```

## Booleans
Booleans in Go are the same as everywhere else and still support short circuiting with logical operations.

```go
x := 5
y := 6
z := 7
// the 'x == z' condition isn't evaluated since 'x < y' already makes the || operation true
if x < y || x == z {
	// do something
}
```

## Strings
Strings are immutable sequences of bytes. Conventionally, Go will interpret strings as being UTF-8 encoded. The actual data stored is then a sequence of Unicode code points (the actual codes for characters, such as \u0041) which are also called runes (see the table under the section on Integers). Being UTF-8 encoded, these runes are only encoded to take up as much space as necessary to specify the Unicode code point, between 1 and 4 bytes.

A substring operation `s[i:j]` can be used to yield a new string between the bytes [i,j) of the string s. Either i or j may be ommitted, with the defaults for `[i:j]` being `[0:len(s)]`.

Comparisons of strings are done with the standard operators (==, <, >, etc.). The comparison is done byte by byte, so the result is the natural lexicographic ordering.

```go
// a string literal
var s string = "i am a string literal"

// a raw string literal
var s2 string = `i am a raw string literal. I won't process escape sequences!\n`
// printing s2 would result in "\n" being directly printed along with the other text
```

### Runes
Go does not contain a 'character' type per say,  but does contain the `rune` alias for `int32`. A shortcut for rune initialization is *rune literals* which look similar to character constants in other languages.

```go
r := '魚'
fmt.Printf("%T\n", r) // "int32"
```


### Unicode & UTF-8
As mentioned before, Go uses UTF-8 encoding for strings which contain Unicode code points that take up 1-4 bytes. Because of this, there are several things we must do differently. An example of this is finding the length of a string. The built-in `len()` function will return the number of bytes in a string, not the number of runes. The `unicode/utf8` library contains many utility functions that can help, such as the one used below.

```go
s := "Hello, 世界"
fmt.Prinln(len(s)) // "13", not what we want
fmt.Println(utf8.RuneCountInString(s)) // "9", much better
```

Range loops over strings will iterate over the decoded UTF-8 runes and return the indices of the bytes where characters start.

```go
s := "Hello, 世界"
for i, r := range s {
	fmt.Printf("Idx: %d\tRune: %q\n", i, r)
}
/*
Idx: 0	Rune: 'H'
Idx: 1	Rune: 'e'
Idx: 2	Rune: 'l'
Idx: 3	Rune: 'l'
Idx: 4	Rune: 'o'
Idx: 5	Rune: ','
Idx: 6	Rune: ' '
Idx: 7	Rune: '世'
Idx: 10	Rune: '界'
*/
```

If a UTF-8 decoder consumes an unexpected input byte, it generates a special Unicode *replacement character* '\uFFFD' (the character of a white question mark inside a black diamond: �).

### Byte Slices
Constructing strings from repeated string concatenations is inefficient since it involves a lot of allocation and copying, so we can utilize byte slices instead. If know we will be doing heavy modification with a string, we can first convert the string to a bytes slice with `[]byte(s)`. This conversion can be done both ways to change to or from a string.

```go
s := "my string"
b := []byte(s)
fmt.Printf("%t\n", s == string(b)) // "true"
```

So that the programmer can choose whether to use a string or a byte slice, there are many parallel functions from both the `strings` and `bytes` packages.

The `bytes` package offers the `bytes.Buffer` type which is optimized for the manipulation of byte slices. The `bytes.Buffer` type has a usable zero value and thus does not require initialization.

```go
var buf bytes.Buffer
// this is permitted even though '3' is a rune literal (int32) since int32
// can be converted to a byte (int8) without changing values for any ASCII
buf.WriteByte('3')
buf.WriteString("this will be appended")
// BAD - this would add the wrong byte
buf.WriteByte('魚')
// GOOD
buf.WriteRune('魚')
```

## Constants
TODO
