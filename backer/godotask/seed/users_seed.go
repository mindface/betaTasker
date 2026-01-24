package seed

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/godotask/infrastructure/db/model"
)

// 日本の名字と名前のリスト
var lastNames = []string{
	"田中", "佐藤", "鈴木", "高橋", "渡辺", "伊藤", "山本", "中村", "小林", "加藤",
	"吉田", "山田", "佐々木", "山口", "松本", "井上", "木村", "林", "斎藤", "清水",
	"山崎", "森", "池田", "橋本", "阿部", "石川", "前田", "藤田", "後藤", "長谷川",
	"村上", "近藤", "石井", "遠藤", "青木", "坂本", "福田", "西村", "太田", "岡田",
	"中島", "藤井", "新井", "福井", "三浦", "藤原", "岡本", "松田", "中川", "竹内",
}

var firstNames = []string{
	"太郎", "次郎", "健一", "誠", "大輔", "拓也", "翔太", "健太", "大樹", "勇",
	"花子", "美咲", "愛", "結衣", "さくら", "美穂", "真理子", "由美", "恵子", "和子",
	"一郎", "二郎", "三郎", "正", "明", "清", "武", "浩", "修", "茂",
	"陽子", "幸子", "裕子", "智子", "直子", "京子", "優子", "絵美", "麻衣", "里奈",
	"健", "勝", "進", "隆", "弘", "豊", "学", "淳", "剛", "純",
}

// ロールのリスト
var roles = []string{
	"trainee",          // 研修生
	"engineer",         // エンジニア
	"senior_engineer",  // シニアエンジニア
	"expert",           // エキスパート
	"supervisor",       // スーパーバイザー
	"manager",          // マネージャー
	"specialist",       // スペシャリスト
	"consultant",       // コンサルタント
}

// ファクター（専門分野）のリスト
var factors = []string{
	"manufacturing",        // 製造
	"quality_control",      // 品質管理
	"advanced_materials",   // 先端材料
	"precision_machining",  // 精密加工
	"tool_management",      // 工具管理
	"production_planning",  // 生産計画
	"safety_management",    // 安全管理
	"maintenance",          // メンテナンス
	"cad_cam_operations",   // CAD/CAM操作
	"measurement_technology", // 測定技術
	"cutting_theory",       // 切削理論
	"automation",           // 自動化
	"lean_manufacturing",   // リーン生産
	"knowledge_transfer",   // 知識伝達
	"basic_operations",     // 基本操作
	"innovation",           // イノベーション
	"research_development", // 研究開発
	"technical_documentation", // 技術文書
}

// プロセス（作業工程）のリスト
var processes = []string{
	"machining",            // 加工
	"precision_machining",  // 精密加工
	"special_machining",    // 特殊加工
	"assembly",             // 組立
	"inspection",           // 検査
	"training",             // 訓練
	"learning",             // 学習
	"planning",             // 計画
	"optimization",         // 最適化
	"analysis",             // 分析
	"design",               // 設計
	"programming",          // プログラミング
	"setup",                // 段取り
	"maintenance_ops",      // メンテナンス作業
	"quality_assurance",    // 品質保証
	"process_improvement",  // 工程改善
	"research",             // 研究
	"documentation",        // 文書化
}

// 評価軸のリスト
var evaluationAxes = []string{
	"skill_level",          // スキルレベル
	"productivity",         // 生産性
	"quality_score",        // 品質スコア
	"efficiency",           // 効率性
	"accuracy",             // 正確性
	"speed",                // スピード
	"innovation_capability", // イノベーション能力
	"knowledge_depth",      // 知識の深さ
	"problem_solving",      // 問題解決力
	"teamwork",             // チームワーク
	"leadership",           // リーダーシップ
	"adaptability",         // 適応力
	"technical_expertise",  // 技術的専門性
	"learning_speed",       // 学習速度
	"reliability",          // 信頼性
}

// 情報量のリスト
var informationAmounts = []string{
	"basic",       // 基本
	"intermediate", // 中級
	"advanced",    // 上級
	"expert",      // エキスパート
	"comprehensive", // 包括的
	"detailed",    // 詳細
	"minimal",     // 最小限
	"extensive",   // 広範囲
	"specialized", // 専門的
	"integrated",  // 統合的
}

// generateUsername - ユーザー名を生成
func generateUsername(lastName, firstName string, index int) string {
	return fmt.Sprintf("%s_%s_%d", romanize(lastName), romanize(firstName), index)
}

