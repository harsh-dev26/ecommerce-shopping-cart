import React, { useState } from 'react';
import { login } from '../services/api';
import './Login.css';

function Login({ onLoginSuccess }) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await login(username, password);
      const { token, user_id } = response.data;
      
      localStorage.setItem('token', token);
      localStorage.setItem('user_id', user_id);
      
      onLoginSuccess();
    } catch (error) {
      window.alert('Invalid username/password');
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2>Login to E-Commerce Store</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Username:</label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label>Password:</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="login-button">Login</button>
        </form>
      </div>
    </div>
  );
}

export default Login;