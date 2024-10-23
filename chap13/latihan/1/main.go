/*
1/ buatkan 2 func yang memiliki parameter context,

	func 1 untuk menampilkan nilai dari context yg dikirim
	func 2 untuk mengolah nilai dari context
	misal nilai int, bisa diolah untuk operasi mtk
*/
package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	parentCtx := context.Background()
	type key string

	ctx1 := context.WithValue(parentCtx, key("num1"), 9)
	ctx2 := context.WithValue(parentCtx, key("num2"), 12)

	printContext(ctx1, key("num1"))
	printContext(ctx2, key("num2"))

	result, err := sumContextValue(ctx1, ctx2, key("num1"), key("num2"))
	if err != nil {
		fmt.Printf("Error message:  %d\n", err)
	} else {
		fmt.Printf("Sum context:  %d\n", result)
	}

	fmt.Println("------------")

}

func printContext(ctx context.Context, key interface{}) {
	value := ctx.Value(key)
	if value == 0 {
		fmt.Println("Context tidak memiliki nilai")
	} else {
		fmt.Printf("Nilai context : %d\n", value)
	}
}

func sumContextValue(ctx1, ctx2 context.Context, key1, key2 interface{}) (int, error) {
	value1, ok1 := ctx1.Value(key1).(int)
	value2, ok2 := ctx2.Value(key2).(int)

	if !ok1 || !ok2 {
		return 0, errors.New("invalid value")
	}

	return value1 + value2, nil
}
