package day7

func gt(ht handType) string {
	switch ht {
	case highCard:
		return "highCard"
	case onePair:
		return "onePair"
	case twoPair:
		return "twoPair"
	case threeOfKind:
		return "threeOfKind"
	case fullHouse:
		return "fullHouse"
	case fourOfKind:
		return "fourOfKind"
	case fiveOfKind:
		return "fiveOfKind"
	default:
		return ""
	}
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

type play struct {
	hand  string
	bid   int
	hType handType
}
