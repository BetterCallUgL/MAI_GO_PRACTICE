package tparallel

import "sync"

type T struct {
	main  chan struct{}
	ch1   chan struct{}
	ch2   chan struct{}
	done  chan struct{}
	flag1 bool
	flag2 bool
	wg    *sync.WaitGroup
}

func (t *T) Parallel() {
	if t.ch2 != nil {
		t.flag2 = true
		t.ch2 <- struct{}{}
		<-t.done
	} else {
		t.ch2 = make(chan struct{})
		t.flag1 = true
		t.main <- struct{}{}
		<-t.ch1
	}
}

func (t *T) Run(subtest func(t *T)) {
	t.wg.Add(1)
	t.ch2 = ch

	go func() {
		subtest(t)
		t.wg.Done()
		if !t.flag2 {
			ch <- struct{}{}
		}
	}()

	<-ch
}

func Run(topTests []func(t *T)) {
	main := make(chan struct{})
	test := make([]*T, len(topTests))
	var wg sync.WaitGroup

	for i, f := range topTests {
		wg.Add(1)
		test[i] = &T{
			main:  main,
			ch1:   make(chan struct{}, 1),
			ch2:   nil,
			flag1: false,
			flag2: false,
			wg:    &wg,
			done:  make(chan struct{}),
		}

		go func(t *T) {
			f(t)
			close(t.done)
			if !t.flag1 {
				main <- struct{}{}
			}
			wg.Done()
		}(test[i])

		<-main
	}

	for i := range test {
		if test[i].flag1 {
			test[i].ch1 <- struct{}{}
		}
	}

	wg.Wait()
}
