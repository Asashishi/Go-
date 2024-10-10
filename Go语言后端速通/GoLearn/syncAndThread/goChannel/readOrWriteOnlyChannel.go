package main

func main() {

	//　声明只写管道
	var writeOnlyChan chan<- int
	writeOnlyChan = make(chan int, 5)
	for i := 0; i < 5; i++ {
		writeOnlyChan <- i
	}
	close(writeOnlyChan)

	//// 声明只读管道
	//var readOnlyChan <-chan int
	//// 初始化只读管道数据
	//if readOnlyChan != nil {
	//	num := <-readOnlyChan
	//	fmt.Println(num)
	//}

	// 一搬将关闭后的双向管道转换为只读管道使用
}
