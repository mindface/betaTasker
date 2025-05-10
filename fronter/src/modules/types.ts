import { Action } from "redux"
import { ActionTypes } from "./actionTypes"
import { Posts } from "../model/posts"

export type Post = {
  post: Posts[]
}

interface GetPostAction extends Action {
  posts: typeof ActionTypes.GetPost
}

interface AddPostAction extends Action {
  posts: typeof ActionTypes.AddPost
}

export type PostActionTypes = GetPostAction | AddPostAction
