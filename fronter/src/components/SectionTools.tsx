"use client"
import SectionMemory from '../components/SectionMemory'

export default function SectionTools() {

  return(
    <div className="section__inner section--tools">
      <div className="section-continer">
        <section className="l-section">
          <div className="controller">
            <h2 className="p-b-2">画面調整</h2>データ構造
            <div className="flex">
              <label htmlFor="" className="label">
                <input type="text" className="input" />
              </label>
              <label htmlFor="" className="label">
                <input type="range" className="input" />
              </label>
            </div>
          </div>
          <div className="iframe-box position-relative max-h-m overflow-y-auto">
            <SectionMemory />
          </div>
          <div className="controller data-controller">モーション関係と情報構造
            <h2 className="p-b-2">データ操作</h2>
            <div className="flex">
              <label htmlFor="target-data" className="label">
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

