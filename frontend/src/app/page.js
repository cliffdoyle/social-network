'use client';

import PrivacyToggle from "../components/privacyToggle";

export default function Home() {
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

