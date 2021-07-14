import apiHandler from './base.service';
// import axios from 'axios';
// const BASE_URL = process.env.VUE_APP_BASE_URL;

const loginHandler = async (username, password) => {
  console.log(username, password);
  axios
    .post(
      BASE_URL + '/login',
      {
        username,
        password
      },
      { withCredentials: true }
    )
    .then(function(response) {
      console.log(response['Set-Cookie']);
    })
    .catch(function(error) {
      console.log(error);
    });
};

const registerHandler = async (email, username, password, comfirm_password) => {
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
