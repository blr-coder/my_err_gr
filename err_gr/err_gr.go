package err_gr

import (
	"fmt"
	"strings"
	"sync"
)

// Написать свою реализацию errgroup которая будет отображать все ошибки

type SomeGroup struct {
	wg sync.WaitGroup
	el ErrorsList
}

func (s *SomeGroup) Go(f func() error) {
	s.wg.Add(1)

	go func() {

		defer s.wg.Done()

		err := f()
		if err != nil {
			s.el = append(s.el, err)
		}

		//s.wg.Done()
	}()
}

// Wait - ожидает работу всех горутин, по завершению возвращает ВСЕ ошибки которые могли случиться
func (s *SomeGroup) Wait() error {
	s.wg.Wait()
	return s.el
}

type ErrorsList []error

func (l ErrorsList) Error() string {

	var s []string

	for _, err := range l {
		s = append(s, fmt.Sprintf("%v", err))

	}
	return fmt.Sprintf("SomeGroupErrors: %v", strings.Join(s, "; "))
}
