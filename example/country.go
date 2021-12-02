package example

// This example shows use of the '-plural' option to set the name of plural
// collections. Because of '-ic', the parser ignores case. The '-using' option
// provides cross-mapping between the country names and their ISO-3166 tags.

//go:generate enumeration -v -type Country -plural Countries -ic -using iso3166Tags

type Country int

const (
	Afghanistan Country = iota
	Aland_Islands
	Albania
	Algeria
	American_Samoa
	Andorra
	Angola
	Anguilla
	Antarctica
	Antigua_and_Barbuda
	Argentina
	Armenia
	Aruba
	Australia
	Austria
	Azerbaijan
	Bahamas
	Bahrain
	Bangladesh
	Barbados
	Belarus
	Belgium
	Belize
	Benin
	Bermuda
	Bhutan
	Bolivia
	Bosnia_and_Herzegovina
	Botswana
	Bouvet_Island
	Brazil
	British_Virgin_Islands
	British_Indian_Ocean_Territory
	Brunei_Darussalam
	Bulgaria
	Burkina_Faso
	Burundi
	Cambodia
	Cameroon
	Canada
	Cape_Verde
	Cayman_Islands
	Central_African_Republic
	Chad
	Chile
	China
	Hong_Kong // Special Administrative Region of China
	Macao     // Special Administrative Region of China
	Christmas_Island
	Cocos_Islands // Cocos (Keeling) Islands
	Colombia
	Comoros
	Congo_Brazzaville // Congo (Brazzaville)
	Congo_DRC         // Democratic Republic of the Congo
	Cook_Islands
	Costa_Rica
	Côte_dIvoire // Côte_d’Ivoire
	Croatia
	Cuba
	Cyprus
	Czech_Republic
	Denmark
	Djibouti
	Dominica
	Dominican_Republic
	Ecuador
	Egypt
	El_Salvador
	Equatorial_Guinea
	Eritrea
	Estonia
	Ethiopia
	Falkland_Islands // Falkland Islands (Malvinas)
	Faroe_Islands
	Fiji
	Finland
	France
	French_Guiana
	French_Polynesia
	French_Southern_Territories
	Gabon
	Gambia
	Georgia
	Germany
	Ghana
	Gibraltar
	Greece
	Greenland
	Grenada
	Guadeloupe
	Guam
	Guatemala
	Guernsey
	Guinea
	Guinea_Bissau
	Guyana
	Haiti
	Heard_Island_and_Mcdonald_Islands
	Holy_See // Vatican City State
	Honduras
	Hungary
	Iceland
	India
	Indonesia
	Iran // Islamic Republic of Iran
	Iraq
	Ireland
	Isle_of_Man
	Israel
	Italy
	Jamaica
	Japan
	Jersey
	Jordan
	Kazakhstan
	Kenya
	Kiribati
	Democratic_Peoples_Republic_of_Korea // Democratic People’s Republic of Korea
	South_Korea                          // Republic of Korea
	Kuwait
	Kyrgyzstan
	Lao_PDR
	Latvia
	Lebanon
	Lesotho
	Liberia
	Libya
	Liechtenstein
	Lithuania
	Luxembourg
	Macedonia // Republic of Macedonia
	Madagascar
	Malawi
	Malaysia
	Maldives
	Mali
	Malta
	Marshall_Islands
	Martinique
	Mauritania
	Mauritius
	Mayotte
	Mexico
	Micronesia // Federated States of Micronesia
	Moldova
	Monaco
	Mongolia
	Montenegro
	Montserrat
	Morocco
	Mozambique
	Myanmar
	Namibia
	Nauru
	Nepal
	Netherlands
	Netherlands_Antilles
	New_Caledonia
	New_Zealand
	Nicaragua
	Niger
	Nigeria
	Niue
	Norfolk_Island
	Northern_Mariana_Islands
	Norway
	Oman
	Pakistan
	Palau
	Palestinian_Territory // Palestinian Territory, Occupied
	Panama
	Papua_New_Guinea
	Paraguay
	Peru
	Philippines
	Pitcairn
	Poland
	Portugal
	Puerto_Rico
	Qatar
	Réunion
	Romania
	Russian_Federation
	Rwanda
	Saint_Barthélemy
	Saint_Helena
	Saint_Kitts_and_Nevis
	Saint_Lucia
	Saint_Martin // Saint-Martin (French part)
	Saint_Pierre_and_Miquelon
	Saint_Vincent_and_Grenadines
	Samoa
	San_Marino
	Sao_Tome_and_Principe
	Saudi_Arabia
	Senegal
	Serbia
	Seychelles
	Sierra_Leone
	Singapore
	Slovakia
	Slovenia
	Solomon_Islands
	Somalia
	South_Africa
	South_Georgia_and_the_South_Sandwich_Islands
	South_Sudan
	Spain
	Sri_Lanka
	Sudan
	Suriname
	Svalbard_and_Jan_Mayen_Islands
	Swaziland // Eswatini
	Sweden
	Switzerland
	Syria  // Syrian Arab Republic
	Taiwan // Taiwan, Republic of China
	Tajikistan
	Tanzania // United Republic of Tanzania
	Thailand
	Timor_Leste
	Togo
	Tokelau
	Tonga
	Trinidad_and_Tobago
	Tunisia
	Turkey
	Turkmenistan
	Turks_and_Caicos_Islands
	Tuvalu
	Uganda
	Ukraine
	United_Arab_Emirates
	United_Kingdom
	United_States_of_America
	United_States_Minor_Outlying_Islands
	Uruguay
	Uzbekistan
	Vanuatu
	Venezuela // Bolivarian Republic of Venezuela
	Viet_Nam
	Virgin_Islands
	Wallis_and_Futuna_Islands
	Western_Sahara
	Yemen
	Zambia
	Zimbabwe
)

