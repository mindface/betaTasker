import type {
  OntologySearchPolicy,
  OntologySearchResponse,
  OntologySearchScoreInput,
} from "@/model/rag";

interface OntologySearchResultsProps {
  response: OntologySearchResponse;
  title?: string;
}

interface ScoreMetric {
  key: keyof OntologySearchScoreInput;
  weightKey: keyof OntologySearchPolicy;
  label: string;
  standard: string;
}

const scoreMetrics: ScoreMetric[] = [
  {
    key: "text_similarity",
    weightKey: "alpha",
    label: "テキスト類似度",
    standard: "ベクトル検索などで算出された意味的な近さ",
  },
  {
    key: "ontology_match",
    weightKey: "beta",
    label: "オントロジー一致度",
    standard: "直接一致 1.00、親概念 0.65、関連概念 0.45 の重みで正規化",
  },
  {
    key: "risk_priority",
    weightKey: "gamma",
    label: "リスク優先度",
    standard: "risk_level、urgency、criticality の最大値を 5.0 で割って正規化",
  },
  {
    key: "graph_centrality",
    weightKey: "delta",
    label: "グラフ中心性",
    standard: "概念の接続数を最大接続数で割り、文書タグまたは一致概念の平均で算出",
  },
];

const formatDecimal = (value: number, digits = 4) =>
  Number.isFinite(value) ? value.toFixed(digits) : "0.0000";

const formatPercent = (value: number) =>
  `${Number.isFinite(value) ? Math.round(value * 100) : 0}%`;

const getNormalizedPolicy = (
  policy: OntologySearchPolicy,
): OntologySearchPolicy => {
  const total = policy.alpha + policy.beta + policy.gamma + policy.delta;

  if (total <= 0) {
    return {
      alpha: 0,
      beta: 0,
      gamma: 0,
      delta: 0,
    };
  }

  return {
    alpha: policy.alpha / total,
    beta: policy.beta / total,
    gamma: policy.gamma / total,
    delta: policy.delta / total,
  };
};

export default function OntologySearchResults({
  response,
  title = "オントロジー検索結果",
}: OntologySearchResultsProps) {
  const normalizedPolicy = getNormalizedPolicy(response.policy);

  if (!response.success) {
    return (
      <section className="rag-search">
        <div className="rag-search__error">検索結果を表示できません。</div>
      </section>
    );
  }

  return (
    <section className="rag-search">
      <div className="rag-search__header">
        <div>
          <h2>{title}</h2>
          {response.message && <p>{response.message}</p>}
        </div>
        <div className="rag-search__count">
          {response.results.length}
          <span>件</span>
        </div>
      </div>

      <div className="rag-search__policy">
        <div>
          <h3>スコア標準</h3>
          <p>
            score = alpha * text_similarity + beta * ontology_match + gamma *
            risk_priority + delta * graph_centrality
          </p>
        </div>
        <div className="rag-search__weight-grid">
          {scoreMetrics.map((metric) => {
            const weight = normalizedPolicy[metric.weightKey];
            return (
              <div className="rag-search__weight-item" key={metric.key}>
                <span>{metric.label}</span>
                <strong>{formatPercent(weight)}</strong>
                <small>{metric.standard}</small>
              </div>
            );
          })}
        </div>
      </div>

      <div className="rag-search__results">
        {response.results.map((result, index) => (
          <article className="rag-search__card" key={result.document_id}>
            <div className="rag-search__card-header">
              <div>
                <span className="rag-search__rank">#{index + 1}</span>
                <h3>{result.document_id}</h3>
              </div>
              <div className="rag-search__score">
                <span>score</span>
                <strong>{formatDecimal(result.score, 6)}</strong>
              </div>
            </div>

            <div className="rag-search__metric-list">
              {scoreMetrics.map((metric) => {
                const inputValue = result.input[metric.key] as number;
                const weight = normalizedPolicy[metric.weightKey];
                const contribution = inputValue * weight;

                return (
                  <div className="rag-search__metric-row" key={metric.key}>
                    <div className="rag-search__metric-text">
                      <span>{metric.label}</span>
                      <small>{metric.standard}</small>
                    </div>
                    <div className="rag-search__metric-values">
                      <span>値 {formatDecimal(inputValue)}</span>
                      <span>重み {formatPercent(weight)}</span>
                      <strong>寄与 {formatDecimal(contribution)}</strong>
                    </div>
                    <div className="rag-search__bar-track">
                      <div
                        className="rag-search__bar-fill"
                        style={{ width: formatPercent(inputValue) }}
                      />
                    </div>
                  </div>
                );
              })}
            </div>
          </article>
        ))}
      </div>
    </section>
  );
}
