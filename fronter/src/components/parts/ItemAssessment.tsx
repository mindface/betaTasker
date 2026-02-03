import React from 'react';
import { Assessment } from "../../model/assessment";

interface ItemAssessmentProps {
  assessment: Assessment;
  onEdit?: (assessment: Assessment) => void;
  onDelete?: (id: number) => void;
}

const ItemAssessment = ({ assessment, onEdit, onDelete }: ItemAssessmentProps) => {
  return (
    <div className="assessment-item card-item">
      <div className="assessment-item__header car-item__header">
        <div className="assessment-item__actions card-item__actions">
          {onEdit && <button onClick={() => onEdit(assessment)} className="btn btn-edit">
            編集
          </button>}
          {onDelete && <button onClick={() => onDelete(assessment.id)} className="btn btn-delete">
            削除
          </button>}
        </div>
      </div>
      <div className="assessment-item__content card-item__content">
        <p>effectiveness_score | {assessment.effectiveness_score}</p>
        <p>effort_score | {assessment.effort_score}</p>
        <div className="assessment-item__score card-item__score">
          qualitative_feedback: {assessment.qualitative_feedback}
        </div>
      </div>
      <div className="assessment-item__footer card-item__footer">
        <span className="date">{assessment.created_at ? new Date(assessment.created_at).toLocaleDateString() : ''}</span>
      </div>
    </div>
  );
};

export default ItemAssessment;
