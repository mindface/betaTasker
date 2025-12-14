import React from 'react'

interface Props {
  title: string
}

function IndexFooter (props: Props)  {
  const { title } = props;

  return (
    <footer className="index-footer">
      <div className="footer--body _flex_s_b_">
        <div className="footer__box bg-white box--my-evaluation">
          <h3 className='title p-b-8'>自主評価</h3>
          <div className="improvement-effect p-b-8">
            情報の改善効果  70%
          </div>
          <div className="improvement-achievement-rate">
            組織目的達成率  80%
          </div>
        </div>
        <div className="footer__box bg-white box--other-evaluation">
          <h3 className='title p-b-8'>他社評価</h3>
          <div className="improvement-effect p-b-8">
            他社でも利用されている技術率  80%
          </div>
          <div className="improvement-achievement-rate">
            同業者の参入障壁率  80%
          </div>
        </div>
        <div className="footer__box box--task-info">
          <h3 className='title p-b-8'>タスク情報</h3>
          <div className="improvement-effect p-b-8">
            Total KPI  32.1%
          </div>
          <div className="improvement-achievement-rate p-b-8">
            Total KGI  42.1%
          </div>
        </div>
        <div className="footer__box bg-white box--task-btn">
          <span className="line"></span>
          <span className="line"></span>
        </div>
      </div>
    </footer>
  )
}

export default IndexFooter
