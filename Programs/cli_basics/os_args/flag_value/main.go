package main

import (
	"flag"
	"fmt"
	"strings"
)

type idsFlag []string

func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}
func main() {
	var ids idsFlag
	flag.Var(&ids, "id", "id to be appended to the list")
	flag.Parse()
	fmt.Println(ids)

}
