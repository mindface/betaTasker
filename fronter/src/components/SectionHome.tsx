"use client";
import React, { useState } from "react";
import Image from "next/image";

import { OntologySearchForm, OntologySearchResults } from "@/components/rag";
import { rankOntologySearchClient } from "@/client/ragApi";
import type {
  OntologySearchRequest,
  OntologySearchResponse,
} from "@/model/rag";

const ontologySearchRequest: OntologySearchRequest = {
  policy: {
    alpha: 0.25,
    beta: 0.4,
    gamma: 0.25,
    delta: 0.1,
  },
  query_concept_ids: ["IND:RISK:SAFETY"],
  concepts: [
    {
      id: "IND:RISK:SAFETY",
      label: "安全リスク",
      definition: "作業安全や設備停止に関わるリスク",
      parent_ids: ["IND:RISK"],
      related_ids: ["IND:ALERT:ANOMALY"],
      metadata: {
        risk_level: 5,
        urgency: 5,
      },
      domain: "industrial",
    },
    {
      id: "IND:RISK",
      label: "産業リスク",
      definition: "産業プロセス全体のリスク分類",
      parent_ids: [],
      related_ids: ["IND:PROCESS"],
      metadata: {
        risk_level: 4,
      },
      domain: "industrial",
    },
    {
      id: "IND:ALERT:ANOMALY",
      label: "異常検知",
      definition: "設備や作業状態の異常を検知する概念",
      parent_ids: ["IND:PROCESS"],
      related_ids: ["IND:RISK:SAFETY"],
      metadata: {
        criticality: 5,
      },
      domain: "industrial",
    },
    {
      id: "IND:PROCESS",
      label: "工程",
      definition: "製造や作業の工程",
      parent_ids: [],
      related_ids: [],
      metadata: {
        risk_level: 2,
      },
      domain: "industrial",
    },
  ],
  documents: [
    {
      id: "DOC:IND:002",
      text: "安全停止と異常検知に関する手順書",
      ontology_tags: ["IND:RISK:SAFETY", "IND:RISK"],
      metadata: {
        risk_level: 5,
      },
    },
    {
      id: "DOC:IND:999",
      text: "設備ログの類似事例と工程監視の記録",
      ontology_tags: ["IND:PROCESS", "IND:ALERT:ANOMALY"],
      metadata: {
        risk_level: 5,
      },
    },
  ],
  text_similarities: {
    "DOC:IND:002": 0.65,
    "DOC:IND:999": 0.95,
  },
};

export default function SectionHome() {
  const [ontologySearchResponse, setOntologySearchResponse] =
    useState<OntologySearchResponse | null>(null);
  const [ontologySearchError, setOntologySearchError] = useState("");
  const [ontologySearchLoading, setOntologySearchLoading] = useState(false);

  const handleRankOntologySearch = async (request: OntologySearchRequest) => {
    setOntologySearchLoading(true);
    setOntologySearchError("");
    setOntologySearchResponse(null);

    const response = await rankOntologySearchClient(request);

    if ("error" in response) {
      setOntologySearchError(response.error.message);
    } else {
      setOntologySearchResponse(response);
    }

    setOntologySearchLoading(false);
  };

  return (
    <div className="home _flex_c_">
      <Image src="/image/informationNetwork.svg" alt="" width={400} height={400} />
      <OntologySearchForm
        defaultValue={ontologySearchRequest}
        loading={ontologySearchLoading}
        onSubmit={handleRankOntologySearch}
      />
      {ontologySearchLoading && <p>RAG検索結果を取得中...</p>}
      {ontologySearchError && <p>{ontologySearchError}</p>}
      {ontologySearchResponse && (
        <OntologySearchResults
          response={ontologySearchResponse}
          title="RAGデモ"
        />
      )}
    </div>
  );
}
