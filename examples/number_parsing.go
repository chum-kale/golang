//parsing nos from strings

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//parsefloat - bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//0 - base, 64 - result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//recognize hex formatted nos
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//uint
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//bas-10 parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//return an error on wrong input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
