package core

type State map[string]interface{}
type Middleware func(Action, *Store)

type Store struct {
	middlewares []Middleware
	reducer     Reducer
	state       State
}

func (s Store) GetState() map[string]interface{} {
	return s.state
}

func (s *Store) Dispatch(action Action) {
	for _, middleware := range s.middlewares {
		middleware(action, s)
	}

	s.state = s.reducer(s.state, action)
}

func (s *Store) ApplyMiddleware(middlewares ...Middleware) {
	s.middlewares = middlewares
}

func CreateStore(reducer Reducer, preloadedState State) Store {
	return Store{reducer: reducer, state: preloadedState}
}
