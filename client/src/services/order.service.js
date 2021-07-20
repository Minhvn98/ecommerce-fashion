import apiHandler from './base.service';

const getOrderById = id => {
  const endpoint = `/order-detail/${id}`;
  return apiHandler.get(endpoint);
};

const getOrdersByStatus = status => {
  const endpoint = `/order?status=${status}`;
  return apiHandler.get(endpoint);
};

export { getOrderById, getOrdersByStatus };
