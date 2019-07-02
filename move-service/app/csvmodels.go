package app

type CSVMoveBase struct {
	ID                   string
	Identifier           string
	GenerationID         string
	TypeID               string
	Power                string
	PP                   string
	Accuracy             string
	Priority             string
	TargetID             string
	DamageClassID        string
	EffectID             string
	EffectChance         string
	ContestTypeID        string
	ContestEffectID      string
	SuperContestEffectID string
}

type CSVMoveMeta struct {
	MoveID         string
	MetaCategoryID string
	MetaAilmentID  string
	MinHits        string
	MaxHits        string
	MinTurns       string
	MaxTurns       string
	Drain          string
	Healing        string
	CriticalRate   string
	AilmentChance  string
	FlinchChance   string
	StatChance     string
}
