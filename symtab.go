package fantasyname

import (
	"fmt"

	st "github.com/s0rg/fantasyname/stringers"
)

var (
	syllables = []fmt.Stringer{
		// a
		st.Literal("ach"), st.Literal("ack"), st.Literal("ad"), st.Literal("age"), st.Literal("ald"),
		st.Literal("ale"), st.Literal("an"), st.Literal("ang"), st.Literal("ar"), st.Literal("ard"),
		st.Literal("as"), st.Literal("ash"), st.Literal("at"), st.Literal("ath"), st.Literal("augh"),
		st.Literal("aw"),
		// b
		st.Literal("ban"), st.Literal("bel"), st.Literal("bur"),
		// c
		st.Literal("cer"), st.Literal("cha"), st.Literal("che"),
		// d
		st.Literal("dan"), st.Literal("dar"), st.Literal("del"), st.Literal("den"), st.Literal("dra"),
		st.Literal("dyn"),
		// e
		st.Literal("ech"), st.Literal("eld"), st.Literal("elm"), st.Literal("em"), st.Literal("en"),
		st.Literal("end"), st.Literal("eng"), st.Literal("enth"), st.Literal("er"), st.Literal("ess"),
		st.Literal("est"), st.Literal("et"),
		// g
		st.Literal("gar"), st.Literal("gha"),
		// h
		st.Literal("hat"), st.Literal("hin"), st.Literal("hon"),
		// i
		st.Literal("ia"), st.Literal("ight"), st.Literal("ild"), st.Literal("im"), st.Literal("ina"),
		st.Literal("ine"), st.Literal("ing"), st.Literal("ir"), st.Literal("is"), st.Literal("iss"),
		st.Literal("it"),
		// k
		st.Literal("kal"), st.Literal("kel"), st.Literal("kim"), st.Literal("kin"),
		// l
		st.Literal("ler"), st.Literal("lor"), st.Literal("lye"),
		// m
		st.Literal("mor"), st.Literal("mos"),
		// n
		st.Literal("nal"), st.Literal("ny"), st.Literal("nys"),
		// o
		st.Literal("old"), st.Literal("om"), st.Literal("on"), st.Literal("or"), st.Literal("orm"),
		st.Literal("os"), st.Literal("ough"),
		// p
		st.Literal("per"), st.Literal("pol"),
		// q
		st.Literal("qua"), st.Literal("que"),
		// r
		st.Literal("rad"), st.Literal("rak"), st.Literal("ran"), st.Literal("ray"), st.Literal("ril"),
		st.Literal("ris"), st.Literal("rod"), st.Literal("roth"), st.Literal("ryn"),
		// s
		st.Literal("sam"), st.Literal("say"), st.Literal("ser"), st.Literal("shy"), st.Literal("skel"),
		st.Literal("sul"),
		// t
		st.Literal("tai"), st.Literal("tan"), st.Literal("tas"), st.Literal("ther"), st.Literal("tia"),
		st.Literal("tin"), st.Literal("ton"), st.Literal("tor"), st.Literal("tur"),
		// u
		st.Literal("um"), st.Literal("und"), st.Literal("unt"), st.Literal("urn"), st.Literal("usk"),
		st.Literal("ust"),
		// v
		st.Literal("ver"), st.Literal("ves"), st.Literal("vor"),
		// w
		st.Literal("war"), st.Literal("wor"),
		// y
		st.Literal("yer"),
	}

	vowels = []fmt.Stringer{
		st.Literal("a"), st.Literal("e"), st.Literal("i"), st.Literal("o"), st.Literal("u"), st.Literal("y"),
	}

	vowelsCombo = []fmt.Stringer{
		st.Literal("a"), st.Literal("e"), st.Literal("i"), st.Literal("o"), st.Literal("u"),
		st.Literal("y"), st.Literal("ae"), st.Literal("ai"), st.Literal("au"), st.Literal("ay"),
		st.Literal("ea"), st.Literal("ee"), st.Literal("ei"), st.Literal("eu"), st.Literal("ey"),
		st.Literal("ia"), st.Literal("ie"), st.Literal("oe"), st.Literal("oi"), st.Literal("oo"),
		st.Literal("ou"), st.Literal("ui"),
	}

	consonants = []fmt.Stringer{
		st.Literal("b"), st.Literal("c"), st.Literal("d"), st.Literal("f"), st.Literal("g"), st.Literal("h"),
		st.Literal("j"), st.Literal("k"), st.Literal("l"), st.Literal("m"), st.Literal("n"), st.Literal("p"),
		st.Literal("q"), st.Literal("r"), st.Literal("s"), st.Literal("t"), st.Literal("v"), st.Literal("w"),
		st.Literal("x"), st.Literal("y"), st.Literal("z"),
	}

	consonantsComboBegin = []fmt.Stringer{
		st.Literal("b"), st.Literal("bl"), st.Literal("br"), st.Literal("c"), st.Literal("ch"), st.Literal("chr"),
		st.Literal("cl"), st.Literal("cr"), st.Literal("d"), st.Literal("dr"), st.Literal("f"), st.Literal("g"),
		st.Literal("h"), st.Literal("j"), st.Literal("k"), st.Literal("l"), st.Literal("ll"), st.Literal("m"),
		st.Literal("n"), st.Literal("p"), st.Literal("ph"), st.Literal("qu"), st.Literal("r"), st.Literal("rh"),
		st.Literal("s"), st.Literal("sch"), st.Literal("sh"), st.Literal("sl"), st.Literal("sm"), st.Literal("sn"),
		st.Literal("st"), st.Literal("str"), st.Literal("sw"), st.Literal("t"), st.Literal("th"), st.Literal("thr"),
		st.Literal("tr"), st.Literal("v"), st.Literal("w"), st.Literal("wh"), st.Literal("y"), st.Literal("z"),
		st.Literal("zh"),
	}

	consonantsComboAny = []fmt.Stringer{
		st.Literal("b"), st.Literal("c"), st.Literal("ch"), st.Literal("ck"), st.Literal("d"), st.Literal("f"),
		st.Literal("g"), st.Literal("gh"), st.Literal("h"), st.Literal("k"), st.Literal("l"), st.Literal("ld"),
		st.Literal("ll"), st.Literal("lt"), st.Literal("m"), st.Literal("n"), st.Literal("nd"), st.Literal("nn"),
		st.Literal("nt"), st.Literal("p"), st.Literal("ph"), st.Literal("q"), st.Literal("r"), st.Literal("rd"),
		st.Literal("rr"), st.Literal("rt"), st.Literal("s"), st.Literal("sh"), st.Literal("ss"), st.Literal("st"),
		st.Literal("t"), st.Literal("th"), st.Literal("v"), st.Literal("w"), st.Literal("y"), st.Literal("z"),
	}

	insults = []fmt.Stringer{
		st.Literal("air"), st.Literal("ankle"), st.Literal("ball"), st.Literal("beef"), st.Literal("bone"),
		st.Literal("bum"), st.Literal("bumble"), st.Literal("bump"), st.Literal("cheese"), st.Literal("clod"),
		st.Literal("clot"), st.Literal("clown"), st.Literal("corn"), st.Literal("dip"), st.Literal("dolt"),
		st.Literal("doof"), st.Literal("dork"), st.Literal("dumb"), st.Literal("face"), st.Literal("finger"),
		st.Literal("foot"), st.Literal("fumble"), st.Literal("goof"), st.Literal("grumble"), st.Literal("head"),
		st.Literal("knock"), st.Literal("knocker"), st.Literal("knuckle"), st.Literal("loaf"), st.Literal("lump"),
		st.Literal("lunk"), st.Literal("meat"), st.Literal("muck"), st.Literal("munch"), st.Literal("nit"),
		st.Literal("numb"), st.Literal("pin"), st.Literal("puff"), st.Literal("skull"), st.Literal("snark"),
		st.Literal("sneeze"), st.Literal("thimble"), st.Literal("twerp"), st.Literal("twit"), st.Literal("wad"),
		st.Literal("wimp"), st.Literal("wipe"),
	}

	mushies = []fmt.Stringer{
		st.Literal("baby"), st.Literal("booble"), st.Literal("bunker"), st.Literal("cuddle"), st.Literal("cuddly"),
		st.Literal("cutie"), st.Literal("doodle"), st.Literal("foofie"), st.Literal("gooble"), st.Literal("honey"),
		st.Literal("kissie"), st.Literal("lover"), st.Literal("lovey"), st.Literal("moofie"), st.Literal("mooglie"),
		st.Literal("moopie"), st.Literal("moopsie"), st.Literal("nookum"), st.Literal("poochie"), st.Literal("poof"),
		st.Literal("poofie"), st.Literal("pookie"), st.Literal("schmoopie"), st.Literal("schnoogle"),
		st.Literal("schnookie"), st.Literal("schnookum"), st.Literal("smooch"), st.Literal("smoochie"),
		st.Literal("smoosh"), st.Literal("snoogle"), st.Literal("snoogy"), st.Literal("snookie"), st.Literal("snookum"),
		st.Literal("snuggy"), st.Literal("sweetie"), st.Literal("woogle"), st.Literal("woogy"), st.Literal("wookie"),
		st.Literal("wookum"), st.Literal("wuddle"), st.Literal("wuddly"), st.Literal("wuggy"), st.Literal("wunny"),
	}

	mushiesEnds = []fmt.Stringer{
		st.Literal("boo"), st.Literal("bunch"), st.Literal("bunny"), st.Literal("cake"), st.Literal("cakes"),
		st.Literal("cute"), st.Literal("darling"), st.Literal("dumpling"), st.Literal("dumplings"), st.Literal("face"),
		st.Literal("foof"), st.Literal("goo"), st.Literal("head"), st.Literal("kin"), st.Literal("kins"),
		st.Literal("lips"), st.Literal("love"), st.Literal("mush"), st.Literal("pie"), st.Literal("poo"),
		st.Literal("pooh"), st.Literal("pook"), st.Literal("pums"),
	}

	consonantsStupid = []fmt.Stringer{
		st.Literal("b"), st.Literal("bl"), st.Literal("br"), st.Literal("cl"), st.Literal("d"), st.Literal("f"),
		st.Literal("fl"), st.Literal("fr"), st.Literal("g"), st.Literal("gh"), st.Literal("gl"), st.Literal("gr"),
		st.Literal("h"), st.Literal("j"), st.Literal("k"), st.Literal("kl"), st.Literal("m"), st.Literal("n"),
		st.Literal("p"), st.Literal("th"), st.Literal("w"),
	}

	syllablesStupid = []fmt.Stringer{
		st.Literal("elch"), st.Literal("idiot"), st.Literal("ob"), st.Literal("og"), st.Literal("ok"),
		st.Literal("olph"), st.Literal("olt"), st.Literal("omph"), st.Literal("ong"), st.Literal("onk"),
		st.Literal("oo"), st.Literal("oob"), st.Literal("oof"), st.Literal("oog"), st.Literal("ook"),
		st.Literal("ooz"), st.Literal("org"), st.Literal("ork"), st.Literal("orm"), st.Literal("oron"),
		st.Literal("ub"), st.Literal("uck"), st.Literal("ug"), st.Literal("ulf"), st.Literal("ult"),
		st.Literal("um"), st.Literal("umb"), st.Literal("ump"), st.Literal("umph"), st.Literal("un"),
		st.Literal("unb"), st.Literal("ung"), st.Literal("unk"), st.Literal("unph"), st.Literal("unt"),
		st.Literal("uzz"),
	}

	symbolMap = map[rune][]fmt.Stringer{
		's': syllables,
		'v': vowels,
		'V': vowelsCombo,
		'c': consonants,
		'B': consonantsComboBegin,
		'C': consonantsComboAny,
		'i': insults,
		'm': mushies,
		'M': mushiesEnds,
		'D': consonantsStupid,
		'd': syllablesStupid,
	}
)
