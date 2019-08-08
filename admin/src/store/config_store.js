import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import api from '../api';

import authReducer from '../reducers/auth_reducer';
import postReducer from '../reducers/post_reducer';
import activePostReducer from '../reducers/active_post_reducer';

export default () => {
	const store = createStore(
		combineReducers({
			auth: authReducer,
			posts: postReducer,
			activePost: activePostReducer
		}),
		applyMiddleware(thunk.withExtraArgument(api))
	);

	return store;
};
