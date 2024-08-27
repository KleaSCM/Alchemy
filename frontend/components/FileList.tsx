
import React, { useEffect, useState } from 'react';
import axiosInstance from '../utils/axiosInstance';
import styles from '../styles/FileList.module.scss';

interface File {
  name: string;
}

const FileList: React.FC = () => {
  const [files, setFiles] = useState<File[]>([]);

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await axiosInstance.get('/files');
        setFiles(response.data.files);
      } catch (error) {
        console.error('Error fetching files:', error);
      }
    };



export default FileList;
