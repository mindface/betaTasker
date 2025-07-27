import React, { use, useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from '../store';
import { loadMemoryAidsByCode } from '../features/memoryAid/memoryAidSlice';
import CommonModal from './parts/CommonModal';

interface Props {
  code: string;
  isOpen?: boolean;
}

const MemoryAidList: React.FC<Props> = ({ code }) => {
  const dispatch = useAppDispatch();
  const { contexts, loading, error } = useAppSelector(state => state.memoryAid);
  const [isModalOpen, setIsModalOpen] = useState(false)

  useEffect(() => {
    if (code) {
      dispatch(loadMemoryAidsByCode(code))
    }
  }, [code, dispatch]);

  useEffect(() => {
    console.log("contexts changed:", contexts)
  } , [contexts])

  if (loading) return <div>Loading...</div>
  if (error) return <div style={{color:'red'}}>Error: {error}</div>

  return (
    <>
      <div className="memory-aid-list-header">
        <h2>Memory Aids</h2>
        <button onClick={() => setIsModalOpen(true)} className="btn btn-primary">Open Memory Aids</button>
      </div>
      <CommonModal title={`Memory Aids for code: ${code}`} isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
        <div className='memory-aid-list max-h-m overflow-y-auto'>
          <div className="select-box">
            <select onChange={(e) => dispatch(loadMemoryAidsByCode(e.target.value))}>
              <option value="MA-C-01">MA-C-01</option>
              <option value="MA-Q-02">MA-Q-03</option>
              <option value="PM-P-03">PM-P-03</option>
              {/* 他のコードも必要に応じて追加 */}
            </select>
          </div>
          {(contexts ?? []).length === 0 ? (
            <div>該当データなし</div>
          ) : (
            <div className="card-list p-8">
              {contexts.map(ctx => (
                <div key={ctx.id} className="card-item">
                  <div className="p-b-5 card-title"><b>対象仕事:</b> {ctx.work_target}</div>
                  <div className="p-b-5"><b>機材:</b> {ctx.machine}</div>
                  <div className="p-b-5"><b>材料:</b> {ctx.material_spec}</div>
                  <div className="p-b-5"><b>目的:</b> {ctx.goal}</div>
                  <div className="p-b-5"><b>Level:</b> {ctx.level}</div>
                  <div className="p-b-5"><b>技術的要素:</b>
                    <ul>
                      {ctx.technical_factors.map(tf => (
                        <li className="p-b-5" key={tf.id}>{tf.tool_spec} / {tf.eval_factors}</li>
                      ))}
                    </ul>
                  </div>
                  <div><b>Knowledge Transformations:(知識の変換:)</b>
                    <ul>
                      {ctx.knowledge_transformations.map(kt => (
                        <li key={kt.id}>{kt.transformation} / {kt.countermeasure}</li>
                      ))}
                    </ul>
                  </div>
                </div>
              ))}
            </div>
            )}
        </div>
      </CommonModal>
    </>
  );
};

export default MemoryAidList;
