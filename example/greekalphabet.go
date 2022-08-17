package example

//go:generate enumeration -v -type GreekAlphabet

// GreekAlphabet: This example shows non-ASCII characters in use. There is also a cross-mapping
// table specified via the '-using' option. So the parser recognises inputs from both
// sets of strings.
//
// See also
// https://unicode.org/charts/PDF/U0370.pdf
// https://en.wikipedia.org/wiki/Greek_alphabet
type GreekAlphabet int

const (
	Αλφα    GreekAlphabet = iota + 1 // text:"alpha"   sql:"\u0391" = Α
	Βήτα                             // text:"beta"    sql:"\u0392"
	Γάμμα                            // text:"gamma"   sql:"\u0393"
	Δέλτα                            // text:"delta"   sql:"\u0394"
	Εψιλον                           // text:"epsilon" sql:"\u0395"
	Ζήτα                             // text:"zeta"    sql:"\u0396"
	Ητα                              // text:"eta"     sql:"\u0397"
	Θήτα                             // text:"theta"   sql:"\u0398"
	Ιώτα                             // text:"iota"    sql:"\u0399"
	Κάππα                            // text:"kappa"   sql:"\u039A"
	Λάμβδα                           // text:"lambda"  sql:"\u039B"
	Μυ                               // text:"mu"      sql:"\u039C"
	Νυ                               // text:"nu"      sql:"\u039D"
	Ξι                               // text:"xi"      sql:"\u039E"
	Ομικρον                          // text:"omicron" sql:"\u039F"
	Πι                               // text:"pi"      sql:"\u03A0"
	Ρώ                               // text:"rho"     sql:"\u03A1"
	Σίγμα                            // text:"sigma"   sql:"\u03A3"
	Ταυ                              // text:"tau"     sql:"\u03A4"
	Υψιλον                           // text:"upsilon" sql:"\u03A5"
	Φι                               // text:"phi"     sql:"\u03A6"
	Χι                               // text:"chi"     sql:"\u03A7"
	Ψι                               // text:"psi"     sql:"\u03A8"
	Ωμέγα                            // text:"omega"   sql:"\u03A9"
	// n.b. there is no u03A2
)
