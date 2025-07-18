<<<<<<< HEAD
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
=======
// import Head from "next/head";
// import Image from "next/image";
// import { Geist, Geist_Mono } from "next/font/google";
// import styles from "@/styles/Home.module.css";
import PrivacyToggle from "../components/privacyToggle";

export default function Home() {
  // This function will be called when the toggle changes
  const handlePrivacyChange = (isPrivate) => {
    console.log("Privacy changed to:", isPrivate ? "Private" : "Public");
  };

  return (
    <div>
      <h1>Profile Privacy</h1>
      <PrivacyToggle initialPrivacy={false} onChange={handlePrivacyChange} />
    </div>
  );
}
>>>>>>> feat/auth
