"use client"
import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../store'
import { loadAssessments, createAssessment, updateAssessment, removeAssessment } from '../features/assessment/assessmentSlice'
import ItemAssessment from "./parts/ItemAssessment"
import AssessmentModal from "./parts/AssessmentModal"
import { AddAssessment, Assessment } from "../model/assessment";

export default function SectionAssessment() {
  const dispatch = useDispatch()
  const { assessments, loading, error } = useSelector((state: RootState) => state.assessment)
  const { isAuthenticated } = useSelector((state: RootState) => state.user)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [editingAssessment, setEditingAssessment] = useState<AddAssessment|Assessment|undefined>()

  useEffect(() => {
    dispatch(loadAssessments())
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
        </div>
        {error && (
          <div className="error-message">
            {error}
          </div>
        )}

        {loading ? (
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
        />
      </div>
    </div>
  )
}
