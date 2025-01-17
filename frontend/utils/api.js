import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

const api = {
  buyProduct: (productValue) =>
    axios.post(`${API_BASE_URL}/purchase`, { productValue }),
  getCashback: (address) =>
    axios.get(`${API_BASE_URL}/cashback?address=${address}`),
};

export default api;