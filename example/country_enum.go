// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.5.2

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"os"
	"strconv"
	"strings"
)

const (
	countryEnumStrings = "AfghanistanAland_IslandsAlbaniaAlgeriaAmerican_SamoaAndorraAngolaAnguillaAntarcticaAntigua_and_BarbudaArgentinaArmeniaArubaAustraliaAustriaAzerbaijanBahamasBahrainBangladeshBarbadosBelarusBelgiumBelizeBeninBermudaBhutanBoliviaBosnia_and_HerzegovinaBotswanaBouvet_IslandBrazilBritish_Virgin_IslandsBritish_Indian_Ocean_TerritoryBrunei_DarussalamBulgariaBurkina_FasoBurundiCambodiaCameroonCanadaCape_VerdeCayman_IslandsCentral_African_RepublicChadChileChinaHong_KongMacaoChristmas_IslandCocos_IslandsColombiaComorosCongo_BrazzavilleCongo_DRCCook_IslandsCosta_RicaCôte_dIvoireCroatiaCubaCyprusCzech_RepublicDenmarkDjiboutiDominicaDominican_RepublicEcuadorEgyptEl_SalvadorEquatorial_GuineaEritreaEstoniaEthiopiaFalkland_IslandsFaroe_IslandsFijiFinlandFranceFrench_GuianaFrench_PolynesiaFrench_Southern_TerritoriesGabonGambiaGeorgiaGermanyGhanaGibraltarGreeceGreenlandGrenadaGuadeloupeGuamGuatemalaGuernseyGuineaGuinea_BissauGuyanaHaitiHeard_Island_and_Mcdonald_IslandsHoly_SeeHondurasHungaryIcelandIndiaIndonesiaIranIraqIrelandIsle_of_ManIsraelItalyJamaicaJapanJerseyJordanKazakhstanKenyaKiribatiDemocratic_Peoples_Republic_of_KoreaSouth_KoreaKuwaitKyrgyzstanLao_PDRLatviaLebanonLesothoLiberiaLibyaLiechtensteinLithuaniaLuxembourgMacedoniaMadagascarMalawiMalaysiaMaldivesMaliMaltaMarshall_IslandsMartiniqueMauritaniaMauritiusMayotteMexicoMicronesiaMoldovaMonacoMongoliaMontenegroMontserratMoroccoMozambiqueMyanmarNamibiaNauruNepalNetherlandsNetherlands_AntillesNew_CaledoniaNew_ZealandNicaraguaNigerNigeriaNiueNorfolk_IslandNorthern_Mariana_IslandsNorwayOmanPakistanPalauPalestinian_TerritoryPanamaPapua_New_GuineaParaguayPeruPhilippinesPitcairnPolandPortugalPuerto_RicoQatarRéunionRomaniaRussian_FederationRwandaSaint_BarthélemySaint_HelenaSaint_Kitts_and_NevisSaint_LuciaSaint_MartinSaint_Pierre_and_MiquelonSaint_Vincent_and_GrenadinesSamoaSan_MarinoSao_Tome_and_PrincipeSaudi_ArabiaSenegalSerbiaSeychellesSierra_LeoneSingaporeSlovakiaSloveniaSolomon_IslandsSomaliaSouth_AfricaSouth_Georgia_and_the_South_Sandwich_IslandsSouth_SudanSpainSri_LankaSudanSurinameSvalbard_and_Jan_Mayen_IslandsSwazilandSwedenSwitzerlandSyriaTaiwanTajikistanTanzaniaThailandTimor_LesteTogoTokelauTongaTrinidad_and_TobagoTunisiaTurkeyTurkmenistanTurks_and_Caicos_IslandsTuvaluUgandaUkraineUnited_Arab_EmiratesUnited_KingdomUnited_States_of_AmericaUnited_States_Minor_Outlying_IslandsUruguayUzbekistanVanuatuVenezuelaViet_NamVirgin_IslandsWallis_and_Futuna_IslandsWestern_SaharaYemenZambiaZimbabwe"
	countryEnumInputs  = "afghanistanaland_islandsalbaniaalgeriaamerican_samoaandorraangolaanguillaantarcticaantigua_and_barbudaargentinaarmeniaarubaaustraliaaustriaazerbaijanbahamasbahrainbangladeshbarbadosbelarusbelgiumbelizebeninbermudabhutanboliviabosnia_and_herzegovinabotswanabouvet_islandbrazilbritish_virgin_islandsbritish_indian_ocean_territorybrunei_darussalambulgariaburkina_fasoburundicambodiacamerooncanadacape_verdecayman_islandscentral_african_republicchadchilechinahong_kongmacaochristmas_islandcocos_islandscolombiacomoroscongo_brazzavillecongo_drccook_islandscosta_ricacôte_divoirecroatiacubacyprusczech_republicdenmarkdjiboutidominicadominican_republicecuadoregyptel_salvadorequatorial_guineaeritreaestoniaethiopiafalkland_islandsfaroe_islandsfijifinlandfrancefrench_guianafrench_polynesiafrench_southern_territoriesgabongambiageorgiagermanyghanagibraltargreecegreenlandgrenadaguadeloupeguamguatemalaguernseyguineaguinea_bissauguyanahaitiheard_island_and_mcdonald_islandsholy_seehondurashungaryicelandindiaindonesiairaniraqirelandisle_of_manisraelitalyjamaicajapanjerseyjordankazakhstankenyakiribatidemocratic_peoples_republic_of_koreasouth_koreakuwaitkyrgyzstanlao_pdrlatvialebanonlesotholiberialibyaliechtensteinlithuanialuxembourgmacedoniamadagascarmalawimalaysiamaldivesmalimaltamarshall_islandsmartiniquemauritaniamauritiusmayottemexicomicronesiamoldovamonacomongoliamontenegromontserratmoroccomozambiquemyanmarnamibianaurunepalnetherlandsnetherlands_antillesnew_caledonianew_zealandnicaraguanigernigerianiuenorfolk_islandnorthern_mariana_islandsnorwayomanpakistanpalaupalestinian_territorypanamapapua_new_guineaparaguayperuphilippinespitcairnpolandportugalpuerto_ricoqatarréunionromaniarussian_federationrwandasaint_barthélemysaint_helenasaint_kitts_and_nevissaint_luciasaint_martinsaint_pierre_and_miquelonsaint_vincent_and_grenadinessamoasan_marinosao_tome_and_principesaudi_arabiasenegalserbiaseychellessierra_leonesingaporeslovakiasloveniasolomon_islandssomaliasouth_africasouth_georgia_and_the_south_sandwich_islandssouth_sudanspainsri_lankasudansurinamesvalbard_and_jan_mayen_islandsswazilandswedenswitzerlandsyriataiwantajikistantanzaniathailandtimor_lestetogotokelautongatrinidad_and_tobagotunisiaturkeyturkmenistanturks_and_caicos_islandstuvaluugandaukraineunited_arab_emiratesunited_kingdomunited_states_of_americaunited_states_minor_outlying_islandsuruguayuzbekistanvanuatuvenezuelaviet_namvirgin_islandswallis_and_futuna_islandswestern_saharayemenzambiazimbabwe"
)

