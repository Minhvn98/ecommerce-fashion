import axios from "axios";
const BASE_URL = process.env.VUE_APP_BASE_URL;

const loginHandler = (username, password) => {
  console.log(username, password);
  axios
    .post(
      BASE_URL + "/login",
      {
        username,
        password
      },
      { withCredentials: true }
    )
    .then(function(response) {
      console.log(response["Set-Cookie"]);
    })
    .catch(function(error) {
      console.log(error);
    });
};

export { loginHandler };
