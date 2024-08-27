import React, { useEffect, useState } from 'react';
import axiosInstance from '../utils/axiosInstance';
import styles from '../styles/FileList.module.scss';

interface FileItem {
  id: string;
  name: string;
  size: number;
  type: string;
}

const FileList: React.FC = () => {
  const [files, setFiles] = useState<FileItem[]>([]);

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await axiosInstance.get('/files');
        setFiles(response.data);
      } catch (error) {
        console.error('Error fetching files', error);
      }
    };

    fetchFiles();
  }, []);

  const handleDelete = async (id: string) => {
    try {
      await axiosInstance.delete(`/files/${id}`);
      setFiles(files.filter(file => file.id !== id));
    } catch (error) {
      console.error('Error deleting file', error);
    }
  };

  return (
    <div className={styles.fileListContainer}>
      <h2 className={styles.title}>Uploaded Files</h2>
      <ul className={styles.fileList}>
        {files.map((file) => (
          <li key={file.id} className={styles.fileItem}>
            <p>{file.name}</p>
            <p>Size: {file.size} bytes</p>
            <p>Type: {file.type}</p>
            <button onClick={() => handleDelete(file.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FileList;
