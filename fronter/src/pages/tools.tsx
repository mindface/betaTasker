"use client"
import * as React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'
import SectionTools from '../components/SectionTools'
import ContainersBase from '../containers/ContainersBase'

const Tools: NextPage = () => {
  return (
    <div className="about">
      <Head>
        <title key="title">index page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <SectionTools />
    </div>
  )
}

export default Tools
