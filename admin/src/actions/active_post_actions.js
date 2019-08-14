import { READ_SINGLE_POST, EDIT_POST, DELETE_POST_PERM, DELETE_POST_TEMP } from './action_types';

export const readSinglePost = (postid) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.get('/post/single?id=' + postid);
			dispatch({ type: READ_SINGLE_POST, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			console.log(error.message);
			reject(error);
		}
	});
};

export const editPost = (postid, data) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			delete data['_id'];
			const resp = await api.put('/post/edit?id=' + postid, data);
			dispatch({ type: EDIT_POST, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};

export const deletePostTemp = (postid) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.delete('/post/delete/temp?id=' + postid);
			dispatch({ type: DELETE_POST_TEMP, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};

export const deletePostPerm = (postid) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.delete('/post/delete/perm?id=' + postid);
			dispatch({ type: DELETE_POST_PERM, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};
