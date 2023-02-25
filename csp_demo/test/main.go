package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 这个链接里面的所有问题都可以分析 https://stackoverflow.com/search?tab=newest&q=%5bdeadlock%5d%20golang

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, 5)
// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)

// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			classId := <-c
// 			fmt.Println("goroutine 1 classId: ", classId)
// 		}
// 		wg.Done()
// 		fmt.Println("goroutine 1")
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			classId := <-c
// 			fmt.Println("goroutine 2 classId: ", classId)
// 		}
// 		wg.Done()
// 		fmt.Println("goroutine 2")
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		for _, ch := range arr {
// 			c <- ch
// 		}
// 		close(c)
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println("all done")
// }

/*
https://segmentfault.com/q/1010000007528186
下面这段程序源自于网络上，对于这段程序产生死锁的原因分析见链接，我也分析了这段程序，和网络上的分析略有区别。
1. 首先按照提问者的说法，如果不注释 go f1() 这段程序运行没有问题，循环输出：call f1...，这种说法是错误的
我们使用 go build -race main.go 编译代码，接着 ./main 会报错
==3080554==ERROR: ThreadSanitizer failed to allocate 0x80000000000 (8796093022208) bytes at address 200000000000 (errno: 12)
failed to reset shadow memory
不能给线程分配内存，也就是这个死循环把内存打爆了，不能再分配新的，这也符合我们对 f1 这段程序的理解

2. 如果注释掉 go f1() 这段代码，然后继续 go build -race main.go 接着 ./main 这段程序直接阻塞住了，为什么会阻塞？哪里发生了阻塞？
主线程发生了阻塞 <-ch 语句阻塞了，并且在主线程中阻塞了，我们把它放入到另外一个 goroutine 中，代码如下：

	func f1() {
		for {
			fmt.Println("call f1...")
		}
	}

	func f2() {
		fmt.Println("call f2...")
	}

	func main() {
		// go f1()
		go f2()
		ch := make(chan int)
		go func() {
			<-ch
		}()
	}

正确
*/
// func f1() {
// 	for {
// 		fmt.Println("call f1...")
// 	}
// }

// func f2() {
// 	fmt.Println("call f2...")
// }

// func main() {
// 	go f1()
// 	go f2()
// 	ch := make(chan int)
// 	<-ch
// }

/*
range for 在有 channel 的情况下等价于

	for{
		if val, ok := <-ch; ok{
			fmt.Println(val)
		}
	}

	因此，当 ch 没有被正确关闭时，它会被阻塞，因此我们应该在 range 之前正确关系 ch
	close ch 应该发生在 channel 的发送方，也就是 ch <- i 这个 goroutine 中
*/
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		// close(ch)  // 不要注释这句话，程序会正确输出
// 	}()
// 	for val := range ch { // 这里实际上如果上面的 goroutine 没有执行，它会被阻塞住，因此不需要 WaitGroup
// 		fmt.Println(val)
// 	}
// }

// 下面代码 main 线程和 classId 阻塞了，主 (main) 线程无论如何不能阻塞
// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, 5)
// 	fin := make(chan int, 3)
// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)
// 	go func() {
// 		for _, ch := range arr {
// 			c <- ch
// 		}
// 		close(c)
// 		wg.Done()
// 	}()
// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				classId := <-c
// 				// time.Sleep(time.Second)
// 				fin <- classId
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait() // 由于上述 classId goroutine 被阻塞了，导致主线程 wg.Wait 被阻塞，只要注释改行，代码就不会报错，但是这样子的设计不符合原意
// }

// 向一个关闭的 channel 拿数据，拿到的都是 0，下面程序证明了这一点
// func main() {
// 	c := make(chan int)
// 	close(c)
// 	wg := new(sync.WaitGroup)
// 	wg.Add(1)
// 	go func() {
// 		if v, ok := <-c; !ok {
// 			fmt.Println("channel has already closed")
// 			fmt.Println("the value of closed channel is: ", v)
// 		} else {
// 			fmt.Println(v)
// 		}
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

/*
1. 无论如何主线程都不能被阻塞
func main() {
	arr := []int{1, 2, 3, 4, 5}
	c := make(chan int, 5)
	fin := make(chan int, 3)
	cl := make(chan int, 1)
	defer close(fin)
	defer close(cl)
	defer close(c)
	go func() {
		for _, ch := range arr {
			c <- ch
		}
	}()
	for i := 0; i < 3; i++ {
		go func() {
			for {
				classId := <-c
				time.Sleep(time.Second)
				fin <- classId
			}
		}()
	}
	go func() {
		for classId := range fin {
			fmt.Printf("classId %d finished \n", classId)
		}
		cl <- 1
	}()
	<-cl // 这里导致主线程阻塞，把这句注释掉
	fmt.Println("all done")
}

由于 cl channel 没有放入数据，这主要是有由于 range fin 这个 goroutine 被阻塞了，
cl <- 1 根本就没有执行，如果注释掉 <-cl 这句话，这个程序能够正常结束。

2. 如果不删除 <-cl ，那么就需要调整 range fin，让 fin 这个channel 的读写等同，就不会阻塞， fin channel 的阻塞归根结底 classId 这三个 goroutine 造成的，
合理设计这些 goroutine 的行为就能让这个程序正确
*/

/*
func main() {
	arr := []int{1, 2, 3, 4, 5}
	c := make(chan int, 5)
	fin := make(chan int, 3)
	cl := make(chan int, 1)
	defer close(fin)
	defer close(cl)
	defer close(c)
	go func() {
		for _, ch := range arr {
			c <- ch
		}
	}()
	for i := 0; i < 3; i++ {
		go func() {
			for {
				classId := <-c
				time.Sleep(time.Second)
				fin <- classId
			}
		}()
	}
	go func() {
		for classId := range fin {
			fmt.Printf("classId %d finished \n", classId)
		}
		cl <- 1
	}()
	<-cl // 这里导致主线程阻塞
	fmt.Println("all done")
}
对上面这段程序的解读，其中不少问题
1. 及时在发送方关闭 channel
go func() {
		for _,ch:=range arr{
			c<-ch
		}
	}()
建议在 for 循环后增加 close(c)，当然这个代码是否关系 c 不影响结果

2. fin 被阻塞，由于 c 已经没有任何数据可读
for i:=0;i<3;i++{
		go func() {
			for{
				classId:=<-c
				time.Sleep(time.Second)
				fin<-classId
			}
		}()
	}

这段代码将开启三个 goroutine，每个 goroutine 都在死循环获取上一个 c channel 中的数据，c 只放了 5 个数据，这显然每个 goroutine 都将被阻塞

3. range 一个被阻塞的 fin channel，这将导致 deadlock，且 range 的 fin 未被正确 close，导致它一致在等
go func() {
		for classId:=range fin{
			fmt.Printf(&#34;classId %d finished \n&#34;,classId)
		}
		cl<-1 // 它不影响
	}()
这个 范围 for 循环实际上做了如下操作:

for{
    if classId, ok := <-ch; ok{
		fmt.Printf(&#34;classId %d finished \n&#34;,classId)
    }else{
		break
    }
}

我们看到这里实际上 fin 迟早会因为 c 没有新的值而被阻塞

而且我们不能在接收方关闭 fin
go func() {
               close(fin) // 这会导致 上面 fin <- chassId 发生错误
		for classId:=range fin{
			fmt.Printf(&#34;classId %d finished \n&#34;,classId)
		}
		cl<-1 // 它不影响
	}()

4. 运行错误提示
你这段代码执行后报错结果错误 stack 也到了下面这段代码上，这和我们上面的结论也是一致的。
go func() {
		for classId := range fin {
			fmt.Printf(&#34;classId %d finished \n&#34;, classId)
		}
		cl <- 1
	}()

5. 错误结果复现
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 这里会阻塞等待读取
		}
	}()
	for val := range ch {
		fmt.Println(val)
	}
}

这段代码报错，改成下面代码后正确

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
*/

// https://stackoverflow.com/questions/36505012/go-fatal-error-all-goroutines-are-asleep-deadlock 这些代码都可以分析下，很好的例子

// func main() {

// 	c1 := make(chan string)
// 	//var c1 chan string
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		fmt.Printf("go routine begin\n")
// 		// time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 		fmt.Printf("go routine done\n")
// 		// defer wg.Done()
// 	}()
// 	fmt.Printf("done c1: %v\n", <-c1)
// 	wg.Wait()
// 	fmt.Printf("out\n")
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, 5)
// 	fin := make(chan int, 3)
// 	cl := make(chan int, 1)
// 	wg := new(sync.WaitGroup)
// 	var once = new(sync.Once)

