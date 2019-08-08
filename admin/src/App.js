import React from 'react';
import { connect } from 'react-redux';
import { fetchAdmin } from './actions/auth_actions';

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
			console.log('x', this.props.auth);
		} catch (error) {
			console.log(error);
		}
		this.setState({ loading: false });
	}

	render() {
		return (
			<div className="App">
				<header>
					{this.state.loading ? (
						<h1>Loading...</h1>
					) : this.props.auth ? (
						<h1>Hello world!</h1>
					) : (
						<h1>Unauth</h1>
					)}
				</header>
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
