package app

type CSVPokemon struct {
	ID             string
	Identifier     string
	SpeciesID      string
	Height         string
	Weight         string
	BaseExperience string
	Order          string
	IsDefault      string
}

type CSVPokemonStat struct {
	PokemonID string
	StatID    string
	BaseStat  string
	Effort    string
}

type CSVPokemonType struct {
	PokemonID string
	TypeID    string
	Slot      string
}
