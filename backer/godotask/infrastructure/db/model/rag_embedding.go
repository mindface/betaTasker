package model

import "time"

// RagDocumentEmbedding stores the RAG document representation generated from
// seed-backed PostgreSQL records and its embedding payload.
type RagDocumentEmbedding struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DocumentID   string    `gorm:"type:varchar(255);uniqueIndex" json:"document_id"`
	SourceTable  string    `gorm:"type:varchar(100);index" json:"source_table"`
	SourceID     string    `gorm:"type:varchar(255);index" json:"source_id"`
	Text         string    `gorm:"type:text" json:"text"`
	OntologyTags JSON      `gorm:"type:jsonb" json:"ontology_tags"`
	Metadata     JSON      `gorm:"type:jsonb" json:"metadata"`
	Embedding    JSON      `gorm:"type:jsonb" json:"embedding"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
