# 3. Data Structures - Arrays, Slices, Maps, and Structs

## 3.1 Arrays

Arrays in Go have fixed size and are value types.

### Array Declaration and Initialization

```go
// Declaration with zero values
var a [5]int
fmt.Println("emp:", a) // [0 0 0 0 0]

// Setting values
a[4] = 100
fmt.Println("set:", a) // [0 0 0 0 100]
fmt.Println("len:", len(a)) // 5

// Declaration with literal
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)

// Compiler infers length
b = [...]int{1, 2, 3, 4, 5}

// Indexed initialization
b = [...]int{100, 3: 400, 500} // [100 0 0 400 500]
```

### Multi-dimensional Arrays

```go
var twoD [2][3]int
for i := range 2 {
    for j := range 3 {
        twoD[i][j] = i + j
    }
}
fmt.Println("2d:", twoD) // [[0 1 2] [1 2 3]]

// Literal initialization
twoD = [2][3]int{
    {1, 2, 3},
    {1, 2, 3},
}
```

## 3.2 Slices - Dynamic Arrays

Slices are more flexible than arrays and are reference types.

### Slice Basics

```go
// Uninitialized slice is nil
var s []string
fmt.Println("uninit:", s, s == nil, len(s) == 0) // true true

// Create slice with make
s = make([]string, 3)
fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // [   ] len:3 cap:3

// Set values
s[0] = "a"
s[1] = "b" 
s[2] = "c"
fmt.Println("set:", s) // [a b c]
```

### Slice Operations

```go
// Append elements
s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("apd:", s) // [a b c d e f]

// Copy slices
c := make([]string, len(s))
copy(c, s)
fmt.Println("cpy:", c)

// Slice syntax [start:end]
l := s[2:5]  // [c d e]
l = s[:5]    // [a b c d e] (from beginning)
l = s[2:]    // [c d e f] (to end)
```

### Slice Declaration Styles

```go
// Literal declaration
t := []string{"g", "h", "i"}

// Comparing slices (Go 1.21+)
t2 := []string{"g", "h", "i"}
if slices.Equal(t, t2) {
    fmt.Println("t == t2")
}
```

### Multi-dimensional Slices

```go
twoD := make([][]int, 3)
for i := range 3 {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := range innerLen {
        twoD[i][j] = i + j
    }
}
fmt.Println("2d:", twoD) // [[0] [1 2] [2 3 4]]
```

## 3.3 Maps - Key-Value Pairs

Maps are Go's hash table implementation.

### Map Basics

```go
// Create empty map
m := make(map[string]int)

// Set values
m["k1"] = 7
m["k2"] = 13
fmt.Println("map:", m) // map[k1:7 k2:13]

// Get values
v1 := m["k1"]
fmt.Println("v1:", v1) // 7

// Non-existent key returns zero value
v3 := m["k3"]
fmt.Println("v3:", v3) // 0

fmt.Println("len:", len(m)) // 2
```

### Map Operations

```go
// Delete key
delete(m, "k2")
fmt.Println("map:", m) // map[k1:7]

// Clear entire map (Go 1.21+)
clear(m)
fmt.Println("map:", m) // map[]

// Check if key exists
_, prs := m["k2"]
fmt.Println("prs:", prs) // false
```

### Map Literals

```go
// Initialize with values
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)

// Comparing maps (Go 1.21+)
n2 := map[string]int{"foo": 1, "bar": 2}
if maps.Equal(n, n2) {
    fmt.Println("n == n2")
}
```

### Map Iteration

```go
kvs := map[string]string{"a": "apple", "b": "banana"}

// Iterate over key-value pairs
for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
}

// Iterate over keys only
for k := range kvs {
    fmt.Println("key:", k)
}
```

## 3.4 Structs - Custom Types

Structs group related data together.

### Basic Struct Definition

