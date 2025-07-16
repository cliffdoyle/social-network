import { useState } from 'react'
import styles from './LoginForm.module.css'


// Component to render the user login form
export default function LoginForm({ switchToRegister }) {
  // State to manage user input and error message
  const [emailOrUsername, setEmailOrUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  // Handle form submission
  const handleLogin = async (e) => {
    e.preventDefault()

    // Send login credentials to backend via the end point
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ identifier: emailOrUsername, password }),
    })

    // If login successful, redirect to dashboard
    if (res.ok) {
      window.location.href = '/dashboard'
    } else {
      // Otherwise, show error message
      const d = await res.json()
      setError(d.message || 'Login failed')
    }
  }

  return (
    <form onSubmit={handleLogin} className={styles.formContainer}>
  <h2>Login</h2>

  {/* Input for username or email */}
  <input
    type="text"
    placeholder="Username or Email"
    value={emailOrUsername}
    onChange={(e) => setEmailOrUsername(e.target.value)}
    required
    className={styles.input}
  />

  {/* Input for password */}
  <input
    type="password"
    placeholder="Password"
    value={password}
    onChange={(e) => setPassword(e.target.value)}
    required
    className={styles.input}
  />

  {/* Error message */}
  {error && <p className={styles.errorMessage}>{error}</p>}

  {/* Submit button */}
  <button className={styles.btn}>Login</button>

  {/* Forgot password and register link */}
  <p className={styles.link}>Forgot Password?</p>

  <p>
    Don’t have an account?{' '}
    <span onClick={switchToRegister} className={styles.link}>
      Register
    </span>
  </p>
</form>
  )
}
