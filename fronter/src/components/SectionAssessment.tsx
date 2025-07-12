"use client"
import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../store'
import { loadAssessments, createAssessment, updateAssessment, removeAssessment } from '../features/assessment/assessmentSlice'
import { loadLearningData } from '../features/learningData/learningDataSlice'
import ItemAssessment from "./parts/ItemAssessment"
import AssessmentModal from "./parts/AssessmentModal"
import MemoryAidList from "./MemoryAidList"
import { AddAssessment, Assessment } from "../model/assessment"
import { loadTasks } from '../features/task/taskSlice'

export default function SectionAssessment() {
  const dispatch = useDispatch()
  const { tasks, taskLoading, taskError } = useSelector((state: RootState) => state.task)
  const { assessments, assessmentLoading, assessmentError } = useSelector((state: RootState) => state.assessment)
  const { memories, memoryLoading, memoryError } = useSelector((state: RootState) => state.memory)
  const { learningData, learningLoading, learningError } = useSelector((state: RootState) => state.learning);
  const { isAuthenticated } = useSelector((state: RootState) => state.user)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [editingAssessment, setEditingAssessment] = useState<AddAssessment|Assessment|undefined>()

  useEffect(() => {
    dispatch(loadAssessments())
    dispatch(loadLearningData())
  }, [dispatch, isAuthenticated])

  const handleAddAssessment = () => {
    setEditingAssessment(undefined)
    setIsModalOpen(true)
  }

  const handleEditAssessment = (assessment: Assessment) => {
    setEditingAssessment(assessment)
    setIsModalOpen(true)
  }

  const handleSaveAssessment = async (assessmentData: AddAssessment | Assessment) => {
    if (editingAssessment) {
      await dispatch(updateAssessment(assessmentData as Assessment))
    } else {
      await dispatch(createAssessment(assessmentData as AddAssessment))
    }
    setIsModalOpen(false)
  }

  const handleDeleteAssessment = async (id: number) => {
    await dispatch(removeAssessment(id))
  }

  useEffect(() => {
    dispatch(loadTasks())
    // dispatch(loadMemories())
  }, [dispatch, isAuthenticated]);

  useEffect(() => {
    if(learningData) {
      console.log("learningData.learningStructure")
      console.log(learningData.learningStructure)
    }
  },[learningData])

  return (
    <div className="section__inner section--assessment">
      <div className="section-container">
        <div className="assessment-header">
          <h2>アセスメント</h2>
          <button
            onClick={() => handleAddAssessment()}
            className="btn btn-primary"
          >
            新規アセスメント
          </button>
          <div className="assessment-header">
            <h2>アセスメント</h2>
            {learningLoading ? (
              <div className="loading">学習構造データ取得中...</div>
            ) : learningError ? (
              <div className="error-message">{learningError}</div>
            ) : learningData ? (
              <div className="learning-structure-info">
                <span className="d-block">学習カテゴリ: {learningData.learningStructure.category}</span>
                <span className="d-block">学習サイクル: {learningData.learningStructure.studyCycle.map((item,index) => <span key={`studyCycle${index}`} className="d-iline-block p-1">{item}</span>)}</span>
              </div>
            ) : null}

            <button
              onClick={() => handleAddAssessment()}
              className="btn btn-primary"
            >
              新規アセスメント
            </button>
          </div>
        </div>
        {assessmentError && (
          <div className="error-message">
            {assessmentError}
          </div>
        )}
        {assessmentLoading ? (
          <div className="loading">読み込み中...</div>
        ) : (
          <div className="assessment-list">
            {assessments.map((assessment: Assessment, index: number) => (
              <ItemAssessment
                key={`assessment-item${index}`}
                assessment={assessment}
                onEdit={(editAssessment: Assessment) => handleEditAssessment(editAssessment)}
                onDelete={() => handleDeleteAssessment(assessment.id)}
              />
            ))}
          </div>
        )}
        <AssessmentModal
          initialData={editingAssessment}
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          onSave={handleSaveAssessment}
          tasks={tasks}
          memories={memories}
        />
      </div>
    </div>
  )
}
