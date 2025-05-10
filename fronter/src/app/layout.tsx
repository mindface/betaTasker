"use client"
import * as React from 'react'
import { Provider } from 'react-redux'
import App, { AppProps } from 'next/app'
import { setupStore } from '../store'

// const store = setupStore()
import IndexLayout from '../layout/IndexLayout'

import '../styles/style.sass'


export default function BaseApp ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <Provider store={setupStore}>
      <IndexLayout>
      <div>
        {children}
      </div>
      </IndexLayout>
    </Provider>
  )
}
