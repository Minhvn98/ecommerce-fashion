import axios from "axios";
const BASE_URL = process.env.VUE_APP_BASE_URL;

const fetchCategories = () => {
  const endPoint = BASE_URL + "/categories";

  return axios.get(endPoint);
};

export { fetchCategories };
