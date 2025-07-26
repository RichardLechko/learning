# 5. Concurrency - Goroutines and Channels

## 5.1 Introduction to Go Concurrency

Go's concurrency model is based on CSP (Communicating Sequential Processes). The main building blocks are goroutines (lightweight threads) and channels (communication pipes).

**Key Principles:**
- "Don't communicate by sharing memory; share memory by communicating"
- Goroutines are cheap - you can have thousands
- Channels provide safe communication between goroutines

## 5.2 Goroutines

Goroutines are functions that run concurrently with other functions.

### Basic Goroutine Usage

From your fetchall example:

```go
func main() {
    start := time.Now()
    ch := make(chan string)
    
    // Start multiple goroutines
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // 'go' keyword starts goroutine
    }
    
    // Wait for all goroutines to complete
    for range os.Args[1:] {
        fmt.Println(<-ch) // Receive from channel
    }
    
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    
    if err != nil {
        ch <- fmt.Sprint(err) // Send error to channel
        return
    }
    
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    
    if err != nil {
        ch <- fmt.Sprint("while reading %s: %v", url, err)
        return
    }
    
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
```

### Goroutine Characteristics

- **Lightweight**: Start with ~2KB stack, grows as needed
- **Multiplexed**: Many goroutines run on fewer OS threads
- **Cooperative**: Go runtime handles scheduling
- **Main goroutine**: `func main()` runs in the main goroutine

## 5.3 Channels

Channels allow goroutines to communicate and synchronize.

### Channel Types

```go
// Unbuffered channel (synchronous)
ch := make(chan int)

// Buffered channel (asynchronous up to buffer size)
ch := make(chan int, 100)

// Directional channels
var sendOnly chan<- int = ch  // Can only send
var recvOnly <-chan int = ch  // Can only receive
```

### Channel Operations

```go
// Send value to channel
ch <- 42

// Receive value from channel
value := <-ch

// Receive with ok flag (check if channel is closed)
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Close channel (only sender should close)
close(ch)
```

### Channel Directions in Function Parameters

```go
// Function can only send to channel
func sender(ch chan<- string) {
    ch <- "hello"
    // ch <- "world" // Can send multiple times
    close(ch) // Sender typically closes
}

// Function can only receive from channel
func receiver(ch <-chan string) {
    for msg := range ch { // Automatically stops when channel closes
        fmt.Println("Received:", msg)
    }
}
```

## 5.4 Synchronization Patterns

### Fan-out/Fan-in Pattern

Distribute work across multiple goroutines and collect results:

```go
func fanOutFanIn() {
    input := make(chan int)
    
    // Fan-out: start multiple workers
    worker1 := startWorker(input)
    worker2 := startWorker(input)
    worker3 := startWorker(input)
    
    // Send work
    go func() {
        for i := 0; i < 100; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Fan-in: merge results
    results := merge(worker1, worker2, worker3)
    
    // Process results
    for result := range results {
        fmt.Println(result)
    }
}

func startWorker(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for n := range input {
            output <- n * n // Example: square the input
        }
    }()
    return output
}
```

### Pipeline Pattern

Chain operations through multiple stages:

```go
func pipeline() {
    // Stage 1: Generate numbers
    numbers := generateNumbers(1, 100)
    
    // Stage 2: Filter even numbers
    evens := filterEvens(numbers)
    
    // Stage 3: Square the numbers
    squares := squareNumbers(evens)
    
    // Consume results
    for square := range squares {
        fmt.Println(square)
    }
}

func generateNumbers(start, end int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := start; i <= end; i++ {
            out <- i
        }
    }()
    return out
}

func filterEvens(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
    }()
    return out
}
```

## 5.5 Select Statement

Select allows a goroutine to wait on multiple channel operations.

### Basic Select

```go
func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "message from ch1"
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "message from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }
}
```

### Select with Timeout

```go
func withTimeout() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "result"
    }()
    
    select {
    case result := <-ch:
        fmt.Println("Got result:", result)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

### Non-blocking Operations

```go
func nonBlocking() {
    messages := make(chan string)
    signals := make(chan bool)
    
    // Non-blocking receive
    select {
    case msg := <-messages:
        fmt.Println("Received message:", msg)
    default:
        fmt.Println("No message received")
    }
    
    // Non-blocking send
    msg := "hello"
    select {
    case messages <- msg:
        fmt.Println("Sent message:", msg)
    default:
        fmt.Println("Cannot send message")
    }
}
```

## 5.6 Synchronization Primitives

### Mutex for Shared State

From your server examples:

```go
import "sync"

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}
```

### WaitGroup for Coordination

```go
import "sync"

func waitGroupExample() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1) // Increment counter
        
        go func(id int) {
            defer wg.Done() // Decrement counter when done
            
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All workers completed")
}
```

## 5.7 Common Concurrency Patterns

### Worker Pool

```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 5; r++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        time.Sleep(time.Second)
        fmt.Printf("Worker %d finished job %d\n", id, j)
        results <- j * 2
    }
}
```

### Rate Limiting

```go
func rateLimiting() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)
    
    // Rate limiter: 1 request per 200ms
    limiter := time.Tick(200 * time.Millisecond)
    
    for req := range requests {
        <-limiter // Wait for limiter
        fmt.Println("Request", req, time.Now())
    }
}
```

## 5.8 Error Handling in Concurrent Code

### Error Channel Pattern

```go
func concurrentWork() error {
    errCh := make(chan error, 1) // Buffered to prevent goroutine leak
    
    go func() {
        // Simulate work that might fail
        if rand.Float32() < 0.5 {
            errCh <- fmt.Errorf("work failed")
            return
        }
        errCh <- nil // Success
    }()
    
    select {
    case err := <-errCh:
        return err
    case <-time.After(5 * time.Second):
        return fmt.Errorf("work timed out")
    }
}
```

### Context for Cancellation

```go
import "context"

func withContext() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    ch := make(chan string)
    
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "work complete"
    }()
    
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-ctx.Done():
        fmt.Println("Work cancelled:", ctx.Err())
    }
}
```

## 5.9 Real-World Example: Web Server

From your Lissajous server:

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        cycles := parseInputs(r)
        lissajous(w, cycles) // Each request handled in separate goroutine
    })
    
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// HTTP server automatically starts goroutine for each request
// Multiple requests can be handled concurrently
```

## 5.10 Best Practices and Pitfalls

### Best Practices

1. **Start goroutines when needed**: Don't pre-create unless necessary
2. **Always have exit strategy**: Ensure goroutines can terminate
3. **Use channels for communication**: Prefer channels over shared memory
4. **Close channels from sender**: Only the sender should close channels
5. **Use select for timeouts**: Prevent indefinite blocking

### Common Pitfalls

1. **Goroutine leaks**: Forgetting to close channels or provide exit conditions
2. **Race conditions**: Accessing shared data without proper synchronization
3. **Deadlocks**: Circular dependencies in channel operations
4. **Channel direction confusion**: Sending on receive-only channels

### Memory Model

Go's memory model guarantees:
- Operations within a single goroutine appear to execute in program order
- Channel operations create happens-before relationships
- Use sync package primitives for other synchronization needs

## 5.11 Key Takeaways

1. **Goroutines**: Lightweight, multiplexed onto OS threads
2. **Channels**: Type-safe communication between goroutines
3. **Select**: Multi-way concurrent control structure
4. **Patterns**: Fan-out/fan-in, pipelines, worker pools
5. **Synchronization**: Channels preferred, but mutexes available

**Philosophy**: "Don't communicate by sharing memory; share memory by communicating"