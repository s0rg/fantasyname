package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	fn "github.com/s0rg/fantasyname"
)

const (
	MIDDLE_EARTH = "(bil|bal|ban|hil|ham|hal|hol|hob|wil|me|or|ol|od|gor|for|fos|tol|ar|fin|ere|leo|vi|bi|bren|thor)" +
		"(|go|orbis|apol|adur|mos|ri|i|na|ole|n)" +
		"(|tur|axia|and|bo|gil|bin|bras|las|mac|grim|wise|l|lo|fo|co|ra|via|" +
		"da|ne|ta|y|wen|thiel|phin|dir|dor|tor|rod|on|rdo|dis)"
	JAPANESE_NAMES_CONSTRAINED = "(aka|aki|bashi|gawa|kawa|furu|fuku|fuji|hana|hara|haru|hashi|hira|hon|hoshi|" +
		"ichi|iwa|kami|kawa|ki|kita|kuchi|kuro|marui|matsu|miya|mori|moto|mura|nabe|naka|nishi|no|da|ta|o|oo|oka|" +
		"saka|saki|sawa|shita|shima|i|suzu|taka|take|to|toku|toyo|ue|wa|wara|wata|yama|yoshi|kei|ko|zawa|zen|sen|" +
		"ao|gin|kin|ken|shiro|zaki|yuki|asa)(||||||||||bashi|gawa|kawa|furu|fuku|fuji|hana|hara|haru|hashi|hira|" +
		"hon|hoshi|chi|wa|ka|kami|kawa|ki|kita|kuchi|kuro|marui|matsu|miya|mori|moto|mura|nabe|naka|nishi|no|da|" +
		"ta|o|oo|oka|saka|saki|sawa|shita|shima|suzu|taka|take|to|toku|toyo|ue|wa|wara|wata|yama|yoshi|kei|ko|" +
		"zawa|zen|sen|ao|gin|kin|ken|shiro|zaki|yuki|sa)"
	JAPANESE_NAMES_DIVERSE = "(a|i|u|e|o|||||)" +
		"(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|" +
		"no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|" +
		"re|ro|ro|ro|wa|wa|wa|wa|wo|wo)(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|" +
		"chi|tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|" +
		"yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)" +
		"(|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|" +
		"no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|" +
		"ro|ro|ro|wa|wa|wa|wa|wo|wo)|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|" +
		"tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|" +
		"yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)(|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|" +
		"shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|" +
		"mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)))(|||n)"
	CHINESE_NAMES = "(zh|x|q|sh|h)(ao|ian|uo|ou|ia)(|(l|w|c|p|b|m)(ao|ian|uo|ou|ia)" +
		"(|n)|-(l|w|c|p|b|m)(ao|ian|uo|ou|ia)(|(d|j|q|l)(a|ai|iu|ao|i)))"
	GREEK_NAMES      = "<s<v|V>(tia)|s<v|V>(os)|B<v|V>c(ios)|B<v|V><c|C>v(ios|os)>"
	HAWAIIAN_NAMES_1 = "((h|k|l|m|n|p|w|')|)(a|e|i|o|u)((h|k|l|m|n|p|w|')|)(a|e|i|o|u)(((h|k|l|m|n|p|w|')|)" +
		"(a|e|i|o|u)|)(((h|k|l|m|n|p|w|')|)(a|e|i|o|u)|)(((h|k|l|m|n|p|w|')|)(a|e|i|o|u)|)(((h|k|l|m|n|p|w|')|)" +
		"(a|e|i|o|u)|)"
	HAWAIIAN_NAMES_2 = "((h|k|l|m|n|p|w|)(a|e|i|o|u|a'|e'|i'|o'|u'|ae|ai|ao|au|oi|ou|eu|ei)" +
		"(k|l|m|n|p|)|)(h|k|l|m|n|p|w|)(a|e|i|o|u|a'|e'|i'|o'|u'|ae|ai|ao|au|oi|ou|eu|ei)(k|l|m|n|p|)"
	OLD_LATIN_PLACE_NAMES = "sv(nia|lia|cia|sia)"
	DRAGONS_PERN          = "<<s|ss>|<VC|vC|B|BVs|Vs>><v|V|v|<v(l|n|r)|vc>>(th)"
	DRAGON_RIDERS         = "c'<s|cvc>"
	POKEMON               = "<i|s>v(mon|chu|zard|rtle)"
	FANTASY_VOWELS_R      = "(|(<B>|s|h|ty|ph|r))(i|ae|ya|ae|eu|ia|i|eo|ai|a)" +
		"(lo|la|sri|da|dai|the|sty|lae|due|li|lly|ri|na|ral|sur|rith)(|(su|nu|sti|llo|ria|))" +
		"(|(n|ra|p|m|lis|cal|deu|dil|suir|phos|ru|dru|rin|raap|rgue))"
	FANTASY_S_A = "(cham|chan|jisk|lis|frich|isk|lass|mind|sond|sund|ass|chad|lirt|und|mar|lis|il|<BVC>)" +
		"(jask|ast|ista|adar|irra|im|ossa|assa|osia|ilsa|<vCv>)(|(an|ya|la|sta|sda|sya|st|nya))"
	FANTASY_H_L = "(ch|ch't|sh|cal|val|ell|har|shar|shal|rel|laen|ral|jh't|alr|ch|ch't|av)" +
		"(|(is|al|ow|ish|ul|el|ar|iel))" +
		"(aren|aeish|aith|even|adur|ulash|alith|atar|aia|erin|aera|ael|ira|iel|ahur|ishul)"
	FANTASY_N_L = "(ethr|qil|mal|er|eal|far|fil|fir|ing|ind|il|lam|quel|quar|quan|qar|pal|mal|yar|um|ard|enn|ey)" +
		"(|(<vc>|on|us|un|ar|as|en|ir|ur|at|ol|al|an))" +
		"(uard|wen|arn|on|il|ie|on|iel|rion|rian|an|ista|rion|rian|cil|mol|yon)"
	FANTASY_K_N = "(taith|kach|chak|kank|kjar|rak|kan|kaj|tach|rskal|kjol|jok|jor|jad|kot|kon|" +
		"knir|kror|kol|tul|rhaok|rhak|krol|jan|kag|ryr)(<vc>|in|or|an|ar|och|un|mar|yk|ja|arn|ir|ros|ror)" +
		"(|(mund|ard|arn|karr|chim|kos|rir|arl|kni|var|an|in|ir|a|i|as))"
	FANTASY_J_G_Z = "(aj|ch|etz|etzl|tz|kal|gahn|kab|aj|izl|ts|jaj|lan|kach|chaj|qaq|jol|ix|az|biq|nam)" +
		"(|(<vc>|aw|al|yes|il|ay|en|tom||oj|im|ol|aj|an|as))" +
		"(aj|am|al|aqa|ende|elja|ich|ak|ix|in|ak|al|il|ek|ij|os|al|im)"
	FANTASY_K_J_Y = "(yi|shu|a|be|na|chi|cha|cho|ksa|yi|shu)" +
		"(th|dd|jj|sh|rr|mk|n|rk|y|jj|th)" +
		"(us|ash|eni|akra|nai|ral|ect|are|el|urru|aja|al|uz|ict|arja|ichi|ural|iru|aki|esh)"
	FANTASY_S_E = "(syth|sith|srr|sen|yth|ssen|then|fen|ssth|kel|syn|est|bess|inth|nen|tin|cor|" +
		"sv|iss|ith|sen|slar|ssil|sthen|svis|s|ss|s|ss)" +
		"(|(tys|eus|yn|of|es|en|ath|elth|al|ell|ka|ith|yrrl|is|isl|yr|ast|iy))" +
		"(us|yn|en|ens|ra|rg|le|en|ith|ast|zon|in|yn|ys)"
)

