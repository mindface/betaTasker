"use client"
import React, { useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { RootState } from "../store"
import { loginRequest, loginSuccess, loginFailure } from "../features/user/userSlice"
import { loginApi, logoutApi, regApi } from "../services/authApi"
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
    try {
      const result = await loginApi(email, password)
      if ('error' in result) {
        dispatch(loginFailure(result.error.message || 'ログイン失敗'))
      } else if('token' in result) {
        dispatch(loginSuccess({ token: result.token, user: result.user }))
        router.push("/")
      }
    } catch (ex) {
      console.error("error")
    }
  }

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault()
    const result = await regApi({ username: userName, email, password, role: "user" })
    if ('error' in result) {
      dispatch(loginFailure(result.error.message || '登録失敗'))
    } else {
      dispatch(loginSuccess({ token: result.token, user: result.user }))
      router.push("/")
    }
  }

  const handleLogout = async () => {
    const result = await logoutApi()
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
          {!loginSwitch && <div className="p-b-8">
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
      </div>
    </div>
  )
}
