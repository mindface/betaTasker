"use client";
import React from "react";
import GenericItemCard from "../parts/GenericItemCard";
import { Assessment } from "../../model/assessment";
import { useItemOperations } from "../../hooks/useItemOperations";

interface AssessmentCardProps {
  assessment: Assessment;
  onRefresh?: () => void;
}

export default function AssessmentCard({
  assessment,
  onRefresh,
}: AssessmentCardProps) {
  const { deleteItem, updateItem } = useItemOperations("assessment", {
    onDeleteSuccess: onRefresh,
    onUpdateSuccess: onRefresh,
  });

  const handleUpdate = async (item: Assessment) => {
    await updateItem(item);
  };

  const renderContent = (assessment: Assessment) => (
    <div className="assessment-details">
      <div className="detail-item">
        <span className="label">タスクID:</span>
        <span className="task-id">{assessment.task_id}</span>
      </div>
      <div className="scores-grid">
        <div className="score-item">
          <span className="score-label">効果:</span>
          <span className="score-value">
            {assessment.effectiveness_score}/100
          </span>
          <div className="score-bar">
            <div
              className="score-fill effectiveness"
              style={{ width: `${assessment.effectiveness_score}%` }}
            />
          </div>
        </div>
        <div className="score-item">
          <span className="score-label">努力:</span>
          <span className="score-value">{assessment.effort_score}/100</span>
          <div className="score-bar">
            <div
              className="score-fill effort"
              style={{ width: `${assessment.effort_score}%` }}
            />
          </div>
        </div>
        <div className="score-item">
          <span className="score-label">インパクト:</span>
          <span className="score-value">{assessment.impact_score}/100</span>
          <div className="score-bar">
            <div
              className="score-fill impact"
              style={{ width: `${assessment.impact_score}%` }}
            />
          </div>
        </div>
      </div>
      {assessment.created_at && (
        <div className="detail-item">
          <span className="label">作成日:</span>
          <span className="created-date">
            {new Date(assessment.created_at).toLocaleDateString("ja-JP")}
          </span>
        </div>
      )}
    </div>
  );

  const renderEditForm = (
    assessment: Assessment,
    onChange: (assessment: Assessment) => void,
  ) => (
    <form className="assessment-edit-form">
      <div className="form-group">
        <label htmlFor="task_id">タスクID</label>
        <input
          type="number"
          id="task_id"
          value={assessment.task_id}
          onChange={(e) =>
            onChange({ ...assessment, task_id: Number(e.target.value) })
          }
        />
      </div>
      <div className="form-group">
        <label htmlFor="effectiveness_score">効果スコア (0-100)</label>
        <input
          type="number"
          id="effectiveness_score"
          value={assessment.effectiveness_score}
          onChange={(e) =>
            onChange({
              ...assessment,
              effectiveness_score: Number(e.target.value),
            })
          }
          min={0}
          max={100}
        />
      </div>
      <div className="form-group">
        <label htmlFor="effort_score">努力スコア (0-100)</label>
        <input
          type="number"
          id="effort_score"
          value={assessment.effort_score}
          onChange={(e) =>
            onChange({ ...assessment, effort_score: Number(e.target.value) })
          }
          min={0}
          max={100}
        />
      </div>
      <div className="form-group">
        <label htmlFor="impact_score">インパクトスコア (0-100)</label>
        <input
          type="number"
          id="impact_score"
          value={assessment.impact_score}
          onChange={(e) =>
            onChange({ ...assessment, impact_score: Number(e.target.value) })
          }
          min={0}
          max={100}
        />
      </div>
      <div className="form-group">
        <label htmlFor="qualitative_feedback">定性的フィードバック</label>
        <textarea
          id="qualitative_feedback"
          value={assessment.qualitative_feedback || ""}
          onChange={(e) =>
            onChange({ ...assessment, qualitative_feedback: e.target.value })
          }
          rows={3}
        />
      </div>
    </form>
  );

  const getTitle = (assessment: Assessment) => {
    return `評価 #${assessment.id} (タスク: ${assessment.task_id})`;
  };

  const getDescription = (assessment: Assessment) => {
    return (
      assessment.qualitative_feedback ||
      `効果: ${assessment.effectiveness_score}, 努力: ${assessment.effort_score}, インパクト: ${assessment.impact_score}`
    );
  };

  return (
    <GenericItemCard
      item={assessment}
      itemType="assessment"
      onDelete={deleteItem}
      onUpdate={handleUpdate}
      renderContent={renderContent}
      renderEditForm={renderEditForm}
      getTitle={getTitle}
      getDescription={getDescription}
      getId={(assessment) => assessment.id}
    />
  );
}
