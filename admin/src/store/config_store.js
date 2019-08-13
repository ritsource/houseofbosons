import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import api from '../api';

import authReducer from '../reducers/auth_reducer';
import postReducer from '../reducers/post_reducer';
import activePostReducer from '../reducers/active_post_reducer';
import topicReducer from '../reducers/topic_reducer';

export default () => {
	const store = createStore(
		combineReducers({
			auth: authReducer,
			posts: postReducer,
			activePost: activePostReducer,
			topics: topicReducer
		}),
		applyMiddleware(thunk.withExtraArgument(api))
	);

	return store;
};
