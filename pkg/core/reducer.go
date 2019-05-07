package core

import "reflect"

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
