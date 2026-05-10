export interface OntologySearchPolicy {
  alpha: number;
  beta: number;
  gamma: number;
  delta: number;
}

export interface OntologyConcept {
  id: string;
  label: string;
  definition: string;
  parent_ids: string[];
  related_ids: string[];
  metadata: Record<string, number | string | boolean>;
  domain: string;
}

export interface OntologyDocument {
  id: string;
  text: string;
  ontology_tags: string[];
  metadata: Record<string, number | string | boolean>;
}

export interface OntologySearchRequest {
  policy?: OntologySearchPolicy;
  query_concept_ids: string[];
  concepts: OntologyConcept[];
  documents: OntologyDocument[];
  text_similarities: Record<string, number>;
}

export interface OntologySearchScoreInput {
  document_id: string;
  text_similarity: number;
  ontology_match: number;
  risk_priority: number;
  graph_centrality: number;
}

export interface OntologySearchResult {
  document_id: string;
  score: number;
  input: OntologySearchScoreInput;
}

export interface OntologySearchResponse {
  message?: string;
  policy: OntologySearchPolicy;
  results: OntologySearchResult[];
  success: boolean;
}