var countryEnumIndex = [...]uint16{0, 11, 24, 31, 38, 52, 59, 65, 73, 83, 102, 111, 118, 123, 132, 139, 149, 156, 163, 173, 181, 188, 195, 201, 206, 213, 219, 226, 248, 256, 269, 275, 297, 327, 344, 352, 364, 371, 379, 387, 393, 403, 417, 441, 445, 450, 455, 464, 469, 485, 498, 506, 513, 530, 539, 551, 561, 574, 581, 585, 591, 605, 612, 620, 628, 646, 653, 658, 669, 686, 693, 700, 708, 724, 737, 741, 748, 754, 767, 783, 810, 815, 821, 828, 835, 840, 849, 855, 864, 871, 881, 885, 894, 902, 908, 921, 927, 932, 965, 973, 981, 988, 995, 1000, 1009, 1013, 1017, 1024, 1035, 1041, 1046, 1053, 1058, 1064, 1070, 1080, 1085, 1093, 1129, 1140, 1146, 1156, 1163, 1169, 1176, 1183, 1190, 1195, 1208, 1217, 1227, 1236, 1246, 1252, 1260, 1268, 1272, 1277, 1293, 1303, 1313, 1322, 1329, 1335, 1345, 1352, 1358, 1366, 1376, 1386, 1393, 1403, 1410, 1417, 1422, 1427, 1438, 1458, 1471, 1482, 1491, 1496, 1503, 1507, 1521, 1545, 1551, 1555, 1563, 1568, 1589, 1595, 1611, 1619, 1623, 1634, 1642, 1648, 1656, 1667, 1672, 1680, 1687, 1705, 1711, 1728, 1740, 1761, 1772, 1784, 1809, 1837, 1842, 1852, 1873, 1885, 1892, 1898, 1908, 1920, 1929, 1937, 1945, 1960, 1967, 1979, 2023, 2034, 2039, 2048, 2053, 2061, 2091, 2100, 2106, 2117, 2122, 2128, 2138, 2146, 2154, 2165, 2169, 2176, 2181, 2200, 2207, 2213, 2225, 2249, 2255, 2261, 2268, 2288, 2302, 2326, 2362, 2369, 2379, 2386, 2395, 2403, 2417, 2442, 2456, 2461, 2467, 2475}

