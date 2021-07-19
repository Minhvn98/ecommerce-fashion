import apiHandler from './base.service';

const getProducts = (offset = 0, limit = 4) => {
  const endpoint = `/products?offset=${offset}&limit=${limit}`;

  return apiHandler.get(endpoint);
};

const getProductById = idProduct => {
  const endpoint = `/products/${idProduct}`;

  return apiHandler.get(endpoint);
};

const getProductsByCategory = (idCategory, offset = 0, limit = 8) => {
  const endpoint = `/categories/${idCategory}/products?offset=${offset}&limit=${limit}`;

  return apiHandler.get(endpoint);
};

const searchProducts = (keyword, offset = 0, limit = 8) => {
  const endpoint = `products/search?text=${keyword}&offset=${offset}&limit=${limit}`;

  return apiHandler.get(endpoint);
};

export { getProducts, getProductsByCategory, getProductById, searchProducts };
