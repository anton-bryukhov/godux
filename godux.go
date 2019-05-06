package godux

type Reducer func(state State, action Action) State

type State interface{}

type Action struct {
	Type    string
	Payload interface{}
}
