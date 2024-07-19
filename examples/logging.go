package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	//standard logger
	log.Println("standard logger")

	//log and flags
	//this sets the output formatconst
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	//file name and line from which log is called
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	//custom logger
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	//set prefix on existing logs
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	//set custom output targets
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	//log output into buf
	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	//slog provides structured format eg - json
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	//slog can contain rbitrary key pairs
	myslog.Info("hello again", "key", "val", "age", 25)
}
