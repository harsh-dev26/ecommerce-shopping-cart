import axios from 'axios';

const API_BASE_URL = "https://ecommerce-backend.onrender.com";

const api = axios.create({
  baseURL: API_BASE_URL,
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const login = (username, password) => {
  return api.post('/users/login', { username, password });
};

export const getItems = () => {
  return api.get('/items');
};

export const addToCart = (item_ids) => {
  return api.post('/carts', { item_ids });
};

export const getCarts = () => {
  return api.get('/carts');
};

export const createOrder = (cart_id) => {
  return api.post('/orders', { cart_id });
};

export const getOrders = () => {
  return api.get('/orders');
};

export default api;