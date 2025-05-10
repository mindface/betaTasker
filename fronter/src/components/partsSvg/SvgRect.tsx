import React, { useState, useEffect, useRef } from 'react'

export default function SvgRect() {

  useEffect(() => {

  },[])

  const onClick = (e:any) => {
    console.log(e.target)
  }

  return(
    <g fill="#c44" stroke="#822" onClick={onClick} strokeWidth="2">
      <circle cx="32" cy="32" r="30" />
      <rect x="40" y="30" width="80" height="50" />
    </g>
  )
}

