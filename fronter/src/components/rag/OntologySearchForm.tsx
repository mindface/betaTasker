import type { OntologySearchRequest } from "@/model/rag";

interface OntologySearchFormProps {
  defaultValue: OntologySearchRequest;
  loading?: boolean;
  onSubmit: (request: OntologySearchRequest) => void;
}

const toConceptIDs = (value: string) =>
  value
    .split(",")
    .map((item) => item.trim())
    .filter(Boolean);

export default function OntologySearchForm({
  defaultValue,
  loading = false,
  onSubmit,
}: OntologySearchFormProps) {
  const concept = defaultValue.concepts[0];
  const primaryDocument = defaultValue.documents[0];
  const secondaryDocument = defaultValue.documents[1];

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);

    const request: OntologySearchRequest = {
      policy: {
        alpha: Number(formData.get("alpha")),
        beta: Number(formData.get("beta")),
        gamma: Number(formData.get("gamma")),
        delta: Number(formData.get("delta")),
      },
      query_concept_ids: toConceptIDs(String(formData.get("query_concept_ids"))),
      concepts: [
        {
          ...concept,
          id: String(formData.get("concept_id")),
          label: String(formData.get("concept_label")),
          definition: String(formData.get("concept_definition")),
          parent_ids: toConceptIDs(String(formData.get("concept_parent_ids"))),
          related_ids: toConceptIDs(String(formData.get("concept_related_ids"))),
          metadata: {
            risk_level: Number(formData.get("concept_risk_level")),
            urgency: Number(formData.get("concept_urgency")),
          },
        },
        ...defaultValue.concepts.slice(1),
      ],
      documents: [
        {
          ...primaryDocument,
          id: String(formData.get("primary_document_id")),
          text: String(formData.get("primary_document_text")),
          ontology_tags: toConceptIDs(
            String(formData.get("primary_document_tags")),
          ),
          metadata: {
            risk_level: Number(formData.get("primary_document_risk_level")),
          },
        },
        {
          ...secondaryDocument,
          id: String(formData.get("secondary_document_id")),
          text: String(formData.get("secondary_document_text")),
          ontology_tags: toConceptIDs(
            String(formData.get("secondary_document_tags")),
          ),
          metadata: {
            risk_level: Number(formData.get("secondary_document_risk_level")),
          },
        },
      ],
      text_similarities: {
        [String(formData.get("primary_document_id"))]: Number(
          formData.get("primary_text_similarity"),
        ),
        [String(formData.get("secondary_document_id"))]: Number(
          formData.get("secondary_text_similarity"),
        ),
      },
    };

    onSubmit(request);
  };

  return (
    <form className="rag-search-form" onSubmit={handleSubmit}>
      <div className="rag-search-form__header">
        <div>
          <h2>RAG検索条件</h2>
          <p>値を入力してランキング結果を確認できます。</p>
        </div>
        <button type="submit" disabled={loading}>
          {loading ? "検索中..." : "検索する"}
        </button>
      </div>

      <fieldset>
        <legend>スコア重み</legend>
        <div className="rag-search-form__grid">
          <label>
            alpha
            <input
              name="alpha"
              type="number"
              step="0.01"
              defaultValue={defaultValue.policy?.alpha}
            />
          </label>
          <label>
            beta
            <input
              name="beta"
              type="number"
              step="0.01"
              defaultValue={defaultValue.policy?.beta}
            />
          </label>
          <label>
            gamma
            <input
              name="gamma"
              type="number"
              step="0.01"
              defaultValue={defaultValue.policy?.gamma}
            />
          </label>
          <label>
            delta
            <input
              name="delta"
              type="number"
              step="0.01"
              defaultValue={defaultValue.policy?.delta}
            />
          </label>
        </div>
      </fieldset>

      <fieldset>
        <legend>検索概念</legend>
        <label>
          query_concept_ids
          <input
            name="query_concept_ids"
            defaultValue={defaultValue.query_concept_ids.join(", ")}
          />
        </label>
        <div className="rag-search-form__grid">
          <label>
            concept id
            <input name="concept_id" defaultValue={concept.id} />
          </label>
          <label>
            label
            <input name="concept_label" defaultValue={concept.label} />
          </label>
          <label>
            parent ids
            <input
              name="concept_parent_ids"
              defaultValue={concept.parent_ids.join(", ")}
            />
          </label>
          <label>
            related ids
            <input
              name="concept_related_ids"
              defaultValue={concept.related_ids.join(", ")}
            />
          </label>
          <label>
            risk_level
            <input
              name="concept_risk_level"
              type="number"
              defaultValue={Number(concept.metadata.risk_level)}
            />
          </label>
          <label>
            urgency
            <input
              name="concept_urgency"
              type="number"
              defaultValue={Number(concept.metadata.urgency)}
            />
          </label>
        </div>
        <label>
          definition
          <textarea
            name="concept_definition"
            rows={2}
            defaultValue={concept.definition}
          />
        </label>
      </fieldset>

      <fieldset>
        <legend>文書</legend>
        {[primaryDocument, secondaryDocument].map((document, index) => {
          const prefix = index === 0 ? "primary" : "secondary";
          const similarity =
            defaultValue.text_similarities[document.id] ?? 0;

          return (
            <div className="rag-search-form__document" key={prefix}>
              <h3>{index === 0 ? "文書 1" : "文書 2"}</h3>
              <div className="rag-search-form__grid">
                <label>
                  document id
                  <input
                    name={`${prefix}_document_id`}
                    defaultValue={document.id}
                  />
                </label>
                <label>
                  text_similarity
                  <input
                    name={`${prefix}_text_similarity`}
                    type="number"
                    step="0.01"
                    defaultValue={similarity}
                  />
                </label>
                <label>
                  ontology_tags
                  <input
                    name={`${prefix}_document_tags`}
                    defaultValue={document.ontology_tags.join(", ")}
                  />
                </label>
                <label>
                  risk_level
                  <input
                    name={`${prefix}_document_risk_level`}
                    type="number"
                    defaultValue={Number(document.metadata.risk_level)}
                  />
                </label>
              </div>
              <label>
                text
                <textarea
                  name={`${prefix}_document_text`}
                  rows={2}
                  defaultValue={document.text}
                />
              </label>
            </div>
          );
        })}
      </fieldset>
    </form>
  );
}
