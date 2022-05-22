package errcustom

import (
	"fmt"
	"strings"
)

type ErrorResp struct {
	ErrList []string
}

func (r ErrorResp) IsEmpty() bool {
	return len(r.ErrList) > 0
}

func (r *ErrorResp) Error() string {
	return strings.Join(r.ErrList, ",")
}

func (r *ErrorResp) AddToErrList(msg string) {
	r.ErrList = append(r.ErrList, msg)
}

func (r *ErrorResp) AddListToErrList(msg []string) {
	r.ErrList = append(r.ErrList, msg...)
}

func (r *ErrorResp) ToError() error {
	return r
}

func (r *ErrorResp) ToErrorAsString() error {
	s := strings.Join(r.ErrList, ",")
	return fmt.Errorf("err: %s", fmt.Errorf(s))
}

func NewErrorResp() *ErrorResp {
	return &ErrorResp{}
}

func NewErrRespFromError(err error) *ErrorResp {
	e, ok := err.(*ErrorResp)
	if ok {
		return e
	}
	return nil
}
