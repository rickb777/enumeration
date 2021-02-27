package example

// https://unicode.org/charts/PDF/U0370.pdf

type GreekAlphabet int

const (
	Α GreekAlphabet = iota // U0391
	Β
	Γ
	Δ
	Ε
	Ζ // U0396
	Η
	Θ
	Ι
	Κ
	Λ // U039b
	Μ
	Ν
	Ξ
	Ο
	Π // U03A0
	Ρ
	// there is no U03A2
	Σ // U03a3
	Τ
	Υ
	Φ
	Χ // U03a7
	Ψ
	Ω
)

var greekStrings = map[GreekAlphabet]string{
	Α: "alpha",
	Β: "beta",
	Γ: "gamma",
	Δ: "delta",
	Ε: "epsilon",
	Ζ: "zeta",
	Η: "eta",
	Θ: "theta",
	Ι: "iota",
	Κ: "kappa",
	Λ: "lambda",
	Μ: "mu",
	Ν: "nu",
	Ξ: "xi",
	Ο: "omicron",
	Π: "pi",
	Ρ: "rho",
	Σ: "sigma",
	Τ: "tau",
	Υ: "upsilon",
	Φ: "phi",
	Χ: "chi",
	Ψ: "psi",
	Ω: "omega",
}