// 	defer close(cl)

// 	wg.Add(1)
// 	go func() {
// 		for _, ch := range arr {
// 			c <- ch
// 		}
// 		close(c)
// 		wg.Done()
// 	}()

// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			for {
// 				if classId, ok := <-c; ok {
// 					// time.Sleep(time.Second)
// 					fin <- classId
// 				} else {
// 					break
// 				}
// 			}
// 			// close(fin)  // 直接 close 会 panic: close of closed channel
// 			once.Do(func() { close(fin) })
// 			wg.Done()
// 		}()
// 	}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for classId := range fin {
// 			fmt.Printf("classId %d finished \n", classId)
// 		}
// 		cl <- 1
// 	}()
// 	<-cl
// 	wg.Wait()
// 	fmt.Println("all done")
// }

// 主线程被阻塞
// func main() {
// 	ch := make(chan int)
// 	ch <- 5
// 	<-ch
// }

// func main() {
// 	f, _ := os.Open("./input1.txt")
// 	scanner := bufio.NewScanner(f)
// 	file1chan := make(chan string)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		// Split the line on a space
// 		parts := strings.Fields(line)

// 		for i := range parts {
// 			file1chan <- parts[i]
// 		}
// 	}
// 	close(file1chan)
// 	print(file1chan)
// }

