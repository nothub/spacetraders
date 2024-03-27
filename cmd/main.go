package main

import (
	"log"
	"os"
	"time"

	"github.com/nothub/spacetraders/internal/rando"
	"github.com/nothub/spacetraders/pkg/st"
)

var elog = log.New(os.Stderr, "", log.Llongfile)
var ilog = log.New(os.Stderr, "", 0)

func main() {

	// this is an unauthenticated request because st.Token is not set
	status, err := st.GetStatus()
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("Status:       %s\n", status.Status)
	ilog.Printf("Version:      %s\n", status.Version)
	ilog.Printf("Last reset:   %s\n", status.ResetDate)
	ilog.Printf("Next reset:   %s\n", status.ServerResets.Next)
	ilog.Printf("Server stats: %+v\n", status.Stats)

	ConfigLoad()
	if Cfg.Token == "" {
		register()
	}
	st.Token = Cfg.Token

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
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("Meta: %++v", meta)
	if len(systems) > 0 {
		ilog.Println("Systems:")
		for _, sys := range systems {
			ilog.Printf("  - symb=%v type=%v sect=%s", sys.Symbol, sys.Type, sys.SectorSymbol)
		}
		sys, err := st.GetSystem(systems[0].Symbol)
		if err != nil {
			elog.Fatalln(err.Error())
		}
		ilog.Printf("System: symb=%v type=%v sect=%s", sys.Symbol, sys.Type, sys.SectorSymbol)
	}

	factions, meta, err := st.ListFactions(20, 1)
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("Meta: %++v", meta)
	if len(factions) > 0 {
		ilog.Println("Factions:")
		for _, fac := range factions {
			ilog.Printf("  - symb=%s (%v)", fac.Symbol, fac.Name)
		}
		fac, err := st.GetFaction(string(factions[0].Symbol))
		if err != nil {
			elog.Fatalln(err.Error())
		}
		ilog.Printf("Faction: %v\t(%v)", fac.Symbol, fac.Name)
	}

	contracts, meta, err := st.ListContracts(20, 1)
	if err != nil {
		elog.Fatalln(err.Error())
	}
	ilog.Printf("Meta: %++v", meta)
	if len(contracts) > 0 {
		ilog.Println("Contracts:")
		for _, ctr := range contracts {
			ilog.Printf("  - id=%s type=%v", ctr.Id, ctr.Type)
		}
		ctr, err := st.GetContract(contracts[0].Id)
		if err != nil {
			elog.Fatalln(err.Error())
		}
		ilog.Printf("Contract: id=%s type=%v", ctr.Id, ctr.Type)
	}

}

func register() {
	ilog.Println("Token is empty, registering new agent...")

	if Cfg.Symbol == "" {
		Cfg.Symbol = rando.String(8)
	}

	if Cfg.Faction == "" {
		Cfg.Faction = st.FactionCosmic
	}

	token, err := st.Register(Cfg.Symbol, Cfg.Faction, Cfg.Email)
	if err != nil {
		elog.Fatalln(err.Error())
	}

	Cfg.Token = token
	Cfg.TokenCreated = time.Now()

	ConfigSave()
}
