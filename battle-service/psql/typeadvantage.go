package psql

func getTypeAdvantage() {
	 switch (atk)
            {
                case Element.Type.NORMAL:
                    switch (def)
                    {
                        case Element.Type.ROCK:
                        case Element.Type.STEEL:
                            return Effect.WEAK;

                        case Element.Type.GHOST:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }
                case Element.Type.FIGHTING:
                    switch (def)
                    {
                        case Element.Type.NORMAL:
                        case Element.Type.ROCK:
                        case Element.Type.STEEL:
                        case Element.Type.ICE:
                        case Element.Type.DARK:
                            return Effect.STRONG;

                        case Element.Type.FLYING:
                        case Element.Type.POISON:
                        case Element.Type.BUG:
                        case Element.Type.PSYCHIC:
                        case Element.Type.FAIRY:
                            return Effect.WEAK;

                        case Element.Type.GHOST:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.FLYING:
                    switch (def)
                    {
                        case Element.Type.FIGHTING:
                        case Element.Type.BUG:
                        case Element.Type.GRASS:
                            return Effect.STRONG;

                        case Element.Type.ROCK:
                        case Element.Type.STEEL:
                        case Element.Type.ELECTRIC:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.POISON:
                    switch (def)
                    {
                        case Element.Type.GRASS:
                        case Element.Type.FAIRY:
                            return Effect.STRONG;

                        case Element.Type.POISON:
                        case Element.Type.GROUND:
                        case Element.Type.ROCK:
                        case Element.Type.GHOST:
                            return Effect.WEAK;

                        case Element.Type.STEEL:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.GROUND:
                    switch (def)
                    {

                        case Element.Type.POISON:
                        case Element.Type.ROCK:
                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                        case Element.Type.ELECTRIC:
                            return Effect.STRONG;

                        case Element.Type.BUG:
                        case Element.Type.GRASS:
                            return Effect.WEAK;

                        case Element.Type.FLYING:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.ROCK:
                    switch (def)
                    {
                        case Element.Type.FLYING:
                        case Element.Type.BUG:
                        case Element.Type.FIRE:
                        case Element.Type.ICE:
                            return Effect.STRONG;

                        case Element.Type.FIGHTING:
                        case Element.Type.GROUND:
                        case Element.Type.STEEL:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.BUG:
                    switch (def)
                    {
                        case Element.Type.GRASS:
                        case Element.Type.PSYCHIC:
                        case Element.Type.DARK:
                            return Effect.STRONG;

                        case Element.Type.FIGHTING:
                        case Element.Type.FLYING:
                        case Element.Type.POISON:
                        case Element.Type.GHOST:
                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.GHOST:
                    switch (def)
                    {
                        case Element.Type.GHOST:
                        case Element.Type.PSYCHIC:
                            return Effect.STRONG;

                        case Element.Type.DARK:
                            return Effect.WEAK;

                        case Element.Type.NORMAL:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.STEEL:
                    switch (def)
                    {
                        case Element.Type.ROCK:
                        case Element.Type.ICE:
                        case Element.Type.FAIRY:
                            return Effect.STRONG;

                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                        case Element.Type.WATER:
                        case Element.Type.ELECTRIC:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.FIRE:
                    switch (def)
                    {
                        case Element.Type.BUG:
                        case Element.Type.STEEL:
                        case Element.Type.GRASS:
                        case Element.Type.ICE:
                            return Effect.STRONG;

                        case Element.Type.ROCK:
                        case Element.Type.FIRE:
                        case Element.Type.WATER:
                        case Element.Type.DRAGON:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.WATER:
                    switch (def)
                    {
                        case Element.Type.GROUND:
                        case Element.Type.ROCK:
                        case Element.Type.FIRE:
                            return Effect.STRONG;

                        case Element.Type.WATER:
                        case Element.Type.GRASS:
                        case Element.Type.DRAGON:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.GRASS:
                    switch (def)
                    {
                        case Element.Type.GROUND:
                        case Element.Type.ROCK:
                        case Element.Type.WATER:
                            return Effect.STRONG;

                        case Element.Type.FLYING:
                        case Element.Type.POISON:
                        case Element.Type.BUG:
                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                        case Element.Type.GRASS:
                        case Element.Type.DRAGON:
                            return Effect.WEAK;


                        default: return Effect.NORMAL;
                    }

                case Element.Type.ELECTRIC:
                    switch (def)
                    {
                        case Element.Type.FLYING:
                        case Element.Type.WATER:
                            return Effect.STRONG;

                        case Element.Type.GRASS:
                        case Element.Type.ELECTRIC:
                        case Element.Type.DRAGON:
                            return Effect.WEAK;

                        case Element.Type.GROUND:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.PSYCHIC:
                    switch (def)
                    {
                        case Element.Type.FIGHTING:
                        case Element.Type.POISON:
                            return Effect.STRONG;

                        case Element.Type.STEEL:
                        case Element.Type.PSYCHIC:
                            return Effect.WEAK;

                        case Element.Type.DARK:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.ICE:
                    switch (def)
                    {
                        case Element.Type.FLYING:
                        case Element.Type.GROUND:
                        case Element.Type.GRASS:
                        case Element.Type.DRAGON:
                            return Effect.STRONG;

                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                        case Element.Type.WATER:
                        case Element.Type.ICE:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.DRAGON:
                    switch (def)
                    {
                        case Element.Type.DRAGON:
                            return Effect.STRONG;

                        case Element.Type.STEEL:
                            return Effect.WEAK;

                        case Element.Type.FAIRY:
                            return Effect.NONE;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.DARK:
                    switch (def)
                    {
                        case Element.Type.GHOST:
                        case Element.Type.PSYCHIC:
                            return Effect.STRONG;

                        case Element.Type.FIGHTING:
                        case Element.Type.DARK:
                        case Element.Type.FAIRY:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                case Element.Type.FAIRY:
                    switch (def)
                    {
                        case Element.Type.FIGHTING:
                        case Element.Type.DRAGON:
                        case Element.Type.DARK:
                            return Effect.STRONG;

                        case Element.Type.POISON:
                        case Element.Type.STEEL:
                        case Element.Type.FIRE:
                            return Effect.WEAK;

                        default: return Effect.NORMAL;
                    }

                default:
                    return Effect.NORMAL;
            }
}