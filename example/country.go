package example

//go:generate enumeration -v -type Country -plural Countries -ic -unsnake -alias iso3166_3LetterCodes

// Country example shows use of the '-plural' option to set the name of plural
// collections. Because of '-ic', the parser ignores case. Because of '-unsnake', the parser
// treats underscores and spaces alike.
//
// When generating JSON & SQL values, the 2-letter ISO-3166 country codes are used. This is
// determined by the `all` tags in comments. The 2-letter codes are also accepted when parsing.
//
// The '-alias' option also allows 3-letter ISO-3166 tags to be parsed as their corresponding
// country using the map iso3166_3LetterCodes.
type Country int

const (
	Afghanistan                                  Country = iota // all:"af"
	Åland_Islands                                               // all:"ax"
	Albania                                                     // all:"al"
	Algeria                                                     // all:"dz"
	American_Samoa                                              // all:"as"
	Andorra                                                     // all:"ad"
	Angola                                                      // all:"ao"
	Anguilla                                                    // all:"ai"
	Antarctica                                                  // all:"aq"
	Antigua_and_Barbuda                                         // all:"ag"
	Argentina                                                   // all:"ar"
	Armenia                                                     // all:"am"
	Aruba                                                       // all:"aw"
	Australia                                                   // all:"au"
	Austria                                                     // all:"at"
	Azerbaijan                                                  // all:"az"
	Bahamas                                                     // all:"bs"
	Bahrain                                                     // all:"bh"
	Bangladesh                                                  // all:"bd"
	Barbados                                                    // all:"bb"
	Belarus                                                     // all:"by"
	Belgium                                                     // all:"be"
	Belize                                                      // all:"bz"
	Benin                                                       // all:"bj"
	Bermuda                                                     // all:"bm"
	Bhutan                                                      // all:"bt"
	Bolivia                                                     // all:"bo"
	Bosnia_and_Herzegovina                                      // all:"ba"
	Botswana                                                    // all:"bw"
	Bouvet_Island                                               // all:"bv"
	Brazil                                                      // all:"br"
	British_Virgin_Islands                                      // all:"vg"
	British_Indian_Ocean_Territory                              // all:"io"
	Brunei_Darussalam                                           // all:"dn"
	Bulgaria                                                    // all:"bg"
	Burkina_Faso                                                // all:"bf"
	Burundi                                                     // all:"bi"
	Cambodia                                                    // all:"kh"
	Cameroon                                                    // all:"cm"
	Canada                                                      // all:"ca"
	Cape_Verde                                                  // all:"cv"
	Cayman_Islands                                              // all:"ky"
	Central_African_Republic                                    // all:"cf"
	Chad                                                        // all:"td"
	Chile                                                       // all:"cl"
	China                                                       // all:"cn"
	Hong_Kong                                                   // all:"hk" -- Special Administrative Region of China
	Macao                                                       // all:"mo" -- Special Administrative Region of China
	Christmas_Island                                            // all:"cx"
	Cocos_Islands                                               // all:"cc" -- Cocos (Keeling) Islands
	Colombia                                                    // all:"co"
	Comoros                                                     // all:"km"
	Congo_Brazzaville                                           // all:"cg" -- Congo (Brazzaville)
	Congo_DRC                                                   // all:"cd" -- Democratic Republic of the Congo
	Cook_Islands                                                // all:"ck"
	Costa_Rica                                                  // all:"cr"
	Côte_dIvoire                                                // all:"ci" -- Côte_d’Ivoire
	Croatia                                                     // all:"hr"
	Cuba                                                        // all:"cu"
	Cyprus                                                      // all:"cy"
	Czechia                                                     // all:"cz"
	Denmark                                                     // all:"dk"
	Djibouti                                                    // all:"dj"
	Dominica                                                    // all:"dm"
	Dominican_Republic                                          // all:"do"
	Ecuador                                                     // all:"ec"
	Egypt                                                       // all:"eg"
	El_Salvador                                                 // all:"sv"
	Equatorial_Guinea                                           // all:"gq"
	Eritrea                                                     // all:"er"
	Estonia                                                     // all:"ee"
	Ethiopia                                                    // all:"et"
	Falkland_Islands                                            // all:"fk" -- Falkland Islands (Malvinas)
	Faroe_Islands                                               // all:"fo"
	Fiji                                                        // all:"fj"
	Finland                                                     // all:"fi"
	France                                                      // all:"fr"
	French_Guiana                                               // all:"gf"
	French_Polynesia                                            // all:"pf"
	French_Southern_Territories                                 // all:"tf"
	Gabon                                                       // all:"ga"
	Gambia                                                      // all:"gm"
	Georgia                                                     // all:"ge"
	Germany                                                     // all:"de"
	Ghana                                                       // all:"gh"
	Gibraltar                                                   // all:"gi"
	Greece                                                      // all:"gr"
	Greenland                                                   // all:"gl"
	Grenada                                                     // all:"gd"
	Guadeloupe                                                  // all:"gp"
	Guam                                                        // all:"gu"
	Guatemala                                                   // all:"gt"
	Guernsey                                                    // all:"gg"
	Guinea                                                      // all:"gn"
	Guinea_Bissau                                               // all:"gw"
	Guyana                                                      // all:"gy"
	Haiti                                                       // all:"ht"
	Heard_Island_and_Mcdonald_Islands                           // all:"hm"
	Holy_See                                                    // all:"va" -- Vatican City State
	Honduras                                                    // all:"hn"
	Hungary                                                     // all:"hu"
	Iceland                                                     // all:"is"
	India                                                       // all:"in"
	Indonesia                                                   // all:"id"
	Iran                                                        // all:"ir" -- Islamic Republic of Iran
	Iraq                                                        // all:"iq"
	Ireland                                                     // all:"ie"
	Isle_of_Man                                                 // all:"im"
	Israel                                                      // all:"il"
	Italy                                                       // all:"it"
	Jamaica                                                     // all:"jm"
	Japan                                                       // all:"jp"
	Jersey                                                      // all:"je"
	Jordan                                                      // all:"jo"
	Kazakhstan                                                  // all:"kz"
	Kenya                                                       // all:"ke"
	Kiribati                                                    // all:"ki"
	Democratic_Peoples_Republic_of_Korea                        // all:"kp" -- Democratic People’s Republic of Korea
	South_Korea                                                 // all:"kr" -- Republic of Korea
	Kuwait                                                      // all:"kw"
	Kyrgyzstan                                                  // all:"kg"
	Lao_PDR                                                     // all:"la"
	Latvia                                                      // all:"lv"
	Lebanon                                                     // all:"lb"
	Lesotho                                                     // all:"ls"
	Liberia                                                     // all:"lr"
	Libya                                                       // all:"ly"
	Liechtenstein                                               // all:"li"
	Lithuania                                                   // all:"lt"
	Luxembourg                                                  // all:"lu"
	North_Macedonia                                             // all:"mk" -- Republic of North_Macedonia
	Madagascar                                                  // all:"mg"
	Malawi                                                      // all:"mw"
	Malaysia                                                    // all:"my"
	Maldives                                                    // all:"mv"
	Mali                                                        // all:"ml"
	Malta                                                       // all:"mt"
	Marshall_Islands                                            // all:"mh"
	Martinique                                                  // all:"mq"
	Mauritania                                                  // all:"mr"
	Mauritius                                                   // all:"mu"
	Mayotte                                                     // all:"yt"
	Mexico                                                      // all:"mx"
	Micronesia                                                  // all:"fm" -- Federated States of Micronesia
	Moldova                                                     // all:"md"
	Monaco                                                      // all:"mc"
	Mongolia                                                    // all:"mn"
	Montenegro                                                  // all:"me"
	Montserrat                                                  // all:"ms"
	Morocco                                                     // all:"ma"
	Mozambique                                                  // all:"mz"
	Myanmar                                                     // all:"mm"
	Namibia                                                     // all:"na"
	Nauru                                                       // all:"nr"
	Nepal                                                       // all:"np"
	Netherlands                                                 // all:"nl",
	Netherlands_Antilles                                        // all:"an" -- obsolete
	New_Caledonia                                               // all:"nc",
	New_Zealand                                                 // all:"nz",
	Nicaragua                                                   // all:"ni",
	Niger                                                       // all:"ne",
	Nigeria                                                     // all:"ng",
	Niue                                                        // all:"nu",
	Norfolk_Island                                              // all:"nf",
	Northern_Mariana_Islands                                    // all:"mp",
	Norway                                                      // all:"no",
	Oman                                                        // all:"om",
	Pakistan                                                    // all:"pk",
	Palau                                                       // all:"pw",
	Palestinian_Territory                                       // all:"ps", -- Palestinian Territory, Occupied
	Panama                                                      // all:"pa",
	Papua_New_Guinea                                            // all:"pg",
	Paraguay                                                    // all:"py",
	Peru                                                        // all:"pe",
	Philippines                                                 // all:"ph",
	Pitcairn                                                    // all:"pn",
	Poland                                                      // all:"pl",
	Portugal                                                    // all:"pt",
	Puerto_Rico                                                 // all:"pr",
	Qatar                                                       // all:"qa",
	Réunion                                                     // all:"re",
	Romania                                                     // all:"ro",
	Russian_Federation                                          // all:"ru",
	Rwanda                                                      // all:"rw",
	Saint_Barthélemy                                            // all:"bl",
	Saint_Helena                                                // all:"sh",
	Saint_Kitts_and_Nevis                                       // all:"kn",
	Saint_Lucia                                                 // all:"lc"
	Saint_Martin                                                // all:"mf" -- Saint-Martin (French part)
	Saint_Pierre_and_Miquelon                                   // all:"pm"
	Saint_Vincent_and_Grenadines                                // all:"vc"
	Samoa                                                       // all:"ws"
	San_Marino                                                  // all:"sm"
	Sao_Tome_and_Principe                                       // all:"st"
	Saudi_Arabia                                                // all:"sa"
	Senegal                                                     // all:"sn"
	Serbia                                                      // all:"rs"
	Seychelles                                                  // all:"sc"
	Sierra_Leone                                                // all:"sl"
	Singapore                                                   // all:"sg"
	Slovakia                                                    // all:"sk"
	Slovenia                                                    // all:"si"
	Solomon_Islands                                             // all:"sb"
	Somalia                                                     // all:"so"
	South_Africa                                                // all:"za"
	South_Georgia_and_the_South_Sandwich_Islands                // all:"gs"
	South_Sudan                                                 // all:"ss"
	Spain                                                       // all:"es"
	Sri_Lanka                                                   // all:"lk"
	Sudan                                                       // all:"sd"
	Suriname                                                    // all:"sr"
	Svalbard_and_Jan_Mayen_Islands                              // all:"sj"
	Swaziland                                                   // all:"sz" -- Eswatini
	Sweden                                                      // all:"se"
	Switzerland                                                 // all:"ch"
	Syria                                                       // all:"sy" -- Syrian Arab Republic
	Taiwan                                                      // all:"tw" -- Taiwan, Republic of China
	Tajikistan                                                  // all:"tj"
	Tanzania                                                    // all:"tz" -- United Republic of Tanzania
	Thailand                                                    // all:"th"
	Timor_Leste                                                 // all:"tl"
	Togo                                                        // all:"tg"
	Tokelau                                                     // all:"tk"
	Tonga                                                       // all:"to"
	Trinidad_and_Tobago                                         // all:"tt"
	Tunisia                                                     // all:"tn"
	Türkiye                                                     // all:"tr"
	Turkmenistan                                                // all:"tm"
	Turks_and_Caicos_Islands                                    // all:"tc"
	Tuvalu                                                      // all:"tv"
	Uganda                                                      // all:"ug"
	Ukraine                                                     // all:"ua"
	United_Arab_Emirates                                        // all:"ae"
	United_Kingdom                                              // all:"gb"
	United_States_of_America                                    // all:"us"
	United_States_Minor_Outlying_Islands                        // all:"um"
	Uruguay                                                     // all:"uy"
	Uzbekistan                                                  // all:"uz"
	Vanuatu                                                     // all:"vu"
	Venezuela                                                   // all:"ve" -- Bolivarian Republic of Venezuela
	Viet_Nam                                                    // all:"vn"
	Virgin_Islands                                              // all:"vi"
	Wallis_and_Futuna_Islands                                   // all:"wf"
	Western_Sahara                                              // all:"eh"
	Yemen                                                       // all:"ye"
	Zambia                                                      // all:"zm"
	Zimbabwe                                                    // all:"zw"
)

