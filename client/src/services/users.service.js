import axios from "axios";
const BASE_URL = process.env.VUE_APP_BASE_URL;

const loginHandler = (username, password) => {
  console.log(username, password);
  axios
    .post(BASE_URL + "/login", {
      username,
      password
    })
    .then(function(response) {
      console.log(response);
    })
    .catch(function(error) {
      console.log(error);
    });
};

export { loginHandler };
