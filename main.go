package main

import (
	"errors"
	"fmt"
	"my_err_gr/err_gr"
	"time"
)

func main() {

	var sg err_gr.SomeGroup
	var err error

	sg.Go(func() error {
		err = someFuncOne()
		if err != nil {
			return err
		}
		return nil
	})
	sg.Go(func() error {
		err = someFuncTwo()
		if err != nil {
			return err
		}
		return nil
	})
	sg.Go(func() error {
		err = someFuncThree()
		if err != nil {
			return err
		}
		return nil
	})

	err = sg.Wait()
	if err != nil {
		fmt.Println("ERRS:", err)
	}
}

func someFuncOne() error {
	fmt.Println("func one ...")
	time.Sleep(2 * time.Second)
	return errors.New("ERR_ONE")
}

func someFuncTwo() error {
	fmt.Println("func two ...")
	return errors.New("ERR_TWO")
}

func someFuncThree() error {
	fmt.Println("func three ...")
	time.Sleep(1 * time.Second)
	return nil
}
