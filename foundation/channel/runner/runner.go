package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

//创建一个Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//添加一个任务到Runner
func (runner *Runner) Add(task ...func(int)) {
	runner.tasks = append(runner.tasks, task...)
}

//开始任务
func (runner *Runner) Start() error {
	//接受所有的中断
	signal.Notify(runner.interrupt, os.Interrupt)
	go func() {
		runner.complete <-runner.run()
	}()

	select {
	case err := <-runner.complete:
		return err
	case <-runner.timeout:
		return ErrTimeout
	}
}

//执行注册的任务,任务执行逻辑
func (runner *Runner) run() error {
	for id, task := range runner.tasks {
		if runner.getInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (runner *Runner) getInterrupt() bool {
	select {
	case <-runner.interrupt:
		signal.Stop(runner.interrupt)
		return true
	default:
		return false
	}

}
