package main

import "fmt"

func main()  {

	//定義すると必ずどっかのアドレスにメモリが格納される

	//勝手に決められたアドレス(0xc0000b2008)のメモリに100が入った

	var n int = 100
	fmt.Println(n) //100

	// &n = nがどのアドレスに入っているかを参照できる
	fmt.Println(&n) //0xc0000b2008

	// pのアドレス(0xc00000e030)のメモリにnのアドレスを入れている(0xc0000b2008)
	var p *int = &n

	fmt.Println(&p) //0xc00000e030
	fmt.Println(p)  //0xc0000b2008

	fmt.Println(*p)

}
