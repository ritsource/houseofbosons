import { FETCH_ADMIN } from './action_types';

export const fetchAdmin = () => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.post('/auth/current_user');
			dispatch({ type: FETCH_ADMIN, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};