```go
type person struct {
    name string
    age  int
}

func main() {
    // Various initialization styles
    fmt.Println(person{"Bob", 20})                    // {Bob 20}
    fmt.Println(person{name: "Alice", age: 30})       // {Alice 30}
    fmt.Println(person{name: "Fred"})                 // {Fred 0} - age gets zero value
    fmt.Println(&person{name: "Ann", age: 40})        // &{Ann 40} - pointer to struct
}
```

### Struct Constructor Pattern

```go
func newPerson(name string, age int) *person {
    p := person{name: name}
    p.age = age
    return &p // Return pointer to struct
}

func main() {
    fmt.Println(newPerson("Jon", 42)) // &{Jon 42}
}
```

### Accessing Struct Fields

```go
s := person{name: "Sean", age: 50}
fmt.Println(s.name) // Sean

// Through pointer
sp := &s
fmt.Println(sp.age) // 50

// Modify through pointer
sp.age = 51
fmt.Println(sp.age) // 51
```

### Anonymous Structs

```go
dog := struct {
    name   string
    isGood bool
}{
    "Rex",
    true,
}
fmt.Println(dog) // {Rex true}
```

## 3.5 Struct Embedding (Composition)

Go uses composition instead of inheritance.

### Basic Embedding

```go
type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base // Embedded struct
    str  string
}
```

### Using Embedded Structs

```go
co := container{
    base: base{
        num: 1,
    },
    str: "some name",
}

// Access embedded fields directly
fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

// Or explicitly
fmt.Println("also num:", co.base.num)

// Embedded methods are promoted
fmt.Println("describe:", co.describe())
```

### Interface Satisfaction through Embedding

```go
type describer interface {
    describe() string
}

// container satisfies describer because base does
var d describer = co
fmt.Println("describer:", d.describe())
```

## 3.6 Practical Examples

### Building Data Structures

From your personal finance tracker:

```go
type Account struct {
    name         string
    accType      string
    balance      float64
    accNumber    int
    transNumber  int
    transHistory map[int]*Transaction
}

type Transaction struct {
    transactionType string
    transactionAmt  float64
    transactionDate time.Time
    accNumber       int
}

func NewAccount(name, accType string, initialBalance float64) *Account {
    return &Account{
        name:         name,
        accType:      accType,
        balance:      initialBalance,
        transNumber:  0,
        transHistory: make(map[int]*Transaction),
        accNumber:    generateAccNum(),
    }
}
```

### State Management with Maps

```go
type ServerState int

const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
}
```

## 3.7 Memory and Performance Considerations

### Arrays vs Slices

```go
// Arrays are value types - copied when passed
func modifyArray(arr [5]int) {
    arr[0] = 100 // Doesn't modify original
}

// Slices are reference types - share underlying array
func modifySlice(slice []int) {
    slice[0] = 100 // Modifies original
}
```

### Slice Capacity and Growth

```go
s := make([]int, 0, 5) // length 0, capacity 5
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    // Capacity doubles when exceeded
}
```

### Map Performance

Maps in Go:
- Average O(1) lookup, insertion, deletion
- No ordering guarantees
- Zero value is `nil` (must use `make` to initialize)

## 3.8 String and Rune Handling

### UTF-8 Strings

```go
const s = "สวัสดี" // Thai text

// len() returns bytes, not characters
fmt.Println("Len:", len(s)) // 18 (not 6)

// Count actual characters
fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6

// Iterate over runes
for idx, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, idx)
}
```

## 3.9 Key Takeaways

1. **Arrays**: Fixed size, value types, less common in Go
2. **Slices**: Dynamic, reference types, preferred over arrays
3. **Maps**: Hash tables, unordered, zero value is nil
4. **Structs**: Custom types, composition over inheritance
5. **Embedding**: Promotes fields and methods from embedded types

**Best Practices:**
- Use slices instead of arrays in most cases
- Initialize maps with `make()` before use
- Use struct embedding for composition
- Consider pointer vs value semantics for structs
- Use constructor functions for complex initialization