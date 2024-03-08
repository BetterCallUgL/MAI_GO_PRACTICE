package tparallel

import "sync"

type T struct {
	main  chan struct{}
	mu    chan struct{}
	done  chan struct{}
	flag1 bool
	wg1   *sync.WaitGroup
	wg2   *sync.WaitGroup
}

func (t *T) Parallel() {
	if t.wg2 != nil {
		t.wg2.Done()
		<-t.done
		t.wg2.Add(1)
	} else {
		t.flag1 = true
		t.main <- struct{}{}
		<-t.mu
	}
}

func (t *T) Run(subtest func(t *T)) {
	t.wg1.Add(1)
	t.done = make(chan struct{})
	tmp := t.done
	var wg sync.WaitGroup
	t.wg2 = &wg
	wg.Add(1)

	go func() {
		subtest(t)
		wg.Done()
		close(tmp)
	}()

	wg.Wait()
}

func Run(topTests []func(t *T)) {
	main := make(chan struct{})
	test := make([]*T, len(topTests))
	mu := make(chan struct{})
	var wg1 sync.WaitGroup

	for i, f := range topTests {
		wg1.Add(1)
		test[i] = &T{
			main:  main,
			mu:    mu,
			flag1: false,
			wg1:   &wg1,
			done:  make(chan struct{}),
			wg2:   nil,
		}

		go func(t *T) {
			tmp := t.done
			f(t)
			if !t.flag1 {
				main <- struct{}{}
			}
			close(tmp)
			wg1.Done()
		}(test[i])

		<-main
	}

	close(mu)
	wg1.Wait()
}
