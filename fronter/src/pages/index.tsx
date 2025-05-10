"use client"
import * as React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'

import ContainersHome from '../containers/ContainersHome'
import SectionHome from '../components/SectionHome'

const Index:NextPage = () => {
  return (
    <div>
      <Head>
        <title key="title">Task Link</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <ContainersHome>
        <SectionHome />
      </ContainersHome>
    </div>
  )
}

export default Index
