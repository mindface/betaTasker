package seed

import (
	"log"
	"fmt"
	"strconv"
	"time"
	"encoding/csv"
	"io"
	"os"
	"strings" 
	"math/rand"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

func classifyScore(score int) string {
	switch {
	case score >= 95:
		return "s_plus" // 卓越（非常に優れている）
	case score >= 90:
		return "s"      // 優秀（文句なし）
	case score >= 85:
		return "a_plus" // 良好（高評価）
	case score >= 80:
		return "a"      // 実用性あり（改善余地あり）
	case score >= 75:
		return "b_plus" // 満足（条件付きで応用可能）
	case score >= 70:
		return "b"
	case score >= 65:
		return "c_plus"
	case score >= 60:
		return "c"
	case score >= 55:
		return "d_plus"
	case score >= 50:
		return "d"
	default:
		return "e"  // 再考・再検証が必要
	}
}

func generateNoteText(scoreClass string, tags []string) string {
	tagText := ""
	if len(tags) > 0 {
		tagText = fmt.Sprintf(" 対象タグ: %s。", strings.Join(tags, ", "))
	}

	templates := map[string]string{
		"s_plus": "この素材は極めて高い評価を得ており、実務レベルでの即時活用が推奨されます。",
		"s":      "高評価の対象であり、すでに業務応用に十分な水準に達しています。",
		"a_plus": "専門的な観点からも優れた成果が得られており、今後の応用が期待されます。",
		"a":      "実用化に十分耐えうる性能を示しており、特定用途への展開が可能です。",
		"b_plus": "使用条件次第では実装可能と判断される水準にあります。",
		"b":      "中程度の評価結果であり、さらなる検証が推奨されます。",
		"c_plus": "限定的な条件下での使用にとどまりそうですが、将来的な改善が期待されます。",
		"c":      "現時点での有用性は限定的であり、慎重な評価が求められます。",
		"d_plus": "多くの課題が残っており、基本的な再評価が必要です。",
		"d":      "導入には大きなハードルが存在し、要再検討項目です。",
		"e":      "評価結果は不十分で、使用は推奨できません。",
	}

	return templates[scoreClass] + tagText
}

func generateTaskDescription(scoreClass string) string {
	desc := map[string]string{
		"s_plus": "即時展開を想定し、実装フェーズへの移行を進める。",
		"s":      "実証フェーズを省略し、直接的な運用評価へと移行する。",
		"a_plus": "応用領域を拡大するための展開計画を立案する。",
		"a":      "プロトタイプ環境にて早期適用を試みる。",
		"b_plus": "部分的な業務活用を見据え、パイロット導入を検討する。",
		"b":      "追加評価と検証を行い、使用可否を精査する。",
		"c_plus": "限定分野にて試行的に導入して経過観察する。",
		"c":      "関連文献を踏まえつつ、慎重な検証計画を策定する。",
		"d_plus": "基本性能の再検証および改善案の策定を行う。",
		"d":      "使用中止とともに代替案の模索を開始する。",
		"e":      "対象から除外し、他の候補へリソースを集中させる。",
	}
	return desc[scoreClass]
}

func generateTaskTitle(scoreClass, baseTitle string) string {
	prefix := map[string]string{
		"s_plus": "最優先評価",
		"s":      "優先評価",
		"a_plus": "高評価素材",
		"a":      "実用検証",
		"b_plus": "試験導入",
		"b":      "検討中素材",
		"c_plus": "条件付き評価",
		"c":      "再検証候補",
		"d_plus": "懸念要素あり",
		"d":      "非推奨候補",
		"e":      "使用不可",
	}
	return fmt.Sprintf("%s: %s", prefix[scoreClass], baseTitle)
}

func generateFactor(scoreClass string) string {
	factors := map[string]string{
		"s_plus": "応用実績",
		"s":      "成熟度",
		"a_plus": "信頼性",
		"a":      "互換性",
		"b_plus": "使用条件",
		"b":      "導入障壁",
		"c_plus": "改善余地",
		"c":      "未検証要素",
		"d_plus": "リスク因子",
		"d":      "技術的不確実性",
		"e":      "不適格要素",
	}
	return factors[scoreClass]
}

func generateProcess(scoreClass string) string {
	processes := map[string]string{
		"s_plus": "導入済→評価完了",
		"s":      "プロト導入→業務適用",
		"a_plus": "評価→試験導入",
		"a":      "適用検討→試験実施",
		"b_plus": "パイロット導入中",
		"b":      "追加検討段階",
		"c_plus": "初期評価段階",
		"c":      "導入見送り検討中",
		"d_plus": "再検討フロー中",
		"d":      "停止中→他技術選定",
		"e":      "評価対象外",
	}
	return processes[scoreClass]
}

func generateEvaluationAxis(scoreClass string) string {
	axes := map[string]string{
		"s_plus": "実装性・効果",
		"s":      "即効性・応用性",
		"a_plus": "適用性・持続性",
		"a":      "効果・効率",
		"b_plus": "有用性・安定性",
		"b":      "導入コスト・効果",
		"c_plus": "試行適性・調整可能性",
		"c":      "妥当性・拡張性",
		"d_plus": "技術課題・適用難度",
		"d":      "問題点・置換可能性",
		"e":      "不可要因・改善不能性",
	}
	return axes[scoreClass]
}

func generateInformationAmount(scoreClass string) string {
	info := map[string]string{
		"s_plus": "多数の応用事例と論文",
		"s":      "十分な文献と業界実績",
		"a_plus": "実証データ多数あり",
		"a":      "プロトデータ一部あり",
		"b_plus": "社内試験記録あり",
		"b":      "社内メモ・少数サンプル",
		"c_plus": "関連資料少数",
		"c":      "参考文献のみ",
		"d_plus": "社外参考事例のみ",
		"d":      "データ未整備",
		"e":      "情報不足",
	}
	return info[scoreClass]
}


func SeedMemoriesModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/memories.csv")
	if err != nil {
		return fmt.Errorf("could not open memories_models.csv: %v", err)
	}

	reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	return fmt.Errorf("could not read CSV: %v", err)
	// }
	defer file.Close()
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var models []model.Memory
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

    id, _ := strconv.Atoi(record[0])
    userID, _ := strconv.Atoi(record[1])

		// スコア生成と分類
		effectiveness := rand.Intn(41) + 60
		scoreClass := classifyScore(effectiveness) // メインの指標に応じて分類

		tags := []string{"3Dプリント", "材料", "Mg合金"}
		// noteText := generateNoteText(scoreClass, tags)
    createdAt, _ := time.Parse("2006-01-02", record[5])
    updatedAt, _ := time.Parse("2006-01-02", record[6])
		t, err := time.Parse(time.RFC3339, record[4])
		var readDate *time.Time
		if err != nil {
				// パースできない場合は nil にする or スキップ
				readDate = nil
		} else {
				readDate = &t
		}

		models = append(models, model.Memory{
			ID:        id,
			UserID:     userID,
			SourceType: "book",
			Title:      record[4],
			Author:     "Researcher",
			Notes:      record[5],
			Tags:       strings.Join(tags, ","),
			ReadStatus: "finished",
			ReadDate:   readDate,
			Factor:            generateFactor(scoreClass),
			Process:           generateProcess(scoreClass),
			EvaluationAxis:    generateEvaluationAxis(scoreClass),
			InformationAmount: generateInformationAmount(scoreClass),
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		})
	}

	// バッチインサート
	if len(models) > 0 {
		if err := db.Create(&models).Error; err != nil {
			return fmt.Errorf("failed to insert optimization models: %w", err)
		}
		fmt.Printf("Successfully seeded %d optimization models\n", len(models))
	}

	log.Printf("✓ Successfully seeded %d optimization models", len(models))
	return nil
}
