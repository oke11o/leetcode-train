package main

import (
	"errors"
	"sync"
)

// У нас есть поток задач которые нужно обрабатывать.
// Напишите библиотеку которая будет принимать на вход задачи и асинхронно обрабатывать их.
// Работать обработчик должен в оперативной памяти.
// В один момент времени выполняется не более N задач,
// и не более M задач могут быть поставлены в очередь.
// Если в очереди нет места, возвращаем ошибку.

var (
	ErrFullQueue       = errors.New("queue is full")
	ErrSchedOnShutdown = errors.New("SchedOnShutdown")
)

func NewSched(workerCnt, queueSize int) *Sched {
	s := Sched{queueSize: queueSize, workerCnt: workerCnt}
	s.input = make(chan Runner, s.queueSize)

	go s.Start()

	return &s
}

type Runner interface {
	Run()
}

type Sched struct {
	workerCnt int
	queueSize int
	input     chan Runner
	wg        sync.WaitGroup
	closed    bool
	mx        sync.Mutex
}

func (s *Sched) Append(r Runner) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.closed {
		return ErrSchedOnShutdown
	}

	select {
	case s.input <- r:
	default:
		return ErrFullQueue

	}

	return nil
}

func (s *Sched) Start() {
	s.wg.Add(s.workerCnt)
	for i := 0; i < s.workerCnt; i++ {
		go func() {
			defer s.wg.Done()
			for r := range s.input {
				r.Run()
			}
		}()
	}
}

func (s *Sched) Shutdown() {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.closed {
		return
	}
	s.closed = true
	close(s.input)

	s.wg.Wait()
	return
}
