'use client';

import React, { useState, useRef, useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  createLabel,
  updateLabel,
  verifyLabel,
  searchLabels,
  loadStatistics,
  suggestQuantification,
  startLabelCreation,
  startLabelEditing,
  closeLabelEditor,
  addAnnotation,
  addConcept,
  removeConcept,
  selectLabel,
  clearSelection,
} from '../../features/heuristics/labelSlice';

const QuantificationLabelManager: React.FC = () => {
  const dispatch = useAppDispatch();
  const {
    labels,
    searchResults,
    currentLabel,
    labelEditor,
    statistics,
    loading,
    error,
    selectedLabels,
  } = useAppSelector((state) => state.label);

  // ローカル状態
  const [activeTab, setActiveTab] = useState<'search' | 'create' | 'statistics'>('search');
  const [searchText, setSearchText] = useState('');
  const [filterDomain, setFilterDomain] = useState('');
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string>('');
  const [annotationMode, setAnnotationMode] = useState(false);
  const [selectedRegion, setSelectedRegion] = useState<{ x: number; y: number; width: number; height: number } | null>(null);

  const fileInputRef = useRef<HTMLInputElement>(null);
  const canvasRef = useRef<HTMLCanvasElement>(null);

  // 初期ロード
  useEffect(() => {
    dispatch(loadStatistics());
    dispatch(searchLabels({ limit: 20 }));
  }, [dispatch]);

  // ファイル選択ハンドラ
  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setImageFile(file);
      const url = URL.createObjectURL(file);
      setPreviewUrl(url);
    }
  };

  // 検索実行
  const handleSearch = () => {
    dispatch(searchLabels({
      text: searchText || undefined,
      domain: filterDomain || undefined,
      limit: 50,
    }));
  };

  // ラベル作成開始
  const handleStartCreation = () => {
    dispatch(startLabelCreation({
      text: searchText,
      imageUrl: previewUrl || undefined,
    }));
    setActiveTab('create');
  };

  // ラベル作成実行
  const handleCreateLabel = async (e: React.FormEvent) => {
    e.preventDefault();
    const formData = new FormData(e.target as HTMLFormElement);

    await dispatch(createLabel({
      text: formData.get('text') as string,
      description: formData.get('description') as string,
      value: parseFloat(formData.get('value') as string),
      unit: formData.get('unit') as string,
      domain: formData.get('domain') as string,
      category: formData.get('category') as string,
      imageFile: imageFile || undefined,
      concepts: labelEditor.concepts,
    }));
  };

  // ラベル検証
  const handleVerifyLabel = async (labelId: string, accurate: boolean) => {
    await dispatch(verifyLabel({
      labelId,
      verification: {
        accurate,
        consistency: true,
        reproducible: true,
        usable: true,
      },
      verifierId: 'current_user', // 実際の実装では現在のユーザーIDを使用
    }));
  };

  // 定量化提案を取得
  const handleGetSuggestion = async () => {
    if (labelEditor.originalText) {
      await dispatch(suggestQuantification({
        text: labelEditor.originalText,
        imageUrl: previewUrl || undefined,
        domain: filterDomain || undefined,
      }));
    }
  };

  // アノテーション追加
  const handleAddAnnotation = (region: { x: number; y: number; width: number; height: number }) => {
    const annotation = {
      id: Date.now().toString(),
      type: 'measurement',
      coordinates: region,
      label: '測定値',
      value: 0,
      unit: 'cm',
      confidence: 0.8,
      createdBy: 'current_user',
      createdAt: new Date().toISOString(),
    };
    
    dispatch(addAnnotation(annotation));
    setSelectedRegion(null);
    setAnnotationMode(false);
  };

  return (
    <div className="max-w-7xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">定量化ラベル管理</h1>

      {/* タブナビゲーション */}
      <div className="flex space-x-1 mb-6 border-b">
        {['search', 'create', 'statistics'].map((tab) => (
          <button
            key={tab}
            onClick={() => setActiveTab(tab as any)}
            className={`px-4 py-2 font-medium text-sm rounded-t-lg ${
              activeTab === tab
                ? 'bg-blue-500 text-white'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            {tab === 'search' && '検索・管理'}
            {tab === 'create' && 'ラベル作成'}
            {tab === 'statistics' && '統計'}
          </button>
        ))}
      </div>

      {/* 検索・管理タブ */}
      {activeTab === 'search' && (
        <div className="space-y-6">
          {/* 検索フィルタ */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold mb-4">検索・フィルタ</h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <label className="block text-sm font-medium mb-1">テキスト検索</label>
                <input
                  type="text"
                  value={searchText}
                  onChange={(e) => setSearchText(e.target.value)}
                  placeholder="例: 小さじ1杯, コップ半分"
                  className="w-full px-3 py-2 border rounded"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-1">ドメイン</label>
                <select
                  value={filterDomain}
                  onChange={(e) => setFilterDomain(e.target.value)}
                  className="w-full px-3 py-2 border rounded"
                >
                  <option value="">全て</option>
                  <option value="cooking">料理</option>
                  <option value="construction">建築</option>
                  <option value="design">デザイン</option>
                  <option value="measurement">測定</option>
                </select>
              </div>
              <div className="flex items-end">
                <button
                  onClick={handleSearch}
                  className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                >
                  検索
                </button>
                <button
                  onClick={handleStartCreation}
                  className="ml-2 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                >
                  新規作成
                </button>
              </div>
            </div>
          </div>

          {/* 検索結果 */}
          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-semibold">
                検索結果 ({searchResults.length}件)
              </h3>
              {selectedLabels.length > 0 && (
                <div className="flex space-x-2">
                  <span className="text-sm text-gray-600">
                    {selectedLabels.length}件選択中
                  </span>
                  <button
                    onClick={() => dispatch(clearSelection())}
                    className="text-sm text-blue-500 hover:underline"
                  >
                    選択解除
                  </button>
                </div>
              )}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {searchResults.map((label) => (
                <div
                  key={label.id}
                  className={`border rounded-lg p-4 hover:shadow-md cursor-pointer ${
                    selectedLabels.includes(label.id) ? 'border-blue-500 bg-blue-50' : ''
                  }`}
                  onClick={() => dispatch(selectLabel(label.id))}
                >
                  <div className="flex items-start space-x-3">
                    <input
                      type="checkbox"
                      checked={selectedLabels.includes(label.id)}
                      onChange={() => dispatch(selectLabel(label.id))}
                      className="mt-1"
                    />
                    
                    {label.visual.thumbnailUrl && (
                      <img
                        src={label.visual.thumbnailUrl}
                        alt={label.linguistic.originalText}
                        className="w-16 h-16 object-cover rounded"
                      />
                    )}
                    
                    <div className="flex-1 min-w-0">
                      <h4 className="font-medium text-gray-900 truncate">
                        {label.linguistic.originalText}
                      </h4>
                      <p className="text-sm text-gray-500 mt-1">
                        {label.quantification.value} {label.quantification.unit}
                      </p>
                      <div className="flex items-center mt-2">
                        <span className={`inline-block w-2 h-2 rounded-full mr-2 ${
                          label.metadata.validated ? 'bg-green-500' : 'bg-yellow-500'
                        }`}></span>
                        <span className="text-xs text-gray-500">
                          信頼度: {Math.round(label.quantification.confidence * 100)}%
                        </span>
                      </div>
                      <div className="flex flex-wrap gap-1 mt-2">
                        <span className="inline-block px-2 py-1 text-xs bg-gray-100 rounded">
                          {label.concept.abstractLevel}
                        </span>
                        {label.concept.semanticTags.slice(0, 2).map((tag, index) => (
                          <span key={index} className="inline-block px-2 py-1 text-xs bg-blue-100 text-blue-700 rounded">
                            {tag}
                          </span>
                        ))}
                      </div>
                    </div>
                  </div>

                  <div className="flex justify-between items-center mt-4">
                    <button
                      onClick={() => dispatch(startLabelEditing(label))}
                      className="text-sm text-blue-500 hover:underline"
                    >
                      編集
                    </button>
                    {!label.metadata.validated && (
                      <div className="flex space-x-2">
                        <button
                          onClick={() => handleVerifyLabel(label.id, true)}
                          className="text-sm text-green-500 hover:underline"
                        >
                          承認
                        </button>
                        <button
                          onClick={() => handleVerifyLabel(label.id, false)}
                          className="text-sm text-red-500 hover:underline"
                        >
                          却下
                        </button>
                      </div>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* ラベル作成タブ */}
      {activeTab === 'create' && (
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-4">新しいラベルの作成</h3>
          
          <form onSubmit={handleCreateLabel} className="space-y-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium mb-1">元テキスト *</label>
                <input
                  name="text"
                  type="text"
                  defaultValue={labelEditor.originalText}
                  required
                  className="w-full px-3 py-2 border rounded"
                  placeholder="例: 小さじ1杯"
                />
              </div>
              
              <div>
                <label className="block text-sm font-medium mb-1">説明 *</label>
                <input
                  name="description"
                  type="text"
                  required
                  className="w-full px-3 py-2 border rounded"
                  placeholder="この表現の詳細説明"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-1">数値 *</label>
                <input
                  name="value"
                  type="number"
                  step="any"
                  required
                  className="w-full px-3 py-2 border rounded"
                  placeholder="5.0"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-1">単位 *</label>
                <input
                  name="unit"
                  type="text"
                  required
                  className="w-full px-3 py-2 border rounded"
                  placeholder="ml"
                />
              </div>

              <div>
                <label className="block text-sm font-medium mb-1">ドメイン *</label>
                <select
                  name="domain"
                  required
                  className="w-full px-3 py-2 border rounded"
                >
                  <option value="cooking">料理</option>
                  <option value="construction">建築</option>
                  <option value="design">デザイン</option>
                  <option value="measurement">測定</option>
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium mb-1">カテゴリ *</label>
                <select
                  name="category"
                  required
                  className="w-full px-3 py-2 border rounded"
                >
                  <option value="volume">体積</option>
                  <option value="weight">重量</option>
                  <option value="length">長さ</option>
                  <option value="area">面積</option>
                  <option value="degree">程度</option>
                </select>
              </div>
            </div>

            {/* 画像アップロード */}
            <div>
              <label className="block text-sm font-medium mb-1">参考画像</label>
              <input
                ref={fileInputRef}
                type="file"
                accept="image/*"
                onChange={handleFileSelect}
                className="hidden"
              />
              <button
                type="button"
                onClick={() => fileInputRef.current?.click()}
                className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300"
              >
                画像を選択
              </button>

              {previewUrl && (
                <div className="mt-4 relative">
                  <img
                    src={previewUrl}
                    alt="Preview"
                    className="max-w-sm rounded"
                  />
                  <button
                    type="button"
                    onClick={() => setAnnotationMode(!annotationMode)}
                    className="mt-2 px-3 py-1 bg-blue-500 text-white rounded text-sm"
                  >
                    {annotationMode ? 'アノテーション終了' : 'アノテーション開始'}
                  </button>
                </div>
              )}
            </div>

            {/* 概念タグ */}
            <div>
              <label className="block text-sm font-medium mb-1">概念タグ</label>
              <div className="flex flex-wrap gap-2 mb-2">
                {labelEditor.concepts.map((concept, index) => (
                  <span
                    key={index}
                    className="inline-flex items-center px-2 py-1 text-sm bg-blue-100 text-blue-700 rounded"
                  >
                    {concept}
                    <button
                      type="button"
                      onClick={() => dispatch(removeConcept(concept))}
                      className="ml-1 text-blue-500 hover:text-blue-700"
                    >
                      ×
                    </button>
                  </span>
                ))}
              </div>
              <input
                type="text"
                placeholder="概念を入力してEnter"
                className="w-full px-3 py-2 border rounded"
                onKeyPress={(e) => {
                  if (e.key === 'Enter') {
                    e.preventDefault();
                    const value = (e.target as HTMLInputElement).value.trim();
                    if (value) {
                      dispatch(addConcept(value));
                      (e.target as HTMLInputElement).value = '';
                    }
                  }
                }}
              />
            </div>

            {/* 提案された定量化 */}
            {labelEditor.suggestedValues.length > 0 && (
              <div>
                <label className="block text-sm font-medium mb-1">提案された値</label>
                <div className="space-y-2">
                  {labelEditor.suggestedValues.map((suggestion, index) => (
                    <div key={index} className="p-3 bg-gray-50 rounded flex justify-between items-center">
                      <span>
                        {suggestion.value} {suggestion.unit}
                        <span className="text-sm text-gray-500 ml-2">
                          (信頼度: {Math.round(suggestion.confidence * 100)}%, ソース: {suggestion.source})
                        </span>
                      </span>
                      <button
                        type="button"
                        onClick={() => {
                          (document.querySelector('[name="value"]') as HTMLInputElement).value = suggestion.value.toString();
                          (document.querySelector('[name="unit"]') as HTMLInputElement).value = suggestion.unit;
                        }}
                        className="text-sm text-blue-500 hover:underline"
                      >
                        使用
                      </button>
                    </div>
                  ))}
                </div>
              </div>
            )}

            <div className="flex space-x-4">
              <button
                type="submit"
                disabled={loading}
                className="px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:opacity-50"
              >
                {loading ? '作成中...' : 'ラベル作成'}
              </button>
              
              <button
                type="button"
                onClick={handleGetSuggestion}
                className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
              >
                提案を取得
              </button>
              
              <button
                type="button"
                onClick={() => dispatch(closeLabelEditor())}
                className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600"
              >
                キャンセル
              </button>
            </div>
          </form>
        </div>
      )}

      {/* 統計タブ */}
      {activeTab === 'statistics' && statistics && (
        <div className="space-y-6">
          {/* 全体統計 */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold mb-4">全体統計</h3>
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div className="text-center">
                <p className="text-2xl font-bold text-blue-500">{statistics.totalLabels}</p>
                <p className="text-sm text-gray-600">総ラベル数</p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-bold text-green-500">
                  {Math.round(statistics.averageMetrics.confidence * 100)}%
                </p>
                <p className="text-sm text-gray-600">平均信頼度</p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-bold text-purple-500">
                  {statistics.quality.verified}
                </p>
                <p className="text-sm text-gray-600">検証済み</p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-bold text-orange-500">
                  {Object.keys(statistics.labelsByDomain).length}
                </p>
                <p className="text-sm text-gray-600">ドメイン数</p>
              </div>
            </div>
          </div>

          {/* ドメイン別分布 */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold mb-4">ドメイン別分布</h3>
            <div className="space-y-2">
              {Object.entries(statistics.labelsByDomain).map(([domain, count]) => (
                <div key={domain} className="flex justify-between items-center">
                  <span className="font-medium">{domain}</span>
                  <div className="flex items-center space-x-2">
                    <div className="w-32 bg-gray-200 rounded-full h-2">
                      <div
                        className="bg-blue-500 h-2 rounded-full"
                        style={{
                          width: `${(count / statistics.totalLabels) * 100}%`
                        }}
                      ></div>
                    </div>
                    <span className="text-sm text-gray-600">{count}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* エラー表示 */}
      {error && (
        <div className="mt-4 p-4 bg-red-100 text-red-700 rounded">
          {error}
        </div>
      )}
    </div>
  );
};

export default QuantificationLabelManager;