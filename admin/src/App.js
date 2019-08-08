import React from 'react';
import { connect } from 'react-redux';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { fetchAdmin } from './actions/auth_actions';
import { FaGoogle } from 'react-icons/fa';
import { serverAdd } from './api';

import LoginBtn from './components/LoginBtn';
import HomePage from './pages/HomePage';
import TopicsPage from './pages/TopicsPage';
import EachPostPage from './pages/EachPostPage';
import NotFoundPage from './pages/NotFoundPage';

class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			loading: true
		};
	}

	async componentDidMount() {
		try {
			await this.props.fetchAdmin();
		} catch (error) {
			console.log(error);
		}
		this.setState({ loading: false });
	}

	render() {
		return (
			<div
				style={{
					minHeight: '100vh',
					display: 'flex',
					flexDirection: 'row',
					justifyContent: 'center',
					alignItems: 'center'
				}}
			>
				<BrowserRouter>
					{this.state.loading ? (
						<h1>Loading...</h1>
					) : this.props.auth ? (
						<Switch>
							<Route path="/" exact component={HomePage} />
							<Route path="/post/:postid" exact component={EachPostPage} />
							<Route path="/topics" exact component={TopicsPage} />
							<Route component={NotFoundPage} />
						</Switch>
					) : (
						<a href={serverAdd + '/api/auth/google'}>
							<LoginBtn>
								<FaGoogle style={{ margin: '0px 10px -2px 0px' }} />
								Login as admin
							</LoginBtn>
						</a>
					)}
				</BrowserRouter>
			</div>
		);
	}
}

const mapStateToProps = (state) => ({
	auth: state.auth
});

const mapDispatchToProps = (dispatch) => ({
	fetchAdmin: () => dispatch(fetchAdmin())
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
