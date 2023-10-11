package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	//ch := make(chan int)
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//for x := range ch {
	//	fmt.Println(x)
	//}

	//ch := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x)
	//	case ch <- i:
	//
	//	}
	//}

	//ticker := time.NewTicker(time.Second)
	//abort := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ticker.C:
	//		fmt.Println(x)
	//	case <-abort:
	//	}
	//}
	//ticker.Stop()

	//goroutine流水线
	//write := make(chan int, 1)
	//write <- 0
	//wg := sync.WaitGroup{}
	//for i := 0; i < 1000000000; i++ {
	//	read := write
	//	write = make(chan int, 1)
	//	wg.Add(1)
	//	go func(read <-chan int, write chan<- int, i int) {
	//		defer wg.Done()
	//		x := <-read
	//		fmt.Printf("现在i=%v x=%v \n", i, x)
	//		write <- x + i
	//	}(read, write, i)
	//}
	//wg.Wait()
	//fmt.Printf("all done \n")

	//两个goroutine通过无缓冲通道互相转发消息，每秒多少次通讯
	//ch := make(chan int)
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	x := 0
	//	for i := 0; i < 100000; i++ {
	//		fmt.Println("i=", i)
	//		if x%2 == 0 {
	//			ch <- 0
	//			fmt.Printf("A  x=%v \n", x)
	//			x += 1
	//		} else {
	//			<-ch
	//			fmt.Printf("A  x=%v \n", x)
	//			x -= 1
	//		}
	//	}
	//
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	x := 1
	//	for {
	//		if x%2 == 0 {
	//			ch <- 0
	//			fmt.Printf("B  x=%v \n", x)
	//			x -= 1
	//		} else {
	//			<-ch
	//			fmt.Printf("B  x=%v \n", x)
	//			x += 1
	//		}
	//	}
	//
	//}()
	//wg.Wait()

	//select的用法
	//	c1 := time.Tick(time.Second)
	//	c2 := make(chan int, 1)
	//	c2 <- 2
	//I:
	//	for {
	//		select {
	//		case <-c1:
	//			fmt.Println("c1")
	//			break I
	//			//case <-c2:
	//			//	fmt.Println("c2")
	//			//
	//			//}
	//		}
	//	}
	//	//close(c1)
	//	close(c2)
	//	fmt.Println("done")

	urlRaw := "https://www.baidu.com/s?wd=%v"
	url := fmt.Sprintf(urlRaw, 10)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err1=%v", err)
		return
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err2=%v", err)
		return
	}
	log.Println(string(ret))

}

func GetFile(filename string) (string, error) {
	return "", nil
}

//为指定文件并行的生成缩略图
func makeThumbnails4(filename []string) error {
	errors := make(chan error)
	for _, f := range filename {
		go func(f string) {
			_, err := GetFile(f)
			errors <- err
		}(f)
	}

	//如果
	for range filename {
		if err := <-errors; err != nil {
			return err //如果err不等于nil，没人消费，goroutine会阻塞
		}
	}
	return nil

}

//为指定文件并行的生成缩略图

//1. 充足容量的channel，2.
//为指定文件并行的生成缩略图
func makeThumbnails5(filename []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filename))
	for _, f := range filename {
		go func(f string) {
			file, err := GetFile(f)
			i := item{
				thumbfile: file,
				err:       err,
			}
			ch <- i
		}(f)
	}

	//如果
	for range filename {
		item := <-ch
		if item.err != nil {
			return thumbfiles, item.err
		} else {
			thumbfiles = append(thumbfiles, item.thumbfile)
		}
	}
	return

}

func makeThumbnails6(filenames []string) int {
	size := make(chan int)
	wg := sync.WaitGroup{}

	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			file, err := GetFile(f)
			if err != nil {
				return
			}
			size <- len(file)
		}(f)
	}

	ret := 0
	for x := range size {
		ret += x
	}
	return 1

}
