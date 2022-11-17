package ref
//无需使用此包 可删
import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)
//由于go build和go run 获取程序地址不一样 写了俩种
func GetLocalFor_go_run() string {
	var res string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		res = path.Dir(filename)
		fmt.Printf("\n")
	}
	fmt.Println(res)
	return res
}
func GetLocalFor_go_buid() string {
	thePath, err := os.Executable()
	res, err := filepath.EvalSymlinks(filepath.Dir(thePath))//evalsymlinks功能：返回链接(快捷方式)所指向的实际文件
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\n")
	}
	fmt.Println(res)
	fmt.Printf("\n")
	return res
}
