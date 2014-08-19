package cli

import (
	"fmt"
	"io"
	"os"
)

type Context struct {
	Args   []string
	Stderr io.Writer
	Stdout io.Writer
}

func NewContext(args []string) *Context {
	return &Context{args, os.Stderr, os.Stdout}
}

func (ctx *Context) Print(objects ...interface{}) {
	fmt.Fprint(ctx.Stdout, objects...)
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
