# 2. Control Flow - Loops, Conditionals, and Switch

## 2.1 For Loops - The Only Loop in Go

Go has only one loop construct: `for`. But it's versatile enough to handle all looping needs.

### Basic For Loop

```go
func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum) // 45
}
```

### For as While Loop

```go
func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum) // 1024
}
```

### Infinite Loop

```go
func main() {
    for {
        fmt.Println("loop")
        break // Need break to exit
    }
}
```

### Range Loop

Range provides a clean way to iterate over data structures:

```go
// With index and value
nums := []int{2, 3, 4}
for i, num := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Value only (ignore index with _)
for _, num := range nums {
    fmt.Printf("Value: %d\n", num)
}

// Index only
for i := range nums {
    fmt.Printf("Index: %d\n", i)
}

// Modern Go: range over integers
for i := range 3 {
    fmt.Println("range", i) // 0, 1, 2
}
```

### Control Flow in Loops

```go
for n := range 6 {
    if n%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(n) // Prints 1, 3, 5
}
```

## 2.2 If Statements

### Basic If/Else

```go
func main() {
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    
    // Multiple conditions
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 is even")
    }
}
```

### If with Short Statement

Go allows a short statement before the condition:

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v // v is available here
    } else {
        fmt.Printf("%g >= %g\n", v, lim) // and here
    }
    // but not here
    return lim
}
```

**Key Point:** Variables declared in the if statement are available in both if and else blocks.

### Complex If Example

```go
func main() {
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

## 2.3 Switch Statements

Go's switch is more powerful than C's - cases don't fall through by default.

### Basic Switch

```go
func main() {
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("macOS.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)
    }
}
```

### Switch with Multiple Values

```go
func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
}
```

### Switch with No Condition (Switch True)

```go
func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
```

### Type Switch

Switch can also switch on types:

```go
func whatAmI(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}

func main() {
    whatAmI(true)  // I'm a bool
    whatAmI(1)     // I'm an int
    whatAmI("hey") // Don't know type string
}
```

## 2.4 Defer Statement

Defer postpones function execution until the surrounding function returns:

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// Output:
// hello
// world
```

### Multiple Defers (LIFO Order)

```go
func main() {
    fmt.Println("counting")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
}
// Output:
// counting
// done
// 9 8 7 6 5 4 3 2 1 0
```

**Important:** Deferred calls are executed in Last-In-First-Out order.

### Practical Defer Usage

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Ensures file is closed when function returns
    
    // Process file...
    return nil
}
```

## 2.5 Advanced Examples

### Newton's Method Square Root

From your Tour of Go exercise:

```go
func Sqrt(x float64) float64 {
    z := 1.0
    
    for {
        newZ := z - (z*z - x) / (2*z)
        
        if math.Abs(newZ - z) < 0.0000001 {
            break
        }
        fmt.Println(z)
        z = newZ
    }
    
    return z
}
```

### Date-Based Switch Logic

```go
func main() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}
```

## 2.6 Comparison with Other Languages

### Go vs C/Java Loops

**Go:**
```go
for i := 0; i < 10; i++ {
    // No parentheses required
    // Opening brace must be on same line
}
```

**C/Java:**
```c
for (int i = 0; i < 10; i++) {
    // Parentheses required
    // Brace placement flexible
}
```

### Go vs C Switch

**Go:**
- No automatic fall-through
- Can switch on any type
- No break needed (but can use explicit fallthrough)

**C:**
- Automatic fall-through
- Only integer types
- Break required to prevent fall-through

## 2.7 Key Patterns and Best Practices

### Early Return Pattern

```go
func processData(data []string) error {
    if len(data) == 0 {
        return errors.New("no data provided")
    }
    
    if !isValid(data) {
        return errors.New("invalid data")
    }
    
    // Main logic here
    return nil
}
```

### Range Patterns

```go
// Map iteration
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
}

// String iteration (runes)
for i, c := range "go" {
    fmt.Println(i, c) // Index and Unicode code point
}
```

## 2.8 Key Takeaways

1. **Single Loop Type**: `for` handles all looping needs
2. **No Parentheses**: Conditions don't need parentheses
3. **Opening Brace Rule**: Must be on same line as statement
4. **Range is Powerful**: Clean iteration over collections
5. **Switch is Clean**: No fall-through by default, works with any type
6. **Defer for Cleanup**: Ensures cleanup happens regardless of return path

**Common Patterns:**
- Use `range` for iterating over data structures
- Use `switch` instead of long if-else chains
- Use `defer` for cleanup operations
- Use short variable declarations in if statements when appropriate