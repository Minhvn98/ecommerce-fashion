import apiHandler from './base.service';

const checkoutHandler = (
  first_name,
  last_name,
  email,
  phone,
  location,
  payment
) => {
  const endpoint = '/checkout';
  
  const payload = {
    first_name,
    last_name,
    email,
    phone,
    location,
    payment
  };

  return apiHandler.post(endpoint, payload);
};

export { checkoutHandler };
