"use client"
import React, { useRef, useState } from 'react'
import ReactDOM from'react-dom'
import { useDispatch, useStore } from 'react-redux'
import { addPostAction } from '../modules/addPostAction'
import { Posts } from '../model/posts'
import { cardCategory } from '../helper/category-Text'

import ContentModal from "./ContentModal";

type Props = {
  title: string
  searchAction: (posts:any,text:string,type:string) => void
}

export default function ContentHeader( props: Props )  {
  const [searchText,setSearchText] = useState('')
  const formValue:Posts = { id:0,title:'',name:'',text:'',disc:'',imgPath:'' }
  const titleRef = useRef('')
  const nameRef = useRef('')
  const textRef = useRef('')
  const discRef = useRef('')
  const imgPathRef = useRef('')

  const dispatch = useDispatch()

  function setInputAction(e:React.ChangeEvent<HTMLInputElement>){
    setSearchText(e.target.value)
  }

  function setSelectAction(e:React.ChangeEvent<HTMLSelectElement>){
    props.searchAction([],e.target.value,'name')
    // dispatch(searchPostCategoryAction(e.target.value))
  }

  function sendAction(){
    props.searchAction([],searchText,'title')
    // dispatch(searchPostAction(searchText))
  }

  function AddSendAction(e:React.MouseEvent<HTMLInputElement, MouseEvent>){
    const element = e.target as HTMLInputElement
    const targetElement = element.parentNode?.previousSibling as HTMLParagraphElement
    formValue["id"] = 0
    formValue["title"] = titleRef.current
    formValue["name"] = nameRef.current
    formValue["text"] = textRef.current
    formValue["disc"] = discRef.current
    formValue["imgPath"] = imgPathRef.current

    if(
      titleRef.current !== '' &&
      nameRef.current !== '' &&
      textRef.current !== '' &&
      discRef.current !== '' &&
      imgPathRef.current !== ''
    ){
      targetElement?.classList.remove('view')
      dispatch(addPostAction(formValue))
      window.location.reload()
    }else if(targetElement !== null){
      targetElement?.classList.add('view')
    }

  }

  function AddSetAction(value:string,sendValue:string){
    switch (value) {
      case "title":
        titleRef.current = sendValue
        break;
      case "name":
        nameRef.current = sendValue
        break;
      case "text":
        textRef.current = sendValue
        break;
      case "disc":
        discRef.current = sendValue
        break;
      case "imgPath":
        imgPathRef.current = sendValue
        break;
    }
  }

  function SelectItem(){
    return (<select className='select'  onChange={e => AddSetAction("name",e.target.value)}>
      <option value='0'>選んでください</option>
      {cardCategory.map((item:{id:number,value:string}) => {
        return (<option key={item.id}>{item.value}</option>)
      })}
    </select>)
  }

  function ModelForm() {
    const formlist = [
      {id:"title",name:"タイトル",id_num:0},
      {id:"name",name:"カテゴリ",id_num:1},
      {id:"text",name:"テキスト",id_num:2},
      {id:"disc",name:"詳細",id_num:3},
      {id:"imgPath",name:"パス",id_num:4}
    ]

    return (
      <div className="fields">
        {formlist.map((item:{id:string,name:string,id_num:number},num:number) => {
        return (
          <p className="field" key={item.id}>
            <label htmlFor="" className="label">{item.name}</label>
            {formlist[num].id !== "disc" && formlist[num].id !== "name" && <input type="text" className='input' onChange={e => AddSetAction(formlist[num].id,e.target?.value)} /> }
            {item.id === 'name' && <SelectItem />}
            {formlist[num].id === "disc" && <textarea className='textarea' onChange={e => AddSetAction(formlist[num].id,e.target?.value)} /> }
          </p>)
        })}
        <p className="field checker">全て入力してください。</p>
        <div className="field"><input type="submit" className='btn' value="新規追加" onClick={(e) => { AddSendAction(e) }} /></div>
      </div>
    )
  }

  function AddAction(){
    const newElements = document.createElement('div')
    newElements.className = 'field-box'
    ReactDOM.render(<ModelForm />, newElements)
    document.body.appendChild(ContentModal(newElements))
  }

  return(
    <div className="content-header">
      <div className="content-header__search _flex_">
        <h3 className="content-header__title">{props.title}</h3>
        <div className="search-box">
          <label htmlFor="header-search" className="label">
            <input type="text" name="" id="header-search" className="input" onChange={setInputAction} />
          </label>
          <input type="submit" className="btn" value="検索" onClick={sendAction} />
        </div>
        <div className="cateogry-box">
          <label htmlFor="header-category" className="label">
            <select name="" id="header-category" className="select" onChange={setSelectAction} >
              <option className="option" value="all">全て</option>
              {cardCategory.map((item:{id:number,value:string}) => {
                return (<option className="option" value={item.value} key={item.id}>{item.value}</option>)
              })}
            </select>
          </label>
        </div>
        <div className="add-box">
          <input type="submit" className="btn" value="追加" onClick={AddAction} />
        </div>
      </div>
    </div>
  )
}
