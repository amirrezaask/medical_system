package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const defaultIndent = 2

type ErrorSeverity string

const (
	ErrorSeverityIgnorable = "ignore"
	ErrorSeverityWarn      = "warn"
	ErrorSeverityError     = "error"
	ErrorSeverityPanic     = "panic"
)

// Errors done right.
type Error struct {
	Operation      string                 `json:"operation"`
	Severity       ErrorSeverity          `json:"severity"`
	ContextualData map[string]interface{} `json:"contextual_data"`
	Inner          error                  `json:"inner"`
}

func (e *Error) Error() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}

func E() *Error {
	return &Error{ContextualData: make(map[string]interface{})}
}

func (e *Error) WhenOperationWas(op string) *Error {
	e.Operation = op
	return e
}

func (e *Error) WithSeverity(s ErrorSeverity) *Error {
	e.Severity = s
	return e
}

func (e *Error) BecauseOf(err error) *Error {
	e.Inner = err
	return e
}

func (e *Error) Knowing(key string, val interface{}) *Error {
	e.ContextualData[key] = val
	return e
}

func (e *Error) PrettyPrint(indent int, writer io.Writer) error {
	indentStr := strings.Repeat(" ", indent*defaultIndent)
	fmt.Fprintf(writer, indentStr+"Operation: %s\n", e.Operation)
	fmt.Fprintf(writer, indentStr+"Severity: %s\n", e.Severity)
	fmt.Fprintf(writer, indentStr+"Context: %+v\n", e.ContextualData)
	if _, isError := e.Inner.(*Error); isError {
		fmt.Fprintf(writer, indentStr+"Inner:\n")
		return e.Inner.(*Error).PrettyPrint(indent+1, writer)
	} else {
		fmt.Fprintf(writer, indentStr+"Inner: "+e.Inner.Error())
	}
	return nil
}
