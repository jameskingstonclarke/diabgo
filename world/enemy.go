package world

type GoblinBehaviour struct {
	// the information about the goblin is here
	angry bool
}

func (goblin GoblinBehaviour) Update() {
	// if withn reach of the player, then get angry
}

func NewGoblin() *Agent{
	return &Agent{
		Tag:        "goblin",
		Inventory:  Inventory{},
		Components: nil,
		Behaviouor: GoblinBehaviour{
			angry: false,
		},
	}
}