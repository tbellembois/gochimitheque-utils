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
	SortedByLengthAtoms []string

	// Atoms is the list of existing atoms
	Atoms = map[string]string{
		"Ac": "actinium",
		"Ag": "silver",
		"Al": "aluminium",
		"Am": "americium",
		"Ar": "argon",
		"As": "arsenic",
		"At": "astatine",
		"Au": "gold",
		"B":  "boron",
		"Ba": "barium",
		"Be": "berylium",
		"Bh": "bohrium",
		"Bi": "bismuth",
		"Bk": "berkelium",
		"Br": "bromine",
		"C":  "carbon",
		"Ca": "calcium",
		"Cd": "cadmium",
		"Ce": "cerium",
		"Cf": "californium",
		"Cl": "chlorine",
		"Cm": "curium",
		"Cn": "copemicium",
		"Co": "cobalt",
		"Cr": "chromium",
		"Cs": "caesium",
		"Cu": "copper",
		"D":  "deuterium",
		"Db": "dubnium",
		"Ds": "darmstadtium",
		"Dy": "dysprosium",
		"Er": "erbium",
		"Es": "einsteinium",
		"Eu": "europium",
		"F":  "fluorine",
		"Fe": "iron",
		"Fm": "fermium",
		"Fr": "francium",
		"Ga": "gallium",
		"Gd": "gadolinium",
		"Ge": "germanium",
		"H":  "hydrogen",
		"He": "helium",
		"Hf": "hafnium",
		"Hg": "mercury",
		"Ho": "holmium",
		"Hs": "hassium",
		"I":  "iodine",
		"In": "indium",
		"Ir": "iridium",
		"K":  "potassium",
		"Kr": "krypton",
		"La": "lanthanum",
		"Li": "lithium",
		"Lr": "lawrencium",
		"Lu": "lutetium",
		"Md": "mendelevium",
		"Mg": "magnesium",
		"Mn": "manganese",
		"Mo": "molybdenum",
		"Mt": "meitnerium",
		"N":  "nitrogen",
		"Na": "sodium",
		"Nb": "niobium",
		"Nd": "neodymium",
		"Ne": "neon",
		"Ni": "nickel",
		"No": "nobelium",
		"Np": "neptunium",
		"O":  "oxygen",
		"Os": "osmium",
		"P":  "phosphorus",
		"Pa": "protactinium",
		"Pb": "lead",
		"Pd": "palladium",
		"Pm": "promethium",
		"Po": "polonium",
		"Pr": "praseodymium",
		"Pt": "platinium",
		"Pu": "plutonium",
		"Ra": "radium",
		"Rb": "rubidium",
		"Re": "rhenium",
		"Rf": "rutherfordium",
		"Rg": "roentgenium",
		"Rh": "rhodium",
		"Rn": "radon",
		"Ru": "ruthenium",
		"S":  "sulfure",
		"Sb": "antimony",
		"Sc": "scandium",
		"Se": "sefeniuo",
		"Sg": "seaborgium",
		"Si": "silicon",
		"Sm": "samarium",
		"Sn": "tin",
		"Sr": "strontium",
		"Ta": "tantalum",
		"Tb": "terbium",
		"Tc": "technetium",
		"Te": "tellurium",
		"Th": "thorium",
		"Ti": "titanium",
		"Tl": "thallium",
		"Tm": "thulium",
		"U":  "uranium",
		"V":  "vanadium",
		"W":  "tungsten",
		"Xe": "xenon",
		"Y":  "yltrium",
		"Yb": "ytterbium",
		"Zn": "zinc",
		"Zr": "zirconium",
	}

	// FormulaRe is the regex matching a chemical formula
	FormulaRe *regexp.Regexp

	// BasicMolRe is the regex matching a chemical formula (atoms and numbers only)
	BasicMolRe *regexp.Regexp

	// OneGroupMolRe is a (AYZ)n molecule like regex
	OneGroupMolRe *regexp.Regexp
)

func init() {

	SortedByLengthAtoms = make([]string, 0, len(Atoms))
	for k := range Atoms {
		SortedByLengthAtoms = append(SortedByLengthAtoms, k)
	}
	// the atom must be sorted by decreasing size
	// to match first Cl before C for example
	sort.Sort(AtomByLength(SortedByLengthAtoms))

	// building the basic molecule regex
	// (atom1|atom2|...)([1-9]*)
	var buf bytes.Buffer
	buf.WriteString("(")
	for _, a := range SortedByLengthAtoms {
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
	for _, a := range SortedByLengthAtoms {
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
