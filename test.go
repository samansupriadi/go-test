package main

import "fmt"

func main(){

	honda := map[string]map[string][]string{
		"mpv":map[string][]string{
			"2000" : {"jazz brio"},
			"3000" : {"civic supra x"},
			
		},

	}

fmt.Println(honda)
fmt.Println(honda["mpv"]["3000"])


m := map[string][]string{ "cat": {"orange", "grey"}, "dog": {"black"}, }

res := append(m["cat"], "koneng")
fmt.Println(res)
fmt.Println(m)

m["nama"] = []string{"saman", "supriadi"}
fmt.Println(m)

}
