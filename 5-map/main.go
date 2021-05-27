package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 声明1
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a %T\n", scoreMap)

	// 声明2
	userInfo := map[string]string{
		"username": "qiaocc",
		"password": "123",
	}
	fmt.Println(userInfo)
	fmt.Println(userInfo["age"])
	//判断键是否存在
	age, ok := userInfo["age"]
	if ok {
		fmt.Println("年龄为", age)
	} else {
		fmt.Println("未找到年龄")
	}

	for k, v := range userInfo {
		fmt.Printf("k=%s v=%s\n", k, v)
	}

	delete(userInfo, "password")
	fmt.Println(userInfo)

	fmt.Println("sortByKeys...")
	sortByKeys()

	// 切片中的元素是map
	var mySlice = make([]map[string]string, 2)
	fmt.Println(mySlice)
	mySlice[0] = map[string]string{
		"username": "tom",
		"password": "123",
	}
	mySlice[1] = map[string]string{
		"username": "tom",
		"password": "123",
	}
	fmt.Println(mySlice)

	var myslice2 = make(map[string][]string, 1)
	myValue := make([]string, 0, 2)
	myValue = append(myValue, "tom", "jerry")
	myslice2["usernames"] = myValue
	fmt.Println(myslice2)

}

func sortByKeys() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i) // 生成stu开头的字符串
		value := rand.Intn(100)         // 生成0-99的随机整数
		scoreMap[key] = value
	}
	//fmt.Println(scoreMap)

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key]) // 注意： 遍历map时的元素顺序与添加键值对的顺序无关。
	}
}
