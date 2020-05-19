# Chapter 4 - Composite Types

In the previous chapter it was mentioned that there were different categories for Go's types. In this chapter, we will cover the composite types which are organizations for groups of data.

First, we will talk about Go's aggregate types, arrays and structs. Aggregate types are types whose values are concatenations of other values in memory.

## Arrays
Arrays have a fixed length and thus are rarely used directly in Go. For initializing an array, we can either use an *array literal* or we can just specify the length, in which case all the arrays values will be given the underlying type's zero value.

```go
// uses zero values
var arr [5]int // [0, 0, 0, 0, 0]
// array literal - missing values up to given length use zero value
arr2 := [3]int{1, 2} // [1, 2, 0]
// array literal - using an ellipsis '...' tells go to infer the length
arr3 := [...]{5, 6, 7, 8} // [5, 6, 7, 8]
// array literal - skip values
arr4 := [...]{5: 1} // [0, 0, 0, 0, 0, 1]

fmt.Println(len(arr3)) // "4"
fmt.Println(arr3[0]) // "5"
```

The array's type is a composition of both the length and the element type.

```go
arr := [2]int{1, 2}
fmt.Printf("%T\n", arr) // "[2]int"
```

If an array's element type is comparable, then the array type is comparable too. A caveat to this is that the length of the two arrays to compare must also be the same. You could not, for example, compare a `[3]int` with a `[4]int`.

In Go, arrays will be passed by value (making a copy). As such, you should not pass large arrays as parameters. We can however take a pointer to an array and use that as a parameter. Slices may be a better tool in this regard, but the pointer approach will still work.

```go
func main() {
	arr := [1]int{0}
	fmt.Println(arr[0]) // "0"
	incIdxZero(&arr)
	fmt.Println(arr[0]) // "1"
}

func incIdxZero(ptr *[1]int) {
	ptr[0]++
n}
```

Using the bracket notation to access an element works the same for both arrays and pointers to arrays. If the type is a pointer, then `a[x]` is shorthand for `(*a)[x]`.

## Slices
Unlike arrays, slices have a variable length. Slices can be written as `[]T`, which looks like an array type without a size. Slices provide access to an underlying array but are directly composed of a pointer, a length, and a capacity. For accessing the length and the capacity, we have the built-in functions `len()` and `cap()`.
* pointer - points to first element of underlying array accessible through the slice
* length - number of slice elements (cannot exceed capacity)
* capacity - number of elements between start of the slice and end of the underlying array

To get a slice from an existing sequence, we can use the *slice operator*, `s[i:j]` where 0<=i<=j<=cap(s). Just like the substring operator of the same form, the defaults for `[i:j]` are `[0:len(s)]`. This creates a new slice of elements i through j-1 of s which may be an array variable, a pointer to an array, or a slice.

