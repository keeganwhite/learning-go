package main

import (
	"fmt"
	"sort"
)

func main() {
	provinces := make(map[string]string)
	fmt.Println(provinces)
	provinces["WC"] = "Western Cape"
	provinces["EC"] = "Eastern Cape"
	provinces["KZN"] = "KwaZulu-Natal"
	fmt.Println(provinces)

	westernCape := provinces["WC"]
	fmt.Println(westernCape)
	delete(provinces, "EC")
	fmt.Println(provinces)

	for k, v := range provinces {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(provinces))
	i := 0
	for k := range provinces {
		keys[i] = k
		i++
	}
	fmt.Println()
	fmt.Println("Unordered representation ", keys)
	sort.Strings(keys)
	fmt.Println("Ordered representation in a slice ", keys)
	fmt.Println()
	for v := range keys {
		fmt.Println(provinces[keys[v]])
	}

}
