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
    loading:false
  }
}

export const deletePostAction = (postId:string) => {
   console.log("postId")
   console.log(postId)
   return (dispatch:any) => {
      dispatch(getPostRequest())
      return fetch(`http://localhost:8080/api/deletebook/${postId}`,{
        method: 'DELETE',
        mode: 'cors',
        headers: {
          'Content-Type': 'text/plain',
          // 'Access-Control-Allow-Origin': '*'
        },
        credentials: 'same-origin',
      }).then((res) => {
        res.json().then((res) => {
          console.log(res)
          dispatch(getPostAction)
        })
      }).catch(err => {
          console.log(err)
          dispatch(getPostFailure(err))
      })
   }
}
