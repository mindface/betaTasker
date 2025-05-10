"use client"
import * as React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'
import SectionBetaTools from '../components/SectionBetaTools'

const betaTools: NextPage = () => {
  return (
    <div className="beta-tools">
      <Head>
        <title key="title">BetaTools page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <SectionBetaTools />
    </div>
  )
}

export default betaTools
