import React, { useState } from 'react';
import axios from 'axios';

function Upload(props) {
  const { isLoggedIn } = props;
  const [file, setFile] = useState(null);
  const [users, setUsers] = useState([]);
  const [uploadStatus, setUploadStatus] = useState(null);

  const onFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    
    const formData = new FormData();
    formData.append('file', file);

    const config = {
      headers: {
        'content-type': 'multipart/form-data'
      },
      withCredentials: true,
    };

    try {
      const response = await axios.post('http://localhost:8080/upload', formData, config);
      console.log(response.data);
      setUploadStatus(<p style={{color: 'green'}}>Successfully uploaded the file</p>);
    } catch (error) {
      console.error(error);
      setUploadStatus(<p style={{color: 'red'}}>Error uploading file</p>);
    }
  };

  return (
    <div>
      <h1>Upload</h1>
      {isLoggedIn ? (
        <>
          <p>You are logged in. Welcome back, let's upload some files!</p>
          <form onSubmit={onSubmit}>
            <input type="file" onChange={onFileChange} />
            <button type="submit">Upload</button>
          </form>
          {uploadStatus}
          <ul>
            {users.map((user) => (
              <li key={user.id}>{user.name}</li>
            ))}
          </ul>
        </>
      ) : (
        <p>Please log in to upload files.</p>
      )}
    </div>
  );
}

export default Upload;
