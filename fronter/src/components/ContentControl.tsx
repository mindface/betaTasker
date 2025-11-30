"use clinet"

type Props = {
  controlAction: (caseAction:string) => void
}

export default function ContentControl( props: Props )  {

  const topAction = () => {
    props.controlAction('top')
  }
  const leftAction = () => {
    props.controlAction('left')
  }
  const rightAction = () => {
    props.controlAction('right')
  }
  const bottomAction = () => {
    props.controlAction('bottom')
  }

  return(
    <div className="viewer-control-box">
      <div className="control-btn">
        <button className="btn" onClick={topAction} >top</button>
        <button className="btn" onClick={leftAction}>left</button>
        <button className="btn" onClick={rightAction}>right</button>
        <button className="btn" onClick={bottomAction}>bottom</button>
      </div>
    </div>
  )
}
