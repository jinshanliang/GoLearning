package main

func main() {
	/*
		var arr = []string{"php", "golang", "java", "vue.3"}
		for _, val := range arr {
			fmt.Println(val)
		} */
	/*
		switch ext := ".css"; ext {
		case ".html":
			fmt.Println("html/xml")
		case ".css":
			fmt.Println("css")
		case ".js":
			fmt.Print(".js")
		default:
			fmt.Print("无法识别")
		} */

	//判断一个数字是奇数还是偶数
	/* 	switch n := 1; n {
	   	case 1, 3, 5, 7, 9:
	   		fmt.Println("这个数是奇数")
	   		fallthrough
	   	case 2, 4, 6, 8, 10:
	   		fmt.Println("这个数是偶数")
	   	default:
	   		fmt.Println("输入格式错误") */

	/* 	for i := 1; i <= 10; i++ {
	   		if i == 2 {

	   		}
	   		fmt.Println(i)
	   	}
				fmt.Println("继续执行") */

	/* 	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue
			}
			fmt.Println("通过")
		}
	} */
	/* 	for i := 0; i < 2; i++ {
			fmt.Println("i")
			goto lable3
		}
	lable3:
		fmt.Print("这个是跳转来的\n")
		fmt.Print("这个也是跳转来的") */
	arr()
}
