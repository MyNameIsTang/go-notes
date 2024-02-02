package demo

import "fmt"

// var flags []uint64

// func fib2(i uint64) (res uint64) {
// 	if flags[i] != 0 {
// 		res = flags[i]
// 		return
// 	}
// 	if i <= 1 {
// 		res = 1
// 	} else {
// 		res = fib2(i-1) + fib2(i-2)
// 	}
// 	flags[i] = res
// 	return
// }

// func fib2Generate() func() uint64 {
// 	yield := make(chan uint64)
// 	go func() {
// 		// count := 0
// 		f := fib()
// 		for {
// 			yield <- uint64(<-f)
// 		}
// 	}()

// 	return func() uint64 {
// 		return <-yield
// 	}
// }

// func InitGeneralLazyEvalution2() {
// 	fib3 := fib2Generate()
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%vth fib: %v\n", i, fib3())
// 	}
// }

func BuildLazyEvalutor2(evalFunc EvalFunc, initState Any) func() Any {
	retStateChan := make(chan Any)
	loopFunc := func() {
		retState := initState
		var retVal Any
		for {
			retVal, retState = evalFunc(retState)
			retStateChan <- retVal
		}
	}
	go loopFunc()
	return func() Any {
		return <-retStateChan
	}
}

func BuildLazyUintEvaluator2(evalFunc EvalFunc, initState Any) func() uint64 {
	func1 := BuildLazyEvalutor2(evalFunc, initState)
	return func() uint64 {
		return func1().(uint64)
	}
}

func InitGeneralLazyEvalution2() {
	evalFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		fir := os[0]
		sec := os[1]
		return sec, []uint64{sec, fir + sec}
	}
	func1 := BuildLazyUintEvaluator2(evalFunc, []uint64{0, 1})
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fib: %v\n", i, func1())
	}
}
