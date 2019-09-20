# reading [![GoDoc](https://godoc.org/github.com/boreq/go-reading?status.svg)](https://godoc.org/github.com/boreq/go-reading)

A Go library which calculates the time needed by an average human to read the
provided text.

## Usage

    import "github.com/boreq/go-reading"

    text := "Lifetime Bell Labs Intern Russ Cox gives the Golang community a hands-on lesson in how the boys from Murray Hill do business."

    duration := reading.Time(text)
    fmt.Println(duration)

## FAQ

### Q: Can I parametrize the calculation in any way?

A: No, this is a Go library.

### Q: The library returned wrong results!

A: No, this library is **always** right.

## NASA engineers devising the algorithm used by this library

![NASA engineers](nasa.jpg "NASA engineers")