// AllCountries lists all 247 values in order.
var AllCountries = []Country{
	Afghanistan, Aland_Islands, Albania, Algeria,
	American_Samoa, Andorra, Angola, Anguilla, Antarctica,
	Antigua_and_Barbuda, Argentina, Armenia, Aruba, Australia,
	Austria, Azerbaijan, Bahamas, Bahrain, Bangladesh,
	Barbados, Belarus, Belgium, Belize, Benin,
	Bermuda, Bhutan, Bolivia, Bosnia_and_Herzegovina, Botswana,
	Bouvet_Island, Brazil, British_Virgin_Islands, British_Indian_Ocean_Territory, Brunei_Darussalam,
	Bulgaria, Burkina_Faso, Burundi, Cambodia, Cameroon,
	Canada, Cape_Verde, Cayman_Islands, Central_African_Republic, Chad,
	Chile, China, Hong_Kong, Macao, Christmas_Island,
	Cocos_Islands, Colombia, Comoros, Congo_Brazzaville, Congo_DRC,
	Cook_Islands, Costa_Rica, Côte_dIvoire, Croatia, Cuba,
	Cyprus, Czech_Republic, Denmark, Djibouti, Dominica,
	Dominican_Republic, Ecuador, Egypt, El_Salvador, Equatorial_Guinea,
	Eritrea, Estonia, Ethiopia, Falkland_Islands, Faroe_Islands,
	Fiji, Finland, France, French_Guiana, French_Polynesia,
	French_Southern_Territories, Gabon, Gambia, Georgia, Germany,
	Ghana, Gibraltar, Greece, Greenland, Grenada,
	Guadeloupe, Guam, Guatemala, Guernsey, Guinea,
	Guinea_Bissau, Guyana, Haiti, Heard_Island_and_Mcdonald_Islands, Holy_See,
	Honduras, Hungary, Iceland, India, Indonesia,
	Iran, Iraq, Ireland, Isle_of_Man, Israel,
	Italy, Jamaica, Japan, Jersey, Jordan,
	Kazakhstan, Kenya, Kiribati, Democratic_Peoples_Republic_of_Korea, South_Korea,
	Kuwait, Kyrgyzstan, Lao_PDR, Latvia, Lebanon,
	Lesotho, Liberia, Libya, Liechtenstein, Lithuania,
	Luxembourg, Macedonia, Madagascar, Malawi, Malaysia,
	Maldives, Mali, Malta, Marshall_Islands, Martinique,
	Mauritania, Mauritius, Mayotte, Mexico, Micronesia,
	Moldova, Monaco, Mongolia, Montenegro, Montserrat,
	Morocco, Mozambique, Myanmar, Namibia, Nauru,
	Nepal, Netherlands, Netherlands_Antilles, New_Caledonia, New_Zealand,
	Nicaragua, Niger, Nigeria, Niue, Norfolk_Island,
	Northern_Mariana_Islands, Norway, Oman, Pakistan, Palau,
	Palestinian_Territory, Panama, Papua_New_Guinea, Paraguay, Peru,
	Philippines, Pitcairn, Poland, Portugal, Puerto_Rico,
	Qatar, Réunion, Romania, Russian_Federation, Rwanda,
	Saint_Barthélemy, Saint_Helena, Saint_Kitts_and_Nevis, Saint_Lucia, Saint_Martin,
	Saint_Pierre_and_Miquelon, Saint_Vincent_and_Grenadines, Samoa, San_Marino, Sao_Tome_and_Principe,
	Saudi_Arabia, Senegal, Serbia, Seychelles, Sierra_Leone,
	Singapore, Slovakia, Slovenia, Solomon_Islands, Somalia,
	South_Africa, South_Georgia_and_the_South_Sandwich_Islands, South_Sudan, Spain, Sri_Lanka,
	Sudan, Suriname, Svalbard_and_Jan_Mayen_Islands, Swaziland, Sweden,
	Switzerland, Syria, Taiwan, Tajikistan, Tanzania,
	Thailand, Timor_Leste, Togo, Tokelau, Tonga,
	Trinidad_and_Tobago, Tunisia, Turkey, Turkmenistan, Turks_and_Caicos_Islands,
	Tuvalu, Uganda, Ukraine, United_Arab_Emirates, United_Kingdom,
	United_States_of_America, United_States_Minor_Outlying_Islands, Uruguay, Uzbekistan, Vanuatu,
	Venezuela, Viet_Nam, Virgin_Islands, Wallis_and_Futuna_Islands, Western_Sahara,
	Yemen, Zambia, Zimbabwe,
}

