import { combineReducers } from 'redux';
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Posts } from '../model/posts'

export interface State {
  posts: Posts[]
}

export const initialState: State = {
  posts: [
   {
   id: 1,
   title: "test01",
   text: "test01",
   name: "test01",
   disc: "test01",
   imgPath: "test01",
  }
  // ,{
  //  id: 2,
  //  text: "test02",
  //  name: "test02",
  //  disc: "test02",
  //  imgPath: "test02",
  // },
  ]
}

export const rootReducer = reducerWithInitialState(initialState)
