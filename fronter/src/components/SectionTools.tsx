"use client"
import { useState } from 'react'
import SectionMemory2 from '../components/SectionMemory2'
import SectionTask2 from '../components/SectionTask2'

export default function SectionTools() {
  const [dataViewType, setDataViewType] = useState("")
  const [dataAnimation, setDataAnimation] = useState(0)

  return(
    <div className="section__inner section--tools">
      <div className="section-continer">
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
              <div className="p-l-16">
                <label
                  htmlFor="data-view-type-1"
                  className="label d-inline-block p-l-8"
                >
                  <input
                    id="data-view-type-1"
                    type="radio"
                    name="data-view-type"
                    value="type1"
                    checked={dataViewType === "type1"}
                    onChange={e => setDataViewType(e.target.value)}
                  />
                  タイプ1
                </label>
                <label
                  htmlFor="data-view-type-2"
                  className="label d-inline-block p-l-8"
                >
                  <input
                    id="data-view-type-2"
                    type="radio"
                    name="data-view-type"
                    value="type2"
                    checked={dataViewType === "type2"}
                    onChange={e => setDataViewType(e.target.value)}
                  />
                  タイプ2
                </label>
                <label
                  htmlFor="data-view-type-3"
                  className="label d-inline-block p-l-8"
                >
                  <input
                    id="data-view-type-3"
                    type="radio"
                    name="data-view-type"
                    value="type3"
                    checked={dataViewType === "type3"}
                    onChange={e => setDataViewType(e.target.value)}
                  />
                  タイプ3
                </label>
                <span className="d-inline-block">{dataViewType}</span>
              </div>
            </div>
          </div>
          <div data-view-type={dataViewType} className="controller-target-box position-relative max-h-m overflow-y-auto">
            <SectionMemory2 />
            <SectionTask2 />
          </div>
          <div className="controller data-controller p-8">
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