// AllCountryEnums lists all 247 values in order.
var AllCountryEnums = enum.IntEnums{
	Afghanistan, Aland_Islands, Albania, Algeria,
	American_Samoa, Andorra, Angola, Anguilla, Antarctica,
	Antigua_and_Barbuda, Argentina, Armenia, Aruba, Australia,
	Austria, Azerbaijan, Bahamas, Bahrain, Bangladesh,
	Barbados, Belarus, Belgium, Belize, Benin,
	Bermuda, Bhutan, Bolivia, Bosnia_and_Herzegovina, Botswana,
	Bouvet_Island, Brazil, British_Virgin_Islands, British_Indian_Ocean_Territory, Brunei_Darussalam,
	Bulgaria, Burkina_Faso, Burundi, Cambodia, Cameroon,
	Canada, Cape_Verde, Cayman_Islands, Central_African_Republic, Chad,
	Chile, China, Hong_Kong, Macao, Christmas_Island,
	Cocos_Islands, Colombia, Comoros, Congo_Brazzaville, Congo_DRC,
	Cook_Islands, Costa_Rica, Côte_dIvoire, Croatia, Cuba,
	Cyprus, Czech_Republic, Denmark, Djibouti, Dominica,
	Dominican_Republic, Ecuador, Egypt, El_Salvador, Equatorial_Guinea,
	Eritrea, Estonia, Ethiopia, Falkland_Islands, Faroe_Islands,
	Fiji, Finland, France, French_Guiana, French_Polynesia,
	French_Southern_Territories, Gabon, Gambia, Georgia, Germany,
	Ghana, Gibraltar, Greece, Greenland, Grenada,
	Guadeloupe, Guam, Guatemala, Guernsey, Guinea,
	Guinea_Bissau, Guyana, Haiti, Heard_Island_and_Mcdonald_Islands, Holy_See,
	Honduras, Hungary, Iceland, India, Indonesia,
	Iran, Iraq, Ireland, Isle_of_Man, Israel,
	Italy, Jamaica, Japan, Jersey, Jordan,
	Kazakhstan, Kenya, Kiribati, Democratic_Peoples_Republic_of_Korea, South_Korea,
	Kuwait, Kyrgyzstan, Lao_PDR, Latvia, Lebanon,
	Lesotho, Liberia, Libya, Liechtenstein, Lithuania,
	Luxembourg, Macedonia, Madagascar, Malawi, Malaysia,
	Maldives, Mali, Malta, Marshall_Islands, Martinique,
	Mauritania, Mauritius, Mayotte, Mexico, Micronesia,
	Moldova, Monaco, Mongolia, Montenegro, Montserrat,
	Morocco, Mozambique, Myanmar, Namibia, Nauru,
	Nepal, Netherlands, Netherlands_Antilles, New_Caledonia, New_Zealand,
	Nicaragua, Niger, Nigeria, Niue, Norfolk_Island,
	Northern_Mariana_Islands, Norway, Oman, Pakistan, Palau,
	Palestinian_Territory, Panama, Papua_New_Guinea, Paraguay, Peru,
	Philippines, Pitcairn, Poland, Portugal, Puerto_Rico,
	Qatar, Réunion, Romania, Russian_Federation, Rwanda,
	Saint_Barthélemy, Saint_Helena, Saint_Kitts_and_Nevis, Saint_Lucia, Saint_Martin,
	Saint_Pierre_and_Miquelon, Saint_Vincent_and_Grenadines, Samoa, San_Marino, Sao_Tome_and_Principe,
	Saudi_Arabia, Senegal, Serbia, Seychelles, Sierra_Leone,
	Singapore, Slovakia, Slovenia, Solomon_Islands, Somalia,
	South_Africa, South_Georgia_and_the_South_Sandwich_Islands, South_Sudan, Spain, Sri_Lanka,
	Sudan, Suriname, Svalbard_and_Jan_Mayen_Islands, Swaziland, Sweden,
	Switzerland, Syria, Taiwan, Tajikistan, Tanzania,
	Thailand, Timor_Leste, Togo, Tokelau, Tonga,
	Trinidad_and_Tobago, Tunisia, Turkey, Turkmenistan, Turks_and_Caicos_Islands,
	Tuvalu, Uganda, Ukraine, United_Arab_Emirates, United_Kingdom,
	United_States_of_America, United_States_Minor_Outlying_Islands, Uruguay, Uzbekistan, Vanuatu,
	Venezuela, Viet_Nam, Virgin_Islands, Wallis_and_Futuna_Islands, Western_Sahara,
	Yemen, Zambia, Zimbabwe,
}

// String returns the literal string representation of a Country, which is
// the same as the const identifier.
func (i Country) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllCountries) {
		return fmt.Sprintf("Country(%d)", i)
	}
	return countryEnumStrings[countryEnumIndex[o]:countryEnumIndex[o+1]]
}

var iso3166TagsInverse = map[string]Country{}

func init() {
	for _, id := range AllCountries {
		v, exists := iso3166Tags[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Country: %s is missing from iso3166Tags\n", id)
		} else {
			k := strings.ToLower(v)
			if _, exists := iso3166TagsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Country: %q is duplicated in iso3166Tags\n", k)
			}
			iso3166TagsInverse[k] = id
		}
	}

	if len(iso3166Tags) != 247 {
		panic(fmt.Sprintf("Country: iso3166Tags has %d items but should have 247", len(iso3166Tags)))
	}

	if len(iso3166Tags) != len(iso3166TagsInverse) {
		panic(fmt.Sprintf("Country: iso3166Tags has %d items but there are only %d distinct items",
			len(iso3166Tags), len(iso3166TagsInverse)))
	}
}

