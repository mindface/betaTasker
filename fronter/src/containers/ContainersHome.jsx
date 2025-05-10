import React from 'react'
import SectionHome from '../components/SectionHome'

class ContainersHome extends React.Component {

  // static getInitialProps({reduxStore,req}){
  //   const isServer = !!req
  //   if(isServer) reduxStore.dispatch(SetTextAction('Re ss'))
  // return {}
  // }

  constructor(props){
    super(props)
    this.state = {
      title: "Index",
      modalcontent:[1,2,3,4]
    }
  }

  componentDidMount(){
  }

  render() {
    return (
      <>
        <div className="container">
          <main className="index-page">
            <SectionHome></SectionHome>
          </main>
        </div>
      </>
    )
  }
}

export default ContainersHome
