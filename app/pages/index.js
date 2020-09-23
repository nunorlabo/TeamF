import Head from 'next/head'
import styles from '../styles/Home.module.scss'
import Layout from '../components/Layout'
import { createRef, useContext, useEffect, useState } from 'react'
import { MainContext } from '../contexts/MainContext'
import { PageTransition } from '../components/PageTransition'
import { useRouter } from 'next/router'
import SPCanvas from '../components/SPCanvas'
import DokodemoInput from '../components/DokodemoInput'
import { TwitterShareButton } from 'react-share'
import { postComment } from '../lib/api'
import { useToasts } from 'react-toast-notifications'
/**
 * ホーム画面のコンポーネント
 * @author Takahiro Nishino
 */
const Home = () => {
  const {
    selectedGenre,
    mode,
    setMode,
    shouldUpdate,
    setShouldUpdate,
    dialogID,
    setDialogID,
    dialog,
    setDialog,
    setSelectedGenre,
    cameBack,
    setCameBack,
    inputOpen,
    setInputOpen
  } = useContext(MainContext)
  // setInterval(() => console.log(mode), 1000)

  const [inputContents, setInputContents] = useState({})
  const [submitting, setSubmitting] = useState(false)
  const [comment, setComment] = useState('')
  const router = useRouter()
  const { addToast } = useToasts()
  if (location.hash.split('#')[1]) {
    // setMode('detail')
  } else {
    // setMode('home')
  }

  /**
   *
   * @type {Array.<{multipleLines: boolean, top: number, left: number, fontSize: number, width: number, height: number, type: string}>} inputVars
   */
  const initialValue = {
    multipleLines: true,
    bottom: 100,
    left: '50%',
    fontSize: 14,
    width: 370,
    height: 100,
    type: 'コメントを投稿しよう！'
  }

  useEffect(() => {
    if (mode === 'home') {
      // もどる
      setInputOpen(false)
    } else if (mode === 'new') {
      // 新規投稿start
      setInputOpen(true)
    }
  }, [mode])

  const submitPost = () => {
    setSubmitting(true)
    console.log('submit!')
    console.log(comment)
    const dialogID = location.hash.split('/')[1]
    console.log(dialogID)
    // submit here
    try {
      postComment(dialogID, comment).then((res) => console.log(res))
    } catch (error) {
      addToast(`サーバーと通信ができませんでした`, { appearance: 'error' })
    }
    console.log('about to post')
    setInputOpen(false)
    setSubmitting(false)
    setComment('')
  }

  const updateInputContent = (type, text) => {
    let old = { ...inputContents }
    old[type] = text
    setInputContents(old)
  }

  return (
    <Layout>
      <div className={styles.container}>
        <Head>
          <title>Home | ScenePicks</title>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <SPCanvas
          selectedGenre={selectedGenre}
          shouldUpdate={shouldUpdate}
          setShouldUpdate={setShouldUpdate}
          router={router}
          dialogID={dialogID}
          setDialogID={setDialogID}
          mode={mode}
          dialog={dialog}
          setDialog={setDialog}
          setSelectedGenre={setSelectedGenre}
          cameBack={cameBack}
          setCameBack={setCameBack}
          setMode={setMode}
        />
        {/* <PageTransition> */}
        {inputOpen && (
          <div>
            {
              <>
                <PageTransition>
                  <DokodemoInput
                    comment={comment}
                    setComment={setComment}
                    {...initialValue}
                    updateInputContent={updateInputContent}
                    submitting={submitting}
                  />
                  <button
                    className={styles.submitButton}
                    onClick={submitPost}
                    type="button"
                    style={{ bottom: `40px`, left: `calc(50%)`, transform: 'translateX(-50%)' }}>
                    送信
                  </button>
                </PageTransition>
              </>
            }
          </div>
        )}
        {/* </PageTransition> */}
      </div>
      {mode === 'detail' && (
        <PageTransition>
          <TwitterShareButton
            url={`${process.env.NEXT_PUBLIC_ENDPOINT_URL}/api${router.asPath}`} // TODO: 自分自身
            className={styles.shareContainer}
            title="scenepicksでセリフをシェア！    " // TODO: dialog の本文など
            hashtags={['scenepicks']} // 考える
          >
            <img src="/twitter.svg" alt="twitter icon" className={styles.twitterIcon} />
          </TwitterShareButton>
        </PageTransition>
      )}
    </Layout>
  )
}

export default Home
