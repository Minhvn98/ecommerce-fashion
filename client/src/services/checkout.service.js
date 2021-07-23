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
  const body = {
    first_name,
    last_name,
    email,
    phone,
    location,
    payment
  };

  return apiHandler.post(endpoint, body);
};
export { checkoutHandler };
