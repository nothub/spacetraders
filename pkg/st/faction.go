package st

import "strconv"

type Faction struct {
	Symbol       FactionSymbol  `json:"symbol"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Headquarters string         `json:"headquarters"`
	Traits       []FactionTrait `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}

// The symbol of the faction.
type FactionSymbol string

const (
	FactionCosmic   FactionSymbol = "COSMIC"
	FactionVoid     FactionSymbol = "VOID"
	FactionGalactic FactionSymbol = "GALACTIC"
	FactionQuantum  FactionSymbol = "QUANTUM"
	FactionDominion FactionSymbol = "DOMINION"
	FactionAstro    FactionSymbol = "ASTRO"
	FactionCorsairs FactionSymbol = "CORSAIRS"
	FactionObsidian FactionSymbol = "OBSIDIAN"
	FactionAegis    FactionSymbol = "AEGIS"
	FactionUnited   FactionSymbol = "UNITED"
	FactionSolitary FactionSymbol = "SOLITARY"
	FactionCobalt   FactionSymbol = "COBALT"
	FactionOmega    FactionSymbol = "OMEGA"
	FactionEcho     FactionSymbol = "ECHO"
	FactionLords    FactionSymbol = "LORDS"
	FactionCult     FactionSymbol = "CULT"
	FactionAncients FactionSymbol = "ANCIENTS"
	FactionShadow   FactionSymbol = "SHADOW"
	FactionEthereal FactionSymbol = "ETHEREAL"
)

type FactionTrait struct {

	// The unique identifier of the trait.
	//
	// (required)
	Symbol FactionTraitSymbol `json:"symbol"`

	// The name of the trait.
	//
	// (required)
	Name string `json:"name"`

	// A description of the trait.
	//
	// (required)
	Description string `json:"description"`
}

// The unique identifier of the trait.
type FactionTraitSymbol string

const (
	FactionTraitBureaucratic            FactionTraitSymbol = "BUREAUCRATIC"
	FactionTraitSecretive               FactionTraitSymbol = "SECRETIVE"
	FactionTraitCapitalistic            FactionTraitSymbol = "CAPITALISTIC"
	FactionTraitIndustrious             FactionTraitSymbol = "INDUSTRIOUS"
	FactionTraitPeaceful                FactionTraitSymbol = "PEACEFUL"
	FactionTraitDistrustful             FactionTraitSymbol = "DISTRUSTFUL"
	FactionTraitWelcoming               FactionTraitSymbol = "WELCOMING"
	FactionTraitSmugglers               FactionTraitSymbol = "SMUGGLERS"
	FactionTraitScavengers              FactionTraitSymbol = "SCAVENGERS"
	FactionTraitRebellious              FactionTraitSymbol = "REBELLIOUS"
	FactionTraitExiles                  FactionTraitSymbol = "EXILES"
	FactionTraitPirates                 FactionTraitSymbol = "PIRATES"
	FactionTraitRaiders                 FactionTraitSymbol = "RAIDERS"
	FactionTraitClan                    FactionTraitSymbol = "CLAN"
	FactionTraitGuild                   FactionTraitSymbol = "GUILD"
	FactionTraitDominion                FactionTraitSymbol = "DOMINION"
	FactionTraitFringe                  FactionTraitSymbol = "FRINGE"
	FactionTraitForsaken                FactionTraitSymbol = "FORSAKEN"
	FactionTraitIsolated                FactionTraitSymbol = "ISOLATED"
	FactionTraitLocalized               FactionTraitSymbol = "LOCALIZED"
	FactionTraitEstablished             FactionTraitSymbol = "ESTABLISHED"
	FactionTraitNotable                 FactionTraitSymbol = "NOTABLE"
	FactionTraitDominant                FactionTraitSymbol = "DOMINANT"
	FactionTraitInescapable             FactionTraitSymbol = "INESCAPABLE"
	FactionTraitInnovative              FactionTraitSymbol = "INNOVATIVE"
	FactionTraitBold                    FactionTraitSymbol = "BOLD"
	FactionTraitVisionary               FactionTraitSymbol = "VISIONARY"
	FactionTraitCurious                 FactionTraitSymbol = "CURIOUS"
	FactionTraitDaring                  FactionTraitSymbol = "DARING"
	FactionTraitExploratory             FactionTraitSymbol = "EXPLORATORY"
	FactionTraitResourceful             FactionTraitSymbol = "RESOURCEFUL"
	FactionTraitFlexible                FactionTraitSymbol = "FLEXIBLE"
	FactionTraitCooperative             FactionTraitSymbol = "COOPERATIVE"
	FactionTraitUnited                  FactionTraitSymbol = "UNITED"
	FactionTraitStrategic               FactionTraitSymbol = "STRATEGIC"
	FactionTraitIntelligent             FactionTraitSymbol = "INTELLIGENT"
	FactionTraitResearchFocused         FactionTraitSymbol = "RESEARCH_FOCUSED"
	FactionTraitCollaborative           FactionTraitSymbol = "COLLABORATIVE"
	FactionTraitProgressive             FactionTraitSymbol = "PROGRESSIVE"
	FactionTraitMilitaristic            FactionTraitSymbol = "MILITARISTIC"
	FactionTraitTechnologicallyAdvanced FactionTraitSymbol = "TECHNOLOGICALLY_ADVANCED"
	FactionTraitAggressive              FactionTraitSymbol = "AGGRESSIVE"
	FactionTraitImperialistic           FactionTraitSymbol = "IMPERIALISTIC"
	FactionTraitTreasureHunters         FactionTraitSymbol = "TREASURE_HUNTERS"
	FactionTraitDexterous               FactionTraitSymbol = "DEXTEROUS"
	FactionTraitUnpredictable           FactionTraitSymbol = "UNPREDICTABLE"
	FactionTraitBrutal                  FactionTraitSymbol = "BRUTAL"
	FactionTraitFleeting                FactionTraitSymbol = "FLEETING"
	FactionTraitAdaptable               FactionTraitSymbol = "ADAPTABLE"
	FactionTraitSelfSufficient          FactionTraitSymbol = "SELF_SUFFICIENT"
	FactionTraitDefensive               FactionTraitSymbol = "DEFENSIVE"
	FactionTraitProud                   FactionTraitSymbol = "PROUD"
	FactionTraitDiverse                 FactionTraitSymbol = "DIVERSE"
	FactionTraitIndependent             FactionTraitSymbol = "INDEPENDENT"
	FactionTraitSelfInterested          FactionTraitSymbol = "SELF_INTERESTED"
	FactionTraitFragmented              FactionTraitSymbol = "FRAGMENTED"
	FactionTraitCommercial              FactionTraitSymbol = "COMMERCIAL"
	FactionTraitFreeMarkets             FactionTraitSymbol = "FREE_MARKETS"
	FactionTraitEntrepreneurial         FactionTraitSymbol = "ENTREPRENEURIAL"
)

func ListFactions(limit int, page int) (factions []Faction, meta Meta, err error) {
	var dto struct {
		Data []Faction `json:"data"`
		Meta Meta      `json:"meta"`
	}

	err = get(BaseUrl+"/factions", map[string]string{
		"limit": strconv.Itoa(limit),
		"page":  strconv.Itoa(page),
	}, &dto)
	if err != nil {
		return factions, meta, err
	}

	return dto.Data, dto.Meta, nil
}

func GetFaction(factionSymbol string) (faction Faction, err error) {
	var dto struct {
		Data Faction `json:"data"`
	}

	err = get(BaseUrl+"/factions/"+factionSymbol, nil, &dto)
	if err != nil {
		return faction, err
	}

	return dto.Data, nil
}
