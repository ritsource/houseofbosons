import axios from 'axios';

export const serverAdd = process.env.NODE_ENV === 'production' ? 'https://houseofbosons.com' : 'http://localhost:8080';

// Axios Instance
const api = axios.create({
	baseURL: serverAdd + '/api',
	headers: {
		'Content-Type': 'application/x-www-form-urlencoded',
		Accept: 'application/json'
	},
	withCredentials: true
});

export default api;
