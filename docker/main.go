package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := `/go/log`
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("warning ", err)
		return
	}
	defer file.Close()
	for {
		time.Sleep(time.Second)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		write.WriteString("http://c.biancheng.net/golang/ \n")

		//Flush将缓存的文件真正写入到文件中
		write.Flush()
	}
}
