"use client"
import { useState } from 'react'
import SectionMemory from '../components/SectionMemory'

export default function SectionTools() {
  const [dataType, setDataType] = useState("")
  const [dataAnimation, setDataAnimation] = useState(0)

  return(
    <div className="section__inner section--tools">
      <div className="section-continer ">
        <section className="l-section">
          <div className="controller controller-top p-8">
            <h2 className="bg-black p-4 d-inline-block">画面調整 (データ構造)</h2>
            <div className="flex p-4">
              <label htmlFor="data-type" className="label p-r-5">
                <input
                  type="text"
                  id="data-type"
                  className="input"
                />
              </label>
              <label htmlFor="data-animation" className="label d-inline-block">
                <input
                  type="range"
                  id="data-animation"
                  className="input"
                  min={0}
                  max={100}
                  value={dataAnimation}
                  onChange={e => setDataAnimation(Number(e.target.value))}
                />
                <span className='d-inline-block '>{dataAnimation}</span>
              </label>
            </div>
          </div>
          <div className="controller-target-box position-relative max-h-m overflow-y-auto">
            <SectionMemory />
          </div>
          <div className="controller data-controller">
            <h2 className="bg-black p-4 d-inline-block">データ操作(モーション関係と情報構造)</h2>
            <div className="flex">
              <label htmlFor="target-data" className="label p-r-5">
                <input
                  type="text" id="target-data"
                  className="input"
                  min={0}
                  max={100}
                  step={1}
                />
              </label>
              <label htmlFor="target-animation" className="label">
                <input
                  type="range"
                  id="target-animation"
                  className="input"
                  min={0}
                  max={100}
                  step={1}
                />
              </label>
            </div>
          </div>
        </section>
      </div>
    </div>
  )
}

