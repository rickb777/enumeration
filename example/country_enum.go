// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.1

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllCountries lists all 247 values in order.
var AllCountries = []Country{
	Afghanistan, Åland_Islands, Albania, Algeria,
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
	Cyprus, Czechia, Denmark, Djibouti, Dominica,
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
	Luxembourg, North_Macedonia, Madagascar, Malawi, Malaysia,
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
	Trinidad_and_Tobago, Tunisia, Türkiye, Turkmenistan, Turks_and_Caicos_Islands,
	Tuvalu, Uganda, Ukraine, United_Arab_Emirates, United_Kingdom,
	United_States_of_America, United_States_Minor_Outlying_Islands, Uruguay, Uzbekistan, Vanuatu,
	Venezuela, Viet_Nam, Virgin_Islands, Wallis_and_Futuna_Islands, Western_Sahara,
	Yemen, Zambia, Zimbabwe,
}

// AllCountryEnums lists all 247 values in order.
var AllCountryEnums = enum.IntEnums{
	Afghanistan, Åland_Islands, Albania, Algeria,
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
	Cyprus, Czechia, Denmark, Djibouti, Dominica,
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
	Luxembourg, North_Macedonia, Madagascar, Malawi, Malaysia,
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
	Trinidad_and_Tobago, Tunisia, Türkiye, Turkmenistan, Turks_and_Caicos_Islands,
	Tuvalu, Uganda, Ukraine, United_Arab_Emirates, United_Kingdom,
	United_States_of_America, United_States_Minor_Outlying_Islands, Uruguay, Uzbekistan, Vanuatu,
	Venezuela, Viet_Nam, Virgin_Islands, Wallis_and_Futuna_Islands, Western_Sahara,
	Yemen, Zambia, Zimbabwe,
}

