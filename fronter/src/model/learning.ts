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

export type LearningData = {
  learningStructure: LearningStructure;
  knowledgePattern: KnowledgePattern;
  feedbackCycle: FeedbackCycle;
};
