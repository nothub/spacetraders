package main

import (
	"github.com/nothub/spacetraders/internal/random"
	"github.com/nothub/spacetraders/pkg/st"
	"log"
	"os"
)

var elog = log.New(os.Stderr, "", log.Llongfile)
var ilog = log.New(os.Stderr, "", 0)

func main() {

	err := ConfigLoad()
	if err != nil {
		elog.Fatalln(err.Error())
	}

	if Cfg.Token == "" {
		ilog.Println("Token is empty, registering new agent...")
		token, err := st.RegisterAnon(random.String(8), st.FactionCosmic)
		if err != nil {
			elog.Fatalln(err.Error())
		}

		Cfg.Token = token
		err = ConfigSave()
		if err != nil {
			elog.Fatalln(err.Error())
		}
	}
	st.Token = Cfg.Token

	status, err := st.GetStatus()
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("Status:       %s\n", status.Status)
	ilog.Printf("Version:      %s\n", status.Version)
	ilog.Printf("Last reset:   %s\n", status.ResetDate)
	ilog.Printf("Next reset:   %s\n", status.ServerResets.Next)
	ilog.Printf("Server stats: %+v\n", status.Stats)

	agent, err := st.GetAgent()
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("AccountId:    %s\n", agent.AccountId)
	ilog.Printf("Symbol:       %s\n", agent.Symbol)
	ilog.Printf("Headquarters: %v\n", agent.Headquarters)
	ilog.Printf("Credits:      %v\n", agent.Credits)
	ilog.Printf("ShipCount:    %v\n", agent.ShipCount)

}
