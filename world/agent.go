package world

/*
An Agent acts as an in game entity
*/
type Agent struct {
	Tag 	string
	Inventory Inventory
	Components []Component
	Behaviouor AgentBehaviour
}

type AgentBehaviour interface{
	Update()
}

func (agent Agent) Update(){
	agent.Behaviouor.Update()
}