type template struct {
	Name string
	Tmpl string
}

var templates = []template{
	{Name: "Middle Earth", Tmpl: MIDDLE_EARTH},
	{Name: "Japanese Names (Constrained)", Tmpl: JAPANESE_NAMES_CONSTRAINED},
	{Name: "Japanese Names (Diverse)", Tmpl: JAPANESE_NAMES_DIVERSE},
	{Name: "Chinese Names", Tmpl: CHINESE_NAMES},
	{Name: "Greek Names", Tmpl: GREEK_NAMES},
	{Name: "Hawaiian Names (1)", Tmpl: HAWAIIAN_NAMES_1},
	{Name: "Hawaiian Names (2)", Tmpl: HAWAIIAN_NAMES_2},
	{Name: "Old Latin Place Names", Tmpl: OLD_LATIN_PLACE_NAMES},
	{Name: "Dragons (Pern)", Tmpl: DRAGONS_PERN},
	{Name: "Dragon Riders", Tmpl: DRAGON_RIDERS},
	{Name: "Pokemon", Tmpl: POKEMON},
	{Name: "Fantasy (Vowels, R, etc.)", Tmpl: FANTASY_VOWELS_R},
	{Name: "Fantasy (S, A, etc.)", Tmpl: FANTASY_S_A},
	{Name: "Fantasy (H, L, etc.)", Tmpl: FANTASY_H_L},
	{Name: "Fantasy (N, L, etc.)", Tmpl: FANTASY_N_L},
	{Name: "Fantasy (K, N, etc.)", Tmpl: FANTASY_K_N},
	{Name: "Fantasy (J, G, Z, etc.)", Tmpl: FANTASY_J_G_Z},
	{Name: "Fantasy (K, J, Y, etc.)", Tmpl: FANTASY_K_J_Y},
	{Name: "Fantasy (S, E, etc.)", Tmpl: FANTASY_S_E},
	{Name: "Funny", Tmpl: "sdD i-i"},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(templates); i++ {
		t := &templates[i]

		gen, err := fn.Compile(t.Tmpl)
		if err != nil {
			log.Fatal(t.Name, err)
		}

		fmt.Println(t.Name)

		for i := 0; i < 10; i++ {
			fmt.Println("\t" + gen.String())
		}
	}
}
