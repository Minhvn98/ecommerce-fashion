import apiHandler from './base.service';

const getCategories = async () => {
  return await apiHandler.get('/categories');
};

export {  getCategories };
