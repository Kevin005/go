package main

/**
http://www.cnblogs.com/liang1101/p/6842230.html
 */

//实现 try catch 例子
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
}

func main() {
	Try(func() {
		panic("foo")
	}, func(e interface{}) {
		print(e)
	})
}
