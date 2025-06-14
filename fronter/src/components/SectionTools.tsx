"use client"
import React, { useEffect } from 'react'

export default function SectionHome() {
  
  const pageInfo = [
    {
      title: "情報をリンクさせるプロセスについて",
      text: "社内で気になる情報内容を調べるツールに",
      subTitle: "subTitle",
      path: "/image/part-01.svg"
    },
    {
      title: "言葉の不一致をなくす",
      text: "社内でドキュメントを作ったが読まれるように変化させる",
      subTitle: "subTitle",
      path: "/image/part-02.svg"
    },
    {
      title: "調査時間を効果的に行う",
      text: "テキストでの確認作業を減らす",
      subTitle: "subTitle",
      path: "/image/part-03.svg"
    }
  ]

  // useEffect(() => {
    // const scroll = new LocomotiveScroll({
    //   el: document.querySelector('[data-scroll-container]'),
    //   smooth: true
    // });
  // },[])

  return(
    <div className="section__inner section--tools">
      <div className="section-continer">
        {pageInfo.map((item) => {
          return (
            <section className="l-section" key={item.path}>
              <div className="content limit-m tools">
                <div className="content-box__title">
                  <h1 className="title">{item.title}</h1>
                </div>
                <div className="content-box__body">
                  <div className="text-box">
                    <p className="text">{item.text}</p>
                  </div>
                  <div className="image-box">
                    <img className='img' src={item.path} alt="" />
                  </div>
                </div>
              </div>
            </section>)
         })}
      </div>
    </div>
  )
}

