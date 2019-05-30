# cancellation

`cancellation` is a package to implement a cancellation object. This is very
similar to how cancellation is modelled around `context.Context`, but with a
dedicated object, and new non-blocking semantic.

## non-blocking cancellation

```go
// create a cancellation object `c` and a cancellation function `cancel`
c, cancel := cancellation.New()

// non-blocking check for cancellation. This prints false
fmt.Println(c.DoneNonBlock())

// cancel the operation, and check again. This prints true
cancel()
fmt.Println(c.DoneNonBlock())
```

## blocking cancellation

```go
c, cancel := cancellation.New()

// delay cancellation by one second
go func() {
    time.Sleep(time.Second)
    cancel()
}

// block until cancellation comes one second later
<-c.Done()
```

## Why `cancel` is not a method?

You may wonder why `cancel` is a function returned by `New` alongside the
`Cancellation` object. The reason is that we want to avoid exposing the
cancellation function to the code that will check for cancellation, to avoid
that it triggers cancellation itself. This is very similar to the
`context.Context` semantic.
