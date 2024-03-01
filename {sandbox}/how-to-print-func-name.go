package main
import (
	"fmt"
	"reflect"
	"runtime"
)

func main () {

    fmt.Println( "/", whoAmI(bar) )
    fmt.Println( "/", whoAmI(foo) )
}

func whoAmI (i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf( i ).Pointer()).Name()
}

func foo () {}
func bar () {}
