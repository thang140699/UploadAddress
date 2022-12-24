package main

import (
	"WeddingUtilities/utilities"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"flag"
	"time"
)

var (
	configPrefix string
	container    *Container

	mode         string
	configSource string
)

func main() {
	flag.Parse()
	fmt.Println(mode)
	defer utilities.TimeTrack(time.Now(), fmt.Sprintf("Wedding API Service"))
	defer func() {
		fmt.Print("ef")
		if e := recover(); e != nil {
			log.Panicln(e)
			main()
		}
	}()

	// load env
	var config Config
	err := utilities.LoadEnvFromFile(&config, configPrefix, configSource)
	if err != nil {
		log.Fatalln(err)

	}
	fmt.Println(config)

	//load container
	container, err = NewContainer(config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server is running at : " + config.Binding)
	http.ListenAndServe(config.Binding, NewAPIv1(container))
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&configPrefix, "configPrefix", "wedding_utilities", "config prefix")
	flag.StringVar(&configSource, "configSource", ".env", "config source")

}
