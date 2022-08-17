package example

//go:generate enumeration -v -type Country -plural Countries -ic

// Country: This example shows use of the '-plural' option to set the name of plural
// collections. Because of '-ic', the parser ignores case. The '-using' option
// provides cross-mapping between the country names and their ISO-3166 tags.
//
// The `all` tags in comments control values used for text/JSON marshalling & SQL storage.
type Country int

const (
	Afghanistan                                  Country = iota // all:"af"
	Aland_Islands                                               // all:"ax"
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
	Czech_Republic                                              // all:"cz"
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
	Macedonia                                                   // all:"mk" -- Republic of Macedonia
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
	Turkey                                                      // all:"tr"
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
