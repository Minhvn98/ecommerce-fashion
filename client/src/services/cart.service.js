import apiHandler from './base.service';

const getCart = () => {
  const endpoint = '/cart';

  return apiHandler
    .get(endpoint)
    .then(res => res)
    .catch(err => err.response);
};

const addToCart = (product_id, quantity) => {
  const endpoint = '/cart';
  console.log(product_id, quantity);

  return apiHandler
    .put(endpoint, { product_id, quantity })
    .then(res => res)
    .catch(err => err.response);
};

const deleteProductCart = product_id => {
  const endpoint = '/cart';
  const quantity = 0;
  console.log(product_id);

  return apiHandler.put(endpoint, { product_id, quantity });
};

const changeQuantityProductCart = (product_id, quantity) => {
  const endpoint = '/cart';
  console.log(product_id, quantity);

  return apiHandler.put(endpoint, { product_id, quantity });
};

export { addToCart, getCart, deleteProductCart, changeQuantityProductCart };
