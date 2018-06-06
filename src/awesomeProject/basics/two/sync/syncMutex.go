package sync
import (
	"fmt"
	"runtime"
	"sync"
)
var (
	countt int32
	wgg    sync.WaitGroup
	mutex sync.Mutex//互斥锁
)
func main() {
	wgg.Add(2)
	go incCountt()
	go incCountt()
	wg.Wait()
	fmt.Println(countt)
}

/**
新声明了一个互斥锁mutex sync.Mutex，这个互斥锁有两个方法，一个是mutex.Lock(),一个是mutex.Unlock(),
这两个之间的区域就是临界区，临界区的代码是安全的。

示例中我们先调用mutex.Lock()对有竞争资源的代码加锁，这样当一个goroutine进入这个区域的时候，其他goroutine就进不来了，
只能等待，一直到调用mutex.Unlock() 释放这个锁为止。

这种方式比较灵活，可以让代码编写者任意定义需要保护的代码范围，也就是临界区。除了原子函数和互斥锁
，Go还为我们提供了更容易在多个goroutine同步的功能，这就是通道chan，我们下篇讲。
 */
func incCountt() {
	defer wgg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()//上锁
		value := countt
		runtime.Gosched()
		value++
		countt = value
		mutex.Unlock()//解锁
	}
}