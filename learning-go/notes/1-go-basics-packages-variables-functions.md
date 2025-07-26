# 1. Go Basics - Packages, Variables, and Functions

## 1.1 Package System and Program Structure

Every Go program starts with a package declaration. The `main` package is special - it creates an executable program.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Key Concepts:**
- `package main` - Creates executable (not library)
- `main()` function - Entry point for execution
- Imports must be used (unused imports cause compilation errors)

### Import Styles

```go
// Individual imports
import "fmt"
import "math"

// Grouped imports (preferred)
import (
    "fmt"
    "math"
    "math/rand"
)
```

### Exported Names

Go uses capitalization to determine visibility:

```go
import "math"

func main() {
    fmt.Println(math.Pi)    // Pi is exported (capitalized)
    // fmt.Println(math.pi) // Error: pi is not exported
}
```

**Rule:** Names starting with capital letters are exported (public), lowercase are not exported (private).

## 1.2 Variables and Declarations

### Variable Declaration Styles

```go
// Explicit type declaration
var i, j int = 1, 2

// Type inference
var c, python, java = true, false, "no!"

// Short variable declaration (inside functions only)
k := 3
name := "Alice"

// Multiple assignment
a, b := swap("hello", "world")
```

### Zero Values

Variables declared without explicit initial values get zero values:

```go
var i int     // 0
var f float64 // 0.0
var b bool    // false
var s string  // ""
```

**Zero Values Example:**
```go
func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}
```

## 1.3 Basic Types

Go's basic types:

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 (represents Unicode code point)

float32 float64

complex64 complex128
```

### Type Conversions

Go requires explicit type conversions:

```go
var x, y int = 3, 4
var f float64 = math.Sqrt(float64(x*x + y*y))
var z uint = uint(f)
```

**Important:** Unlike C, Go has no implicit type conversions between numeric types.

## 1.4 Constants

Constants are declared with the `const` keyword:

```go
const Pi = 3.14

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
```

### Numeric Constants

Numeric constants are high-precision values:

```go
const (
    Big   = 1 << 100  // Very large number
    Small = Big >> 99 // Big / 2
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
    fmt.Println(needInt(Small))     // Works: Small fits in int
    fmt.Println(needFloat(Small))   // Works: Small converted to float64
    fmt.Println(needFloat(Big))     // Works: Big converted to float64
}
```

**Key Point:** Constants take the type needed by the context.

## 1.5 Functions

### Basic Function Syntax

```go
func add(x int, y int) int {
    return x + y
}

// Shortened parameter syntax (when types are same)
func add(x, y int) int {
    return x + y
}
```

### Multiple Return Values

Go functions can return multiple values:

```go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b) // world hello
}
```

### Named Return Values

Functions can name their return values:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // "naked" return
}
```

**Named returns** act like variables defined at the top of the function.

### Variadic Functions

Functions can accept variable number of arguments:

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(1, 2)        // [1 2] 3
    sum(1, 2, 3)     // [1 2 3] 6
    
    nums := []int{1, 2, 3, 4}
    sum(nums...)     // Unpack slice
}
```

## 1.6 Type Inference

When declaring variables without explicit types, Go infers the type:

```go
var i = 42           // int
var f = 3.142        // float64
var g = 0.867 + 0.5i // complex128

// In function context
func main() {
    v := 42       // int
    fmt.Printf("v is of type %T\n", v)
}
```

## 1.7 Practical Examples

### Command-Line Arguments (from the-go-programming-language)

Go provides command-line arguments through `os.Args`:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args[0] is the program name
    // os.Args[1:] are the arguments
    
    for i, arg := range os.Args {
        fmt.Printf("Arg %d: %s\n", i, arg)
    }
}
```

### String Joining Example

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    
    fmt.Println(s)
}
```

## 1.8 Key Takeaways

1. **Package System**: Every file starts with `package`, `main` package creates executables
2. **Imports**: Must use all imports, group them for readability
3. **Variables**: Multiple declaration styles, zero values, type inference
4. **Functions**: Can return multiple values, support named returns
5. **Types**: Explicit conversions required, no implicit casting
6. **Constants**: High-precision, take type from context

**Best Practices:**
- Use `:=` for new variables inside functions
- Group imports and declarations
- Use descriptive variable names
- Prefer explicit over implicit when clarity matters