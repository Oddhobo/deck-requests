package deck

const deckEndpoint string = "https://deckofcardsapi.com/api/deck"

/* Card represents a card*/
type Card struct {
	FaceImage string `json:"image"`
	Value     string `json:"value"`
	Suit      string `json:"suit"`
	Code      string `json:"code"`
}

/* Hand is a slice of cards*/
type Hand struct {
	Cards []Card
}

type Deck struct {
	DeckURI string `json:"deck_id"`
}
