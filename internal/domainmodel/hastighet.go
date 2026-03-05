package domainmodel

// Tagkategori contains standard train category constants.
var (
	TagkategoriA = "A"
	TagkategoriB = "B"
	TagkategoriC = "C"
	TagkategoriS = "S"
)

// Hastighetsprofil groups speed sections by category.
type Hastighetsprofil struct {
	Kategori string
	Strackor Entiteter[*Hastighetsstracka]
}

// NewHastighetsprofil creates a new Hastighetsprofil.
func NewHastighetsprofil(kategori string) *Hastighetsprofil {
	return &Hastighetsprofil{
		Kategori: kategori,
	}
}

// Hastighetsstracka represents a speed section with a spread across a link.
type Hastighetsstracka struct {
	*EntitetMedUtbredning
	Kategori  string
	Hastighet int
}

// NewHastighetsstracka creates a new Hastighetsstracka.
func NewHastighetsstracka(kategori string, hastighet int, lankutbredning *Lankutbredning) *Hastighetsstracka {
	return &Hastighetsstracka{
		EntitetMedUtbredning: NewEntitetMedUtbredning(lankutbredning),
		Kategori:             kategori,
		Hastighet:            hastighet,
	}
}
