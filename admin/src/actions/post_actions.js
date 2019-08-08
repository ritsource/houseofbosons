import { CREATE_POST, READ_ALL_POST } from './action_types';

/*
post data holds default values
for creating a new post
*/
const postdata = {
	id_str: '',
	title: 'Title',
	description: '',
	formatted_date: '',
	doc_type: 0,
	md_src: '',
	html_src: '',
	thumbnail: '',
	created_at: '',
	topics: [],
	is_featured: false,
	is_public: false,
	is_deleted: false,
	is_series: false
};

export const createPost = (id_str) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.post('/post/all', { ...postdata, id_str });
			dispatch({ type: CREATE_POST, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const readPosts = (limit = 10) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const state = getState();
			const skip = state.posts.length;

			const resp = await api.get(`/post/all?skip=${skip}&limit=${limit}`);
			dispatch({ type: READ_ALL_POST, skip: skip, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};
