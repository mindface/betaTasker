import React, { use, useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from '../store';
import { loadMemoryAidsByCode } from '../features/memoryAid/memoryAidSlice';

interface Props {
  code: string;
}

const MemoryAidList: React.FC<Props> = ({ code }) => {
  const dispatch = useAppDispatch();
  const { contexts, loading, error } = useAppSelector(state => state.memoryAid);

  useEffect(() => {
    if (code) {
      dispatch(loadMemoryAidsByCode(code));
    }
  }, [code, dispatch]);

  useEffect(() => {
    console.log("contexts changed:", contexts);
  } , [contexts]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div style={{color:'red'}}>Error: {error}</div>;

  return (
    <div className='memory-aid-list'>
      <h3>Memory Aids for code: {code}</h3>
      <div className="select-box">
        <select onChange={(e) => dispatch(loadMemoryAidsByCode(e.target.value))}>
          <option value="MA-C-01">MA-C-01</option>
          <option value="MA-Q-02">MA-Q-03</option>
          <option value="PM-P-03">PM-P-03</option>
          {/* 他のコードも必要に応じて追加 */}
        </select>
      </div>
      {contexts.length === 0 ? (
        <div>該当データなし</div>
      ) : (
        contexts.map(ctx => (
          <div key={ctx.id} style={{border:'1px solid #ccc', margin:'8px', padding:'8px'}}>
            <div><b>WorkTarget:</b> {ctx.work_target}</div>
            <div><b>Machine:</b> {ctx.machine}</div>
            <div><b>Material:</b> {ctx.material_spec}</div>
            <div><b>Goal:</b> {ctx.goal}</div>
            <div><b>Level:</b> {ctx.level}</div>
            <div><b>Technical Factors:</b>
              <ul>
                {ctx.technical_factors.map(tf => (
                  <li key={tf.id}>{tf.tool_spec} / {tf.eval_factors}</li>
                ))}
              </ul>
            </div>
            <div><b>Knowledge Transformations:</b>
              <ul>
                {ctx.knowledge_transformations.map(kt => (
                  <li key={kt.id}>{kt.transformation} / {kt.countermeasure}</li>
                ))}
              </ul>
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default MemoryAidList;
