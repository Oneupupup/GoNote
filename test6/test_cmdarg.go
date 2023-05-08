package test6

import (
	"flag"
	"fmt"
	"os"
)

func CmdArgs(){
	fmt.Printf("命令行参数有 %d 个\n",len(os.Args))
	for i,v := range os.Args {
		fmt.Printf("args[%d]=%s\n",i,v)
	}
	fmt.Println()
	vPtr := flag.Bool("v",false,"show version")

	var useName string
	flag.StringVar(&useName,"u","root","用户名")
	flag.Func("f","",func(value string) error {
		fmt.Println("自定义参数：",value)
		return nil
	})

	flag.Parse()
	if *vPtr {
		fmt.Println("version 1.0.0")
	}
	fmt.Println("用户名：",useName)

	for i, v := range flag.Args() {
		fmt.Printf("参数%d=%s\n",i,v)
	}
}