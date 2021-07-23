import apiHandler from './base.service';

const getOrderById = id => {
  const endpoint = `/order-detail/${id}`;
  return apiHandler.get(endpoint);
};

const getOrdersByStatus = status => {
  const endpoint = `/order?status=${status}`;
  return apiHandler.get(endpoint);
};

const getAllOrder = () => {
  const endpoint = `/orders`;
  return apiHandler.get(endpoint);
};

const updateStatusOrder = (id, status) => {
  const endpoint = `/order/${id}`;
  return apiHandler.put(endpoint, { status: status });
};
export { getOrderById, getOrdersByStatus, getAllOrder, updateStatusOrder };
