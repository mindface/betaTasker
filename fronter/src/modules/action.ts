import { Posts } from '../model/posts'

export type Action = {
  type: String,
  payload: any,
  rePayload: Posts[] | any,
  loading: boolean
}

export type textAction = {
  type: String,
  searchText: string,
  searchCategpryText: string,
}

export const GET_POST_REQUST = "GET_POST_REQUST"
export const GET_POST_SUCCESS = "GET_POST_SUCCESS"
export const GET_POST_FAILURE = "GET_POST_FAILURE"
