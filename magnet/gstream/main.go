package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gst/go-glib/glib"
	_ "github.com/go-gst/go-glib/glib"
	"github.com/go-gst/go-gst/gst"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Pipeline string cannot be empty")
		os.Exit(1)
	}
	//Initialize GStreamer with the arguments passed to the program. Gstreamer
	// and the bindings will automatically pop off any handlerd arguments leaving
	// nothing but a pipeline string
	gst.Init(&os.Args)
	//Create a main loop. This is only required when utilizing signals via the bindings
	mainLoop := glib.NewMainLoop(glib.MainContextDefault(), false)
	//Build a pipeline string from the cli arguments
	pipelineString := strings.Join(os.Args[1:], " ")
	//Get GStream create a pipeline from the parsed launch syntax on the cli
	pipeline, err := gst.NewPipelineFromString(pipelineString)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//Add a message handler to the pipeline bus, printing interesting information to the console
	pipeline.GetPipelineBus().AddWatch(func(msg *gst.Message) bool {
		switch msg.Type() {
		case gst.MessageEOS:
			pipeline.BlockSetState(gst.StateNull)
			mainLoop.Quit()
		case gst.MessageError:
			err := msg.ParseError()
			fmt.Println("ERROR:", err.Error())
			if debug := err.DebugString(); debug != "" {
				fmt.Println("DEBUG:", debug)
			}
			mainLoop.Quit()
		default:
			fmt.Println(msg)
		}
		return true
	})
	//Start the pipeline
	pipeline.SetState(gst.StatePlaying)
	//Block and iterate on the main loop
	mainLoop.Run()
}
