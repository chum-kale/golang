package main
//embedding = inheritance

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

//container embeds a base
//embedding looks like field w/o a name
type container struct {
	base
	str string
}

func main() {
	//initialize embedding explicitly
	co := container {
		base: base {
			num: 1
		},
		str: "some name",
	}
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	//full path for embedded var
	fmt.Println("also num:", co.base.num)

	//invoking method of embedded receiver
	fmt.Println("describe:", co.describe())


	//embedding structs with methods will allow interfaces to inherit other structs
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

}