package cli

import (
	"fmt"
	"io"
)

type Context struct {
	Args    []string
	Stderr  io.Writer
	Stdout  io.Writer
	Version string
}

func (ctx *Context) Print(objects ...interface{}) {
	fmt.Fprint(ctx.Stdout, objects...)
}

func (ctx *Context) Printf(format string, objects ...interface{}) {
	fmt.Fprintf(ctx.Stdout, format, objects...)
}

func (ctx *Context) Println(objects ...interface{}) {
	fmt.Fprintln(ctx.Stdout, objects...)
}

func (ctx *Context) ErrPrint(objects ...interface{}) {
	fmt.Fprint(ctx.Stderr, objects...)
}

func (ctx *Context) ErrPrintln(objects ...interface{}) {
	fmt.Fprintln(ctx.Stderr, objects...)
}
