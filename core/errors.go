package core

import "fmt"

// ERRORS //

type Error struct {
	ErrorName string
	Details   string
	PosStart  *Position
	PosEnd    *Position
}

func (e Error) asString() string {
	result := fmt.Sprintf("%s: %s", e.ErrorName, e.Details)
	result += fmt.Sprintf(" File %s, line %d", e.PosStart.Fn, e.PosStart.Ln+1)
	return result
}

// RTERROR //
type RTError struct {
	Error
	Context *Context
}

func (e *RTError) AsString() string {
	result := e.Context.GenerateTraceback()
	result += fmt.Sprintf("%s: %s\n\n", e.ErrorName, e.Details)
	return result
}

// RTRESULT
// RTResult is a type that represents the result of an operation
type RTResult struct {
	Value interface{}
	Error error
}

// NewRTResult creates a new RTResult instance
func NewRTResult() *RTResult {
	return &RTResult{}
}

// Success sets the result value and returns it
func (r *RTResult) success(value interface{}) *RTResult {
	r.Value = value
	return r
}

// Failure sets the error and returns it
func (r *RTResult) failure(err error) *RTResult {
	r.Error = err
	return r
}

// Register handles the result of another RTResult, propagating errors if necessary
func (r *RTResult) register(res *RTResult) interface{} {
	if res.Error != nil {
		r.Error = res.Error
	}
	return res.Value
}

func (r *RTResult) successWithReturn(value interface{}) *RTResult {
	r.Value = value
	r.Error = nil
	return r // Acts as a return carrier
}

// InvalidSyntaxError struct inheriting from Error
type InvalidSyntaxError struct {
	Error
}

func NewInvalidSyntaxError(posStart, posEnd *Position, details string) *InvalidSyntaxError {
	return &InvalidSyntaxError{
		Error: Error{
			PosStart:  posStart,
			PosEnd:    posEnd,
			ErrorName: "Invalid Syntax",
			Details:   details,
		},
	}
}

// IllegalCharError struct inheriting from Error
type IllegalCharError struct {
	Error
}

// test
func NewIllegalCharError(posStart, posEnd *Position, details string) *IllegalCharError {
	return &IllegalCharError{
		Error: Error{
			PosStart:  posStart,
			PosEnd:    posEnd,
			ErrorName: "Illegal Character",
			Details:   details,
		},
	}
}
