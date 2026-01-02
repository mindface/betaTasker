"use client"
import * as React from 'react'
import { usePathname } from 'next/navigation'

// const store = setupStore()
import CoreLayout from './layouts/CoreLayout'

import '../styles/style.sass'


export default function Layout ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const path = usePathname();
  console.log('layout path:', path);
  return (
    <html lang="ja">
      <body>
        <CoreLayout>
          {children}
        </CoreLayout>
      </body>
    </html>
  )
}
