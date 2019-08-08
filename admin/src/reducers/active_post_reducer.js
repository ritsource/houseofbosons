import { DELETE_POST_PERM, DELETE_POST_TEMP, EDIT_POST, READ_SINGLE_POST } from '../actions/action_types';

export default (state = null, action) => {
	switch (action.type) {
		case READ_SINGLE_POST:
			return action.data;
		case EDIT_POST:
			return action.data;
		case DELETE_POST_TEMP:
			return action.data;
		case DELETE_POST_PERM:
			return action.data;
		default:
			return state;
	}
};
