// frontend/components/PrivacyToggle.js
import { useState } from "react";

export default function PrivacyToggle({ initialPrivacy, onChange }) {
  const [isPrivate, setIsPrivate] = useState(initialPrivacy);

  const handleToggle = () => {
    setIsPrivate((prev) => !prev);
    onChange && onChange(!isPrivate);
  };

  return (
    <label>
      <input
        type="checkbox"
        checked={isPrivate}
        onChange={handleToggle}
      />
      {isPrivate ? "Private Profile" : "Public Profile"}
    </label>
  );
}