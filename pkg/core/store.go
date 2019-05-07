package core

type State map[string]interface{}

type Store struct {
	reducer Reducer
	state   State
}

func (s Store) GetState() map[string]interface{} {
	return s.state
}

func (s *Store) Dispatch(action Action) {
	s.state = s.reducer(s.state, action)
}

func CreateStore(reducer Reducer, preloadedState State) Store {
	return Store{reducer: reducer, state: preloadedState}
}
