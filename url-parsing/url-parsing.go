package main

import (
	"fmt"
	"net"
	"net/url"
)

// type URL struct {
// 	Scheme      string
// 	Opaque      string    // encoded opaque data
// 	User        *Userinfo // username and password information
// 	Host        string    // host or host:port
// 	Path        string    // path (relative paths may omit leading slash)
// 	RawPath     string    // encoded path hint (see EscapedPath method)
// 	OmitHost    bool      // do not emit empty host (authority)
// 	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
// 	RawQuery    string    // encoded query values, without '?'
// 	Fragment    string    // fragment for references, without '#'
// 	RawFragment string    // encoded fragment hint (see EscapedFragment method)
// }

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	fmt.Println(u.Query())
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
