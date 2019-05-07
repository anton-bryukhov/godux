package godux

import (
	"reflect"
)

type Action struct {
	Type    string
	Payload interface{}
}

type State map[string]interface{}
type Reducer func(State, Action) State

func CombineReducers(reducers map[string]interface{}) Reducer {
	return func(state State, action Action) State {
		nextState := State{}

		for key, reducer := range reducers {
			reducerValue := reflect.ValueOf(reducer)
			prevState := state[key]
			reducerArguments := []reflect.Value{
				reflect.ValueOf(prevState),
				reflect.ValueOf(action),
			}

			nextState[key] = reducerValue.Call(reducerArguments)[0].Interface()
		}

		return nextState
	}
}

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
