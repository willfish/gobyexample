package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	printFlags := func(name string, flag int) {
		fmt.Printf("%-15s: %5d (binary: %07b)\n", name, flag, flag)
	}

	// Ldate = 1 << iota         // the date in the local time zone: 2009/01/23
	// Ltime                     // the time in the local time zone: 01:23:23
	// Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	// Llongfile                 // full file name and line number: /a/b/c/d.go:23
	// Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
	// LUTC                      // if Ldate or Ltime is set, use UTC rather than the local time zone
	// Lmsgprefix                // move the "prefix" from the beginning of the line to before the message
	// LstdFlags = Ldate | Ltime // initial values for the standard logger

	flag := log.Ldate
	printFlags("Ldate", flag)

	flag = log.Ltime
	printFlags("Ltime", flag)

	flag = log.Lmicroseconds
	printFlags("Lmicroseconds", flag)

	flag = log.Llongfile
	printFlags("Llongfile", flag)

	flag = log.Lshortfile
	printFlags("Lshortfile", flag)

	flag = log.LUTC
	printFlags("LUTC", flag)

	flag = log.Lmsgprefix
	printFlags("Lmsgprefix", flag)

	flag = log.LstdFlags
	printFlags("LstdFlags", flag)

	flag = log.LstdFlags | log.Lmicroseconds | log.Lshortfile | log.LUTC | log.Lmsgprefix
	log.SetFlags(flag)

	checkFlag := func(f int, name string) {
		if flag&f != 0 {
			fmt.Printf("%s is set\n", name)
		} else {
			fmt.Printf("%s is not set\n", name)
		}
	}

	checkFlag(log.Ldate, "Ldate")
	checkFlag(log.Ltime, "Ltime")
	checkFlag(log.Lmicroseconds, "Lmicroseconds")
	checkFlag(log.Llongfile, "Llongfile")
	checkFlag(log.Lshortfile, "Lshortfile")
	checkFlag(log.LUTC, "LUTC")
	checkFlag(log.Lmsgprefix, "Lmsgprefix")

	// log.Println("with micro")
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}
