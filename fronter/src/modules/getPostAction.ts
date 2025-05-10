import { Action, GET_POST_REQUST, GET_POST_SUCCESS, GET_POST_FAILURE } from './action'
import { Posts } from '../model/posts'

const getPostRequest = (): Action => {
  return {
    type: GET_POST_REQUST,
    payload: [],
    rePayload: [],
    loading: true
  }
}

const getPostSuccess = (json:Posts[]): Action => {
  return {
    type: GET_POST_SUCCESS,
    payload: json,
    rePayload: json,
    loading: false
  }
}

const getPostFailure = (error:object): Action => {
  return {
    type: GET_POST_FAILURE,
    payload: error,
    rePayload: [],
    loading: false
  }
}

export const getPostAction = () => {
   return (dispatch:any) => {
      dispatch(getPostRequest())
      return fetch('http://localhost:8080/api/book',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
        },
        credentials: 'same-origin',
      }).then((res) => {
        res.json().then((res) => {
          setTimeout(() => {
            dispatch(getPostSuccess(res.list))
          },500)
        })
      }).catch(err => {
        dispatch(getPostFailure(err))
      })
   }
}
