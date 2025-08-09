"use client"
import React, { useRef } from 'react'
import ReactDOM from'react-dom'
import { Posts } from '../../model/posts'
import { useDispatch } from 'react-redux'
import Image from 'next/image'
import ItemCardSearchLink from "./ItemCardSearchLink"
import ContentViewItemModal from "./ContentViewItemModal"
import { deletePostAction } from '../../modules/deletePostAction'
import { updatePostAction } from '../../modules/updatePostAction'

import { categoryText } from '../../helper/category-Text'

type Props = {
  posts: Posts;
  key: number;
}

export default function ItemCard( props: Props )  {
  const cardInfo = props.posts
  const dispatch = useDispatch()
  const ItemsRef = () => {
    const list = [];
    for (const [key,value] of Object.entries(cardInfo)) {
      list.push({id:key,value:value})
      AddSetAction(key,value)
    }
    return list;
  }

  const titleRef = useRef("")
  const nameRef = useRef("")
  const textRef = useRef("")
  const discRef = useRef("")
  // const [disc,setDisc] = useState('')
  const imgPathRef = useRef("")
  const formValue:Posts = {id:0,title:'',name:'',text:'',disc:'',imgPath:''}

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

  function FormList() {
    return (<div className='form-box'>
      <ul className='list'>
        {ItemsRef().length && ItemsRef().map((item:{id:string,value:string}) => {
          return (
            <li className='item' key={item.id}>
              <p className='caption'><span className='aid'>{categoryText(item.id)}</span>{item.id === "id" && <span className='id'>{item.value}</span>}</p>
                {item.id === 'title' && <input type='text' className='input' defaultValue={props.posts.title} onChange={(e:React.ChangeEvent<HTMLInputElement>) => AddSetAction(item.id,e.target?.value)} />}
                {item.id === 'name' && <input type='text' className='input' defaultValue={props.posts.name} onChange={(e:React.ChangeEvent<HTMLInputElement>) => AddSetAction(item.id,e.target?.value)} />}
                {item.id === 'text' && <input type='text' className='input' defaultValue={props.posts.text} onChange={(e:React.ChangeEvent<HTMLInputElement>) => AddSetAction(item.id,e.target?.value)} />}
                {item.id === 'disc' && <textarea className='input textarea' defaultValue={props.posts.disc} onChange={(e:React.ChangeEvent<HTMLTextAreaElement>) => AddSetAction(item.id,e.target?.value) } />}
                {item.id === 'imgPath' && <input type='text' className='input' defaultValue={props.posts.imgPath} onChange={(e:React.ChangeEvent<HTMLInputElement>) => AddSetAction(item.id,e.target?.value)} />}
            </li>)
        })}
        <li className='item'><button className="btn" onClick={updateAcion}>更新</button></li>
      </ul>
    </div>)
  }

  function modalViewAcion(cardInfo:Posts):void{
    const viewBox = document.createElement('div')
    viewBox.className = 'view-box'
    document.body.appendChild(ContentViewItemModal(viewBox))
    ReactDOM.render(<FormList />, viewBox)
  }

  function modalSearchAcion(title:string | undefined):void{
    const searchBox = document.createElement('div')
    searchBox.className = 'search-box'
    document.body.appendChild(ContentViewItemModal(searchBox))
    ReactDOM.render(<ItemCardSearchLink title={title ? title : 'none'} />, searchBox)
  }

  function updateAcion(){
    formValue["id"] = props.posts.id
    formValue["title"] = titleRef.current
    formValue["name"] = nameRef.current
    formValue["text"] = textRef.current
    formValue["disc"] = discRef.current
    formValue["imgPath"] = imgPathRef.current

    dispatch(updatePostAction(formValue))
    window.location.reload()
  }

  function deleteAcion(cardInfo:Posts){
    if(confirm("このタスクは取り消せません。削除しますか。")){
      dispatch(deletePostAction(String(cardInfo.id)))
    }
  }

  return(
    <div className="card">
      <h4 className="card__title">{cardInfo.title}</h4>
      <div className="card__body">
        <div className="category"><span className="aid">{cardInfo.name}</span></div>
        <div className="sentence"><p className='aid'>詳細</p><p className="sentence-text">{cardInfo.disc}</p></div>
      </div>
      <div className="card__btns">
        <div className="btns">
          <button className="btn" onClick={e => modalViewAcion(cardInfo)}>
            <Image
              src="/image/look.svg"
              alt="表示"
              width={20}
              height={20}
            />
          </button>
          <button className="btn" onClick={e => modalSearchAcion(cardInfo.title)}>
            <Image
              src="/image/link.svg"
              alt="編集"
              width={20}
              height={20}
            />
          </button>
          <button className="btn" onClick={e =>  deleteAcion(cardInfo)}>
            <Image
              src="/image/delete.svg"
              alt="削除"
              width={20}
              height={20}
            />
          </button>
        </div>
      </div>
    </div>
  )
}
