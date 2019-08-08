import { FETCH_ADMIN } from '../actions/action_types';

export default (state = null, action) => {
	switch (action.type) {
		case FETCH_ADMIN:
			return action.data;
		default:
			return state;
	}
};