var iso3166_3LetterCodes = map[string]Country{
	"afg": Afghanistan,
	"axa": Åland_Islands,
	"alb": Albania,
	"dza": Algeria,
	"asm": American_Samoa,
	"and": Andorra,
	"ago": Angola,
	"aia": Anguilla,
	"ata": Antarctica,
	"atg": Antigua_and_Barbuda,
	"arg": Argentina,
	"arm": Armenia,
	"abw": Aruba,
	"aus": Australia,
	"aut": Austria,
	"aze": Azerbaijan,
	"bhs": Bahamas,
	"bhr": Bahrain,
	"bgd": Bangladesh,
	"brb": Barbados,
	"blr": Belarus,
	"bel": Belgium,
	"blz": Belize,
	"ben": Benin,
	"bmu": Bermuda,
	"btn": Bhutan,
	"bol": Bolivia,
	"bih": Bosnia_and_Herzegovina,
	"bwa": Botswana,
	"bvt": Bouvet_Island,
	"bra": Brazil,
	"vgb": British_Virgin_Islands,
	"iot": British_Indian_Ocean_Territory,
	"brn": Brunei_Darussalam,
	"bgr": Bulgaria,
	"bfa": Burkina_Faso,
	"boi": Burundi,
	"khh": Cambodia,
	"cmr": Cameroon,
	"can": Canada,
	"cpv": Cape_Verde,
	"cym": Cayman_Islands,
	"caf": Central_African_Republic,
	"tcd": Chad,
	"chl": Chile,
	"chn": China,
	"hkg": Hong_Kong,
	"mac": Macao,
	"cxr": Christmas_Island,
	"cck": Cocos_Islands,
	"col": Colombia,
	"com": Comoros,
	"cog": Congo_Brazzaville,
	"cod": Congo_DRC,
	"cok": Cook_Islands,
	"cri": Costa_Rica,
	"civ": Côte_dIvoire,
	"hrv": Croatia,
	"cub": Cuba,
	"cyp": Cyprus,
	"cze": Czechia,
	"dnk": Denmark,
	"dji": Djibouti,
	"dma": Dominica,
	"dom": Dominican_Republic,
	"ecu": Ecuador,
	"egy": Egypt,
	"slv": El_Salvador,
	"gnq": Equatorial_Guinea,
	"eri": Eritrea,
	"est": Estonia,
	"eth": Ethiopia,
	"flk": Falkland_Islands,
	"fro": Faroe_Islands,
	"fji": Fiji,
	"fin": Finland,
	"fra": France,
	"guf": French_Guiana,
	"pyf": French_Polynesia,
	"atf": French_Southern_Territories,
	"gab": Gabon,
	"gmb": Gambia,
	"geo": Georgia,
	"deu": Germany,
	"gha": Ghana,
	"gib": Gibraltar,
	"grc": Greece,
	"grl": Greenland,
	"grd": Grenada,
	"glp": Guadeloupe,
	"gum": Guam,
	"gtm": Guatemala,
	"ggy": Guernsey,
	"gin": Guinea,
	"gnb": Guinea_Bissau,
	"guy": Guyana,
	"hti": Haiti,
	"hmd": Heard_Island_and_Mcdonald_Islands,
	"vat": Holy_See,
	"hnd": Honduras,
	"hun": Hungary,
	"isl": Iceland,
	"ind": India,
	"idn": Indonesia,
	"irn": Iran,
	"irq": Iraq,
	"irl": Ireland,
	"imn": Isle_of_Man,
	"isr": Israel,
	"ita": Italy,
	"jam": Jamaica,
	"jpn": Japan,
	"jey": Jersey,
	"jor": Jordan,
	"kaz": Kazakhstan,
	"ken": Kenya,
	"kir": Kiribati,
	"prk": Democratic_Peoples_Republic_of_Korea,
	"kor": South_Korea,
	"kwt": Kuwait,
	"kgz": Kyrgyzstan,
	"lao": Lao_PDR,
	"lva": Latvia,
	"lbn": Lebanon,
	"lso": Lesotho,
	"lbr": Liberia,
	"lby": Libya,
	"lie": Liechtenstein,
	"ltu": Lithuania,
	"lux": Luxembourg,
	"mkd": North_Macedonia,
	"mdg": Madagascar,
	"mwi": Malawi,
	"mys": Malaysia,
	"mdv": Maldives,
	"mli": Mali,
	"mlt": Malta,
	"mhl": Marshall_Islands,
	"mtq": Martinique,
	"mrt": Mauritania,
	"mus": Mauritius,
	"myt": Mayotte,
	"mex": Mexico,
	"fsm": Micronesia,
	"mda": Moldova,
	"mco": Monaco,
	"mng": Mongolia,
	"mne": Montenegro,
	"msr": Montserrat,
	"mar": Morocco,
	"moz": Mozambique,
	"mmr": Myanmar,
	"nam": Namibia,
	"nru": Nauru,
	"npl": Nepal,
	"nld": Netherlands,
	"ant": Netherlands_Antilles,
	"ncl": New_Caledonia,
	"nzl": New_Zealand,
	"nic": Nicaragua,
	"ner": Niger,
	"nga": Nigeria,
	"niu": Niue,
	"nfk": Norfolk_Island,
	"mnp": Northern_Mariana_Islands,
	"nor": Norway,
	"omn": Oman,
	"pak": Pakistan,
	"plw": Palau,
	"pse": Palestinian_Territory,
	"pan": Panama,
	"png": Papua_New_Guinea,
	"pry": Paraguay,
	"per": Peru,
	"phl": Philippines,
	"pcn": Pitcairn,
	"pol": Poland,
	"prt": Portugal,
	"pri": Puerto_Rico,
	"qat": Qatar,
	"reu": Réunion,
	"rou": Romania,
	"rus": Russian_Federation,
	"rwa": Rwanda,
	"blm": Saint_Barthélemy,
	"shn": Saint_Helena,
	"kna": Saint_Kitts_and_Nevis,
	"lca": Saint_Lucia,
	"mfa": Saint_Martin,
	"ppm": Saint_Pierre_and_Miquelon,
	"vct": Saint_Vincent_and_Grenadines,
	"wsm": Samoa,
	"smr": San_Marino,
	"stp": Sao_Tome_and_Principe,
	"sau": Saudi_Arabia,
	"sen": Senegal,
	"srb": Serbia,
	"syc": Seychelles,
	"sle": Sierra_Leone,
	"sgp": Singapore,
	"svk": Slovakia,
	"svn": Slovenia,
	"slb": Solomon_Islands,
	"som": Somalia,
	"zaf": South_Africa,
	"sgs": South_Georgia_and_the_South_Sandwich_Islands,
	"ssd": South_Sudan,
	"esp": Spain,
	"lka": Sri_Lanka,
	"sdn": Sudan,
	"sur": Suriname,
	"sjm": Svalbard_and_Jan_Mayen_Islands,
	"swz": Swaziland,
	"swe": Sweden,
	"che": Switzerland,
	"syr": Syria,
	"twn": Taiwan,
	"tjk": Tajikistan,
	"tza": Tanzania,
	"tha": Thailand,
	"tls": Timor_Leste,
	"tgo": Togo,
	"tkl": Tokelau,
	"ton": Tonga,
	"ttd": Trinidad_and_Tobago,
	"tun": Tunisia,
	"tur": Türkiye,
	"tkm": Turkmenistan,
	"tca": Turks_and_Caicos_Islands,
	"tuv": Tuvalu,
	"uga": Uganda,
	"ukr": Ukraine,
	"are": United_Arab_Emirates,
	"gbr": United_Kingdom,
	"usa": United_States_of_America,
	"umi": United_States_Minor_Outlying_Islands,
	"ury": Uruguay,
	"uzb": Uzbekistan,
	"vut": Vanuatu,
	"ven": Venezuela,
	"vnm": Viet_Nam,
	"vir": Virgin_Islands,
	"wlf": Wallis_and_Futuna_Islands,
	"esh": Western_Sahara,
	"yem": Yemen,
	"zmb": Zambia,
	"zwe": Zimbabwe,
}
