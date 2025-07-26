# 4. Functions, Methods, and Interfaces

## 4.1 Functions Deep Dive

### Function Basics

```go
func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    res := plus(1, 2)
    fmt.Println("1 + 2 =", res) // 3
    
    res = plusPlus(1, 2, 3)
    fmt.Println("1 + 2 + 3 =", res) // 6
}
```

### Multiple Return Values

```go
func vals() (int, int) {
    return 3, 7
}

func main() {
    a, b := vals()
    fmt.Println(a) // 3
    fmt.Println(b) // 7
    
    // Ignore values with blank identifier
    _, c := vals()
    fmt.Println(c) // 7
}
```

### Variadic Functions

Functions that accept variable number of arguments:

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
    
    // Unpack slice as arguments
    nums := []int{1, 2, 3, 4}
    sum(nums...)     // [1 2 3 4] 10
}
```

### Closures

Functions that capture variables from their outer scope:

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    
    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
    fmt.Println(nextInt()) // 3
    
    // Different closure instance
    newInts := intSeq()
    fmt.Println(newInts()) // 1
}
```

### Recursion

```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7)) // 5040
    
    // Recursive closure
    var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    
    fmt.Println(fib(7)) // 13
}
```

## 4.2 Methods

Methods are functions with receiver arguments that allow you to define behavior on types.

### Basic Method Syntax

```go
type rect struct {
    width, height int
}

// Method with value receiver
func (r rect) area() int {
    return r.width * r.height
}

// Method with pointer receiver
func (r *rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}
    
    fmt.Println("area:", r.area())   // 50
    fmt.Println("perim:", r.perim()) // 30
    
    // Pointer to struct
    rp := &r
    fmt.Println("area:", rp.area())   // 50 (Go automatically dereferences)
    fmt.Println("perim:", rp.perim()) // 30
}
```

### Value vs Pointer Receivers

**Value Receiver**: Operates on a copy of the struct
```go
func (r rect) scale(factor int) {
    r.width *= factor  // Modifies copy, not original
    r.height *= factor
}
```

**Pointer Receiver**: Operates on the original struct
```go
func (r *rect) scale(factor int) {
    r.width *= factor  // Modifies original
    r.height *= factor
}
```

**When to use pointer receivers:**
- Method needs to modify the receiver
- Receiver is large (avoid copying)
- Consistency (if some methods use pointers, all should)

## 4.3 Interfaces

Interfaces define method signatures. Types implement interfaces implicitly.

### Basic Interface

```go
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

// rect implements geometry
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// circle implements geometry  
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
```

### Using Interfaces

```go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    
    measure(r) // rect{3 4}, 12, 14
    measure(c) // circle{5}, 78.54, 31.42
}
```

### Type Assertions

Check if interface holds a specific type:

```go
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    
    detectCircle(r) // (no output)
    detectCircle(c) // circle with radius 5
}
```

### Empty Interface

`interface{}` can hold any type:

```go
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    describe(42)      // (42, int)
    describe("hello") // (hello, string)
    describe(true)    // (true, bool)
}
```

## 4.4 Pointers

Go has pointers but no pointer arithmetic.

### Basic Pointer Operations

```go
// Pass by value - creates a copy
func zeroval(ival int) {
    ival = 0
}

// Pass by pointer - modifies original
func zeroptr(iptr *int) {
    *iptr = 0 // Dereference pointer to change value
}

func main() {
    i := 1
    fmt.Println("initial:", i) // 1
    
    zeroval(i)
    fmt.Println("zeroval:", i) // 1 (unchanged)
    
    zeroptr(&i) // Pass address of i
    fmt.Println("zeroptr:", i) // 0 (changed)
    
    fmt.Println("pointer:", &i) // Memory address
}
```

### Pointers vs Values in Methods

