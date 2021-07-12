import axios from "axios";
const BASE_URL = process.env.VUE_APP_BASE_URL;

const fetchProducts = (limit, offset) => {
  const endPoint = `${BASE_URL}/products?offset=${offset}&limit=${limit}`;

  return axios.get(endPoint);
};

const fetchProductById = idProduct => {
  const endPoint = BASE_URL + "/products/" + idProduct;

  return axios.get(endPoint);
};

const fetchProductByCategory = idCategory => {
  const endPoint =
    BASE_URL + "/categories/" + idCategory + "/products?limit=4&offset=0";

  return axios.get(endPoint);
};

export { fetchProducts, fetchProductById, fetchProductByCategory };
