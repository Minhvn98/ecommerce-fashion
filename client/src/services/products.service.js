import axios from "axios";
const BASE_URL = process.env.VUE_APP_BASE_URL;

const fetchProducts = () => {
  const endPoint = BASE_URL + "/products";

  return axios.get(endPoint);
};

const fetchProductById = idProduct => {
  const endPoint = BASE_URL + "/products/" + idProduct;

  return axios.get(endPoint);
};

const fetchProductByCategory = idCategory => {
  const endPoint = BASE_URL + "/products/categories/" + idCategory;

  return axios.get(endPoint);
};

export { fetchProducts, fetchProductById, fetchProductByCategory };