const (
	countryEnumStrings = "AfghanistanÅland IslandsAlbaniaAlgeriaAmerican SamoaAndorraAngolaAnguillaAntarcticaAntigua and BarbudaArgentinaArmeniaArubaAustraliaAustriaAzerbaijanBahamasBahrainBangladeshBarbadosBelarusBelgiumBelizeBeninBermudaBhutanBoliviaBosnia and HerzegovinaBotswanaBouvet IslandBrazilBritish Virgin IslandsBritish Indian Ocean TerritoryBrunei DarussalamBulgariaBurkina FasoBurundiCambodiaCameroonCanadaCape VerdeCayman IslandsCentral African RepublicChadChileChinaHong KongMacaoChristmas IslandCocos IslandsColombiaComorosCongo BrazzavilleCongo DRCCook IslandsCosta RicaCôte dIvoireCroatiaCubaCyprusCzechiaDenmarkDjiboutiDominicaDominican RepublicEcuadorEgyptEl SalvadorEquatorial GuineaEritreaEstoniaEthiopiaFalkland IslandsFaroe IslandsFijiFinlandFranceFrench GuianaFrench PolynesiaFrench Southern TerritoriesGabonGambiaGeorgiaGermanyGhanaGibraltarGreeceGreenlandGrenadaGuadeloupeGuamGuatemalaGuernseyGuineaGuinea BissauGuyanaHaitiHeard Island and Mcdonald IslandsHoly SeeHondurasHungaryIcelandIndiaIndonesiaIranIraqIrelandIsle of ManIsraelItalyJamaicaJapanJerseyJordanKazakhstanKenyaKiribatiDemocratic Peoples Republic of KoreaSouth KoreaKuwaitKyrgyzstanLao PDRLatviaLebanonLesothoLiberiaLibyaLiechtensteinLithuaniaLuxembourgNorth MacedoniaMadagascarMalawiMalaysiaMaldivesMaliMaltaMarshall IslandsMartiniqueMauritaniaMauritiusMayotteMexicoMicronesiaMoldovaMonacoMongoliaMontenegroMontserratMoroccoMozambiqueMyanmarNamibiaNauruNepalNetherlandsNetherlands AntillesNew CaledoniaNew ZealandNicaraguaNigerNigeriaNiueNorfolk IslandNorthern Mariana IslandsNorwayOmanPakistanPalauPalestinian TerritoryPanamaPapua New GuineaParaguayPeruPhilippinesPitcairnPolandPortugalPuerto RicoQatarRéunionRomaniaRussian FederationRwandaSaint BarthélemySaint HelenaSaint Kitts and NevisSaint LuciaSaint MartinSaint Pierre and MiquelonSaint Vincent and GrenadinesSamoaSan MarinoSao Tome and PrincipeSaudi ArabiaSenegalSerbiaSeychellesSierra LeoneSingaporeSlovakiaSloveniaSolomon IslandsSomaliaSouth AfricaSouth Georgia and the South Sandwich IslandsSouth SudanSpainSri LankaSudanSurinameSvalbard and Jan Mayen IslandsSwazilandSwedenSwitzerlandSyriaTaiwanTajikistanTanzaniaThailandTimor LesteTogoTokelauTongaTrinidad and TobagoTunisiaTürkiyeTurkmenistanTurks and Caicos IslandsTuvaluUgandaUkraineUnited Arab EmiratesUnited KingdomUnited States of AmericaUnited States Minor Outlying IslandsUruguayUzbekistanVanuatuVenezuelaViet NamVirgin IslandsWallis and Futuna IslandsWestern SaharaYemenZambiaZimbabwe"
	countryEnumInputs  = "afghanistanåland islandsalbaniaalgeriaamerican samoaandorraangolaanguillaantarcticaantigua and barbudaargentinaarmeniaarubaaustraliaaustriaazerbaijanbahamasbahrainbangladeshbarbadosbelarusbelgiumbelizebeninbermudabhutanboliviabosnia and herzegovinabotswanabouvet islandbrazilbritish virgin islandsbritish indian ocean territorybrunei darussalambulgariaburkina fasoburundicambodiacamerooncanadacape verdecayman islandscentral african republicchadchilechinahong kongmacaochristmas islandcocos islandscolombiacomoroscongo brazzavillecongo drccook islandscosta ricacôte divoirecroatiacubacyprusczechiadenmarkdjiboutidominicadominican republicecuadoregyptel salvadorequatorial guineaeritreaestoniaethiopiafalkland islandsfaroe islandsfijifinlandfrancefrench guianafrench polynesiafrench southern territoriesgabongambiageorgiagermanyghanagibraltargreecegreenlandgrenadaguadeloupeguamguatemalaguernseyguineaguinea bissauguyanahaitiheard island and mcdonald islandsholy seehondurashungaryicelandindiaindonesiairaniraqirelandisle of manisraelitalyjamaicajapanjerseyjordankazakhstankenyakiribatidemocratic peoples republic of koreasouth koreakuwaitkyrgyzstanlao pdrlatvialebanonlesotholiberialibyaliechtensteinlithuanialuxembourgnorth macedoniamadagascarmalawimalaysiamaldivesmalimaltamarshall islandsmartiniquemauritaniamauritiusmayottemexicomicronesiamoldovamonacomongoliamontenegromontserratmoroccomozambiquemyanmarnamibianaurunepalnetherlandsnetherlands antillesnew caledonianew zealandnicaraguanigernigerianiuenorfolk islandnorthern mariana islandsnorwayomanpakistanpalaupalestinian territorypanamapapua new guineaparaguayperuphilippinespitcairnpolandportugalpuerto ricoqatarréunionromaniarussian federationrwandasaint barthélemysaint helenasaint kitts and nevissaint luciasaint martinsaint pierre and miquelonsaint vincent and grenadinessamoasan marinosao tome and principesaudi arabiasenegalserbiaseychellessierra leonesingaporeslovakiasloveniasolomon islandssomaliasouth africasouth georgia and the south sandwich islandssouth sudanspainsri lankasudansurinamesvalbard and jan mayen islandsswazilandswedenswitzerlandsyriataiwantajikistantanzaniathailandtimor lestetogotokelautongatrinidad and tobagotunisiatürkiyeturkmenistanturks and caicos islandstuvaluugandaukraineunited arab emiratesunited kingdomunited states of americaunited states minor outlying islandsuruguayuzbekistanvanuatuvenezuelaviet namvirgin islandswallis and futuna islandswestern saharayemenzambiazimbabwe"
	countryTextStrings = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
	countryTextInputs  = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
	countryJSONStrings = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
	countryJSONInputs  = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
	countrySQLStrings  = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
	countrySQLInputs   = "afaxaldzasadaoaiaqagaramawauatazbsbhbdbbbybebzbjbmbtbobabwbvbrvgiodnbgbfbikhcmcacvkycftdclcnhkmocxcccokmcgcdckcrcihrcucyczdkdjdmdoecegsvgqereeetfkfofjfifrgfpftfgagmgedeghgigrglgdgpgugtgggngwgyhthmvahnhuisinidiriqieimilitjmjpjejokzkekikpkrkwkglalvlblslrlyliltlumkmgmwmymvmlmtmhmqmrmuytmxfmmdmcmnmemsmamzmmnanrnpnlanncnzninengnunfmpnoompkpwpspapgpypephpnplptprqarerorurwblshknlcmfpmvcwssmstsasnrsscslsgsksisbsozagssseslksdsrsjszsechsytwtjtzthtltgtktotttntrtmtctvuguaaegbusumuyuzvuvevnviwfehyezmzw"
)

