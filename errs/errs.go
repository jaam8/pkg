package errs

import (
	"fmt"
	"runtime"
	"strings"
)

func WithOp(err error, comments ...string) error {
	if err == nil {
		return nil
	}

	op := FnName(1)

	if len(comments) > 0 {
		return fmt.Errorf("%s [%s] -> %w", op, strings.Join(comments, " "), err)
	}

	return fmt.Errorf("%s -> %w", op, err)
}

func FnName(skip ...int) string {
	funcName := "UNKNOWN"
	sk := 1

	if len(skip) != 0 {
		sk = skip[0] + 1
	}

	pc, _, _, ok := runtime.Caller(sk)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	s := strings.Split(funcName, ".")
	if len(s) < 2 {
		return s[len(s)-1]
	}

	fnName := s[len(s)-1]
	pkg := strings.ReplaceAll(s[len(s)-2], "(", "")
	pkg = strings.ReplaceAll(pkg, ")", "")
	pkg = strings.ReplaceAll(pkg, "*", "")

	return fmt.Sprintf("%s.%s", pkg, fnName)
}