// func print(in <-chan string) {
// 	for str := range in {
// 		fmt.Printf("%s\n", str)
// 	}
// }

// unbuffered channel 在同一个线程中会被阻塞
// func main() {
// 	c := make(chan int)
// 	c <- 12
// 	time.Sleep(2 * time.Second)
// 	ch := <-c
// 	fmt.Println(ch)
// }

// buffered 只会在两种情况下阻塞，channel 为空，或者 channel 满了之后还在写入也会阻塞
// func main() {
// 	c := make(chan int, 1)
// 	c <- 12
// 	time.Sleep(2 * time.Second)
// 	ch := <-c
// 	fmt.Println(ch)
// }

// buffer channel 为空阻塞
// func main() {
// 	c := make(chan int, 1)
// 	// c <- 12
// 	time.Sleep(2 * time.Second)
// 	ch := <-c // 这里阻塞住了，因为没有线程写入 channel 数据
// 	fmt.Println(ch)
// }

// buffered 满了之后在同一个线程中会被阻塞
// func main() {
// 	c := make(chan int, 1)
// 	c <- 12
// 	// c <- 13 // 这里阻塞住了，因为没有线程来读上一个数据
// 	time.Sleep(2 * time.Second)
// 	ch := <-c
// 	fmt.Println(ch)
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, 5)
// 	fin := make(chan int, 3)
// 	wg := new(sync.WaitGroup)

// 	wg.Add(1)
// 	go func() {
// 		for _, ch := range arr {
// 			c <- ch
// 		}
// 		close(c)
// 		wg.Done()
// 	}()

// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for classId := range c {
// 				fin <- classId
// 			}
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(fin)
// 	}()
// 	for classId := range fin {
// 		fmt.Printf("classId %d finished \n", classId)
// 	}
// 	fmt.Println("all done")
// }

// func main() {
// 	c1 := make(chan string)
// 	wg := new(sync.WaitGroup)

// 	wg.Add(1)
// 	go func() {
// 		fmt.Printf("go routine begin\n")
// 		c1 <- "one"
// 		fmt.Printf("go routine done\n")
// 		wg.Done()
// 	}()
// 	fmt.Printf("done c1: %v\n", <-c1)
// 	wg.Wait()
// 	fmt.Printf("out\n")
// }

// func main() {
// 	c1 := make(chan string)

// 	go func() {
// 		fmt.Printf("go routine begin\n")
// 		c1 <- "one"
// 		fmt.Printf("go routine done\n")
// 	}()
// 	select {
// 	case res := <-c1:
// 		fmt.Printf("done c1: %v\n", res)
// 	}
// 	fmt.Printf("out\n")
// }

func main() {
	f, _ := os.Open("./input1.txt")
	scanner := bufio.NewScanner(f)
	file1chan := make(chan string)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line on a space
		parts := strings.Fields(line)

		for i := range parts {
			file1chan <- parts[i]
		}
	}
	close(file1chan)
	print(file1chan)
}

func print(in <-chan string) {
	for str := range in {
		fmt.Printf("%s\n", str)
	}
}
