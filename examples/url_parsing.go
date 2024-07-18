package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	//url includes - scheme, authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//check for errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	//scheme
	fmt.Println(u.Scheme)

	//authentication info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	//hostname and port
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	//path and fragment
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	//params in k=v format - rawquery
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)

	//query params into map
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
