import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { Provider } from 'react-redux'
import { authStore } from './store/authStore.ts'
import { ToastContainer } from 'react-toastify'

createRoot(document.getElementById('root')!).render(
  <Provider store={authStore}>
    <StrictMode>
      <App />
      <ToastContainer />
    </StrictMode>,
  </Provider>
)
