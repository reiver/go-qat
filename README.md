# go-qat

Package **qat** provides tools for the **QAT** (**Q&A** / **FAQ**) format, for the Go programming-language (golang).

Q.A.T. (pronounced **kat**) is short for: **question-answer text** format.

An example QAT file looks like this:

```
Q: What is a goroutine in Go?

A: A goroutine is a lightweight thread managed by the Go runtime.

Q: What is the main package?

A: The main package defines an executable program. The program starts execution in main().
```

An example QAT file with multi-line answers looks like this:

```
Q: What is a goroutine in Go?

A:
	A goroutine is a lightweight thread managed by the Go runtime.
	
	Here is an example:
	
		go fn()

Q: What is the main package?

A: The main package defines an executable program. The program starts execution in main().
```

An example QAT file with multi-line question looks like this:

```
Q:
	What is:
	
	5 + 2
	
	?

A: 7

Q:
	What is:
	
	5 - 2
	
	?

A: 3

```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-qat

[![GoDoc](https://godoc.org/github.com/reiver/go-qat?status.svg)](https://godoc.org/github.com/reiver/go-qat)


## Import

To import package **qat** use `import` code like the following:
```
import "github.com/reiver/go-qat"
```

## Installation

To install package **qat** do the following:
```
GOPROXY=direct go get github.com/reiver/go-qat
```

## Author

Package **qat** was written by [Charles Iliya Krempeaux](http://reiver.link)
