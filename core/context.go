package core

import "fmt"

// CONTEXT //
type Context struct {
	DisplayName    string
	Parent         *Context
	ParentEntryPos *Position
	Symbol_Table   *SymbolTable
}

func (c *Context) GenerateTraceback() string {
	result := ""
	pos := c.ParentEntryPos
	ctx := c

	for ctx != nil {
		result = fmt.Sprintf("File %s, line %d, in %s\n", pos.Fn, pos.Ln+1, ctx.DisplayName) + result // !!!!
		pos = ctx.ParentEntryPos
		ctx = ctx.Parent
	}

	return "Traceback (most recent call last): \n" + result
}
