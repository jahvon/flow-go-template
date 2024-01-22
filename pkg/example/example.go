package example

import (
	"flag"
	"fmt"
	"os"

	. "flow-go-template/common"
)

func main() {
	execName := os.Getenv("FLOW_EXECUTABLE_NAME")
	fmt.Printf("Example | %s\n", execName)

	flags := flag.NewFlagSet(execName, flag.ExitOnError)
	flags.Usage = func() {
		fmt.Printf("%s\nPowered by the Example package", execName)
		flags.PrintDefaults()
	}
	flags.Parse(os.Args[1:])

	arg1 := flags.String("arg1", "true", "arg1 description")
	arg2 := flags.Bool("arg2", true, "arg2 description")
	processArgs(arg1, arg2)
}

func processArgs(arg1 *string, arg2 *bool) {
	// arg1Val, ok := HasValue(arg1)
	// if !ok {
	// 	fmt.Fprint(os.Stderr, "arg1 not provided. using default value")
	// }
	// arg2Val, ok := HasValue(arg2)
	// if !ok {
	// 	fmt.Fprint(os.Stderr, "arg2 not provided. using default value")
	// }

	// if arg1Val == "hello" && arg2Val == true {
	// 	journeyOne(arg1Val, arg2Val)
	// } else {
	// 	journeyTwo(arg1Val, arg2Val)
	// }
	journeyOne(*arg1, *arg2)
	journeyOne(*arg1, *arg2)
}

func journeyOne(arg1 string, arg2 bool) {

}

func journeyTwo(arg1 string, arg2 bool) {

}
