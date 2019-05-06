package godux

type Reducer func(state State, action Action) State

type State interface{}

type Action struct {
	Type    string
	Payload interface{}
}

type Store struct {
	state   State
	reducer Reducer
}

func (s Store) GetState() State {
	return s.state
}

func CreateStore(initialState State, reducer Reducer) Store {
	return Store{initialState, reducer}
}
