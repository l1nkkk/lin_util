package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func demoAs() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func (e *MyError) Is(target error) bool {
	return target.Error() == e.Error()
}

func main() {
	// conclosion:
	// 1. == 判断的是对象是否相等，与其字符串值无关
	// 2. Is 可以被自定义
	// 3. Unwrap 怎能解链式的情况，不能解树状

	// New && fmt.Errors() && %w
	{
		fmt.Println("test1====")
		err1 := errors.New("error1")
		err2 := fmt.Errorf("error2 [%w]", err1)
		fmt.Println(err1)
		fmt.Println(err2)

		err3 := errors.New("error2 [error1]")
		if err2 != err3 {
			fmt.Println("err2 != err3 is true")
		}
		// 1. == 判断的是对象是否相等，与其字符串值无关
		if err2.Error() == err3.Error() {
			fmt.Println("err2.Error() == err3.Error() is true")
		}
		err4 := errors.New("error1")
		if err4 == err1 {
			fmt.Println("err4 == err1 is true")
		}
	}
	{
		// 2. Is 可以被自定义
		fmt.Println("test2====")
		err1 := &MyError{"error1"}
		if errors.Is(err1, &MyError{}) {
			fmt.Println("err1 is MyError")
		}
		if errors.Is(err1, &MyError{"error1"}) {
			fmt.Println("err1 is MyError(error1)")
		}
	}
	{
		// join
		fmt.Println("test3====")
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		err3 := errors.Join(err1, err2)
		if errors.Is(err3, err1) && errors.Is(err3, err2) {
			fmt.Println("err3 is err1 & err2")
		}
	}
	{
		// 3. Unwrap 怎能解链式的情况，不能解树状
		fmt.Println("test4====")
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		err3 := fmt.Errorf("%w, %w", err1, err2)
		errUnwrap := errors.Unwrap(err3)
		// output: nil
		fmt.Println(errUnwrap)

		err4 := fmt.Errorf("%w", err1)
		errUnwrap = errors.Unwrap(err4)
		// output: error1
		fmt.Println(errUnwrap)
	}
	{
		fmt.Println("test5====")
		err1 := &MyError{"myerror"}
		err2 := errors.New("error")
		err3 := fmt.Errorf("%w %w", err1, err2)
		var myError *MyError
		if errors.As(err3, &myError) {
			// output: myerror
			fmt.Println(myError)
		}

		var perr *fs.PathError
		if errors.As(err3, &perr) {
			fmt.Println(perr.Path)
		}
	}
}
