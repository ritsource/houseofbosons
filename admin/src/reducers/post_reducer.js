import { CREATE_POST, READ_ALL_POST } from '../actions/action_types';

export default (state = [], action) => {
	switch (action.type) {
		case CREATE_POST:
			return [ ...state, action.data ];
		case READ_ALL_POST:
			return [ ...state.slice(0, action.skip), ...action.data ];
		default:
			return state;
	}
};
