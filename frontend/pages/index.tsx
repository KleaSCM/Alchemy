import React from 'react';
import FileUpload from '../components/FileUpload';
import styles from '../styles/Home.module.scss';

const HomePage: React.FC = () => {
  return (
    <div className={styles.container}>
      <h1>File Management System</h1>
      <FileUpload />
  
    </div>
  );
};

export default HomePage;