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

const getUserByCookie = () => {
  const endpoint = '/user/token';
  return apiHandler.get(endpoint);
};

const logout = () => {
  const endpoint = '/logout';
  return apiHandler.get(endpoint);
};

export { loginHandler, registerHandler, getUserByCookie, logout };
