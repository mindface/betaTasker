import React from 'react'
import Link from 'next/link'
import menu from '../json/menu.json'

interface titleType{
  title: string
}

class BaseHeader extends React.Component<titleType> {
  constructor(props:titleType){
    super(props)
    this.state = {
      title: "BaseHeader"
    }
  }

  render() {
    return (
      <header className="base-header">
        <div className="header--body _flex_s_b_">
          <h3 className="header__title">
            { this.props.title }
          </h3>
          <nav className="nav">
            <ul className="nav--list _flex_">
              {/* menu.jsonのリンクを動的に追加 */}
              {menu.map((item: { title: string; path: string }, idx: number) => (
                <li className="nav__item" key={`menu${idx}`}>
                  <Link href={item.path}>
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

export default BaseHeader
