package player

type Player struct {
	Name           string
	State          PlayerState
	IsServingFirst bool
}

type PlayerState struct {
	IsHitting bool
	IsWaiting bool
}

func NewPlayer(name string, isServingFirst bool) *Player {
	var isHitting bool
	if isServingFirst {
		isHitting = true
	} else {
		isHitting = false
	}
	return &Player{
		Name: name,
		State: PlayerState{
			IsHitting: isHitting,
		},
		IsServingFirst: isServingFirst,
	}
}
