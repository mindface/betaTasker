import React from 'react'
import Link from 'next/link'

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
            <li className="nav__item">
                <Link href="/">
                  home
                </Link>
              </li>
              <li className="nav__item">
                <Link href="/photo">
                  photo
                </Link>
              </li>
              <li className="nav__item">
                <Link href="/about">
                  about
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      </header>
    )
  }
}

export default BaseHeader
