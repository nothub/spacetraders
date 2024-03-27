package st

// Register creates a new agent and ties it to an account. The agent symbol
// must consist of a 3-14 character string, and will be used to represent your
// agent. This symbol will prefix the symbol of every ship you own. Agent
// symbols will be cast to all uppercase characters.
//
// This new agent will be tied to a starting faction of your choice, which
// determines your starting location, and will be granted an authorization
// token, a contract with their starting faction, a command ship that can fly
// across space with advanced capabilities, a small probe ship that can be used
// for reconnaissance, and 150,000 credits.
//
// Email can be empty. If set, is used to reserve a call sign between resets.
func Register(symbol string, faction FactionSymbol, email string) (token string, err error) {
	var dto struct {
		Data struct {
			// We just care about the Token, all other Response data is noise.
			// Remove comments below if you need noise data in this function.
			//Agent    Agent    `json:"agent"`
			//Contract Contract `json:"contract"`
			//Faction  Faction  `json:"faction"`
			//Ship     Ship     `json:"ship"`
			Token string `json:"token"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/register", nil, struct {
		Symbol  string        `json:"symbol"`
		Email   string        `json:"email"`
		Faction FactionSymbol `json:"faction"`
	}{
		Symbol:  symbol,
		Email:   email,
		Faction: faction,
	}, &dto)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return dto.Data.Token, nil
}
