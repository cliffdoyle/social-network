// components/RegisterForm.js
import { useState } from 'react'

// Component to render the user registration form
export default function RegisterForm({ switchToLogin }) {
  // Form state to hold all input values
  const [form, setForm] = useState({
    firstName: '',
    lastName: '',
    dob: '',
    email: '',
    password: '',
    confirmPassword: '',
    nickname: '',
    about: '',
    avatar: null,
  })

  // Error message state
  const [error, setError] = useState('')

  // Handle input changes, including file uploads
  const handleChange = (e) => {
    const { name, value, files } = e.target
    setForm({
      ...form,
      [name]: files ? files[0] : value,
    })
  }

  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault()

    // Check if password and confirm password match
    if (form.password !== form.confirmPassword) {
      return setError("Passwords do not match")
    }

    // Create FormData object for sending multipart/form-data
    const data = new FormData()
    Object.entries(form).forEach(([key, value]) => {
      if (value) data.append(key, value)
    })

    // Send registration request to backend
    const res = await fetch('http://localhost:8080/register', {
      method: 'POST',
      body: data,
      credentials: 'include',
    })

    // Handle response
    if (res.ok) {
      alert("Registration successful")
      switchToLogin()
    } else {
      const d = await res.json()
      setError(d.message || "Registration failed")
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-3">
      <h2 className="text-xl font-bold">Register</h2>

      {/* Input fields for required and optional info */}
      <input name="firstName" onChange={handleChange} placeholder="First Name" required className="input" />
      <input name="lastName" onChange={handleChange} placeholder="Last Name" required className="input" />
      <input type="date" name="dob" onChange={handleChange} required className="input" />
      <input type="email" name="email" onChange={handleChange} placeholder="Email" required className="input" />
      <input type="password" name="password" onChange={handleChange} placeholder="Password" required className="input" />
      <input type="password" name="confirmPassword" onChange={handleChange} placeholder="Confirm Password" required className="input" />

      {/* Optional fields */}
      <input type="file" name="avatar" onChange={handleChange} className="input" />
      <input name="nickname" onChange={handleChange} placeholder="Nickname (optional)" className="input" />
      <textarea name="about" onChange={handleChange} placeholder="About Me (optional)" className="input"></textarea>

      {/* Error message display */}
      {error && <p className="text-red-500">{error}</p>}

      {/* Submit button and switch link */}
      <button className="btn">Register</button>
      <p>
        Already have an account?{' '}
        <span onClick={switchToLogin} className="text-blue-500 cursor-pointer">
          Login
        </span>
      </p>
    </form>
  )
}