var (
	countryEnumIndex = [...]uint16{0, 11, 25, 32, 39, 53, 60, 66, 74, 84, 103, 112, 119, 124, 133, 140, 150, 157, 164, 174, 182, 189, 196, 202, 207, 214, 220, 227, 249, 257, 270, 276, 298, 328, 345, 353, 365, 372, 380, 388, 394, 404, 418, 442, 446, 451, 456, 465, 470, 486, 499, 507, 514, 531, 540, 552, 562, 575, 582, 586, 592, 599, 606, 614, 622, 640, 647, 652, 663, 680, 687, 694, 702, 718, 731, 735, 742, 748, 761, 777, 804, 809, 815, 822, 829, 834, 843, 849, 858, 865, 875, 879, 888, 896, 902, 915, 921, 926, 959, 967, 975, 982, 989, 994, 1003, 1007, 1011, 1018, 1029, 1035, 1040, 1047, 1052, 1058, 1064, 1074, 1079, 1087, 1123, 1134, 1140, 1150, 1157, 1163, 1170, 1177, 1184, 1189, 1202, 1211, 1221, 1236, 1246, 1252, 1260, 1268, 1272, 1277, 1293, 1303, 1313, 1322, 1329, 1335, 1345, 1352, 1358, 1366, 1376, 1386, 1393, 1403, 1410, 1417, 1422, 1427, 1438, 1458, 1471, 1482, 1491, 1496, 1503, 1507, 1521, 1545, 1551, 1555, 1563, 1568, 1589, 1595, 1611, 1619, 1623, 1634, 1642, 1648, 1656, 1667, 1672, 1680, 1687, 1705, 1711, 1728, 1740, 1761, 1772, 1784, 1809, 1837, 1842, 1852, 1873, 1885, 1892, 1898, 1908, 1920, 1929, 1937, 1945, 1960, 1967, 1979, 2023, 2034, 2039, 2048, 2053, 2061, 2091, 2100, 2106, 2117, 2122, 2128, 2138, 2146, 2154, 2165, 2169, 2176, 2181, 2200, 2207, 2215, 2227, 2251, 2257, 2263, 2270, 2290, 2304, 2328, 2364, 2371, 2381, 2388, 2397, 2405, 2419, 2444, 2458, 2463, 2469, 2477}
	countryTextIndex = [...]uint16{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198, 200, 202, 204, 206, 208, 210, 212, 214, 216, 218, 220, 222, 224, 226, 228, 230, 232, 234, 236, 238, 240, 242, 244, 246, 248, 250, 252, 254, 256, 258, 260, 262, 264, 266, 268, 270, 272, 274, 276, 278, 280, 282, 284, 286, 288, 290, 292, 294, 296, 298, 300, 302, 304, 306, 308, 310, 312, 314, 316, 318, 320, 322, 324, 326, 328, 330, 332, 334, 336, 338, 340, 342, 344, 346, 348, 350, 352, 354, 356, 358, 360, 362, 364, 366, 368, 370, 372, 374, 376, 378, 380, 382, 384, 386, 388, 390, 392, 394, 396, 398, 400, 402, 404, 406, 408, 410, 412, 414, 416, 418, 420, 422, 424, 426, 428, 430, 432, 434, 436, 438, 440, 442, 444, 446, 448, 450, 452, 454, 456, 458, 460, 462, 464, 466, 468, 470, 472, 474, 476, 478, 480, 482, 484, 486, 488, 490, 492, 494}
	countryJSONIndex = [...]uint16{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198, 200, 202, 204, 206, 208, 210, 212, 214, 216, 218, 220, 222, 224, 226, 228, 230, 232, 234, 236, 238, 240, 242, 244, 246, 248, 250, 252, 254, 256, 258, 260, 262, 264, 266, 268, 270, 272, 274, 276, 278, 280, 282, 284, 286, 288, 290, 292, 294, 296, 298, 300, 302, 304, 306, 308, 310, 312, 314, 316, 318, 320, 322, 324, 326, 328, 330, 332, 334, 336, 338, 340, 342, 344, 346, 348, 350, 352, 354, 356, 358, 360, 362, 364, 366, 368, 370, 372, 374, 376, 378, 380, 382, 384, 386, 388, 390, 392, 394, 396, 398, 400, 402, 404, 406, 408, 410, 412, 414, 416, 418, 420, 422, 424, 426, 428, 430, 432, 434, 436, 438, 440, 442, 444, 446, 448, 450, 452, 454, 456, 458, 460, 462, 464, 466, 468, 470, 472, 474, 476, 478, 480, 482, 484, 486, 488, 490, 492, 494}
	countrySQLIndex  = [...]uint16{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198, 200, 202, 204, 206, 208, 210, 212, 214, 216, 218, 220, 222, 224, 226, 228, 230, 232, 234, 236, 238, 240, 242, 244, 246, 248, 250, 252, 254, 256, 258, 260, 262, 264, 266, 268, 270, 272, 274, 276, 278, 280, 282, 284, 286, 288, 290, 292, 294, 296, 298, 300, 302, 304, 306, 308, 310, 312, 314, 316, 318, 320, 322, 324, 326, 328, 330, 332, 334, 336, 338, 340, 342, 344, 346, 348, 350, 352, 354, 356, 358, 360, 362, 364, 366, 368, 370, 372, 374, 376, 378, 380, 382, 384, 386, 388, 390, 392, 394, 396, 398, 400, 402, 404, 406, 408, 410, 412, 414, 416, 418, 420, 422, 424, 426, 428, 430, 432, 434, 436, 438, 440, 442, 444, 446, 448, 450, 452, 454, 456, 458, 460, 462, 464, 466, 468, 470, 472, 474, 476, 478, 480, 482, 484, 486, 488, 490, 492, 494}
)