// romanize - 簡易的なローマ字変換
func romanize(name string) string {
	romanMap := map[string]string{
		"田中": "tanaka", "佐藤": "sato", "鈴木": "suzuki", "高橋": "takahashi", "渡辺": "watanabe",
		"伊藤": "ito", "山本": "yamamoto", "中村": "nakamura", "小林": "kobayashi", "加藤": "kato",
		"吉田": "yoshida", "山田": "yamada", "佐々木": "sasaki", "山口": "yamaguchi", "松本": "matsumoto",
		"井上": "inoue", "木村": "kimura", "林": "hayashi", "斎藤": "saito", "清水": "shimizu",
		"山崎": "yamazaki", "森": "mori", "池田": "ikeda", "橋本": "hashimoto", "阿部": "abe",
		"石川": "ishikawa", "前田": "maeda", "藤田": "fujita", "後藤": "goto", "長谷川": "hasegawa",
		"村上": "murakami", "近藤": "kondo", "石井": "ishii", "遠藤": "endo", "青木": "aoki",
		"坂本": "sakamoto", "福田": "fukuda", "西村": "nishimura", "太田": "ota", "岡田": "okada",
		"中島": "nakajima", "藤井": "fujii", "新井": "arai", "福井": "fukui", "三浦": "miura",
		"藤原": "fujiwara", "岡本": "okamoto", "松田": "matsuda", "中川": "nakagawa", "竹内": "takeuchi",
		"太郎": "taro", "次郎": "jiro", "健一": "kenichi", "誠": "makoto", "大輔": "daisuke",
		"拓也": "takuya", "翔太": "shota", "健太": "kenta", "大樹": "daiki", "勇": "isamu",
		"花子": "hanako", "美咲": "misaki", "愛": "ai", "結衣": "yui", "さくら": "sakura",
		"美穂": "miho", "真理子": "mariko", "由美": "yumi", "恵子": "keiko", "和子": "kazuko",
		"一郎": "ichiro", "二郎": "niro", "三郎": "saburo", "正": "tadashi", "明": "akira",
		"清": "kiyoshi", "武": "takeshi", "浩": "hiroshi", "修": "osamu", "茂": "shigeru",
		"陽子": "yoko", "幸子": "sachiko", "裕子": "yuko", "智子": "tomoko", "直子": "naoko",
		"京子": "kyoko", "優子": "yuko", "絵美": "emi", "麻衣": "mai", "里奈": "rina",
		"健": "ken", "勝": "masaru", "進": "susumu", "隆": "takashi", "弘": "hiroshi",
		"豊": "yutaka", "学": "manabu", "淳": "atsushi", "剛": "tsuyoshi", "純": "jun",
	}

	if roman, ok := romanMap[name]; ok {
		return roman
	}
	return "user"
}

// hashPassword - パスワードをハッシュ化
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// SeedAdminUser - 管理者ユーザーを作成
func SeedAdminUser(db *gorm.DB) error {
	log.Println("Starting admin user seeding...")

	// 固定パスワード（本番環境では環境変数から読み込むことを推奨）
	adminPassword := "Admin@123456" // 変更可能
	
	hashedPassword, err := hashPassword(adminPassword)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	adminUser := model.User{
		ID:                1,
		Username:          "admin",
		Email:             "admin@example.com",
		PasswordHash:      hashedPassword,
		Role:              "admin",
		CreatedAt:         time.Now().Add(-365 * 24 * time.Hour),
		UpdatedAt:         time.Now(),
		IsActive:          true,
		Factor:            "system_management",
		Process:           "administration",
		EvaluationAxis:    "leadership",
		InformationAmount: "comprehensive",
	}

	// 既存の管理者をチェック
	var existingAdmin model.User
	if err := db.Where("id = ? OR username = ? OR email = ?", 1, "admin", "admin@example.com").First(&existingAdmin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 管理者が存在しない場合は作成
			if err := db.Create(&adminUser).Error; err != nil {
				return fmt.Errorf("failed to create admin user: %w", err)
			}
			fmt.Println("✓ Successfully created admin user")
			fmt.Printf("  Username: %s\n", adminUser.Username)
			fmt.Printf("  Email: %s\n", adminUser.Email)
			fmt.Printf("  Password: %s\n", adminPassword)
			fmt.Println("  ⚠️  Please change the password after first login!")
		} else {
			return fmt.Errorf("failed to check existing admin: %w", err)
		}
	} else {
		fmt.Println("⚠️  Admin user already exists, skipping creation")
		fmt.Printf("  Existing admin ID: %d\n", existingAdmin.ID)
		fmt.Printf("  Existing admin username: %s\n", existingAdmin.Username)
	}

	return nil
}

