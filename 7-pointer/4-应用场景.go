package main

import "fmt"

func main() {

	v := New()
	fmt.Println("方法一")
	fmt.Printf("函数外, v的地址: %p, config:%s\n", &v, v.configFile)
	v.SetConfigFile("config")
	fmt.Printf("函数外, v的地址: %p, config:%s\n", &v, v.configFile)

	configName := "config"
	fmt.Println("方法二")
	fmt.Printf("函数外, v的地址: %p, config:%s\n", &v, v.configFile)
	v.SetConfigFile2(configName)
	fmt.Printf("函数外, v的地址: %p, config:%s\n", &v, v.configFile)

}

type Viper struct {
	keyDelim   string
	configFile string
}

func New() *Viper {
	return &Viper{
		keyDelim:   ".",
		configFile: "",
	}
}

// SetConfigFile 函数的参数in string类型
// string采用值传递, 即形成那和实参分配不同的内存地址
// 这种情况下, 修改函数内的实参, 并不会影响形参
func (v Viper) SetConfigFile(in string) {
	v.configFile = in
	fmt.Printf("函数内, v的地址: %p, config=%v\n", &v, v.configFile)
}

func (v *Viper) SetConfigFile2(in string) {
	v.configFile = in
	fmt.Printf("函数内, v的地址: %p, config=%v, config变了\n", v, v.configFile)
}
