package psql

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

type TypeAdvantage string

const (
	TypeAdvantageStrong TypeAdvantage = "TypeAdvantageStrong"
	TypeAdvantageWeak   TypeAdvantage = "TypeAdvantageWeak"
	TypeAdvantageNormal TypeAdvantage = "TypeAdvantageNormal"
	TypeAdvantageNone   TypeAdvantage = "TypeAdvantageNone"
)

func getTypeAdvantage(atk pokemon.Type, def pokemon.Type) TypeAdvantage {
	switch atk {
	case pokemon.Type_NORMAL:
		switch def {
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_STEEL:
			return TypeAdvantageWeak

		case pokemon.Type_GHOST:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}
	case pokemon.Type_FIGHTING:
		switch def {
		case pokemon.Type_NORMAL:
			fallthrough
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_ICE:
			fallthrough
		case pokemon.Type_DARK:
			return TypeAdvantageStrong

		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_PSYCHIC:
			fallthrough
		case pokemon.Type_FAIRY:
			return TypeAdvantageWeak

		case pokemon.Type_GHOST:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_FLYING:
		switch def {
		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_GRASS:
			return TypeAdvantageStrong

		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_ELECTRIC:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_POISON:
		switch def {
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_FAIRY:
			return TypeAdvantageStrong

		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_GROUND:
			fallthrough
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_GHOST:
			return TypeAdvantageWeak

		case pokemon.Type_STEEL:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_GROUND:
		switch def {

		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_ELECTRIC:
			return TypeAdvantageStrong

		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_GRASS:
			return TypeAdvantageWeak

		case pokemon.Type_FLYING:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_ROCK:
		switch def {
		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_ICE:
			return TypeAdvantageStrong

		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_GROUND:
			fallthrough
		case pokemon.Type_STEEL:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_BUG:
		switch def {
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_PSYCHIC:
			fallthrough
		case pokemon.Type_DARK:
			return TypeAdvantageStrong

		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_GHOST:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_GHOST:
		switch def {
		case pokemon.Type_GHOST:
			fallthrough
		case pokemon.Type_PSYCHIC:
			return TypeAdvantageStrong

		case pokemon.Type_DARK:
			return TypeAdvantageWeak

		case pokemon.Type_NORMAL:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_STEEL:
		switch def {
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_ICE:
			fallthrough
		case pokemon.Type_FAIRY:
			return TypeAdvantageStrong

		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_WATER:
			fallthrough
		case pokemon.Type_ELECTRIC:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_FIRE:
		switch def {
		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_ICE:
			return TypeAdvantageStrong

		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_WATER:
			fallthrough
		case pokemon.Type_DRAGON:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_WATER:
		switch def {
		case pokemon.Type_GROUND:
			fallthrough
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_FIRE:
			return TypeAdvantageStrong

		case pokemon.Type_WATER:
			fallthrough
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_DRAGON:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_GRASS:
		switch def {
		case pokemon.Type_GROUND:
			fallthrough
		case pokemon.Type_ROCK:
			fallthrough
		case pokemon.Type_WATER:
			return TypeAdvantageStrong

		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_BUG:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_DRAGON:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_ELECTRIC:
		switch def {
		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_WATER:
			return TypeAdvantageStrong

		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_ELECTRIC:
			fallthrough
		case pokemon.Type_DRAGON:
			return TypeAdvantageWeak

		case pokemon.Type_GROUND:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_PSYCHIC:
		switch def {
		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_POISON:
			return TypeAdvantageStrong

		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_PSYCHIC:
			return TypeAdvantageWeak

		case pokemon.Type_DARK:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_ICE:
		switch def {
		case pokemon.Type_FLYING:
			fallthrough
		case pokemon.Type_GROUND:
			fallthrough
		case pokemon.Type_GRASS:
			fallthrough
		case pokemon.Type_DRAGON:
			return TypeAdvantageStrong

		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			fallthrough
		case pokemon.Type_WATER:
			fallthrough
		case pokemon.Type_ICE:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_DRAGON:
		switch def {
		case pokemon.Type_DRAGON:
			return TypeAdvantageStrong

		case pokemon.Type_STEEL:
			return TypeAdvantageWeak

		case pokemon.Type_FAIRY:
			return TypeAdvantageNone

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_DARK:
		switch def {
		case pokemon.Type_GHOST:
			fallthrough
		case pokemon.Type_PSYCHIC:
			return TypeAdvantageStrong

		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_DARK:
			fallthrough
		case pokemon.Type_FAIRY:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	case pokemon.Type_FAIRY:
		switch def {
		case pokemon.Type_FIGHTING:
			fallthrough
		case pokemon.Type_DRAGON:
			fallthrough
		case pokemon.Type_DARK:
			return TypeAdvantageStrong

		case pokemon.Type_POISON:
			fallthrough
		case pokemon.Type_STEEL:
			fallthrough
		case pokemon.Type_FIRE:
			return TypeAdvantageWeak

		default:
			return TypeAdvantageNormal
		}

	default:
		return TypeAdvantageNormal
	}
}
