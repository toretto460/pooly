# Pooly

Pooly is a golang package to take a pool of workers under control. The API is pretty straightforward:

## Example

You can read a full working example here
```golang
func waitSignal(done func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		done()
	}()
}

func main() {
	ctx, done := context.WithCancel(context.Background())
	waitSignal(done)

	p := pooly.New(ctx, 4)

	p.RunFunc(func() {
		time.Sleep(time.Second * 2)
		fmt.Print("Job 1 Done\n")
	})

	p.RunFunc(func() {
		time.Sleep(time.Second * 10)
		fmt.Print("Job 2 Done\n")
	})

	p.Wait()
}
```

# Package API

## Create a new pool of workers
```golang
p := pooly.New(
    ctx, // The context
    4,   // the number of workers
)
```

## Send a func to the pool
```golang
p.RunFunc(func() {
    // Submit a func to the pool
})
```

## Wait for the pool shutdown (triggered by the context)
```golang
p.Wait() // Wait for the shutdown sequence after the Context closing.
```