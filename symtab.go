package fantasyname

import (
	"fmt"

	"github.com/s0rg/fantasyname/stringers"
)

var (
	syllables = []string{
		// a
		"ach", "ack", "ad", "age", "ald", "ale", "an", "ang", "ar", "ard", "as", "ash", "at", "ath", "augh", "aw",
		// b
		"ban", "bel", "bur",
		// c
		"cer", "cha", "che",
		// d
		"dan", "dar", "del", "den", "dra", "dyn",
		// e
		"ech", "eld", "elm", "em", "en", "end", "eng", "enth", "er", "ess", "est", "et",
		// g
		"gar", "gha",
		// h
		"hat", "hin", "hon",
		// i
		"ia", "ight", "ild", "im", "ina", "ine", "ing", "ir", "is", "iss", "it",
		// k
		"kal", "kel", "kim", "kin",
		// l
		"ler", "lor", "lye",
		// m
		"mor", "mos",
		// n
		"nal", "ny", "nys",
		// o
		"old", "om", "on", "or", "orm", "os", "ough",
		// p
		"per", "pol",
		// q
		"qua", "que",
		// r
		"rad", "rak", "ran", "ray", "ril", "ris", "rod", "roth", "ryn",
		// s
		"sam", "say", "ser", "shy", "skel", "sul",
		// t
		"tai", "tan", "tas", "ther", "tia", "tin", "ton", "tor", "tur",
		// u
		"um", "und", "unt", "urn", "usk", "ust",
		// v
		"ver", "ves", "vor",
		// w
		"war", "wor",
		// y
		"yer",
	}

	vowels = []string{
		"a", "e", "i", "o", "u", "y",
	}

	vowelsCombo = []string{
		"ae", "ai", "au", "ay", "ea", "ee", "ei", "eu",
		"ey", "ia", "ie", "oe", "oi", "oo", "ou", "ui",
	}

	consonantsBase = []string{
		"b", "c", "d", "f", "g", "h", "k", "l", "m",
		"n", "p", "r", "s", "t", "v", "w", "y", "z",
	}

	consonantsComboBase = []string{
		"ch", "ll", "ph", "sh", "st", "th",
	}

	consonantsBegin = []string{
		"bl", "br", "chr", "cl", "cr", "dr", "qu", "rh", "sch",
		"sl", "sm", "sn", "str", "sw", "thr", "tr", "wh", "zh",
	}

	consonantsAny = []string{
		"ck", "gh", "ld", "lt", "nd", "nn", "nt", "rd", "rr", "rt", "ss",
	}

	consX = "x"
	consJ = "j"
	consQ = "q"

	consonants = merge(consonantsBase, []string{consJ, consQ, consX})

	consonantsComboBegin = merge(
		append(consonantsBase, consJ),
		consonantsComboBase,
		consonantsBegin,
	)

	consonantsComboAny = merge(
		append(consonantsBase, consQ),
		consonantsComboBase,
		consonantsAny,
	)

	consonantsStupid = []string{
		"b", "d", "f", "g", "h", "j", "k", "m", "n", "p", "w", "bl", "br", "cl",
		"fl", "fr", "gh", "gl", "gr", "kl", "th",
	}

	insults = []string{
		"air", "ankle", "ball", "beef", "bone", "bum", "bumble", "bump", "cheese",
		"clod", "clot", "clown", "corn", "dip", "dolt", "doof", "dork", "dumb", "face",
		"finger", "foot", "fumble", "goof", "grumble", "head", "knock", "knocker",
		"knuckle", "loaf", "lump", "lunk", "meat", "muck", "munch", "nit", "numb", "pin",
		"puff", "skull", "snark", "sneeze", "thimble", "twerp", "twit", "wad", "wimp", "wipe",
	}

	mushies = []string{
		"baby", "booble", "bunker", "cuddle", "cuddly", "cutie", "doodle", "foofie",
		"gooble", "honey", "kissie", "lover", "lovey", "moofie", "mooglie", "moopie",
		"moopsie", "nookum", "poochie", "poof", "poofie", "pookie", "schmoopie",
		"schnoogle", "schnookie", "schnookum", "smooch", "smoochie", "smoosh", "snoogle",
		"snoogy", "snookie", "snookum", "snuggy", "sweetie", "woogle", "woogy", "wookie",
		"wookum", "wuddle", "wuddly", "wuggy", "wunny",
	}

	mushiesEnds = []string{
		"boo", "bunch", "bunny", "cake", "cakes", "cute", "darling", "dumpling",
		"dumplings", "face", "foof", "goo", "head", "kin", "kins", "lips", "love",
		"mush", "pie", "poo", "pooh", "pook", "pums",
	}

	syllablesStupid = []string{
		"elch", "idiot", "ob", "og", "ok", "olph", "olt", "omph", "ong", "onk", "oo",
		"oob", "oof", "oog", "ook", "ooz", "org", "ork", "orm", "oron", "ub", "uck",
		"ug", "ulf", "ult", "um", "umb", "ump", "umph", "un", "unb", "ung", "unk",
		"unph", "unt", "uzz",
	}

	symbolMap = map[rune][]fmt.Stringer{
		's': convert(syllables),
		'v': convert(vowels),
		'V': convert(merge(vowels, vowelsCombo)),
		'c': convert(consonants),
		'B': convert(consonantsComboBegin),
		'C': convert(consonantsComboAny),
		'i': convert(insults),
		'm': convert(mushies),
		'M': convert(mushiesEnds),
		'D': convert(consonantsStupid),
		'd': convert(syllablesStupid),
	}
)

func convert(in []string) (rv []fmt.Stringer) {
	rv = make([]fmt.Stringer, len(in))

	for i, s := range in {
		rv[i] = stringers.Literal(s)
	}

	return rv
}

func merge(a []string, b ...[]string) (rv []string) {
	rv = a

	for _, v := range b {
		rv = append(rv, v...)
	}

	return rv
}
