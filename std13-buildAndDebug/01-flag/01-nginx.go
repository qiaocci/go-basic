package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h    bool
	v, V bool
	t, T bool

	q *bool

	s, p, c, g string
)

func init() {
	flag.BoolVar(&h, "h", false, "show help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "supress non-error messages during configuration test")

	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration")

	flag.Usage = usage
}
func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	flag.Parse()
	if h {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.21.1
Usage: nginx [-?hvVtTq] [-s signal] [-p prefix]
             [-e filename] [-c filename] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
