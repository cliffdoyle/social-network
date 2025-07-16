// components/RegisterForm.js
import { useState } from 'react'
import styles from './RegisterForm.module.css'


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
    <form onSubmit={handleSubmit} className={styles.formContainer}>
    <h2>Register</h2>
  
    <input name="firstName" onChange={handleChange} placeholder="First Name" required className={styles.input} />
    <input name="lastName" onChange={handleChange} placeholder="Last Name" required className={styles.input} />
    <input type="date" name="dob" onChange={handleChange} required className={styles.input} />
    <input type="email" name="email" onChange={handleChange} placeholder="Email" required className={styles.input} />
    <input type="password" name="password" onChange={handleChange} placeholder="Password" required className={styles.input} />
    <input type="password" name="confirmPassword" onChange={handleChange} placeholder="Confirm Password" required className={styles.input} />
    <input type="file" name="avatar" onChange={handleChange} className={styles.input} />
    <input name="nickname" onChange={handleChange} placeholder="Nickname (optional)" className={styles.input} />
    <textarea name="about" onChange={handleChange} placeholder="About Me (optional)" className={styles.input}></textarea>
  
    {error && <p className={styles.errorMessage}>{error}</p>}
  
    <button className={styles.btn}>Register</button>
    <p>
      Already have an account?{' '}
      <span onClick={switchToLogin} className={styles.textLink}>
        Login
      </span>
    </p>
  </form>
  )
}
