import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  // 認証が必要なパス一覧
  const protectedPaths = [
    '/dashboard',
    '/task',
    '/memory',
    '/assessment',
    '/book',
    '/relation',
    '/tools'
  ];
  const { pathname } = request.nextUrl;

  // ログインページやAPIは除外
  if (pathname.startsWith('/login') || pathname.startsWith('/api')) {
    return NextResponse.next();
  }

  // 保護対象パスかどうか
  const isProtected = protectedPaths.some((path) => pathname.startsWith(path));
  if (!isProtected) {
    return NextResponse.next();
  }

  // cookieからtokenを取得
  const token = request.cookies.get('token')?.value;
  if (!token) {
    const loginUrl = new URL('/login', request.url);
    return NextResponse.redirect(loginUrl);
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    // 保護したいパスを列挙
    '/dashboard/:path*',
    '/task/:path*',
    '/memory/:path*',
    '/assessment/:path*',
    '/book/:path*',
    '/relation/:path*',
    '/tools/:path*',
  ],
};
