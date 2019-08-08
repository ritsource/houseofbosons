import React from 'react';
import { connect } from 'react-redux';

import { createPost, readPosts } from '../actions/post_actions';

import PostList from '../components/PostList';

class HomePage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			loading: true,
			errorMsg: false
		};
	}

	async componentDidMount() {
		try {
			await this.props.readPosts();
			console.log(this.props.posts);
		} catch (error) {
			console.log(error);
			this.setState({ errorMsg: error.message });
		}
		this.setState({ loading: false });
	}

	render() {
		return (
			<div>
				<h1>Posts</h1>
				<PostList posts={this.props.posts} />
			</div>
		);
	}
}

const mapStateToProps = ({ posts }) => ({ posts });

const mapDispatchToProps = (dispatch) => ({
	createPost: (...args) => dispatch(createPost(...args)),
	readPosts: (...args) => dispatch(readPosts(...args))
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