```go
type Account struct {
    balance float64
}

// Value receiver - can't modify original
func (a Account) depositValue(amount float64) {
    a.balance += amount // Modifies copy only
}

// Pointer receiver - modifies original
func (a *Account) deposit(amount float64) {
    a.balance += amount // Modifies original
}

func main() {
    acc := Account{balance: 100}
    
    acc.depositValue(50)
    fmt.Println(acc.balance) // 100 (unchanged)
    
    acc.deposit(50)
    fmt.Println(acc.balance) // 150 (changed)
}
```

## 4.5 Advanced Interface Patterns

### Interface Composition

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

### Common Go Interfaces

```go
// Stringer - custom string representation
type Stringer interface {
    String() string
}

func (ss ServerState) String() string {
    return stateName[ss]
}

// Error interface
type error interface {
    Error() string
}
```

## 4.6 Real-World Examples

### From Your Personal Finance Tracker

```go
type finance interface {
    GetBalance() float64
    Deposit(amount float64)
    Withdraw(amount float64) bool
    SeeTransHistory()
    SeeAccountDetailsPretty()
    GetAccountNumber()
}

type Account struct {
    name         string
    accType      string
    balance      float64
    accNumber    int
    transNumber  int
    transHistory map[int]*Transaction
}

// Account implements finance interface
func (a *Account) GetBalance() float64 {
    return a.balance
}

func (a *Account) Deposit(depositAmt float64) {
    depositTrans := NewTransaction("Deposit", depositAmt, a.GetAccountNumber())
    a.balance += depositAmt
    a.transNumber++
    a.transHistory[a.transNumber] = depositTrans
}

func (a *Account) Withdraw(withdrawAmt float64) bool {
    if a.balance < withdrawAmt {
        fmt.Println("Insufficient funds")
        return false
    }
    withdrawTrans := NewTransaction("Withdraw", withdrawAmt*-1, a.GetAccountNumber())
    a.balance -= withdrawAmt
    a.transNumber++
    a.transHistory[a.transNumber] = withdrawTrans
    return true
}
```

### State Machine with Enums

```go
type ServerState int

const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
```

## 4.7 Function Types and Higher-Order Functions

### Functions as Types

```go
type mathFunc func(int, int) int

func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

func applyMath(a, b int, fn mathFunc) int {
    return fn(a, b)
}

func main() {
    result1 := applyMath(3, 4, add)      // 7
    result2 := applyMath(3, 4, multiply) // 12
}
```

### Method Values and Expressions

```go
r := rect{width: 3, height: 4}

// Method value - binds receiver
areaFunc := r.area
fmt.Println(areaFunc()) // 12

// Method expression - receiver as first parameter  
areaMethod := rect.area
fmt.Println(areaMethod(r)) // 12
```

## 4.8 Error Handling Patterns

### Custom Errors

```go
type ValidationError struct {
    Field string
    Value interface{}
    Msg   string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %v - %s", 
        e.Field, e.Value, e.Msg)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field: "age",
            Value: age,
            Msg:   "must be non-negative",
        }
    }
    return nil
}
```

### Error Checking Pattern

```go
func processData() error {
    data, err := readFile("input.txt")
    if err != nil {
        return fmt.Errorf("failed to read file: %w", err)
    }
    
    result, err := transform(data)
    if err != nil {
        return fmt.Errorf("failed to transform data: %w", err)
    }
    
    err = writeFile("output.txt", result)
    if err != nil {
        return fmt.Errorf("failed to write file: %w", err)
    }
    
    return nil
}
```

## 4.9 Key Takeaways

1. **Functions**: First-class values, support closures and recursion
2. **Methods**: Functions with receivers, use pointers when modifying or for large structs
3. **Interfaces**: Implicit implementation, define behavior contracts
4. **Pointers**: Reference types, no arithmetic, essential for mutable operations
5. **Error Handling**: Use error interface, wrap errors for context

**Best Practices:**
- Use pointer receivers for methods that modify the receiver
- Keep interfaces small and focused (often just 1-2 methods)
- Return errors explicitly rather than using exceptions
- Use composition over inheritance through embedding
- Prefer explicit error handling over panics