var iso3166Tags = map[Country]string{
	Afghanistan:                          "af",
	Aland_Islands:                        "ax",
	Albania:                              "al",
	Algeria:                              "dz",
	American_Samoa:                       "as",
	Andorra:                              "ad",
	Angola:                               "ao",
	Anguilla:                             "ai",
	Antarctica:                           "aq",
	Antigua_and_Barbuda:                  "ag",
	Argentina:                            "ar",
	Armenia:                              "am",
	Aruba:                                "aw",
	Australia:                            "au",
	Austria:                              "at",
	Azerbaijan:                           "az",
	Bahamas:                              "bs",
	Bahrain:                              "bh",
	Bangladesh:                           "bd",
	Barbados:                             "bb",
	Belarus:                              "by",
	Belgium:                              "be",
	Belize:                               "bz",
	Benin:                                "bj",
	Bermuda:                              "bm",
	Bhutan:                               "bt",
	Bolivia:                              "bo",
	Bosnia_and_Herzegovina:               "ba",
	Botswana:                             "bw",
	Bouvet_Island:                        "bv",
	Brazil:                               "br",
	British_Virgin_Islands:               "vg",
	British_Indian_Ocean_Territory:       "io",
	Brunei_Darussalam:                    "dn",
	Bulgaria:                             "bg",
	Burkina_Faso:                         "bf",
	Burundi:                              "bi",
	Cambodia:                             "kh",
	Cameroon:                             "cm",
	Canada:                               "ca",
	Cape_Verde:                           "cv",
	Cayman_Islands:                       "ky",
	Central_African_Republic:             "cf",
	Chad:                                 "td",
	Chile:                                "cl",
	China:                                "cn",
	Hong_Kong:                            "hk",
	Macao:                                "mo",
	Christmas_Island:                     "cx",
	Cocos_Islands:                        "cc",
	Colombia:                             "co",
	Comoros:                              "km",
	Congo_Brazzaville:                    "cg",
	Congo_DRC:                            "cd",
	Cook_Islands:                         "ck",
	Costa_Rica:                           "cr",
	Côte_dIvoire:                         "ci",
	Croatia:                              "hr",
	Cuba:                                 "cu",
	Cyprus:                               "cy",
	Czech_Republic:                       "cz",
	Denmark:                              "dk",
	Djibouti:                             "dj",
	Dominica:                             "dm",
	Dominican_Republic:                   "do",
	Ecuador:                              "ec",
	Egypt:                                "eg",
	El_Salvador:                          "sv",
	Equatorial_Guinea:                    "gq",
	Eritrea:                              "er",
	Estonia:                              "ee",
	Ethiopia:                             "et",
	Falkland_Islands:                     "fk",
	Faroe_Islands:                        "fo",
	Fiji:                                 "fj",
	Finland:                              "fi",
	France:                               "fr",
	French_Guiana:                        "gf",
	French_Polynesia:                     "pf",
	French_Southern_Territories:          "tf",
	Gabon:                                "ga",
	Gambia:                               "gm",
	Georgia:                              "ge",
	Germany:                              "de",
	Ghana:                                "gh",
	Gibraltar:                            "gi",
	Greece:                               "gr",
	Greenland:                            "gl",
	Grenada:                              "gd",
	Guadeloupe:                           "gp",
	Guam:                                 "gu",
	Guatemala:                            "gt",
	Guernsey:                             "gg",
	Guinea:                               "gn",
	Guinea_Bissau:                        "gw",
	Guyana:                               "gy",
	Haiti:                                "ht",
	Heard_Island_and_Mcdonald_Islands:    "hm",
	Holy_See:                             "va",
	Honduras:                             "hn",
	Hungary:                              "hu",
	Iceland:                              "is",
	India:                                "in",
	Indonesia:                            "id",
	Iran:                                 "ir",
	Iraq:                                 "iq",
	Ireland:                              "ie",
	Isle_of_Man:                          "im",
	Israel:                               "il",
	Italy:                                "it",
	Jamaica:                              "jm",
	Japan:                                "jp",
	Jersey:                               "je",
	Jordan:                               "jo",
	Kazakhstan:                           "kz",
	Kenya:                                "ke",
	Kiribati:                             "ki",
	Democratic_Peoples_Republic_of_Korea: "kp",
	South_Korea:                          "kr",
	Kuwait:                               "kw",
	Kyrgyzstan:                           "kg",
	Lao_PDR:                              "la",
	Latvia:                               "lv",
	Lebanon:                              "lb",
	Lesotho:                              "ls",
	Liberia:                              "lr",
	Libya:                                "ly",
	Liechtenstein:                        "li",
	Lithuania:                            "lt",
	Luxembourg:                           "lu",
	Macedonia:                            "mk",
	Madagascar:                           "mg",
	Malawi:                               "mw",
	Malaysia:                             "my",
	Maldives:                             "mv",
	Mali:                                 "ml",
	Malta:                                "mt",
	Marshall_Islands:                     "mh",
	Martinique:                           "mq",
	Mauritania:                           "mr",
	Mauritius:                            "mu",
	Mayotte:                              "yt",
	Mexico:                               "mx",
	Micronesia:                           "fm",
	Moldova:                              "md",
	Monaco:                               "mc",
	Mongolia:                             "mn",
	Montenegro:                           "me",
	Montserrat:                           "ms",
	Morocco:                              "ma",
	Mozambique:                           "mz",
	Myanmar:                              "mm",
	Namibia:                              "na",
	Nauru:                                "nr",
	Nepal:                                "np",
	Netherlands:                          "nl",
	Netherlands_Antilles:                 "an", // obsolete
	New_Caledonia:                        "nc",
	New_Zealand:                          "nz",
	Nicaragua:                            "ni",
	Niger:                                "ne",
	Nigeria:                              "ng",
	Niue:                                 "nu",
	Norfolk_Island:                       "nf",
	Northern_Mariana_Islands:             "mp",
	Norway:                               "no",
	Oman:                                 "om",
	Pakistan:                             "pk",
	Palau:                                "pw",
	Palestinian_Territory:                "ps",
	Panama:                               "pa",
	Papua_New_Guinea:                     "pg",
	Paraguay:                             "py",
	Peru:                                 "pe",
	Philippines:                          "ph",
	Pitcairn:                             "pn",
	Poland:                               "pl",
	Portugal:                             "pt",
	Puerto_Rico:                          "pr",
	Qatar:                                "qa",
	Réunion:                              "re",
	Romania:                              "ro",
	Russian_Federation:                   "ru",
	Rwanda:                               "rw",
	Saint_Barthélemy:                     "bl",
	Saint_Helena:                         "sh",
	Saint_Kitts_and_Nevis:                "kn",
	Saint_Lucia:                          "lc",
	Saint_Martin:                         "mf",
	Saint_Pierre_and_Miquelon:            "pm",
	Saint_Vincent_and_Grenadines:         "vc",
	Samoa:                                "ws",
	San_Marino:                           "sm",
	Sao_Tome_and_Principe:                "st",
	Saudi_Arabia:                         "sa",
	Senegal:                              "sn",
	Serbia:                               "rs",
	Seychelles:                           "sc",
	Sierra_Leone:                         "sl",
	Singapore:                            "sg",
	Slovakia:                             "sk",
	Slovenia:                             "si",
	Solomon_Islands:                      "sb",
	Somalia:                              "so",
	South_Africa:                         "za",
	South_Georgia_and_the_South_Sandwich_Islands: "gs",
	South_Sudan:                          "ss",
	Spain:                                "es",
	Sri_Lanka:                            "lk",
	Sudan:                                "sd",
	Suriname:                             "sr",
	Svalbard_and_Jan_Mayen_Islands:       "sj",
	Swaziland:                            "sz",
	Sweden:                               "se",
	Switzerland:                          "ch",
	Syria:                                "sy",
	Taiwan:                               "tw",
	Tajikistan:                           "tj",
	Tanzania:                             "tz",
	Thailand:                             "th",
	Timor_Leste:                          "tl",
	Togo:                                 "tg",
	Tokelau:                              "tk",
	Tonga:                                "to",
	Trinidad_and_Tobago:                  "tt",
	Tunisia:                              "tn",
	Turkey:                               "tr",
	Turkmenistan:                         "tm",
	Turks_and_Caicos_Islands:             "tc",
	Tuvalu:                               "tv",
	Uganda:                               "ug",
	Ukraine:                              "ua",
	United_Arab_Emirates:                 "ae",
	United_Kingdom:                       "gb",
	United_States_of_America:             "us",
	United_States_Minor_Outlying_Islands: "um",
	Uruguay:                              "uy",
	Uzbekistan:                           "uz",
	Vanuatu:                              "vu",
	Venezuela:                            "ve",
	Viet_Nam:                             "vn",
	Virgin_Islands:                       "vi",
	Wallis_and_Futuna_Islands:            "wf",
	Western_Sahara:                       "eh",
	Yemen:                                "ye",
	Zambia:                               "zm",
	Zimbabwe:                             "zw",
}
