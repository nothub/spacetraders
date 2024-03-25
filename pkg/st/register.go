package st

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

	err = Post(BaseUrl+"/register", struct {
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

// RegisterAnon will save agent and registration data to config automatically.
func RegisterAnon(symbol string, faction FactionSymbol) (token string, err error) {
	return Register(symbol, faction, "")
}
