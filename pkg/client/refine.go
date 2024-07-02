package st

type RefineYield string

const (
	RefineYieldIron     = RefineYield(TradeGoodIron)
	RefineYieldCopper   = RefineYield(TradeGoodCopper)
	RefineYieldSilver   = RefineYield(TradeGoodSilver)
	RefineYieldGold     = RefineYield(TradeGoodGold)
	RefineYieldAluminum = RefineYield(TradeGoodAluminum)
	RefineYieldPlatinum = RefineYield(TradeGoodPlatinum)
	RefineYieldUranite  = RefineYield(TradeGoodUranite)
	RefineYieldMeritium = RefineYield(TradeGoodMeritium)
	RefineYieldFuel     = RefineYield(TradeGoodFuel)
)
