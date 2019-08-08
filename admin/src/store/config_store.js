import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import api from '../api';

import authReducer from '../reducers/auth_reducer';

export default () => {
	const store = createStore(
		combineReducers({
			auth: authReducer
		}),
		applyMiddleware(thunk.withExtraArgument(api))
	);

	return store;
};
