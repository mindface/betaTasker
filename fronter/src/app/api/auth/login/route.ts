import { NextRequest, NextResponse } from 'next/server';

export async function POST(req: NextRequest) {
  const { email, password } = await req.json();
  console.log('Login request received:', { email, password });
  if (!email || !password) {
    return NextResponse.json(
      { message: 'Email and password are required' },
      { status: 400 }
    );
  }

  const res = await fetch('http://localhost:8080/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  const data = await res.json();
  console.log('Login response from backend:', data);

  if (res.ok && data.token) {
    const response = NextResponse.json({ message: 'Login successful', token: data.token, user: data.user }, { status: 200 });

    response.cookies.set({
      name: 'token',
      value: data.token,
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production', // 開発中は false でも可
      sameSite: 'strict',
      path: '/',
      maxAge: 60 * 60 * 24,
    });

    return response;
  }

  return NextResponse.json(
    { message: 'Invalid credentials or login failed' },
    { status: 401 }
  );
}
