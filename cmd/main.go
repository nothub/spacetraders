package main

import (
	"log"
	"os"
	"time"

	"github.com/nothub/spacetraders/internal/random"
	"github.com/nothub/spacetraders/pkg/st"
)

var elog = log.New(os.Stderr, "", log.Llongfile)
var ilog = log.New(os.Stderr, "", 0)

func main() {

	ConfigLoad()

	if Cfg.Symbol == "" {
		Cfg.Symbol = random.String(8)
	}

	if Cfg.Faction == "" {
		Cfg.Faction = st.FactionCosmic
	}

	if Cfg.Token == "" {
		ilog.Println("Token is empty, registering new agent...")

		var err error
		if Cfg.Email == "" {
			Cfg.Token, err = st.Register(Cfg.Symbol, Cfg.Faction, "")
		} else {
			Cfg.Token, err = st.Register(Cfg.Symbol, Cfg.Faction, Cfg.Email)
		}
		if err != nil {
			elog.Fatalln(err.Error())
		}

		Cfg.Created = time.Now()
		ConfigSave()
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

	systems, meta, err := st.ListSystems(20, 1)
	ilog.Printf("Meta: %++v", meta)
	ilog.Println("Systems:")
	for _, sys := range systems {
		ilog.Printf("  - %v\t(%v)", sys.Symbol, sys.Type)
	}
	system, err := st.GetSystem(systems[0].Symbol)
	ilog.Printf("System: %v\t(%v)", system.Symbol, system.Type)

}
