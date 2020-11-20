package enum_test

import (
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/enum"
	"github.com/rickb777/enumeration/example"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	RegisterTestingT(t)

	days := example.AllDayEnums.Strings()

	Ω(days).Should(HaveLen(7))
	Ω(strings.Join(days, "|")).Should(Equal("Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday"))
}

func TestOrdinals(t *testing.T) {
	RegisterTestingT(t)

	days := enum.Enums{example.Wednesday, example.Friday}.Ordinals()

	Ω(days).Should(HaveLen(2))
	Ω(days[0]).Should(Equal(3))
	Ω(days[1]).Should(Equal(5))
}
