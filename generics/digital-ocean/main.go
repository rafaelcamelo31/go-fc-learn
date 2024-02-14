package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card interface {
	fmt.Stringer
	Name() string
}

type TradingCard struct {
	CollectableName string
}

func (tc *TradingCard) Name() string {
	return tc.String()
}

func NewTradingCard(collectableName string) *TradingCard {
	return &TradingCard{
		CollectableName: collectableName,
	}
}

func (tc *TradingCard) String() string {
	return tc.CollectableName
}

type PlayingCard struct {
	Suit string
	Rank string
}

func (pc *PlayingCard) Name() string {
	return pc.String()
}

func NewPlayingCard(suite string, card string) *PlayingCard {
	return &PlayingCard{
		Suit: suite,
		Rank: card,
	}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

type Deck[C Card] struct {
	cards []C
}

func NewTradingCardDeck() *Deck[*TradingCard] {
	collectables := []string{
		"Magic The Gathering",
		"Pokemon",
		"Yugioh",
	}

	deck := &Deck[*TradingCard]{}
	for _, collectable := range collectables {
		deck.AddCard(NewTradingCard(collectable))
	}
	return deck
}

func NewPlayingCardDeck() *Deck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func (d *Deck[C]) AddCard(card C) {
	d.cards = append(d.cards, card)
}

func (d *Deck[C]) RandomCard() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

func printCard[C Card](card C) {
	fmt.Println("card name:", card.Name())
}

func main() {
	playingDeck := NewPlayingCardDeck()
	tradingDeck := NewTradingCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard := playingDeck.RandomCard()
	fmt.Printf("drew card: %s\n", playingCard)

	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)

	fmt.Printf("--- drawing playing card ---\n")
	tradingCard := tradingDeck.RandomCard()
	fmt.Printf("drew card: %s\n", tradingCard)
	fmt.Printf("card collectable name: %s\n", tradingCard.CollectableName)

	fmt.Printf("--- printing cards ---\n")
	printCard[*PlayingCard](playingCard)
	printCard(tradingCard)
}
