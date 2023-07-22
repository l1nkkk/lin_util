package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

func f1(level int) {
	if level == 0 {
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, true)
		stackStr := string(buf[:n])
		fmt.Println("Stack Trace:")
		fmt.Println(stackStr)

		// transfer to []string
		stackTemp := lo.Compact(lo.Map(strings.Split(stackStr, "\n"), func(src string, index int) string { return strings.TrimSpace(src) }))
		var stack []string
		for i := 1; i < len(stackTemp)-1; i += 2 {
			stack = append(stack, fmt.Sprintf("%s -> %s", stackTemp[i], stackTemp[i+1]))
		}

		for _, res := range stack {
			fmt.Println(res)
		}
		fmt.Println(stack)
		return
	}
	f1(level - 1)
	return
}

func main() {
	f1(3)
}
