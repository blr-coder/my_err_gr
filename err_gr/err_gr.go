package err_gr

import (
	"bytes"
	"fmt"
	"sync"
)

// Написать свою реализацию errgroup которая будет отображать все ошибки

type SomeGroup struct {
	wg sync.WaitGroup
	el ErrorsList
	m  sync.Mutex
}

func (s *SomeGroup) Go(f func() error) {
	s.wg.Add(1)

	go func() {

		defer s.wg.Done()

		err := f()
		if err != nil {
			s.TryAppendErr(err)
		}
	}()
}

// Wait - ожидает работу всех горутин, по завершению возвращает ВСЕ ошибки которые могли случиться
func (s *SomeGroup) Wait() error {
	s.wg.Wait()
	return s.el
}

type ErrorsList []error

func (l ErrorsList) Error() string {

	//ss := strings.Builder{}

	bb := bytes.NewBuffer([]byte{})

	for _, err := range l {
		//ss.WriteString(fmt.Sprintf("%v", err))

		_, _ = fmt.Fprintf(bb, "%v; ", err)
	}
	//return ss.String()
	return bb.String()
}

func (s *SomeGroup) TryAppendErr(e error) {
	s.m.Lock()
	s.el = append(s.el, e)
	s.m.Unlock()
}
