import apiHandler from './base.service';

const getProducts = async (offset = 0, limit = 4) => {
  const endpoint = `/products?offset=${offset}&limit=${limit}`;

  return await apiHandler.get(endpoint);
};

const getProductById = idProduct => {
  const endpoint = `/products/${idProduct}`;

  return apiHandler.get(endpoint);
};

const getProductsByCategory = async (idCategory, offset = 0, limit = 8) => {
  const endpoint = `/categories/${idCategory}/products?offset=${offset}&limit=${limit}`;

  return await apiHandler.get(endpoint);
};

const searchProducts = async (keyword, offset = 0, limit = 8) => {
  const endpoint = `products/search?text=${keyword}&offset=${offset}&limit=${limit}`;
  
  return await apiHandler.get(endpoint);
};

export {
  getProducts,
  getProductsByCategory,
  getProductById,
  searchProducts
};
