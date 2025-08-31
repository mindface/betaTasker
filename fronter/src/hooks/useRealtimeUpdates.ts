/**
 * ヒューリスティクス機能用リアルタイム更新フック
 * WebSocketを使用してリアルタイムデータ更新を提供
 */

import { useEffect, useState, useCallback, useRef } from 'react';
import { useDispatch } from 'react-redux';
import { 
  addPattern, 
  addInsight, 
  updateAnalysis,
  setAnalysisStatus 
} from '../features/heuristics/heuristicsSlice';

interface RealtimeMessage {
  type: 'PATTERN_DETECTED' | 'INSIGHT_GENERATED' | 'ANALYSIS_COMPLETED' | 'ANALYSIS_FAILED';
  data: any;
  timestamp: string;
}

interface RealtimeOptions {
  autoReconnect?: boolean;
  reconnectInterval?: number;
  maxReconnectAttempts?: number;
}

export const useRealtimeUpdates = (options: RealtimeOptions = {}) => {
  const dispatch = useDispatch();
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);
  const [connectionStatus, setConnectionStatus] = useState<'connecting' | 'connected' | 'disconnected' | 'error'>('disconnected');
  const [lastMessage, setLastMessage] = useState<RealtimeMessage | null>(null);
  const [error, setError] = useState<string | null>(null);
  
  const reconnectAttempts = useRef(0);
  const reconnectTimeout = useRef<NodeJS.Timeout | null>(null);
  const heartbeatInterval = useRef<NodeJS.Timeout | null>(null);
  
  const {
    autoReconnect = true,
    reconnectInterval = 5000,
    maxReconnectAttempts = 10
  } = options;

  // WebSocket接続の確立
  const connect = useCallback(() => {
    try {
      const wsUrl = process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8080/ws/heuristics';
      const ws = new WebSocket(wsUrl);
      
      setConnectionStatus('connecting');
      
      ws.onopen = () => {
        console.log('WebSocket接続が確立されました');
        setIsConnected(true);
        setConnectionStatus('connected');
        setError(null);
        reconnectAttempts.current = 0;
        
        // ハートビートの開始
        heartbeatInterval.current = setInterval(() => {
          if (ws.readyState === WebSocket.OPEN) {
            ws.send(JSON.stringify({ type: 'ping', timestamp: Date.now() }));
          }
        }, 30000); // 30秒ごと
      };
      
      ws.onmessage = (event) => {
        try {
          const message: RealtimeMessage = JSON.parse(event.data);
          setLastMessage(message);
          
          // メッセージタイプに基づく処理
          switch (message.type) {
            case 'PATTERN_DETECTED':
              dispatch(addPattern(message.data.pattern));
              break;
              
            case 'INSIGHT_GENERATED':
              dispatch(addInsight(message.data.insight));
              break;
              
            case 'ANALYSIS_COMPLETED':
              dispatch(updateAnalysis(message.data.analysis));
              dispatch(setAnalysisStatus('completed'));
              break;
              
            case 'ANALYSIS_FAILED':
              dispatch(setAnalysisStatus('failed'));
              break;
              
            default:
              console.log('未対応のメッセージタイプ:', message.type);
          }
        } catch (parseError) {
          console.error('メッセージの解析に失敗しました:', parseError);
        }
      };
      
      ws.onclose = (event) => {
        console.log('WebSocket接続が閉じられました:', event.code, event.reason);
        setIsConnected(false);
        setConnectionStatus('disconnected');
        
        // ハートビートの停止
        if (heartbeatInterval.current) {
          clearInterval(heartbeatInterval.current);
          heartbeatInterval.current = null;
        }
        
        // 自動再接続
        if (autoReconnect && reconnectAttempts.current < maxReconnectAttempts) {
          reconnectAttempts.current += 1;
          console.log(`再接続を試行中... (${reconnectAttempts.current}/${maxReconnectAttempts})`);
          
          reconnectTimeout.current = setTimeout(() => {
            connect();
          }, reconnectInterval);
        } else if (reconnectAttempts.current >= maxReconnectAttempts) {
          setConnectionStatus('error');
          setError('最大再接続回数に達しました');
        }
      };
      
      ws.onerror = (event) => {
        console.error('WebSocketエラーが発生しました:', event);
        setConnectionStatus('error');
        setError('WebSocket接続エラー');
      };
      
      setSocket(ws);
      
    } catch (err) {
      console.error('WebSocket接続の確立に失敗しました:', err);
      setConnectionStatus('error');
      setError('接続の確立に失敗しました');
    }
  }, [dispatch, autoReconnect, reconnectInterval, maxReconnectAttempts]);

  // WebSocket接続の切断
  const disconnect = useCallback(() => {
    if (socket) {
      socket.close(1000, 'ユーザーによる切断');
      setSocket(null);
    }
    
    // タイマーのクリーンアップ
    if (reconnectTimeout.current) {
      clearTimeout(reconnectTimeout.current);
      reconnectTimeout.current = null;
    }
    
    if (heartbeatInterval.current) {
      clearInterval(heartbeatInterval.current);
      heartbeatInterval.current = null;
    }
    
    setIsConnected(false);
    setConnectionStatus('disconnected');
    reconnectAttempts.current = 0;
  }, [socket]);

  // メッセージの送信
  const sendMessage = useCallback((message: any) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(message));
      return true;
    }
    return false;
  }, [socket]);

  // 分析の開始
  const startAnalysis = useCallback((analysisId: string) => {
    return sendMessage({
      type: 'START_ANALYSIS',
      analysisId,
      timestamp: Date.now()
    });
  }, [sendMessage]);

  // パターン検出の開始
  const startPatternDetection = useCallback((params: {
    user_id: string;
    data_type: string;
    period: string;
  }) => {
    return sendMessage({
      type: 'START_PATTERN_DETECTION',
      params,
      timestamp: Date.now()
    });
  }, [sendMessage]);

  // インサイト生成の開始
  const startInsightGeneration = useCallback((params: {
    user_id: string;
    insight_type: string;
  }) => {
    return sendMessage({
      type: 'START_INSIGHT_GENERATION',
      params,
      timestamp: Date.now()
    });
  }, [sendMessage]);

  // 接続状態の監視
  useEffect(() => {
    connect();
    
    return () => {
      disconnect();
    };
  }, [connect, disconnect]);

  // 接続状態の変更を監視
  useEffect(() => {
    if (connectionStatus === 'error' && autoReconnect) {
      const timeout = setTimeout(() => {
        if (reconnectAttempts.current < maxReconnectAttempts) {
          connect();
        }
      }, reconnectInterval * 2);
      
      return () => clearTimeout(timeout);
    }
  }, [connectionStatus, autoReconnect, reconnectInterval, maxReconnectAttempts, connect]);

  // 手動再接続
  const reconnect = useCallback(() => {
    disconnect();
    reconnectAttempts.current = 0;
    setTimeout(connect, 1000);
  }, [disconnect, connect]);

  return {
    isConnected,
    connectionStatus,
    lastMessage,
    error,
    sendMessage,
    startAnalysis,
    startPatternDetection,
    startInsightGeneration,
    reconnect,
    disconnect
  };
};

// デフォルトのリアルタイム更新フック
export const useHeuristicsRealtime = () => {
  return useRealtimeUpdates({
    autoReconnect: true,
    reconnectInterval: 5000,
    maxReconnectAttempts: 10
  });
};
