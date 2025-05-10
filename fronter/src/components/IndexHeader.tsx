import React from 'react'
import Link from 'next/link'

interface titleType{
  title: string
}

class IndexHeader extends React.Component<titleType> {
  links: {id:number,path:string,text:string}[]
  constructor(props:titleType){
    super(props)
    this.state = {
      title: "Index Header"
    }
    this.links = [
      {id:1,path:"/",text:"home"},
      {id:2,path:"/tools",text:"tools"},
      {id:3,path:"/betaTools",text:"beta tools"},
      {id:4,path:"/visualer",text:"visualer"},
      {id:5,path:"/tasks",text:"tasks"},
    ]
  }

  render() {
    return (
      <header className="index-header">
        <div className="header--body _flex_s_b_">
          <Link href="/tools" >
              <h1 className="header__title">
                { this.props.title }
              </h1>
          </Link>
          <nav className="g-nav">
            <ul className="list _flex_">
              {this.links.map((item) => (
                <li className="item" key={item.id}>
                  <Link href={item.path} className="link">
                    {item.text}
                  </Link>
                </li>
                ))
              }
            </ul>
          </nav>
        </div>
      </header>
    )
  }
}

export default IndexHeader
