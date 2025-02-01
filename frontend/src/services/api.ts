import axios from 'axios';

const api = axios.create({
  baseURL: '/api/v1',
});

export const fetchBlocks = async () => {
  const { data } = await api.get('/collector/blocks');
  return data;
};

export const fetchTransactions = async () => {
  const { data } = await api.get('/collector/transactions');
  return data;
};

export const fetchStatus = async () => {
  const { data } = await api.get('/collector/status');
  return data;
}; 