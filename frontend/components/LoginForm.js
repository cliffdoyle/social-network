import { useState } from 'react'

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
    <form onSubmit={handleLogin} className="space-y-4">
      <h2 className="text-xl font-bold">Login</h2>

      {/* Input for username or email */}
      <input
        type="text"
        placeholder="Username or Email"
        value={emailOrUsername}
        onChange={(e) => setEmailOrUsername(e.target.value)}
        required
        className="input"
      />

      {/* Input for password */}
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
        className="input"
      />

      {/* Error message */}
      {error && <p className="text-red-500">{error}</p>}

      {/* Submit button */}
      <button className="btn">Login</button>

      {/* Forgot password and register link */}
      <p className="text-sm text-right text-blue-600 cursor-pointer">
        Forgot Password?
      </p>

      <p>
        Don't have an account?{' '}
        <span onClick={switchToRegister} className="text-blue-600 cursor-pointer">
          Register
        </span>
      </p>
    </form>
  )
}
