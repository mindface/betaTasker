"use client"
import React, { useState, useEffect, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { loginRequest, loginSuccess, loginFailure } from '../modules/userReducer';
import { loginApi, logoutApi } from '../services/authApi';
import { fetchMemoriesService } from '../services/memoryApi';

export default function SectionLogin() {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    dispatch(loginRequest());
    const result = await loginApi(email, password);
    console.log(result);
    // if (result.token) {
    //   dispatch(loginSuccess(result.token));
    // } else {
    //   dispatch(loginFailure(result.error || 'ログイン失敗'));
    // }
  };

  const handleLogout = async () => {
    const result = await logoutApi();
    console.log(result);
  }

  const getMemory = async () => {
    const result = await fetchMemoriesService();
    console.log(result);
  }

  return (
    <div className="section__inner section--tools">
      <div className="section-continer">
        <div className="tools-header">
          <h2>ログイン</h2>
        </div>
        <form onSubmit={handleLogin} className="tools__body">
          <div>
            <label>メールアドレス</label>
            <input type="email" value={email} onChange={e => setEmail(e.target.value)} required />
          </div>
          <div>
            <label>パスワード</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />
          </div>
          <button type="submit" disabled={user.loading}>ログイン</button>
          {user.error && <div style={{color:'red'}}>{user.error}</div>}
          {user.isAuthenticated && <div style={{color:'green'}}>ログイン成功</div>}
        </form>
      </div>
      <button onClick={handleLogout}>ログアウト</button>
      <button onClick={getMemory}>getMemory</button>
    </div>
  );
}
