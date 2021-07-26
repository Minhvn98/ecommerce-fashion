import apiHandler from './base.service';

const loginHandler = (username, password) => {
  const endpoint = '/login';
  const payload = {
    username,
    password
  };
  return apiHandler.post(endpoint, payload);
};

const registerHandler = (email, username, password, comfirm_password) => {
  const endpoint = '/register';

  const payload = {
    email,
    username,
    password,
    comfirm_password
  };

  return apiHandler.post(endpoint, payload);
};

const getUserByCookie = () => {
  const endpoint = '/user/token';
  return apiHandler.get(endpoint);
};

const logout = () => {
  const endpoint = '/logout';
  return apiHandler.get(endpoint);
};

export { loginHandler, registerHandler, getUserByCookie, logout };
