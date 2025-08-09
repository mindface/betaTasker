"use client"
import React, { useState,useRef, useEffect } from 'react'
import { useSelector } from 'react-redux'
import Image from 'next/image'
import { Posts } from '../model/posts'

import ContentTask01 from "./ContentTask01"
import ContentHeader from "./ContentTasksHeader"
import ItemCard from "./parts/ItemCard"
import Loading from "./Loading"
import { RootState } from "../store"

export interface State {
  status:String,
  posts: Posts[],
  rePosts: Posts[],
  lastUpdated: number,
}

export default function SectionTaskView() {
  const [switchClass,setSwitchClass] = useState('')
  const [posts,setPosts] = useState<Posts[]>([])
  const postsSwtch = useRef<boolean>(false)
  const postsState = useRef<Posts[]>([])
  const [isClient, setIsClient] = useState(false)
  const postsStateSelector = useSelector((state:RootState) => {
    postsState.current = state.post.rePosts
    return state.post.rePosts ? state.post.rePosts : []
  })
  const currentLoding = useSelector((state:RootState) => {
    return state.post.loading
  })

  function viewSwitch(type:string){
    setSwitchClass(type)
  }

  function reSetPosts(posts:Posts[],text:string,type:string){
    if(text === ''){
      setPosts(posts)
    }
    postsSwtch.current = true
    const rePoster = postsState.current.filter((item:Posts) => {
      let post = item.title
      if(type !== 'title') {
        post = item.name
      }
      if( post?.indexOf(text) != -1 ){
        return item
      }
    })
    setPosts(rePoster)
  }

  useEffect(() => {
    setIsClient(true)
  },[])

  return(
    <ContentTask01>
      <>
        <ContentHeader searchAction={reSetPosts} title="タスク形式"></ContentHeader>
        <div className="section__inner section--task">
          <div className="view-switch">
            <button className="btn" onClick={e => viewSwitch('list')}>
              <Image
                src="/image/list.svg"
                alt=""
                width={20}
                height={20}
              />
            </button>
            <button className="btn" onClick={e => viewSwitch('cardview')}>
              <Image
                src="/image/card.svg"
                alt=""
                width={20}
                height={20}
              />
            </button>
            <button className="btn" onClick={e => viewSwitch('category')}>
              <Image
                src="/image/link.svg"
                alt=""
                width={20}
                height={20}
              />
            </button>
          </div>
          {isClient &&
            <div className={`task-box _flex_ ${switchClass}`}>
              {currentLoding && <Loading />}
              {!postsSwtch.current && postsStateSelector.map((item:Posts) => {
                return (<ItemCard posts={item} key={Number(item.id)}></ItemCard>)
              })}
              {postsSwtch.current && posts.map((item:Posts) => {
                return (<ItemCard posts={item} key={Number(item.id)}></ItemCard>)
              })}
            </div>
          }
        </div>
      </>
    </ContentTask01>
  )
}
