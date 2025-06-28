import { NextResponse } from 'next/server';
import learning from './learning.json'

export type LearningStructure = {
  category: string;
  studySpan: string[];
  studyCycle: string[];
  studyCycleTime: string;
  basicInfo: string;
  applicationInfo: string;
};

export type KnowledgePattern = {
  Name: string;
  technicalHistory: string;
  expertiseKnowledge: string;
  customaryPractice: string;
  contactTimeItem: string[];
};

export type FeedbackCycle = {
  cycleName: string;
  influencingFactors: string[];
  reconfigLogic: string[];
};

const learningStructure: LearningStructure = {
  category: learning.learningStructure.category,
  studySpan: learning.learningStructure.studySpan,
  studyCycle: learning.learningStructure.studyCycle,
  studyCycleTime: learning.learningStructure.studyCycleTime,
  basicInfo: learning.learningStructure.basicInfo,
  applicationInfo: learning.learningStructure.applicationInfo
};

const knowledgePattern: KnowledgePattern = {
  Name: learning.knowledgePattern.Name,
  technicalHistory: learning.knowledgePattern.technicalHistory.join(" → "),
  expertiseKnowledge: learning.knowledgePattern.expertiseKnowledge.join("；"),
  customaryPractice: learning.knowledgePattern.customaryPractice.join("；"),
  contactTimeItem: learning.knowledgePattern.contactTimeItem
};

const feedbackCycle: FeedbackCycle = {
  cycleName: learning.feedbackCycle.cycleName,
  influencingFactors: learning.feedbackCycle.influencingFactors,
  reconfigLogic: learning.feedbackCycle.reconfigLogic
};

export async function GET() {
  return NextResponse.json({
    learningStructure,
    knowledgePattern,
    feedbackCycle
  });
}
