import axios from 'axios';

const API_URL = import.meta.env.VITE_APP_API_URL;

export const signUp = async (username, email, country_code, phone, password, confirm_password) => {
  try {
    const {data: results} = await axios.post(`${API_URL}/auth/signup`, {
      username: username,
      email: email,
      country_code: country_code,
      phone: phone,
      password: password,
      confirm_password: confirm_password,
    });

    return results;
  } catch (error) {
    throw error;
  }
}

export const signIn = async (username, password) => {
  try {
    const {data: results} = await axios.post(`${API_URL}/auth/signin`, {
      username_or_email: username,
      password: password,
    });

    return results;
  } catch (error) {
    throw error;
  }
}