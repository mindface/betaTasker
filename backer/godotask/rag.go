package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/godotask/rag"
	_ "github.com/mattn/go-sqlite3"
)

// Python側と同じ構造に合わせた構造体
type Document struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	FullText  string `json:"full_text"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

func ragHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "クエリパラメータ 'q' が必要です。例: /rag?q=保険", http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`
		SELECT summary, full_text FROM documents
		WHERE full_text LIKE ? COLLATE NOCASE OR summary LIKE ? COLLATE NOCASE
		LIMIT 3
	`, "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, "検索に失敗しました", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	context := ""
	for rows.Next() {
		var summary, fullText string
		if err := rows.Scan(&summary, &fullText); err == nil {
			context += summary + "\n" + fullText + "\n\n"
		}
	}

	if context == "" {
		http.Error(w, "関連文書が見つかりませんでした。", http.StatusNotFound)
		return
	}

	prompt := "以下の文書コンテキストを元に、ユーザーの質問に日本語で詳しく答えてください。\n\n" +
		"[文書コンテキスト]:\n" + context[:3000] + "\n\n" +
		"[質問]: " + query

	answer, err := rag.QueryOllama(prompt)
	if err != nil {
		http.Error(w, "Ollamaとの通信に失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"answer": answer})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "クエリパラメータ 'q' が必要です。例: /search?q=保険", http.StatusBadRequest)
		return
	}
	log.Println(query)
	db, err := sql.Open("sqlite3", "./pdfdb/documents.db")
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	rows, err := db.Query(`
		SELECT id, title, summary, full_text
		FROM documents
		WHERE full_text LIKE ? OR summary LIKE ?
	`, "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, "検索に失敗しました", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []Document
	for rows.Next() {
		var doc Document
		err := rows.Scan(&doc.ID, &doc.Title, &doc.Summary, &doc.FullText)
		if err == nil {
			results = append(results, doc)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	// SQLiteデータベースに接続
	var err error
	db, err = sql.Open("sqlite3", "./pdfdb/documents.db")
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	// defer db.Close()

	// GET /documents で全件取得
	http.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, full_text, summary, created_at FROM documents ORDER BY id DESC")
		if err != nil {
			http.Error(w, "データベースエラー: "+err.Error(), http.StatusInternalServerError)
			returns
		}
		defer rows.Close()

		var documents []Document
		for rows.Next() {
			var doc Document
			err := rows.Scan(&doc.ID, &doc.Title, &doc.FullText, &doc.Summary, &doc.CreatedAt)
			if err != nil {
				http.Error(w, "データ読み取りエラー: "+err.Error(), http.StatusInternalServerError)
				return
			}
			documents = append(documents, doc)
		}

		// JSONレスポンスとして返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(documents)
	})
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/rag", ragHandler)

	log.Println("📡 サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
