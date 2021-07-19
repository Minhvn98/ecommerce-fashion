import apiHandler from './base.service';

const loginHandler = (username, password) => {
  console.log(username, password);
  const endpoint = '/login';
  const body = {
    username,
    password
  };
  return apiHandler.post(endpoint, body);
};

const registerHandler = (email, username, password, comfirm_password) => {
  const endpoint = '/register';
  const body = {
    email,
    username,
    password,
    comfirm_password
  };

  return apiHandler.post(endpoint, body);
};
export { loginHandler, registerHandler };
