package main

import "fmt"

type any interface{}
type evalFunc func(any) (any, any)

func main() {
	evenFunc := func(state any) (any, any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := buildLazyEvaluatorInt(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

func buildLazyEvaluatorInt(evalFunc2 evalFunc, initState any) func() int {
	ef := buildLazyEvaluator(evalFunc2, initState)
	return func() int {
		return ef().(int)
	}
}

func buildLazyEvaluator(evalFunc evalFunc, initState any) func() any {
	retValChan := make(chan any)
	loopFunc := func() {
		var actState any = initState
		var retVal any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	retFunc := func() any {
		return <-retValChan
	}

	go loopFunc()
	return retFunc
}
