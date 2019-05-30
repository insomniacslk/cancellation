# interruption

`interruption` is a package to implement generic types of interruption, like
cancellation or pausing. This is very similar in principle to how cancellation
is modelled in `context.Context`, but with a dedicated object (no bag-of-values)
and new non-blocking semantic.

See https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation for
background.

## non-blocking interruption

```go
// create an interruption object `c` and an interruption function `interrupt`
c, interrupt := interruption.New()

// non-blocking check for interruption. This prints false
fmt.Println(c.DoneNonBlock())

// interrupt the operation, and check again. This prints true
interrupt()
fmt.Println(c.DoneNonBlock())
```

## blocking interruption

```go
c, interrupt := interruption.New()

// delay interruption by one second
go func() {
    time.Sleep(time.Second)
    interrupt()
}

// block until interruption comes one second later
<-c.Done()
```

## Why `interupt` is a function and not a method?

You may wonder why `interrupt` is a function returned by `New` alongside the
`Interruption` object. The reason is that we want to avoid exposing the
interruption function to the code that is target of the interruption, otherwise
it may trigger its own interruption, and of any object bound to this
interruption object. This is very similar to the `context.Context` semantic.
