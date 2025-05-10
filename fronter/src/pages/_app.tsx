"use client"
import * as React from 'react'
import { Provider } from 'react-redux'
import { AppProps } from 'next/app'
import { setupStore } from '../store'

import IndexLayout from '../layout/IndexLayout'

import '../styles/style.sass'


export default function BaseApp ({ Component, pageProps }: AppProps) {
  return (
    <Provider store={setupStore}>
      <IndexLayout>
      <div>
        <Component {...pageProps} />
      </div>
      </IndexLayout>
    </Provider>
  )
}
