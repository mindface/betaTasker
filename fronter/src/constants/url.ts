
const domainAndHost = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

const URLs = {
  login: `${domainAndHost}/api/login`,
  register: `${domainAndHost}/api/register`,
  assessment: `${domainAndHost}/api/assessment`,
  assessmentsForTaskUser: `${domainAndHost}/api/assessmentsForTaskUser`,
  heuristicsAnalyze: `${domainAndHost}/api/heuristics/analyze`,
  heuristicsInsights: `${domainAndHost}/api/heuristics/insights`,
  heuristicsPatterns: `${domainAndHost}/api/heuristics/patterns`,
  heuristicsPattern: `${domainAndHost}/api/heuristics/pattern`,
  heuristicsTrack: `${domainAndHost}/api/heuristics/track`,
  heuristicsInsight: `${domainAndHost}/api/heuristics/insight`,
  memory: `${domainAndHost}/api/memory`,
  user: `${domainAndHost}/api/user`,
  task: `${domainAndHost}/api/task`,
  project: `${domainAndHost}/api/project`,
  test: `${domainAndHost}/api/test`,
  heuristics_track: `${domainAndHost}/api/heuristics/track`,
  learning: `${domainAndHost}/api/learning`,
  memoryAid: `${domainAndHost}/api/memory_aid`,
  knowledgePattern: `${domainAndHost}/api/knowledge_pattern`,
  teachingFreeControl: `${domainAndHost}/api/teaching_free_control`,
  qualitativeLabel: `${domainAndHost}/api/qualitative_label`,
  processOptimization: `${domainAndHost}/api/process_optimization`,
  languageOptimization: `${domainAndHost}/api/language_optimization`,
}

export { URLs };