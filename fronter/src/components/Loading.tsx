import React from 'react'

interface ComponentState{}

class Loading extends React.Component<ComponentState> {
  constructor(props:ComponentState){
    super(props)
  }

  render(): JSX.Element {
    return (
      <div className="loading">
        loading
      </div>
    )
  }
}

export default Loading
