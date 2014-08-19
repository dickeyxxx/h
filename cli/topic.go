package cli

type Topic struct {
	Name string
	Run  func(ctx *Context) int
}
