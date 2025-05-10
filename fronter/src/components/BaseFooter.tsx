import React from 'react'

interface titleType{
   title: string
}

class BaseFooter extends React.Component<titleType> {
  constructor(props:titleType){
    super(props)
    this.state = {
      title: "&copy; Footer"
    }
  }

  render(): JSX.Element {
    return (
      <footer className="base-footer">
        <div className="footer--body">
          <h3 className="footer__title">
            { this.props.title }
          </h3>
        </div>
      </footer>
    )
  }
}

export default BaseFooter
