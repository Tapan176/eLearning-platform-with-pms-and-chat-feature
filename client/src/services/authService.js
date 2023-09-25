// authService.js - Authentication service for handling user authentication

import api from './api';

// Example function to log in a user
export async function login(username, password) {
  try {
    const response = await api.post('/auth/login', { username, password });
    return response.data;
  } catch (error) {
    throw error;
  }
}

// Example function to log out a user
export async function logout() {
  try {
    const response = await api.post('/auth/logout');
    return response.data;
  } catch (error) {
    throw error;
  }
}

// Implement other authentication-related functions as needed
