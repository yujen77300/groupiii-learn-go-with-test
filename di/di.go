package di

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	// 直接輸出到標準輸出，具有side effect，不能簡單地從函數返回值中得到這些輸出，進行驗證。
	// 建議引入interface進行測試
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}
func main() {
	Greet(os.Stdout, "Elodie")
}
