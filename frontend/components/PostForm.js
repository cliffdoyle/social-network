import React, { useState } from 'react';

const PostForm = ({ onPostCreated }) => {
  // Post type state
  const [postType, setPostType] = useState('text'); 

  // Form fields
  const [title, setTitle] = useState('');
  const [text, setText] = useState('');
  const [link, setLink] = useState('');
  const [file, setFile] = useState(null);

  // UI states
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  const [loading, setLoading] = useState(false);

  // Privacy settings
  const [privacy, setPrivacy] = useState('public');
  const [audience, setAudience] = useState('');

  const handleFileChange = (e) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      // file type validation
      const validTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif'];
      if (!validTypes.includes(selectedFile.type)) {
        setError('Only JPEG, PNG, and GIF images are allowed.');
        setFile(null);
        return;
      }

      // Check file size (limit to 10MB)
      const maxSize = 10 * 1024 * 1024; // 10MB in bytes
      if (selectedFile.size > maxSize) {
        setError('File size must be less than 10MB.');
        setFile(null);
        return;
      }

      setError(null);
      setSuccess(null);
      setFile(selectedFile);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);
    setLoading(true);

    // Validate required fields
    if (!title.trim()) {
      setError('Title is required');
      setLoading(false);
      return;
    }

    // Validate based on post type
    if (postType === 'text' && !text.trim()) {
      setError('Text content is required for text posts');
      setLoading(false);
      return;
    }

    if (postType === 'link' && !link.trim()) {
      setError('URL is required for link posts');
      setLoading(false);
      return;
    }

    if (postType === 'image' && !file) {
      setError('Image is required for image posts');
      setLoading(false);
      return;
    }

    // Validate private post audience
    if (privacy === 'private' && !audience.trim()) {
      setError('Private posts must specify an audience');
      setLoading(false);
      return;
    }

    // Handle different content types
    let content = null;
    let mediaUrl = null;
    let mediaType = null;

    if (postType === 'text') {
      content = text.trim();
    } else if (postType === 'link') {
      content = link.trim();
    } else if (postType === 'image' && file) {
      // show error since upload isn't implemented
      setError('Image upload is not yet implemented. Please use text or link posts for now.');
      setLoading(false);
      return;
    }

    try {
      // Send post data as JSON 
      const postData = {
        groupId: null,
        title: title.trim() || null,
        content: content ? content.trim() : null,
        mediaUrl: mediaUrl || null,
        mediaType: mediaType || null,
        privacy,
        audience: privacy === 'private' ? audience.split(',').map(id => id.trim()).filter(id => id) : [],
      };

      const res = await fetch('/api/create-post', {
        method: 'POST',
        credentials: 'include', // Include session cookie
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(postData),
      });
 

      if (!res.ok) {
        const errorData = await res.json().catch(() => ({}));

        // Handle validation errors from backend
        if (res.status === 422) {
          throw new Error(errorData.message || 'Validation failed');
        }

        throw new Error(errorData.message || 'Failed to create post');
      }

      const responseData = await res.json();

      // Success - reset form
      setTitle('');
      setText('');
      setLink('');
      setFile(null);
      setAudience('');
      setPrivacy('public');
      setPostType('text');
      setSuccess('Post created successfully!');

      // Call callback 
      if (onPostCreated) {
        onPostCreated();
      }

    } catch (err) {
      setError(err.message || 'Error creating post.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{
      border: '1px solid #ccc',
      borderRadius: '8px',
      backgroundColor: 'white',
      overflow: 'hidden'
    }}>
      {/* Header */}
      <div style={{
        padding: '16px',
        borderBottom: '1px solid #eee',
        backgroundColor: '#f8f9fa'
      }}>
        <h3 style={{ margin: 0, fontSize: '18px', fontWeight: '600' }}>
          Create a post
        </h3>
      </div>

      {/* Post Type Tabs */}
      <div style={{
        display: 'flex',
        borderBottom: '1px solid #eee',
        backgroundColor: '#f8f9fa'
      }}>
        {[
          { type: 'text', label: '📝 Text', icon: '📝' },
          { type: 'link', label: '🔗 Link', icon: '🔗' },
          { type: 'image', label: '📷 Image', icon: '📷' }
        ].map(({ type, label, icon }) => (
          <button
            key={type}
            type="button"
            onClick={() => setPostType(type)}
            style={{
              flex: 1,
              padding: '12px 16px',
              border: 'none',
              backgroundColor: postType === type ? 'white' : 'transparent',
              borderBottom: postType === type ? '2px solid #0079d3' : '2px solid transparent',
              cursor: 'pointer',
              fontWeight: postType === type ? '600' : '400',
              color: postType === type ? '#0079d3' : '#666'
            }}
            disabled={loading}
          >
            {label}
          </button>
        ))}
      </div>

      <form onSubmit={handleSubmit} style={{ padding: '16px' }}>
        {/* Title Input */}
        <div style={{ marginBottom: '16px' }}>
          <input
            type="text"
            value={title}
            onChange={e => setTitle(e.target.value)}
            placeholder="Title"
            style={{
              width: '100%',
              padding: '12px',
              border: '1px solid #ccc',
              borderRadius: '4px',
              fontSize: '16px',
              fontWeight: '500'
            }}
            disabled={loading}
            required
          />
        </div>

        {/* Dynamic Content Area Based on Post Type */}
        {postType === 'text' && (
          <div style={{ marginBottom: '16px' }}>
            <textarea
              value={text}
              onChange={e => setText(e.target.value)}
              placeholder="Text (optional)"
              rows={6}
              style={{
                width: '100%',
                padding: '12px',
                border: '1px solid #ccc',
                borderRadius: '4px',
                fontSize: '14px',
                resize: 'vertical',
                fontFamily: 'inherit'
              }}
              disabled={loading}
            />
          </div>
        )}

        {postType === 'link' && (
          <div style={{ marginBottom: '16px' }}>
            <input
              type="url"
              value={link}
              onChange={e => setLink(e.target.value)}
              placeholder="Url"
              style={{
                width: '100%',
                padding: '12px',
                border: '1px solid #ccc',
                borderRadius: '4px',
                fontSize: '14px'
              }}
              disabled={loading}
            />
          </div>
        )}

        {postType === 'image' && (
          <div style={{ marginBottom: '16px' }}>
            <div style={{
              border: '2px dashed #ccc',
              borderRadius: '4px',
              padding: '40px',
              textAlign: 'center',
              backgroundColor: '#f8f9fa'
            }}>
              <input
                type="file"
                accept="image/jpeg,image/jpg,image/png,image/gif"
                onChange={handleFileChange}
                disabled={true}
                style={{ display: 'none' }}
                id="image-upload"
              />
              <label htmlFor="image-upload" style={{ cursor: 'not-allowed' }}>
                <div style={{ fontSize: '48px', marginBottom: '8px' }}>📷</div>
                <div style={{ color: '#999', fontSize: '14px' }}>
                  Image upload coming soon
                </div>
              </label>
            </div>
          </div>
        )}

        {/* Privacy Settings */}
        <div style={{ marginBottom: '16px' }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            gap: '12px',
            marginBottom: '8px'
          }}>
            <span style={{ fontSize: '14px', fontWeight: '500' }}>Privacy:</span>
            <select
              value={privacy}
              onChange={e => setPrivacy(e.target.value)}
              disabled={loading}
              style={{
                padding: '6px 12px',
                border: '1px solid #ccc',
                borderRadius: '4px',
                fontSize: '14px'
              }}
            >
              <option value="public"> Public</option>
              <option value="followers">Followers Only</option>
              <option value="private"> Private</option>
            </select>
          </div>

          {privacy === 'private' && (
            <input
              type="text"
              value={audience}
              onChange={e => setAudience(e.target.value)}
              placeholder="Specify audience (comma-separated user IDs)"
              disabled={loading}
              style={{
                width: '100%',
                padding: '8px 12px',
                border: '1px solid #ccc',
                borderRadius: '4px',
                fontSize: '14px',
                backgroundColor: '#f8f9fa'
              }}
            />
          )}
        </div>

        {/* Error and Success Messages */}
        {error && (
          <div style={{
            color: '#d93025',
            marginBottom: '16px',
            padding: '12px',
            backgroundColor: '#fce8e6',
            border: '1px solid #d93025',
            borderRadius: '4px',
            fontSize: '14px'
          }}>
            {error}
          </div>
        )}

        {success && (
          <div style={{
            color: '#137333',
            marginBottom: '16px',
            padding: '12px',
            backgroundColor: '#e6f4ea',
            border: '1px solid #137333',
            borderRadius: '4px',
            fontSize: '14px'
          }}>
            {success}
          </div>
        )}

        {/* Submit Button */}
        <div style={{
          display: 'flex',
          justifyContent: 'flex-end',
          paddingTop: '16px',
          borderTop: '1px solid #eee'
        }}>
          <button
            type="submit"
            disabled={loading || !title.trim() || (
              (postType === 'text' && !text.trim()) ||
              (postType === 'link' && !link.trim()) ||
              (postType === 'image' && !file)
            )}
            style={{
              padding: '8px 24px',
              backgroundColor: loading || !title.trim() ? '#ccc' : '#0079d3',
              color: 'white',
              border: 'none',
              borderRadius: '20px',
              cursor: loading || !title.trim() ? 'not-allowed' : 'pointer',
              fontSize: '14px',
              fontWeight: '600',
              minWidth: '80px'
            }}
          >
            {loading ? 'Posting...' : 'Post'}
          </button>
        </div>
      </form>
    </div>
  );
};

export default PostForm;