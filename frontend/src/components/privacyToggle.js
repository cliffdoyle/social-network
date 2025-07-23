'use client';

import { useState } from 'react';

export default function PrivacyToggle({ initialPrivacy = false, onChange }) {
  const [isPrivate, setIsPrivate] = useState(initialPrivacy);

  const handleToggle = () => {
    const newPrivacy = !isPrivate;
    setIsPrivate(newPrivacy);
    if (onChange) {
      onChange(newPrivacy);
    }
  };

  return (
    <div className="privacy-toggle">
      <label>
        <input
          type="checkbox"
          checked={isPrivate}
          onChange={handleToggle}
        />
        Private Profile
      </label>
      <p>Status: {isPrivate ? 'Private' : 'Public'}</p>
    </div>
  );
}