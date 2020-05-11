package main

import (
	"fmt"
	"os"
	"sync"
	"wget/wget"
)

func main() {
	//fmt.Println("os.Args:", os.Args)
	osArgs:=[]string{"./wgetBin","http://doc.oss-cn-shanghai-shltjyy-d01-a.procloud.10010sh.cn/65c0ebd8-26b2-4d8c-aef6-4dbd8c36ae76/A5D26A3D94A0488187281A12E1793614/63DABD9197DD41F3A036C0A01F285859"}
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		dir := "video" + fmt.Sprint(i)
		go func() {
			err, i := wget.WgetCli(dir,osArgs)
			if err != nil {
				fmt.Printf("Error: %+v. %d\n", err, i)
				os.Exit(1)
			}
			fmt.Println(dir + "success")
			wg.Done()
		}()
	}
	wg.Wait()
}
