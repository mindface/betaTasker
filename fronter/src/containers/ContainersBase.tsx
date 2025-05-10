import React from 'react'

type Props = {
  children?: React.ReactNode
}

function ContainersHome (props:Props) {

  return (
    <div className="container l-container">
      <main className="base-l">
        { props.children }
      </main>
    </div>
  )
}

export default ContainersHome
