package main

import "github.com/go-omnibus/proof"

func main() {
	p := proof.New()
	p.SetDivision(proof.TimeDivision)
	p.SetTimeUnit(proof.Day)
	p.SetEncoding(proof.JSONEncoder)
	//p.CloseConsoleDisplay()

	p.SetInfoFile("./logs/application.log")
	p.SetErrorFile("./logs/application_err.log")

	p.Run()

	proof.Info("info level test", proof.With("Trace", "123123123"))
	proof.Debug("debug level test", proof.With("user_id", "123"))
	proof.Warnf("asdass %s", "asd")

}
