"use client"
import React, { useState } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { loginRequest, loginSuccess, loginFailure } from '../features/user/userSlice';
import { loginApi, logoutApi } from '../services/authApi';
import { fetchMemoriesService } from '../services/memoryApi';

export default function SectionLogin() {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user);
  const [loginSwitch, setLoginSwitch] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    dispatch(loginRequest());
    const result = await loginApi(email, password);
    if (result.token && result.user) {
      dispatch(loginSuccess({ token: result.token, user: result.user }));
    } else {
      dispatch(loginFailure(result.error || 'ログイン失敗'));
    }
  };

  const handleLogout = async () => {
    const result = await logoutApi();
    console.log(result);
  }

  const getMemory = async () => {
    const result = await fetchMemoriesService();
    console.log(result);
  }

  const switchAction = () => {
    setLoginSwitch(!loginSwitch)
  }

  return (
    <div className="section__inner section--tools">
      <div className="section-continer">
        <div className="tools-header">
          <h2>{ loginSwitch ? "ログイン画面" : "新規登録" }</h2>
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
      <div className="swith-text" onClick={switchAction}>{ loginSwitch ? "新規登録へ移行" : "ログインへ移行" }</div>
      <button onClick={getMemory}>getMemory</button>
    </div>
  );
}
