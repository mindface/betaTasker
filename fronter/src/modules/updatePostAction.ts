import { Action, GET_POST_REQUST, GET_POST_SUCCESS, GET_POST_FAILURE } from './action'
import { Posts } from '../model/posts'

export interface State {
  status:String,
  posts: Posts[],
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

const getPostSuccess = (json:Posts[]): Action => {
  return {
    type: GET_POST_SUCCESS,
    payload: json,
    rePayload: [],
    loading:false
  }
}

const getPostFailure = (error:object): Action => {
  return {
    type: GET_POST_FAILURE,
    payload: error,
    rePayload: [],
    loading:false
  }
}

export const updatePostAction = (data:Posts) => {
   return (dispatch:any) => {
      dispatch(getPostRequest())
      return fetch(`http://localhost:8080/api/updatebook/${data.id}`,{
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
        },
        credentials: 'same-origin',
        body: JSON.stringify(data)
      }).then((res) => {
        res.json().then((res) => {
          console.log(res)
        })
      }).catch(err => {
        dispatch(getPostFailure(err))
      })
   }
}
