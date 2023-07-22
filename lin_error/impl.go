package lin_error

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

var _ linError = (*errorImpl)(nil)

const Partition = " | "

// errorImpl error wrap, linklist structure
type errorImpl struct {
	err   error     // wrap error, next err
	code  ErrorCode // error code
	msg   string    // error msg
	stack []string  // stack
}

func (impl *errorImpl) Error() string {
	if impl == nil {
		return ""
	}
	if impl.err != nil {
		return fmt.Sprintf("%s%s%s", impl.msg, Partition, impl.err)
	}
	return impl.msg
}

func (impl *errorImpl) Code() ErrorCode {
	if impl == nil {
		return NoError
	}
	return impl.code
}

func (impl *errorImpl) Stack() string {
	if impl == nil {
		return ""
	}
	var res string
	if impl.stack != nil {
		for _, s := range impl.stack {
			res = fmt.Sprintf("%s%s\n", res, s)
		}
	} else if impl.err != nil {
		if next, ok := impl.err.(linError); ok {
			return next.Stack()
		}
	}
	return res
}

func (impl *errorImpl) WithMessage(format string, a ...any) linError {
	return &errorImpl{
		err: impl,
		msg: fmt.Sprintf(format, a...),
	}
}

func (impl *errorImpl) WithStack(format string, a ...any) linError {
	rtn := &errorImpl{
		err: impl,
		msg: fmt.Sprintf(format, a...),
	}
	// get runtime stack
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stackStr := string(buf[:n])

	// transfer to []string
	stackTemp := lo.Compact(lo.Map(strings.Split(stackStr, "\n"), func(src string, index int) string { return strings.TrimSpace(src) }))
	var stack []string
	for i := 3; i < len(stackTemp)-1; i += 2 {
		stack = append(stack, fmt.Sprintf("%s -> %s", stackTemp[i], stackTemp[i+1]))
	}
	rtn.stack = stack
	return rtn
}
