package world

type World struct {
	Agents []*Agent
}

var (
	WorldInstance *World
)


func (world *World) Update(){
	for _, agent := range world.Agents {
		agent.Update()
	}
}

func (world *World) AddAgent(agent *Agent){
	world.Agents = append(world.Agents, agent)
}

func (world *World) GetAgent(tag string) *Agent{
	for _, agent := range world.Agents {
		if agent.Tag == tag{
			return agent
		}
	}
	return nil
}


func NewWorld() *World {

	WorldInstance = &World{}
	WorldInstance.AddAgent(NewAgent("player"))

	return WorldInstance
}