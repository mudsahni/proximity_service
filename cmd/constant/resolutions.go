package constant

type H3Resolution int

const (
	H3Resolution9  = "h3indexresolution9"
	H3Resolution12 = "h3indexresolution12"
)

var H3ResolutionStrings = map[int]string{
	9:  H3Resolution9,
	12: H3Resolution12,
}
var H3ResolutionValues = map[string]int{
	H3Resolution9:  9,
	H3Resolution12: 12,
}
