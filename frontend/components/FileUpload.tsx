import React, { useState } from 'react';
import axiosInstance from '../utils/axiosInstance'; 

const FileUpload: React.FC = () => {
  const [fileName, setFileName] = useState<string>('');

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setFileName(file.name);
    }
  };

  const handleFileUpload = async () => {
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    const file = fileInput?.files?.[0];
    if (!file) return;

    const formData = new FormData();
    formData.append('file', file);

    try {
      await axiosInstance.post('/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      setFileName(''); // Clear the file name after upload
      fileInput.value = ''; // Clear the file input after upload
    } catch (error) {
      console.error('Error uploading file', error);
    }
  };

  return (
    <div>
      <input
        id="fileInput"
        type="file"
        onChange={handleFileChange}
        accept="*/*"
      />
      <button onClick={handleFileUpload} disabled={!fileName}>
        Upload
      </button>
      {fileName && <p>Selected file: {fileName}</p>}
    </div>
  );
};

export default FileUpload;
