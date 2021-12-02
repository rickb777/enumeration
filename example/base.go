package example

// This example demonstrates using floating point values instead of integers.
// These are float32 but could be float64; the only ither restriction is that
// no two values can be the same number.

//go:generate enumeration -v -f -type Base -lc

type Base float32

// Nucleotide Molecular Weights, g/mol
const A Base = 331.2
const C Base = 307.2
const G Base = 347.2
const T Base = 322.2
