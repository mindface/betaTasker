"use client"
import React, { ChangeEvent, useRef, useState, useEffect } from 'react'
import LinkData from '../json/link.json'

type Props = {
  title: string;
}

type UseState = {
  selectTitle: string;
  urlText: string;
  radioSelect: string;
  targetSite: string;
  checkUrl: string;
  urlCopy: {id:string;value:string,view:boolean}[]
}

export default function ItemCardSearchLink( props: Props )  {
  const useLinker = useRef('https://www.google.com/search?q=')
  const [state,setState] = useState<UseState>({
    selectTitle: '',
    urlText: '',
    radioSelect: '',
    targetSite: '',
    checkUrl: '',
    urlCopy: [],
  })

  const changeHandler = (e:ChangeEvent<HTMLInputElement>)=> {
    e.preventDefault()
    setState({...state, [e.target.name]: e.target.value})
  }

  function viewPage() {
    let setOpenUrl = state.urlText
    if(state.selectTitle !== ''){
      setOpenUrl += `${state.selectTitle}`
    }
    if(state.urlText !== ''){
      console.log(setOpenUrl)
      window.open(setOpenUrl)
    }else {
      alert('urlを確認してください。')
    }
  }

  function listener(e:ClipboardEvent,url:string) {
    if(e.clipboardData) {
      e.clipboardData.setData("text/plain",url)
      e.preventDefault()
    }
  }

  function urlCopy(e:any,id:number,url:string) {
    setState({...state,urlText:url})
    console.log(state.urlText)
    document.addEventListener('copy',(e:ClipboardEvent) => listener(e,url))
    document.execCommand('copy')
    const copyedElement = document.createElement('span')
    copyedElement.className = 'aid'
    copyedElement.innerText = 'コピーしました。'
    if(e.target) {
      const aidText = e.target.getElementsByClassName('aid')[0]
      aidText.classList.add('view')
      window.setTimeout(() => {
        aidText.classList.remove('view')
      },400)
    }
  }

  useEffect(() => {
    state.selectTitle = props.title
  })

  return (
    <div className='search-link-box'>
      <div className="search-list">
        {LinkData.map( (item) => {
          state.urlCopy.push({id:String(item.id),value:item.title,view:false})
          return (<div className="search-item" key={item.id} onClick={(e:React.MouseEvent) => urlCopy(e,item.id,item.link)}>
              <span className="url">{item.title}</span> 
              <span className='aid'>コピーしました。</span>
            </div>)
        } )}
      </div>

      <ul className='list'>
       <li className='item'>
          <label htmlFor="select-title" className="label">調査文字列</label>
          <input type="text" className='input' name="selectTitle" id="select-title" defaultValue={props.title} onChange={changeHandler} />
        </li>
        <li className='item'>
          <label htmlFor="target-site" className="label">特定のサイトを確認する[googleのみ]</label>
          <input type="text" className='input' name="targetSite" id="target-site" onChange={changeHandler} />
        </li>
        <li className='item'>
          <label htmlFor="url-text" className="label">検索URL</label>
          <input type="text" className='input' name="urlText" id="url-text" value={state.urlText} onChange={changeHandler} />
        </li>
      </ul>
      <div className="btn-box">
        <button className='btn' onClick={viewPage}>別タブで検索する</button>
      </div>
    </div>
  )
}
