package example

// This example shows non-ASCII characters in use. There is also a cross-mapping
// table specified via the '-using' option. So the parser recognises inputs from both
// sets of strings.

// See also
// https://unicode.org/charts/PDF/U0370.pdf
// https://en.wikipedia.org/wiki/Greek_alphabet

//go:generate enumeration -v -type GreekAlphabet -using greekStrings -marshaltext tag

type GreekAlphabet int

const (
	Αλφα GreekAlphabet = iota + 1 // U0391 = Α
	Βήτα
	Γάμμα
	Δέλτα
	Εψιλον
	Ζήτα // U0396
	Ητα
	Θήτα
	Ιώτα
	Κάππα
	Λάμβδα // U039b
	Μυ
	Νυ
	Ξι
	Ομικρον
	Πι // U03A0
	Ρώ
	// there is no U03A2
	Σίγμα // U03a3
	Ταυ
	Υψιλον
	Φι
	Χι // U03a7
	Ψι
	Ωμέγα
)

var greekStrings = map[GreekAlphabet]string{
	Αλφα:    "alpha",
	Βήτα:    "beta",
	Γάμμα:   "gamma",
	Δέλτα:   "delta",
	Εψιλον:  "epsilon",
	Ζήτα:    "zeta",
	Ητα:     "eta",
	Θήτα:    "theta",
	Ιώτα:    "iota",
	Κάππα:   "kappa",
	Λάμβδα:  "lambda",
	Μυ:      "mu",
	Νυ:      "nu",
	Ξι:      "xi",
	Ομικρον: "omicron",
	Πι:      "pi",
	Ρώ:      "rho",
	Σίγμα:   "sigma",
	Ταυ:     "tau",
	Υψιλον:  "upsilon",
	Φι:      "phi",
	Χι:      "chi",
	Ψι:      "psi",
	Ωμέγα:   "omega",
}
