package main

import (
	"fmt"

	"github.com/go-omnibus/proof"
)

type customType int
type testStruct struct {
	S string
	V *map[string]int
	I interface{}
}

func main() {
	p := proof.New()
	p.SetDivision(proof.TimeDivision)
	p.SetTimeUnit(proof.Day)
	p.SetEncoding(proof.ConsoleEncoder)
	p.SetCaller(true)
	p.SetCapitalColor(true)
	//p.CloseConsoleDisplay()

	p.SetInfoFile("./logs/application.log")
	p.SetErrorFile("./logs/application_err.log")

	p.Run()

	defer proof.Sync()

	proof.Info("info level test", proof.With("Trace", "123123123"))
	proof.Debug("debug level test", proof.With("user_id", "123"))
	proof.Warnf("asdass %s", "asd")
	proof.Errorf("test err", proof.WithError(fmt.Errorf("err test")))

	proof.Infof("msg", proof.Render("body", testStruct{
		S: "hello",
		V: &map[string]int{"foo": 0, "bar": 1},
		I: customType(42),
	}))

}
