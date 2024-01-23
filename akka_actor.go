package akka_models

type Actor struct {
	state int
	mailbox chan int
}

func NewActor(initialState int) *Actor {
	return &Actor{
		state: initialState,
		mailbox: make(chan int),
	}
}

func (a *Actor) ProcessMessage(msg int) {
	a.state += msg
}

func (a *Actor) Run() {
	for msg := range a.mailbox {
		a.ProcessMessage(msg)
	}
}

type System struct {
	actors []*Actor
}

func NewSystem(numActors int) *System {
	var system System
	for i := 1; i <= numActors; i++ {
		actor := NewActor(i)
		system.actors = append(system.actors, actor)
		go actor.Run()
	}
	return &system
}

func (s *System) SendMessage(msg int) {
	actorIndex := msg % len(s.actors)
	s.actors[actorIndex].mailbox <- msg
}

func main() {

	system := NewSystem(3)

	for i := 1; i <= 5;  i++ {
		go system.SendMessage(i)
	}

}
