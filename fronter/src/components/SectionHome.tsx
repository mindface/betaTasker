"use client"
import React, { useState, useEffect, useRef } from 'react'
import { useDispatch, useStore, useSelector } from 'react-redux';
import { State } from '../modules/reducer'
import Image from 'next/image'

interface titleType{
  title: string
}
export default function SectionHome() {
  const store = useStore()
  const dispatch = useDispatch()
  const [post_data, setPostData] = useState([]);
  const post_d = useRef(null);
  const d_name = useRef<HTMLInputElement | null>(null);
  const d_text = useRef<HTMLInputElement | null>(null);
  const d_disc = useRef<HTMLInputElement | null>(null);
  const d_imgPath = useRef<HTMLInputElement | null>(null);
  const postsState = useSelector((payload:{state:State}) => {
    return payload.state.posts
  })

  useEffect(() => {
    console.log(postsState)
  },[])


  return(
    <div className="home _flex_c_">
      <Image
        src="/image/home.svg"
        alt="ズズズ..."
        width={400}
        height={400}
      />
    </div>
  )
}

