package hw06pipelineexecution

import "sync"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func merge(cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup // HL
	out := make(chan interface{})

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan interface{}) {
		for n := range c {
			out <- n
		}
		wg.Done() // HL
	}
	wg.Add(len(cs)) // HL
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait() // HL
		close(out)
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	for _, stage := range stages {
		//in = merge(in, done)
		in = stage(in)
	}
	return in
}
