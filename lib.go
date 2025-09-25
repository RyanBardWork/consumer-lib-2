package cl2

import (
	"context"
	"errors"
	"fmt"

	"github.com/RyanBardWork/common-lib"
)

func DoStuff(ctx context.Context, e1, e2, e3 error) {
	fmt.Println("consumer1:")

	fmt.Printf("ctx one: %v\n", common.GetContextOne(ctx))
	fmt.Printf("ctx two: %v\n", common.GetContextTwo(ctx))

	fmt.Printf("e1 is SentinelError: %v\n", errors.Is(e1, common.ExampleSentinelError))
	structErr := &common.StructError{}
	fmt.Printf("e2 is StructError: %v\n", errors.As(e2, &structErr))
	fmt.Printf("e3 is SentinelError: %v\n", errors.Is(e3, common.ExampleSentinelError))
	fmt.Printf("e3 is StructError: %v\n", errors.As(e3, &structErr))
}