// Ordinal returns the ordinal number of a Country. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Country) Ordinal() int {
	switch v {
	case Afghanistan:
		return 0
	case Åland_Islands:
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
	case Czechia:
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
	case North_Macedonia:
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
	case Türkiye:
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

// String returns the literal string representation of a Country, which is
// the same as the const identifier but without prefix or suffix.
func (v Country) String() string {
	o := v.Ordinal()
	return v.toString(o, countryEnumStrings, countryEnumIndex[:])
}

func (v Country) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllCountries) {
		return fmt.Sprintf("Country(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Country is one of the defined constants.
func (v Country) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Country) Int() int {
	return int(v)
}

var invalidCountryValue = func() Country {
	var v Country
	for {
		if !slices.Contains(AllCountries, v) {
			return v
		}
		v++
	} // AllCountries is a finite set so loop will terminate eventually
}()

// CountryOf returns a Country based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Country is returned.
func CountryOf(v int) Country {
	if 0 <= v && v < len(AllCountries) {
		return AllCountries[v]
	}
	return invalidCountryValue
}

// Parse parses a string to find the corresponding Country, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsCountry.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Country)
//	err := v.Parse(s)
//	...  etc
func (v *Country) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := countryTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Country) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Country(num)
		return v.IsValid()
	}
	return false
}

func (v *Country) parseFallback(in, s string) error {
	if v.parseString(s, countryEnumInputs, countryEnumIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = iso3166_3LetterCodes[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised country")
}

func (v *Country) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllCountries[j-1]
			return true
		}
		i0 = i1
	}
	*v, ok = iso3166_3LetterCodes[s]
	return ok
}

// countryTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var countryTransformInput = func(in string) string {
	return strings.ToLower(strings.ReplaceAll(in, "_", " "))
}

// AsCountry parses a string to find the corresponding Country, accepting either one of the string values or
// a number. The input representation is determined by countryMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsCountry(s string) (Country, error) {
	var v = new(Country)
	err := v.Parse(s)
	return *v, err
}

// MustParseCountry is similar to AsCountry except that it panics on error.
// The input case does not matter.
func MustParseCountry(s string) Country {
	v, err := AsCountry(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Country) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Country) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Country) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, countryTextStrings, countryTextIndex[:]), nil
}

func (v Country) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Country) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Country) invalidError() error {
	return fmt.Errorf("%d is not a valid country", v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Country) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, countryJSONStrings, countryJSONIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Country) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, countryJSONStrings, countryJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Country) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Country) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := countryTransformInput(in)

	if v.parseString(s, countryTextInputs, countryTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Country) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Country) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := countryTransformInput(in)

	if v.parseString(s, countryJSONInputs, countryJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, countryEnumInputs, countryEnumIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = iso3166_3LetterCodes[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised country")
}

// countryMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var countryMarshalNumber = func(v Country) string {
	return strconv.FormatInt(int64(v), 10)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Country) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Country(x)
		return v.errorIfInvalid()
	case float64:
		*v = Country(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful country", value, value)
	}

	return v.scanParse(s)
}

func (v Country) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v *Country) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := countryTransformInput(in)

	if v.parseString(s, countrySQLInputs, countrySQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Value converts the Country to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Country) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, countrySQLStrings, countrySQLIndex[:]), nil
}
