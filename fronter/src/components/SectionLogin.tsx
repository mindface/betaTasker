"use client"
import React, { useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { RootState } from "../store"
import { loginRequest, loginSuccess, loginFailure } from "../features/user/userSlice"
import { loginApi, logoutApi, regApi } from "../services/authApi"
import { fetchMemoriesService } from "../services/memoryApi"
import { useRouter } from "next/navigation"

export default function SectionLogin() {
  const router = useRouter()
  const dispatch = useDispatch()
  const user = useSelector((state: RootState) => state.user)
  const [loginSwitch, setLoginSwitch] = useState(true)
  const [email, setEmail] = useState("")
  const [userName, setUserName] = useState("")
  const [password, setPassword] = useState("")

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    dispatch(loginRequest())

    const result = await loginApi(email, password)
    if (result.token && result.user) {
      dispatch(loginSuccess({ token: result.token, user: result.user }))
      router.push("/")
    } else {
      dispatch(loginFailure(result.error || 'ログイン失敗'))
    }
  }

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault()
    const result = await regApi({ username: userName, email, password, role: "user" })
    if (result.token && result.user) {
      dispatch(loginSuccess({ token: result.token, user: result.user }))
      router.push("/")
    } else {
      dispatch(loginFailure(result.error || '登録失敗'))
    }
  }

  const handleLogout = async () => {
    const result = await logoutApi()
  }

  const getMemory = async () => {
    const result = await fetchMemoriesService()
  }

  const switchAction = () => {
    setLoginSwitch(!loginSwitch)
  }

  return (
    <div className="section__inner section--tools base-width">
      <div className="p-b-8">
        <div className="swith-text pointer text-btn" onClick={switchAction}>{ loginSwitch ? "新規登録へ移行" : "ログインへ移行" }</div>
      </div>
      <div className="section p-8">
        <div className="tools-header">
          <h2 className="p-b-8">{ loginSwitch ? "ログイン画面" : "新規登録" }</h2>
        </div>
        <form onSubmit={handleLogin} className="tools__body">
          <div className="p-b-8">
            <label>メールアドレス</label>
            <input type="email" value={email} onChange={e => setEmail(e.target.value)} required />
          </div>
          <div className="p-b-8">
            <label>パスワード</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />
          </div>
          {userName && <div className="p-b-8">
            <label>ユーザー名</label>
            <input type="text" value={userName} onChange={e => setUserName(e.target.value)} required />
          </div>}
          {loginSwitch ? (
            <button type="submit" disabled={user.loading}>ログイン</button>
          ) : (
            <button type="submit" disabled={user.loading} onClick={handleRegister}>新規登録</button>
          )}
          {user.error && <div style={{color:"red"}}>{user.error}</div>}
          {user.isAuthenticated && <div style={{color:"green"}}>ログイン成功</div>}
        </form>
      </div>
      <div className="p-8">
        { user.isAuthenticated && (
          <div>
            <p className="p-b-8">
              <button onClick={handleLogout}>ログアウト</button>
            </p>
          </div>
        )}
        <button onClick={getMemory}>getMemory</button>
      </div>
    </div>
  )
}
