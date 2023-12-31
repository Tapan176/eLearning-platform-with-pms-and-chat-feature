// api.js - API service for making HTTP requests to the backend

import axios from 'axios';

const baseURL = 'https://localhost:8080'; // Change this to your backend API URL

const api = axios.create({
  baseURL,
});

// Interceptors for handling authentication, headers, etc. can be added here

export default api;