// Tag returns the string representation of a Country.
func (i Country) Tag() string {
	s, ok := iso3166Tags[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}

// Ordinal returns the ordinal number of a Country.
func (i Country) Ordinal() int {
	switch i {
	case Afghanistan:
		return 0
	case Aland_Islands:
		return 1
	case Albania:
		return 2
	case Algeria:
		return 3
	case American_Samoa:
		return 4
	case Andorra:
		return 5
	case Angola:
		return 6
	case Anguilla:
		return 7
	case Antarctica:
		return 8
	case Antigua_and_Barbuda:
		return 9
	case Argentina:
		return 10
	case Armenia:
		return 11
	case Aruba:
		return 12
	case Australia:
		return 13
	case Austria:
		return 14
	case Azerbaijan:
		return 15
	case Bahamas:
		return 16
	case Bahrain:
		return 17
	case Bangladesh:
		return 18
	case Barbados:
		return 19
	case Belarus:
		return 20
	case Belgium:
		return 21
	case Belize:
		return 22
	case Benin:
		return 23
	case Bermuda:
		return 24
	case Bhutan:
		return 25
	case Bolivia:
		return 26
	case Bosnia_and_Herzegovina:
		return 27
	case Botswana:
		return 28
	case Bouvet_Island:
		return 29
	case Brazil:
		return 30
	case British_Virgin_Islands:
		return 31
	case British_Indian_Ocean_Territory:
		return 32
	case Brunei_Darussalam:
		return 33
	case Bulgaria:
		return 34
	case Burkina_Faso:
		return 35
	case Burundi:
		return 36
	case Cambodia:
		return 37
	case Cameroon:
		return 38
	case Canada:
		return 39
	case Cape_Verde:
		return 40
	case Cayman_Islands:
		return 41
	case Central_African_Republic:
		return 42
	case Chad:
		return 43
	case Chile:
		return 44
	case China:
		return 45
	case Hong_Kong:
		return 46
	case Macao:
		return 47
	case Christmas_Island:
		return 48
	case Cocos_Islands:
		return 49
	case Colombia:
		return 50
	case Comoros:
		return 51
	case Congo_Brazzaville:
		return 52
	case Congo_DRC:
		return 53
	case Cook_Islands:
		return 54
	case Costa_Rica:
		return 55
	case Côte_dIvoire:
		return 56
	case Croatia:
		return 57
	case Cuba:
		return 58
	case Cyprus:
		return 59
	case Czech_Republic:
		return 60
	case Denmark:
		return 61
	case Djibouti:
		return 62
	case Dominica:
		return 63
	case Dominican_Republic:
		return 64
	case Ecuador:
		return 65
	case Egypt:
		return 66
	case El_Salvador:
		return 67
	case Equatorial_Guinea:
		return 68
	case Eritrea:
		return 69
	case Estonia:
		return 70
	case Ethiopia:
		return 71
	case Falkland_Islands:
		return 72
	case Faroe_Islands:
		return 73
	case Fiji:
		return 74
	case Finland:
		return 75
	case France:
		return 76
	case French_Guiana:
		return 77
	case French_Polynesia:
		return 78
	case French_Southern_Territories:
		return 79
	case Gabon:
		return 80
	case Gambia:
		return 81
	case Georgia:
		return 82
	case Germany:
		return 83
	case Ghana:
		return 84
	case Gibraltar:
		return 85
	case Greece:
		return 86
	case Greenland:
		return 87
	case Grenada:
		return 88
	case Guadeloupe:
		return 89
	case Guam:
		return 90
	case Guatemala:
		return 91
	case Guernsey:
		return 92
	case Guinea:
		return 93
	case Guinea_Bissau:
		return 94
	case Guyana:
		return 95
	case Haiti:
		return 96
	case Heard_Island_and_Mcdonald_Islands:
		return 97
	case Holy_See:
		return 98
	case Honduras:
		return 99
	case Hungary:
		return 100
	case Iceland:
		return 101
	case India:
		return 102
	case Indonesia:
		return 103
	case Iran:
		return 104
	case Iraq:
		return 105
	case Ireland:
		return 106
	case Isle_of_Man:
		return 107
	case Israel:
		return 108
	case Italy:
		return 109
	case Jamaica:
		return 110
	case Japan:
		return 111
	case Jersey:
		return 112
	case Jordan:
		return 113
	case Kazakhstan:
		return 114
	case Kenya:
		return 115
	case Kiribati:
		return 116
	case Democratic_Peoples_Republic_of_Korea:
		return 117
	case South_Korea:
		return 118
	case Kuwait:
		return 119
	case Kyrgyzstan:
		return 120
	case Lao_PDR:
		return 121
	case Latvia:
		return 122
	case Lebanon:
		return 123
	case Lesotho:
		return 124
	case Liberia:
		return 125
	case Libya:
		return 126
	case Liechtenstein:
		return 127
	case Lithuania:
		return 128
	case Luxembourg:
		return 129
	case Macedonia:
		return 130
	case Madagascar:
		return 131
	case Malawi:
		return 132
	case Malaysia:
		return 133
	case Maldives:
		return 134
	case Mali:
		return 135
	case Malta:
		return 136
	case Marshall_Islands:
		return 137
	case Martinique:
		return 138
	case Mauritania:
		return 139
	case Mauritius:
		return 140
	case Mayotte:
		return 141
	case Mexico:
		return 142
	case Micronesia:
		return 143
	case Moldova:
		return 144
	case Monaco:
		return 145
	case Mongolia:
		return 146
	case Montenegro:
		return 147
	case Montserrat:
		return 148
	case Morocco:
		return 149
	case Mozambique:
		return 150
	case Myanmar:
		return 151
	case Namibia:
		return 152
	case Nauru:
		return 153
	case Nepal:
		return 154
	case Netherlands:
		return 155
	case Netherlands_Antilles:
		return 156
	case New_Caledonia:
		return 157
	case New_Zealand:
		return 158
	case Nicaragua:
		return 159
	case Niger:
		return 160
	case Nigeria:
		return 161
	case Niue:
		return 162
	case Norfolk_Island:
		return 163
	case Northern_Mariana_Islands:
		return 164
	case Norway:
		return 165
	case Oman:
		return 166
	case Pakistan:
		return 167
	case Palau:
		return 168
	case Palestinian_Territory:
		return 169
	case Panama:
		return 170
	case Papua_New_Guinea:
		return 171
	case Paraguay:
		return 172
	case Peru:
		return 173
	case Philippines:
		return 174
	case Pitcairn:
		return 175
	case Poland:
		return 176
	case Portugal:
		return 177
	case Puerto_Rico:
		return 178
	case Qatar:
		return 179
	case Réunion:
		return 180
	case Romania:
		return 181
	case Russian_Federation:
		return 182
	case Rwanda:
		return 183
	case Saint_Barthélemy:
		return 184
	case Saint_Helena:
		return 185
	case Saint_Kitts_and_Nevis:
		return 186
	case Saint_Lucia:
		return 187
	case Saint_Martin:
		return 188
	case Saint_Pierre_and_Miquelon:
		return 189
	case Saint_Vincent_and_Grenadines:
		return 190
	case Samoa:
		return 191
	case San_Marino:
		return 192
	case Sao_Tome_and_Principe:
		return 193
	case Saudi_Arabia:
		return 194
	case Senegal:
		return 195
	case Serbia:
		return 196
	case Seychelles:
		return 197
	case Sierra_Leone:
		return 198
	case Singapore:
		return 199
	case Slovakia:
		return 200
	case Slovenia:
		return 201
	case Solomon_Islands:
		return 202
	case Somalia:
		return 203
	case South_Africa:
		return 204
	case South_Georgia_and_the_South_Sandwich_Islands:
		return 205
	case South_Sudan:
		return 206
	case Spain:
		return 207
	case Sri_Lanka:
		return 208
	case Sudan:
		return 209
	case Suriname:
		return 210
	case Svalbard_and_Jan_Mayen_Islands:
		return 211
	case Swaziland:
		return 212
	case Sweden:
		return 213
	case Switzerland:
		return 214
	case Syria:
		return 215
	case Taiwan:
		return 216
	case Tajikistan:
		return 217
	case Tanzania:
		return 218
	case Thailand:
		return 219
	case Timor_Leste:
		return 220
	case Togo:
		return 221
	case Tokelau:
		return 222
	case Tonga:
		return 223
	case Trinidad_and_Tobago:
		return 224
	case Tunisia:
		return 225
	case Turkey:
		return 226
	case Turkmenistan:
		return 227
	case Turks_and_Caicos_Islands:
		return 228
	case Tuvalu:
		return 229
	case Uganda:
		return 230
	case Ukraine:
		return 231
	case United_Arab_Emirates:
		return 232
	case United_Kingdom:
		return 233
	case United_States_of_America:
		return 234
	case United_States_Minor_Outlying_Islands:
		return 235
	case Uruguay:
		return 236
	case Uzbekistan:
		return 237
	case Vanuatu:
		return 238
	case Venezuela:
		return 239
	case Viet_Nam:
		return 240
	case Virgin_Islands:
		return 241
	case Wallis_and_Futuna_Islands:
		return 242
	case Western_Sahara:
		return 243
	case Yemen:
		return 244
	case Zambia:
		return 245
	case Zimbabwe:
		return 246
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Country) Int() int {
	return int(i)
}

// CountryOf returns a Country based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Country is returned.
func CountryOf(i int) Country {
	if 0 <= i && i < len(AllCountries) {
		return AllCountries[i]
	}
	// an invalid result
	return Afghanistan + Aland_Islands + Albania + Algeria + American_Samoa + Andorra + Angola + Anguilla + Antarctica + Antigua_and_Barbuda + Argentina + Armenia + Aruba + Australia + Austria + Azerbaijan + Bahamas + Bahrain + Bangladesh + Barbados + Belarus + Belgium + Belize + Benin + Bermuda + Bhutan + Bolivia + Bosnia_and_Herzegovina + Botswana + Bouvet_Island + Brazil + British_Virgin_Islands + British_Indian_Ocean_Territory + Brunei_Darussalam + Bulgaria + Burkina_Faso + Burundi + Cambodia + Cameroon + Canada + Cape_Verde + Cayman_Islands + Central_African_Republic + Chad + Chile + China + Hong_Kong + Macao + Christmas_Island + Cocos_Islands + Colombia + Comoros + Congo_Brazzaville + Congo_DRC + Cook_Islands + Costa_Rica + Côte_dIvoire + Croatia + Cuba + Cyprus + Czech_Republic + Denmark + Djibouti + Dominica + Dominican_Republic + Ecuador + Egypt + El_Salvador + Equatorial_Guinea + Eritrea + Estonia + Ethiopia + Falkland_Islands + Faroe_Islands + Fiji + Finland + France + French_Guiana + French_Polynesia + French_Southern_Territories + Gabon + Gambia + Georgia + Germany + Ghana + Gibraltar + Greece + Greenland + Grenada + Guadeloupe + Guam + Guatemala + Guernsey + Guinea + Guinea_Bissau + Guyana + Haiti + Heard_Island_and_Mcdonald_Islands + Holy_See + Honduras + Hungary + Iceland + India + Indonesia + Iran + Iraq + Ireland + Isle_of_Man + Israel + Italy + Jamaica + Japan + Jersey + Jordan + Kazakhstan + Kenya + Kiribati + Democratic_Peoples_Republic_of_Korea + South_Korea + Kuwait + Kyrgyzstan + Lao_PDR + Latvia + Lebanon + Lesotho + Liberia + Libya + Liechtenstein + Lithuania + Luxembourg + Macedonia + Madagascar + Malawi + Malaysia + Maldives + Mali + Malta + Marshall_Islands + Martinique + Mauritania + Mauritius + Mayotte + Mexico + Micronesia + Moldova + Monaco + Mongolia + Montenegro + Montserrat + Morocco + Mozambique + Myanmar + Namibia + Nauru + Nepal + Netherlands + Netherlands_Antilles + New_Caledonia + New_Zealand + Nicaragua + Niger + Nigeria + Niue + Norfolk_Island + Northern_Mariana_Islands + Norway + Oman + Pakistan + Palau + Palestinian_Territory + Panama + Papua_New_Guinea + Paraguay + Peru + Philippines + Pitcairn + Poland + Portugal + Puerto_Rico + Qatar + Réunion + Romania + Russian_Federation + Rwanda + Saint_Barthélemy + Saint_Helena + Saint_Kitts_and_Nevis + Saint_Lucia + Saint_Martin + Saint_Pierre_and_Miquelon + Saint_Vincent_and_Grenadines + Samoa + San_Marino + Sao_Tome_and_Principe + Saudi_Arabia + Senegal + Serbia + Seychelles + Sierra_Leone + Singapore + Slovakia + Slovenia + Solomon_Islands + Somalia + South_Africa + South_Georgia_and_the_South_Sandwich_Islands + South_Sudan + Spain + Sri_Lanka + Sudan + Suriname + Svalbard_and_Jan_Mayen_Islands + Swaziland + Sweden + Switzerland + Syria + Taiwan + Tajikistan + Tanzania + Thailand + Timor_Leste + Togo + Tokelau + Tonga + Trinidad_and_Tobago + Tunisia + Turkey + Turkmenistan + Turks_and_Caicos_Islands + Tuvalu + Uganda + Ukraine + United_Arab_Emirates + United_Kingdom + United_States_of_America + United_States_Minor_Outlying_Islands + Uruguay + Uzbekistan + Vanuatu + Venezuela + Viet_Nam + Virgin_Islands + Wallis_and_Futuna_Islands + Western_Sahara + Yemen + Zambia + Zimbabwe + 1
}

// IsValid determines whether a Country is one of the defined constants.
func (i Country) IsValid() bool {
	switch i {
	case Afghanistan, Aland_Islands, Albania, Algeria,
		American_Samoa, Andorra, Angola, Anguilla, Antarctica,
		Antigua_and_Barbuda, Argentina, Armenia, Aruba, Australia,
		Austria, Azerbaijan, Bahamas, Bahrain, Bangladesh,
		Barbados, Belarus, Belgium, Belize, Benin,
		Bermuda, Bhutan, Bolivia, Bosnia_and_Herzegovina, Botswana,
		Bouvet_Island, Brazil, British_Virgin_Islands, British_Indian_Ocean_Territory, Brunei_Darussalam,
		Bulgaria, Burkina_Faso, Burundi, Cambodia, Cameroon,
		Canada, Cape_Verde, Cayman_Islands, Central_African_Republic, Chad,
		Chile, China, Hong_Kong, Macao, Christmas_Island,
		Cocos_Islands, Colombia, Comoros, Congo_Brazzaville, Congo_DRC,
		Cook_Islands, Costa_Rica, Côte_dIvoire, Croatia, Cuba,
		Cyprus, Czech_Republic, Denmark, Djibouti, Dominica,
		Dominican_Republic, Ecuador, Egypt, El_Salvador, Equatorial_Guinea,
		Eritrea, Estonia, Ethiopia, Falkland_Islands, Faroe_Islands,
		Fiji, Finland, France, French_Guiana, French_Polynesia,
		French_Southern_Territories, Gabon, Gambia, Georgia, Germany,
		Ghana, Gibraltar, Greece, Greenland, Grenada,
		Guadeloupe, Guam, Guatemala, Guernsey, Guinea,
		Guinea_Bissau, Guyana, Haiti, Heard_Island_and_Mcdonald_Islands, Holy_See,
		Honduras, Hungary, Iceland, India, Indonesia,
		Iran, Iraq, Ireland, Isle_of_Man, Israel,
		Italy, Jamaica, Japan, Jersey, Jordan,
		Kazakhstan, Kenya, Kiribati, Democratic_Peoples_Republic_of_Korea, South_Korea,
		Kuwait, Kyrgyzstan, Lao_PDR, Latvia, Lebanon,
		Lesotho, Liberia, Libya, Liechtenstein, Lithuania,
		Luxembourg, Macedonia, Madagascar, Malawi, Malaysia,
		Maldives, Mali, Malta, Marshall_Islands, Martinique,
		Mauritania, Mauritius, Mayotte, Mexico, Micronesia,
		Moldova, Monaco, Mongolia, Montenegro, Montserrat,
		Morocco, Mozambique, Myanmar, Namibia, Nauru,
		Nepal, Netherlands, Netherlands_Antilles, New_Caledonia, New_Zealand,
		Nicaragua, Niger, Nigeria, Niue, Norfolk_Island,
		Northern_Mariana_Islands, Norway, Oman, Pakistan, Palau,
		Palestinian_Territory, Panama, Papua_New_Guinea, Paraguay, Peru,
		Philippines, Pitcairn, Poland, Portugal, Puerto_Rico,
		Qatar, Réunion, Romania, Russian_Federation, Rwanda,
		Saint_Barthélemy, Saint_Helena, Saint_Kitts_and_Nevis, Saint_Lucia, Saint_Martin,
		Saint_Pierre_and_Miquelon, Saint_Vincent_and_Grenadines, Samoa, San_Marino, Sao_Tome_and_Principe,
		Saudi_Arabia, Senegal, Serbia, Seychelles, Sierra_Leone,
		Singapore, Slovakia, Slovenia, Solomon_Islands, Somalia,
		South_Africa, South_Georgia_and_the_South_Sandwich_Islands, South_Sudan, Spain, Sri_Lanka,
		Sudan, Suriname, Svalbard_and_Jan_Mayen_Islands, Swaziland, Sweden,
		Switzerland, Syria, Taiwan, Tajikistan, Tanzania,
		Thailand, Timor_Leste, Togo, Tokelau, Tonga,
		Trinidad_and_Tobago, Tunisia, Turkey, Turkmenistan, Turks_and_Caicos_Islands,
		Tuvalu, Uganda, Ukraine, United_Arab_Emirates, United_Kingdom,
		United_States_of_America, United_States_Minor_Outlying_Islands, Uruguay, Uzbekistan, Vanuatu,
		Venezuela, Viet_Nam, Virgin_Islands, Wallis_and_Futuna_Islands, Western_Sahara,
		Yemen, Zambia, Zimbabwe:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Country, accepting one of the string values or
// a number. The input representation is determined by countryMarshalTextRep. It is used by AsCountry.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Country)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Country) Parse(s string) error {
	return v.parse(s, countryMarshalTextRep)
}

func (v *Country) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := strings.ToLower(in)

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseIdentifier(s) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised country")
}

// parseNumber attempts to convert a decimal value
func (v *Country) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Country(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Country) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllCountries) {
		*v = AllCountries[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in iso3166TagsInverse
func (v *Country) parseTag(s string) (ok bool) {
	*v, ok = iso3166TagsInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *Country) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(countryEnumIndex); j++ {
		i1 := countryEnumIndex[j]
		p := countryEnumInputs[i0:i1]
		if s == p {
			*v = AllCountries[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsCountry parses a string to find the corresponding Country, accepting either one of the string values or
// a number. The input representation is determined by countryMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsCountry(s string) (Country, error) {
	var i = new(Country)
	err := i.Parse(s)
	return *i, err
}

// countryMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var countryMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to countryMarshalTextRep.
func (i Country) MarshalText() (text []byte, err error) {
	return i.marshalText(countryMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to countryMarshalTextRep.
func (i Country) MarshalJSON() ([]byte, error) {
	return i.marshalText(countryMarshalTextRep, true)
}

func (i Country) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		bs = []byte(strconv.FormatInt(int64(i), 10))
	case enum.Ordinal:
		bs = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		if quoted {
			bs = i.quotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = []byte(i.quotedString(i.String()))
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

func (i Country) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Country) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Country) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// countryStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var countryStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Country) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if countryStoreRep == enum.Ordinal {
			*i = CountryOf(int(v))
		} else {
			*i = Country(v)
		}
	case float64:
		*i = Country(v)
	case []byte:
		err = i.parse(string(v), countryStoreRep)
	case string:
		err = i.parse(v, countryStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful country", value, value)
	}

	return err
}

// Value converts the Country to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Country) Value() (driver.Value, error) {
	switch countryStoreRep {
	case enum.Number:
		return int64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
