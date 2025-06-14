import React from 'react'
import Link from 'next/link'
import menu from '../json/menu.json'

interface titleType{
  title: string
}

class IndexHeader extends React.Component<titleType> {
  constructor(props:titleType){
    super(props)
    this.state = {
      title: "Index Header"
    }
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
              {menu.map((item: { title: string; path: string }, idx: number) => (
                <li className="item" key={`menu${idx}`}>
                  <Link href={item.path} className="link">
                    {item.title}
                  </Link>
                </li>
              ))}
            </ul>
          </nav>
        </div>
      </header>
    )
  }
}

export default IndexHeader
