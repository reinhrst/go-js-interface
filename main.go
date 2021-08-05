package main

import (
	"runtime"
	"syscall/js"
	"time"
	"webinterface/mylib"
)

func AddInts(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		panic(`Expect two arguments: the numbers to be added`)
	}
	a := args[0].Int()
	b := args[1].Int()
	return mylib.AddInts(a, b)
}

func GetAddIntsFunction(this js.Value, args []js.Value) interface{} {
	return js.FuncOf(AddInts)
}

func Fib(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	return Fib(x-1) + Fib(x-2)
}

func AsyncFib(this js.Value, args []js.Value) interface{} {
	fibnr := args[0].Int()
	callback := args[1]

	go func() {
		time.Sleep(1 * time.Millisecond)
		result := Fib(fibnr)
		callback.Invoke(result)
	}()
	return nil
}

func SumAndProdInts(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		panic(`Expect one arguments: an array of integers`)
	}
	var numbers []int
	for i := 0; i < args[0].Length(); i++ {
		numbers = append(numbers, args[0].Index(i).Int())
	}
	sum := 0
	prod := 1
	for _, number := range numbers {
		sum += number
		prod *= number
	}
	return []interface{}{sum, prod}
}

func CombineName(this js.Value, args []js.Value) interface{} {
	var name string
	if args[0].Get("first").IsUndefined() {
		if args[0].Get("last").IsUndefined() {
			name = "<anonymous>"
		} else {
			name = "Mr/Mrs/Ms. " + args[0].Get("last").String()
		}
	} else {
		if args[0].Get("last").IsUndefined() {
			name = args[0].Get("first").String() + " X."
		} else {
			name = args[0].Get("first").String() + " " + args[0].Get("last").String()
		}
	}
	return map[string]interface{}{
		"full name": name,
	}
}

func main() {
	js.Global().Set("mylib", js.ValueOf(map[string]interface{}{
		"AddInts":            js.FuncOf(AddInts),
		"GetAddIntsFunction": js.FuncOf(GetAddIntsFunction),
		"AsyncFib":           js.FuncOf(AsyncFib),
		"SumAndProdInts":     js.FuncOf(SumAndProdInts),
		"CombineName":        js.FuncOf(CombineName),
	}))
	js.Global().Get("console").Call("log", "hello", "from", "go", runtime.Version())
	c := make(chan struct{}, 0)
	<-c
}
