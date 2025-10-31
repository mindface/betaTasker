import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(req: NextRequest) {
  const cookieStore = await cookies();
  const token = cookieStore.get('token')?.value;
  const { searchParams } = new URL(req.url);
  const code = searchParams.get('code') || '';

  const res = await fetch(`http://localhost:8080/api/memory/aid/${code}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });
  const data = await res.json();
  return NextResponse.json({ contexts: data.contexts, status: 200 });
}
