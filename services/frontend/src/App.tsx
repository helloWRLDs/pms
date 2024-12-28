import { useEffect } from 'react'
import useAuth from './hooks/useAuth'
import { Button } from './components/ui/Button'
import { Modal } from './components/ui/Modal'
import { useModal } from './hooks/useModal'
import { useLoading } from './hooks/useToast'
import { infoToast } from './utils/toast'
import { Header } from './components/ui/Header'
import { Gi3dGlasses } from "react-icons/gi";

function App() {

  const {isAuthenticated} = useAuth()
  useEffect(() => {
    console.log(isAuthenticated)
  }, [])

  const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

  const [someModal, toggleSomeModal, closeSomeModal] = useModal(false)

  return (
    <div className='container'>
      <Header logoURL='./src/assets/logo.png'>

      </Header>


      <Button size='md' color='primary-1' icon={Gi3dGlasses} onClick= {async() => {
        const { done } = useLoading('loading users...')
        await sleep(5000)
        done('fetched users', true)
      }}>Yo</Button>
      <Button size='md' onClick={ () => {
        toggleSomeModal()
        infoToast("user created")
      }}>Toggle Modal</Button>
      <Modal visible={someModal} title='Some modal title' onClose={closeSomeModal}>
        Some modal
      </Modal>
    </div>
  )
}

export default App
