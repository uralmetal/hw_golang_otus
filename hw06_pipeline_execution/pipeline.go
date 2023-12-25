package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func terminateStage(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for v := range in {
			_, ok := <-done
			if ok {
				out <- v
			} else {
				close(out)
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if done != nil {
		in = terminateStage(in, done)
	}
	for _, stage := range stages {
		in = stage(in)
	}
	return in
}