// SeedUsers - 99人の一般ユーザーデータのシード（ID 2-100）
func SeedUsers(db *gorm.DB) error {
	log.Println("Starting users seeding (99 regular users)...")

	users := []model.User{}
	baseTime := time.Now()

	// ID 2から100まで（ID 1は管理者用に予約）
	for i := 2; i <= 300; i++ {
		lastName := lastNames[(i-2)%len(lastNames)]
		firstName := firstNames[(i-2)%len(firstNames)]
		username := generateUsername(lastName, firstName, i)
		email := fmt.Sprintf("%s@example.com", username)
		
		// デフォルトパスワード（テスト用）
		defaultPassword := fmt.Sprintf("User@%d", i)
		hashedPassword, err := hashPassword(defaultPassword)
		if err != nil {
			return fmt.Errorf("failed to hash password for user %d: %w", i, err)
		}
		
		// ロールの決定（経験に応じて段階的に）
		var role string
		if i <= 20 {
			role = roles[0] // trainee
		} else if i <= 40 {
			role = roles[1] // engineer
		} else if i <= 60 {
			role = roles[2] // senior_engineer
		} else if i <= 75 {
			role = roles[3] // expert
		} else if i <= 85 {
			role = roles[4] // supervisor
		} else if i <= 95 {
			role = roles[5] // manager
		} else {
			role = roles[6] // specialist
		}

		// ファクターとプロセスの決定
		factor := factors[(i-2)%len(factors)]
		process := processes[(i-2)%len(processes)]
		
		// 評価軸と情報量の決定
		evaluationAxis := evaluationAxes[(i-2)%len(evaluationAxes)]
		informationAmount := informationAmounts[(i-2)%len(informationAmounts)]

		// アカウント作成日（過去30日〜180日の範囲）
		daysAgo := 30 + ((i - 2) * 150 / 99)
		createdAt := baseTime.Add(-time.Duration(daysAgo) * 24 * time.Hour)
		
		// 更新日（作成日から現在までの範囲）
		hoursAgo := (i - 2) % 48
		updatedAt := baseTime.Add(-time.Duration(hoursAgo) * time.Hour)

		// アクティブ状態（95%はアクティブ）
		isActive := i <= 95

		user := model.User{
			ID:                uint(i),
			Username:          username,
			Email:             email,
			PasswordHash:      hashedPassword,
			Role:              role,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
			IsActive:          isActive,
			Factor:            factor,
			Process:           process,
			EvaluationAxis:    evaluationAxis,
			InformationAmount: informationAmount,
		}

		users = append(users, user)
	}

	// バッチインサート
	if len(users) > 0 {
		if err := db.CreateInBatches(&users, 100).Error; err != nil {
			return fmt.Errorf("failed to insert users: %w", err)
		}
		fmt.Printf("✓ Successfully seeded %d regular users (ID: 2-100)\n", len(users))
		fmt.Println("  Default password format: User@{ID}")
		fmt.Println("  Example: User@2, User@3, User@4, ...")
	}

	// 統計情報を表示
	printUserStatistics(users)
	
	return nil
}

// printUserStatistics - ユーザー統計を表示
func printUserStatistics(users []model.User) {
	fmt.Println("\n=== User Statistics ===")

	// ロール別集計
	roleCount := make(map[string]int)
	for _, user := range users {
		roleCount[user.Role]++
	}
	fmt.Println("Role Distribution:")
	for role, count := range roleCount {
		fmt.Printf("  %s: %d users\n", role, count)
	}
	
	// アクティブユーザー数
	activeCount := 0
	for _, user := range users {
		if user.IsActive {
			activeCount++
		}
	}
	fmt.Printf("\nActive Users: %d / %d (%.1f%%)\n", activeCount, len(users), float64(activeCount)/float64(len(users))*100)
	
	// ファクター別集計（上位5件）
	factorCount := make(map[string]int)
	for _, user := range users {
		factorCount[user.Factor]++
	}
	fmt.Println("\nFactor Distribution (Top 5):")
	count := 0
	for factor, cnt := range factorCount {
		if count < 5 {
			fmt.Printf("  %s: %d users\n", factor, cnt)
			count++
		}
	}
	
	// プロセス別集計（上位5件）
	processCount := make(map[string]int)
	for _, user := range users {
		processCount[user.Process]++
	}
	fmt.Println("\nProcess Distribution (Top 5):")
	count = 0
	for process, cnt := range processCount {
		if count < 5 {
			fmt.Printf("  %s: %d users\n", process, cnt)
			count++
		}
	}
	
	fmt.Println("======================\n")
}

// SeedAllUsers - 管理者と一般ユーザーを一括でシード
func SeedAllUsers(db *gorm.DB) error {
	// 管理者ユーザーを作成
	if err := SeedAdminUser(db); err != nil {
		return fmt.Errorf("failed to seed admin user: %w", err)
	}

	// 一般ユーザーを作成
	if err := SeedUsers(db); err != nil {
		return fmt.Errorf("failed to seed regular users: %w", err)
	}

	fmt.Println("\n✓ All users seeded successfully!")
	return nil
}

// CleanUsers - ユーザーデータをクリーンアップ（開発用）
func CleanUsers(db *gorm.DB) error {
	log.Println("Cleaning existing users...")
	
	if err := db.Where("id <= ?", 100).Delete(&model.User{}).Error; err != nil {
		return fmt.Errorf("failed to clean users: %w", err)
	}
	
	fmt.Println("✓ Successfully cleaned users")
	return nil
}

// GetUserCount - ユーザー数を取得
func GetUserCount(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}
	return count, nil
}

// VerifyPassword - パスワード検証用ヘルパー関数（テスト用）
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}