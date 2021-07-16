import apiHandler from './base.service';

const getCategories = () => {
  return apiHandler.get('/categories');
};

export { getCategories };
