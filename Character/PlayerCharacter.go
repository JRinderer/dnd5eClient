package Character

type PlayerCharacter struct {
	CharacterName string
	PlayerName    string
	Invent        Inventory
}

type Inventory struct {
	Items []string
}
