package test6

import (
	"GoNote/util"
	"fmt"
	"os"
	"bufio"
)

func TestFile1() {
	fmt.Println("TestFile")
	util.MkdirWithFilePath("test7/test9/test_file.go")
}


func TestFile2() {
	// 文件属性
	DirEntrys, err := os.ReadDir("E:\\Learning\\Go\\GoNote\\test1")
	if err != nil {
		panic(err)
	}
	for _, DirEntry := range DirEntrys {
		fmt.Println(DirEntry.Name(), DirEntry.IsDir())
	}
}

func TestFile3() {
	// 打开文件
	file, err := os.OpenFile("E:\\Learning\\Go\\GoNote\\README.md", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func TestFile4() {
	// 无缓冲区读取文件
	file, err := os.ReadFile("E:\\Learning\\Go\\GoNote\\README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	os.WriteFile("E:\\Learning\\Go\\GoNote\\README.txt",file,0775);
}

// 带缓冲区读取文件
func TestFile5() {
	f5 ,err := os.OpenFile("E:\\Learning\\Go\\GoNote\\f5",os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		panic(err)
	}
	// defer : 延迟执行，函数结束前执行
	defer f5.Close()
	// 写入文件
	wirter := bufio.NewWriter(f5)
	for i:=1;i < 4;i++ {
		filename := fmt.Sprintf("f%v",i)
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		data = append(data,'\n')
		wirter.Write(data)		// 写入缓冲区
	}
	wirter.Flush()				// 将缓冲区数据写入文件
}