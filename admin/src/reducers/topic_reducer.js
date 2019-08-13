import { CREATE_TOPIC, DELETE_TOPIC, EDIT_TOPIC, READ_ALL_TOPICS } from '../actions/action_types';

export default (state = [], action) => {
	switch (action.type) {
		case CREATE_TOPIC:
			return [ ...state, action.data ];
		case READ_ALL_TOPICS:
			return action.data;
		case EDIT_TOPIC:
			return [ ...state.filter(({ _id }) => _id !== action.data._id), ...action.data ];
		case DELETE_TOPIC:
			return [ ...state.filter(({ _id }) => _id !== action.topicid) ];
		default:
			return state;
	}
};
