import apiHandler from './base.service';
// import axios from 'axios';
// const BASE_URL = process.env.VUE_APP_BASE_URL;

const loginHandler = async (username, password) => {
  console.log(username, password);
  const endpoint = '/login';
  const body = {
    username,
    password
  };
  return await apiHandler.post(endpoint, body);
};


const registerHandler = async (email, username, password, comfirm_password) => {
  const endpoint = '/register'
  const body = {
    email,
    username,
    password,
    comfirm_password
  }
  
  return apiHandler.post(endpoint, body)
}
export { loginHandler, registerHandler };
