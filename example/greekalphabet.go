package example

// See also
// https://unicode.org/charts/PDF/U0370.pdf
// https://en.wikipedia.org/wiki/Greek_alphabet

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
