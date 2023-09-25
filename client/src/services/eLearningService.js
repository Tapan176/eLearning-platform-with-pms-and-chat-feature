// eLearningService.js - Service for eLearning-related operations

import api from './api';

// Example function to fetch courses
export async function getCourses() {
  try {
    const response = await api.get('/elearning/courses');
    return response.data;
  } catch (error) {
    throw error;
  }
}

// Implement other eLearning-related functions as needed
