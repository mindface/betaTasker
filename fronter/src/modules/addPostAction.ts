import { Action, GET_POST_REQUST, GET_POST_SUCCESS, GET_POST_FAILURE } from './action'
import { Posts } from '../model/posts'
import { getPostAction } from './getPostAction'

const getPostRequest = (): Action => {
  return {
    type: GET_POST_REQUST,
    payload: [],
    rePayload: [],
    loading:true
  }
}

const getPostSuccess = (json:Posts[]): Action => {
  return {
    type: GET_POST_SUCCESS,
    payload: json,
    rePayload: json,
    loading:false
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

export const addPostAction = (sendData:Posts) => {
   return (dispatch:any) => {
      dispatch(getPostRequest())
      return fetch('http://localhost:8080/api/book',{
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'text/plain',
          // 'Access-Control-Allow-Origin': '*'
        },
        credentials: 'same-origin',
        body: JSON.stringify(sendData)
      }).then((res) => {
        res.json().then((res) => {
          dispatch(getPostSuccess(res))
          console.log(JSON.stringify(res))
          dispatch(getPostAction)
        })
      }).catch(err => {
          console.log(err)
          dispatch(getPostFailure(err))
      })
   }
}
