import { Action, textAction, GET_POST_REQUST, GET_POST_SUCCESS, GET_POST_FAILURE } from './action'
import { Posts } from '../model/posts'

export interface State {
  status:String,
  posts: Posts[],
  rePosts: Posts[],
  lastUpdated: number,
}

const getPostRequest = (): Action => {
  return {
    type: GET_POST_REQUST,
    payload: null,
    rePayload: [],
    loading:false
  }
}

const getPostSuccess = (json:Posts[],rejson:Posts[]): Action => {
  return {
    type: GET_POST_SUCCESS,
    payload: json,
    rePayload: rejson,
    loading:false
  }
}

const setPostTextSuccess = (text:string,rejson:string): textAction => {
  return {
    type: GET_POST_SUCCESS,
    searchText: text,
    searchCategpryText: rejson,
  }
}

const getPostFailure = (error:string): Action => {
  return {
    type: GET_POST_FAILURE,
    payload: error,
    rePayload: [],
    loading:false
  }
}

export const searchPostAction = (text:string) => {
  return (dispatch:any,state:any) => {
    state().state.searchText = text
    dispatch(setPostTextSuccess(text,text))
  }
}

export const searchPostCategoryAction = (text:string) => {
  return (dispatch:any,state:any) => {
    console.log(text)
    try {
     const posts = state().state.posts
     if(text === 'all'){
      dispatch(getPostSuccess(posts,posts))
      return
     }
     const reText = posts.filter((item:Posts) => {
       if( item.name?.indexOf(text) != -1 ){
         return item
       }
     })
     dispatch(getPostSuccess(posts,reText))
    }catch{
      dispatch(getPostFailure('error'))
    }
  }
}
