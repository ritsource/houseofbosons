import { CREATE_TOPIC, DELETE_TOPIC, EDIT_TOPIC, READ_ALL_TOPICS } from './action_types';

export const createTopic = (title) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.post('/topic/new', { title });
			dispatch({ type: CREATE_TOPIC, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};

export const readTopics = () => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.get('/topic/all');
			dispatch({ type: READ_ALL_TOPICS, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};

export const editTopic = (topicid, title) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.put('/topic/edit?id=' + topicid, { title });
			dispatch({ type: EDIT_TOPIC, data: resp.data });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};

export const deleteTopic = (topicid) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const resp = await api.delete('/topic/delete?id=' + topicid);
			dispatch({ type: DELETE_TOPIC, topicid: topicid });
			resolve(resp.data);
		} catch (error) {
			reject(error);
		}
	});
};
