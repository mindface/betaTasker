import React, { useState, useEffect } from 'react';
import { AddAssessment, Assessment } from "../../model/assessment";
import Cookies from 'js-cookie';
import CommonModal from "./CommonModal";
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../../store'
import { getAssessmentsForTaskUser } from "../../features/assessment/assessmentSlice";

interface AssessmentListModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (assessmentData: AddAssessment | Assessment) => void;
  taskId?: number;
}

const setCheker = ['user_id', 'task_id', 'effectiveness_score', 'effort_score', 'impact_score'];
const AssessmentListModal: React.FC<AssessmentListModalProps> = ({ isOpen, onClose, onSave, taskId }) => {
  const dispatch = useDispatch();
  const { assessments, assessmentLoading, assessmentError } = useSelector((state: RootState) => state.assessment)
  const { user } = useSelector((state: RootState) => state.user)
  const [formData, setFormData] = useState<AddAssessment | Assessment | undefined>();
  const [openMemoryId, setOpenMemoryId] = useState<number | null>(null);

  useEffect(() => {
    // if (assessmentId) {
    // } else {
    //   setFormData({
    //     task_id: 0,
    //     user_id: 0,
    //     effectiveness_score: 0,
    //     effort_score: 0,
    //     impact_score: 0,
    //     qualitative_feedback: '',
    //   });
    // }

    if(taskId !== -1) {
      dispatch(getAssessmentsForTaskUser({ userId: 1, taskId: taskId || 0 }))
    }
  }, [dispatch, user, taskId]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData(prev => prev ? {
      ...prev,
      [name]: setCheker.includes(name) ? Number(value) : value
    } : prev);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData) return;
    // cookieからuser_id取得
    const userIdStr = Cookies.get('user_id');
    const user_id = userIdStr ? Number(userIdStr) : 0;
    await onSave({ ...formData, user_id });
  };

  if (!isOpen) return null;

  return (
    <CommonModal isOpen={isOpen} onClose={onClose} title="アセスメントリスト">
      <div className="Assessment-list__box">
        {/* 全メモリー一覧を表示（タイトルクリックで詳細トグル） */}
        {assessmentLoading ? (
          <div className="loading">アセスメント取得中...</div>
        ) : assessmentError ? (
          <div className="error-message">{assessmentError}</div>
        ) : assessments && assessments.length > 0 ? (
          <ul className="assessment-list card-list">
            {assessments.map((assessment) => (
              <li key={assessment.id} className="assessment-item card-item">
                <div className="assessment-title card-title" onClick={() => setOpenMemoryId(assessment.id)}>
                  {assessment.qualitative_feedback || 'アセスメント'}
                </div>
                {openMemoryId === assessment.id && (
                  <div className="assessment-details card-details">
                    <p><b>効果スコア:</b> {assessment.effectiveness_score}</p>
                    <p><b>努力スコア:</b> {assessment.effort_score}</p>
                    <p><b>影響スコア:</b> {assessment.impact_score}</p>
                    <p><b>フィードバック:</b> {assessment.qualitative_feedback}</p>
                  </div>
                )}
              </li>
            ))}
          </ul>
        ) : (
          <div className="no-assessments">アセスメントがありません。</div>
        )}
        {/* 紐づくメモリー情報を表示 */}
        {/* {relatedMemory && (
          <div className="related-memory-info" style={{margin: '1em 0', padding: '0.5em', background: '#f6f8fa', borderRadius: 6}}>
            <div><b>関連メモ:</b> {relatedMemory.title}</div>
            <div><b>内容:</b> {relatedMemory.notes}</div>
            <div><b>タグ:</b> {relatedMemory.tags}</div>
            <div><b>ステータス:</b> {relatedMemory.read_status}</div>
          </div>
        )} */}
        <div className="form-actions">
          <button type="button" onClick={onClose} className="btn">閉じる</button>
        </div>
      </div>
    </CommonModal>
  );
};

export default AssessmentListModal;
