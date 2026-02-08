"use client";
import React, { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store";
import {
  loadAssessments,
  getAssessmentsLimit,
  createAssessment,
  updateAssessment,
  removeAssessment,
} from "../features/assessment/assessmentSlice";
import { loadKnowledgePatterns } from "../features/knowledge_pattern/knowledgePatternSlice";
import { loadLearningData } from "../features/learning_data/learningDataSlice";
import ItemAssessment from "./parts/ItemAssessment";
import AssessmentModal from "./parts/AssessmentModal";
import PageNation from "./parts/PageNation";
import { AddAssessment, Assessment } from "../model/assessment";
import { loadTasks } from "../features/task/taskSlice";
import { HeuristicsDashboard } from "./heuristics";

export default function SectionAssessment() {
  const dispatch = useDispatch();
  const { tasks, taskLoading, taskError } = useSelector(
    (state: RootState) => state.task,
  );
  const { assessments, assessmentLoading, assessmentError, assessmentsPage, assessmentsLimit, assessmentsTotal, assessmentsTotalPages } = useSelector(
    (state: RootState) => state.assessment,
  );
  // TODO　APIの調整後に再実装を考慮
  const {
    knowledgePatterns,
    knowledgePatternsError,
    knowledgePatternsLoading,
  } = useSelector((state: RootState) => state.knowledgePattern);
  const { memories, memoryLoading, memoryError } = useSelector(
    (state: RootState) => state.memory,
  );
  const { learningData, learningLoading, learningError } = useSelector(
    (state: RootState) => state.learning,
  );
  const { isAuthenticated } = useSelector((state: RootState) => state.user);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingAssessment, setEditingAssessment] = useState<
    AddAssessment | Assessment | undefined
  >();
  const [showHeuristics, setShowHeuristics] = useState(false);

  useEffect(() => {
    // dispatch(loadAssessments());
    dispatch(getAssessmentsLimit({ page: 1, limit: 20 }));
    dispatch(loadLearningData());
    dispatch(loadKnowledgePatterns());
  }, [dispatch, isAuthenticated]);

  const handleAddAssessment = () => {
    setEditingAssessment(undefined);
    setIsModalOpen(true);
  };

  const handleEditAssessment = (assessment: Assessment) => {
    setEditingAssessment(assessment);
    setIsModalOpen(true);
  };

  const handleSaveAssessment = async (
    assessmentData: AddAssessment | Assessment,
  ) => {
    if (editingAssessment) {
      await dispatch(updateAssessment(assessmentData as Assessment));
    } else {
      await dispatch(createAssessment(assessmentData as AddAssessment));
    }
    setIsModalOpen(false);
  };

  const handleDeleteAssessment = async (id: number) => {
    await dispatch(removeAssessment(id));
  };

  const handlePageChange = (newPage: number) => {
    dispatch(getAssessmentsLimit({ page: newPage, limit: 20 }));
  };

  useEffect(() => {
    dispatch(loadTasks());
    // dispatch(loadMemories())
  }, [dispatch, isAuthenticated]);

  useEffect(() => {
    if (learningData) {
      console.log("learningData.learningStructure");
      console.log(learningData.learningStructure);
    }
  }, [learningData]);

  return (
    <div className="section__inner section--assessment">
      <div className="section-container">
        <div className="assessment-header">
          <h2>アセスメントセクション</h2>
          <div className="assessment-header p-8">
            {learningLoading ? (
              <div className="loading">学習構造データ取得中...</div>
            ) : learningError ? (
              <div className="error-message">{learningError}</div>
            ) : learningData ? (
              <div className="learning-structure-info">
                <span className="d-block">
                  学習カテゴリ: {learningData.learningStructure.category}
                </span>
                <span className="d-block p-4">
                  学習サイクル:{" "}
                  {learningData.learningStructure.studyCycle.map(
                    (item, index) => (
                      <span
                        key={`studyCycle${index}`}
                        className="d-iline-block p-4"
                      >
                        {item}
                      </span>
                    ),
                  )}
                </span>
              </div>
            ) : null}
            <div style={{ display: "flex", gap: "10px" }}>
              <button
                onClick={() => handleAddAssessment()}
                className="btn btn-primary"
              >
                新規アセスメント
              </button>
              <button
                onClick={() => setShowHeuristics(!showHeuristics)}
                className="btn btn-secondary"
                style={{
                  backgroundColor: showHeuristics ? "#667eea" : "#6c757d",
                  color: "white",
                  border: "none",
                  padding: "8px 16px",
                  borderRadius: "4px",
                  cursor: "pointer",
                }}
              >
                {showHeuristics ? "アセスメント表示" : "ヒューリスティック分析"}
              </button>
            </div>
          </div>
        </div>
        {showHeuristics ? (
          <HeuristicsDashboard />
        ) : (
          <>
            {assessmentError && (
              <div className="error-message">{assessmentError}</div>
            )}
            {assessmentLoading ? (
              <div className="loading">読み込み中...</div>
            ) : (
              <>
              <div className="assessment-list card-list">
                {assessments.map((assessment: Assessment, index: number) => (
                  <ItemAssessment
                    key={`assessment-item${index}`}
                    assessment={assessment}
                    onEdit={(editAssessment: Assessment) =>
                      handleEditAssessment(editAssessment)
                    }
                    onDelete={() => handleDeleteAssessment(assessment.id)}
                  />
                ))}
              </div>
              <PageNation
                page={assessmentsPage}
                limit={assessmentsLimit}
                totalPages={assessmentsTotalPages}
                onChange={(newPage: number) => {
                  handlePageChange(newPage);
                }}
              />
              </>
            )}
            <AssessmentModal
              initialData={editingAssessment}
              isOpen={isModalOpen}
              onClose={() => setIsModalOpen(false)}
              onSave={handleSaveAssessment}
              tasks={tasks}
              memories={memories}
            />
          </>
        )}
      </div>
    </div>
  );
}