Slicing is permitted beyond `len(s)` but not past `cap(s)`. This effectively means that you can access values further into the underlying array from a slice (past the original slice's length) so long as you do not go past the end of the underlying array.

```go
// slice literal
days := []string{"m", "tu", "w", "th", "f", "sa", "su"}
// try making some slices
weekDays := days[:5] // ["m", "tu", "w", "th", "f"]
weekendDays := days[5:] // ["sa", "su"]
// can slice past len(weekDays)
allDays := weekDays[:7] // ["m", "tu", "w", "th", "f", "sa", "su"]
```

Since slices contain a pointer to the underlying array, they can be copied (such as implicitly through function parameters) and mutated elsewhere.

Unlike arrays, slices are not comparable. Byte slices may be compared with the `bytes.Equal` function, but we will have to write our own comparisons if using other types. Slice comparability is omitted from Go because slice elements are indirect and may overlay. On top of this, the `==` operator tests *reference identity* for reference types, which we may not want. The one exception to slice comparability is testing against `nil`. The zero value of a slice is `nil`, which contains no underlying array and has a length and capacity of 0. Since non-nil slices can still have a length and capacity 0, it is best to use `len(s) == 0` to check if a slice is empty.

`make()` can be used to create slices with a given length and capacity.

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

### Append
Since slices have variable-length, we need a way to add to them. This is where the built-in function `append()` comes in. A call to append() will append a given value to the end of the slice and return the new slice. The append operation may need to allocate a new array that can hold all the values, so we cannot be sure if the underlying array before the call to append() is the same as the new underlying array. To accomodate this, it is usual to do `s = append(s, val)`. This is demonstrated below.

```go
var runes []rune
newRunes := append(runes, 'H')
fmt.Println(string(runes)) // ""
fmt.Println(string(newRunes)) // "H"
```

Additionally, we have the built-in function `copy(dest, src)` which will append all the values from `src` to the end of `dest`.

### Modeling Stacks & Queues with Slices

The below operations are flawed. Popping from the stack does not yield a value other than the shrunken slice and the dequeue operation could be leaving memory at the front of the underlying array that won't be freed. Nonetheless, these show how slices might be used.

```go
// push to stack
stack = append(stack, val)
// pop from stack
stack = stack[:len(stack) - 1]

// add to queue
queue = append(queue, val)
// remove from queue
queue = queue[1:]
```

## Maps
In Go, the map type represents a hash table which maps a comparable type `K` (key) to any type `V` (value). A map may be written via the construct `map[K]V`. It is especially important for type `K` to be comparable so that the the map can test if given keys exist. As with slices, maps themselves are not comparable, except to `nil`.

The zero value for maps is `nil` which does permit operations such as lookup, delete, len, and range, but does not permit storing such as `m[some-key] = some-value`.

```go
var m1 map[string]int // nil
m2 := make(map[string]int) // empty map

// map literal
ages := map[string]int{
	"jeremy": 23,
	"charles": 19,
}
// Note that composite literals requires a "trailing comma"
```

Item lookup is done with the format `m[key]`. If the map does not contain the provided key, the zero value of the map's values is returned. This is helpful for being able to skip lookups followed by initialization or assignment.

When assigning, the key-value pair will either be initialized within the map or will update to the new value.

Key-value pairs may be deleted with the built-in function `delete(m, value)`.

```go
// continuing from previous example

fmt.Println(ages["charles"]) // "19"
fmt.Println(ages["bob"]) // "0"

delete(ages, "charles")
ages["charles"]++
fmt.Println(ages["charles"]) // 1
```

Keys in Go's maps are not asserted to be sorted in any particular order. This namely comes into play when iterating over a map's key-value pairs with a `range` loop, which returns both keys and values. The following is a common pattern that accounts for both of these points:

```go
// still using the 'ages' map from earlier

names := make([]string, 0, len(ages)) // slice with just enough capacity
                                      // for all the names
for name := range ages {
	names = append(names, name)
}
// we have to sort things ourselves
sort.Strings(names)
for _, name := range names {
	fmt.Printf("%s\t%d\n", name, ages[name])
}
```

Since item lookup returns the type's zero value if not present, we need a way to distinguish betwen a key that doesn't exist and an existing key whose value is the type's zero value (such as 0 for int). Recall from an earlier chapter that many functions return a `ok` boolean value.

```go
// verbose
age, ok := ages["john"]
if !ok {
	// key "john" did not exist
}
// terse
if age, ok := ages["john"]; !ok {
	// key "john" did not exist
}
```

A map's keys must be comparable but sometimes we might want to use a key that isn't comparable, or perhaps we want to implement a custom equality definition. A method around the limitation of maps is to have some helper function `k` that maps each would-be key (such as our incomparable type) to a comparable type (namely a string). With this, we can call `k(would-be-key)` to get the actual key into our map. Altogether, this might look like `m[k(would-be-key)]`.

## Structs
Structs are unique in that they allow the aggregation of arbitrary types. Values stored within a struct are called *fields* and may have any type. In this regard, structs are thought of as heterogeneous whereas arrays and slices are considered homogeneous.

Struct fields may be accessed via dot notation `myStruct.SomeField`. When initializing a struct, the fields will follow expected behavior and default to their type's zero value.

There is also the *empty struct*, `struct{}` which has size zero and carries no information.

```go
// declaration of a named struct
type MyStruct struct {
	Field1 int
	Field2 int
	// above fields could also be written as:
	// Field1, Field2 int
	Field3 rune
	Field4 string
}

// initialize with zero values
var s1 MyStruct
fmt.Printf("%t\n", s1.Field3 == 0) // "true"

// initializing when we know the order of fields
s2 := MyStruct{3, 5, 'R', "a string"}

// initializing explicitly
// fields may be omitted, which will be initialized to their zero values
s3 := MyStruct{
	Field1: 3,
	Field2: 5,
	Field3: 'R',
	Field4: "a string",
}
// composite literal still requires "trailing comma"

// shorthand for creating a struct and taking its address
ptr := &MyStruct{1, 2, 'G', "a string"}

// the empty struct
emptyStruct := struct{}
```

Just like how arrays and pointers to arrays seem to be treated the same when using bracket notation to access elements, structs and pointers to structs seem to be treated the same when using the dot notation to access fields.

```go
type Person struct {
	Name string
	Age int
}

me := Person{
	Name: "my name",
	Age: 21,
}

fmt.Println(me.Age) // "21"

mePtr *= &me
fmtPrintln(mePtr.Age) // "21"
// mePtr.Age is the same as (*mePtr).Age
```

Structs are comparable if all the fields of the struct are comparable. This means that they could even be used for map keys.

### Struct Embedding

Aggregate types cannot contain themselves directly, but a workaround for structs is that a struct may contain a pointer to its own type. This is useful for more complex creations such as recursive data structures.

```go
type Tree struct {
	Value int
	Left, Right *tree
}
```

If you only need a single field of a struct type, you can use an `anonymous field` where the field's name can overlap with the struct's name. Anonymous fields provide a syntactic shortcut where we can, using dot notation, access a chain of fields more easily. This can be a bit quirky since the type name is used as the implicit field name - this means that the visibility of the field depends on the type.

```go
type Point struct {
	X, Y int
}

type Circle struct {
	Point // Point is 'embedded' in Cricle
	Radius int
}

c := Circle{
	Point: Point{
		X: 5,
		Y: 7,
	},
	Radius: 3
}

fmt.Println(c.X) // "5", c.X is shorthand for c.Point.X
// if Point were called 'point' instead (not exported), we could
// still use c.X from outside the declaring package but couldn't
// explicitly do c.Point.X
```
