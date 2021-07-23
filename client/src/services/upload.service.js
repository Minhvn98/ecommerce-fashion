import apiHandler from './base.service';

const uploadFile = data => {
  const endpoint = `/admin/product/upload`;

  return apiHandler.post(endpoint, data, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
};

export { uploadFile };
