package global

import (
	"bytes"
	"regexp"
	"sort"
)

// AtomByLength is a string slice sorter.
type AtomByLength []string

func (s AtomByLength) Len() int           { return len(s) }
func (s AtomByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s AtomByLength) Less(i, j int) bool { return len(s[i]) > len(s[j]) }

var (

	// Atoms is the list of existing atoms
	Atoms = map[string]string{
		"H":  "hydrogen",
		"He": "helium",
		"Li": "lithium",
		"Be": "berylium",
		"B":  "boron",
		"C":  "carbon",
		"N":  "nitrogen",
		"O":  "oxygen",
		"F":  "fluorine",
		"Ne": "neon",
		"Na": "sodium",
		"Mg": "magnesium",
		"Al": "aluminium",
		"Si": "silicon",
		"P":  "phosphorus",
		"S":  "sulfure",
		"Cl": "chlorine",
		"Ar": "argon",
		"K":  "potassium",
		"Ca": "calcium",
		"Sc": "scandium",
		"Ti": "titanium",
		"V":  "vanadium",
		"Cr": "chromium",
		"Mn": "manganese",
		"Fe": "iron",
		"Co": "cobalt",
		"Ni": "nickel",
		"Cu": "copper",
		"Zn": "zinc",
		"Ga": "gallium",
		"Ge": "germanium",
		"As": "arsenic",
		"Se": "sefeniuo",
		"Br": "bromine",
		"Kr": "krypton",
		"Rb": "rubidium",
		"Sr": "strontium",
		"Y":  "yltrium",
		"Zr": "zirconium",
		"Nb": "niobium",
		"Mo": "molybdenum",
		"Tc": "technetium",
		"Ru": "ruthenium",
		"Rh": "rhodium",
		"Pd": "palladium",
		"Ag": "silver",
		"Cd": "cadmium",
		"In": "indium",
		"Sn": "tin",
		"Sb": "antimony",
		"Te": "tellurium",
		"I":  "iodine",
		"Xe": "xenon",
		"Cs": "caesium",
		"Ba": "barium",
		"Hf": "hafnium",
		"Ta": "tantalum",
		"W":  "tungsten",
		"Re": "rhenium",
		"Os": "osmium",
		"Ir": "iridium",
		"Pt": "platinium",
		"Au": "gold",
		"Hg": "mercury",
		"Tl": "thallium",
		"Pb": "lead",
		"Bi": "bismuth",
		"Po": "polonium",
		"At": "astatine",
		"Rn": "radon",
		"Fr": "francium",
		"Ra": "radium",
		"Rf": "rutherfordium",
		"Db": "dubnium",
		"Sg": "seaborgium",
		"Bh": "bohrium",
		"Hs": "hassium",
		"Mt": "meitnerium",
		"Ds": "darmstadtium",
		"Rg": "roentgenium",
		"Cn": "copemicium",
		"La": "lanthanum",
		"Ce": "cerium",
		"Pr": "praseodymium",
		"Nd": "neodymium",
		"Pm": "promethium",
		"Sm": "samarium",
		"Eu": "europium",
		"Gd": "gadolinium",
		"Tb": "terbium",
		"Dy": "dysprosium",
		"Ho": "holmium",
		"Er": "erbium",
		"Tm": "thulium",
		"Yb": "ytterbium",
		"Lu": "lutetium",
		"Ac": "actinium",
		"Th": "thorium",
		"Pa": "protactinium",
		"U":  "uranium",
		"Np": "neptunium",
		"Pu": "plutonium",
		"Am": "americium",
		"Cm": "curium",
		"Bk": "berkelium",
		"Cf": "californium",
		"Es": "einsteinium",
		"Fm": "fermium",
		"Md": "mendelevium",
		"No": "nobelium",
		"Lr": "lawrencium",
		"D":  "deuterium",
	}

	// FormulaRe is the regex matching a chemical formula
	FormulaRe *regexp.Regexp

	// BasicMolRe is the regex matching a chemical formula (atoms and numbers only)
	BasicMolRe *regexp.Regexp

	// OneGroupMolRe is a (AYZ)n molecule like regex
	OneGroupMolRe *regexp.Regexp
)

func init() {

	sortedAtoms := make([]string, 0, len(Atoms))
	for k := range Atoms {
		sortedAtoms = append(sortedAtoms, k)
	}
	// the atom must be sorted by decreasing size
	// to match first Cl before C for example
	sort.Sort(AtomByLength(sortedAtoms))

	// building the basic molecule regex
	// (atom1|atom2|...)([1-9]*)
	var buf bytes.Buffer
	buf.WriteString("(")
	for _, a := range sortedAtoms {
		buf.WriteString(a)
		buf.WriteString("|")
	}
	// removing the last |
	buf.Truncate(buf.Len() - 1)
	buf.WriteString(")")
	buf.WriteString("([0-9]+)*")
	BasicMolRe = regexp.MustCompilePOSIX(buf.String())

	// building the one group molecule regex
	buf.Reset()
	buf.WriteString("(?:\\(|\\[)")
	buf.WriteString("((?:[")
	for _, a := range sortedAtoms {
		buf.WriteString(a)
		buf.WriteString("|")
	}
	// removing the last |
	buf.Truncate(buf.Len() - 1)
	buf.WriteString("]+[0-9]*)+)")
	buf.WriteString("(?:\\)|\\])")
	buf.WriteString("([0-9]*)")

	OneGroupMolRe = regexp.MustCompile(buf.String())

	FormulaRe = regexp.MustCompile(`[A-Za-z0-9,\^]+`)

}
