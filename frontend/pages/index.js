import { useState } from 'react'
import LoginForm from '../components/LoginForm'
import RegisterForm from '../components/RegisterForm'

export default function Home() {
  const [showLogin, setShowLogin] = useState(true)

  return (
    <div className="max-w-md mx-auto mt-20 p-6 shadow-md border rounded-md">
      {showLogin ? (
        <LoginForm switchToRegister={() => setShowLogin(false)} />
      ) : (
        <RegisterForm switchToLogin={() => setShowLogin(true)} />
      )}
    </div>
  )
}
