import React from 'react';
import { connect } from 'react-redux';

import { createPost, readPosts } from '../actions/post_actions';

import PostList from '../components/PostList';
import CreatePostModal from '../components/CreatePostModal';

class HomePage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			loading: true,
			errorMsg: false,
			modal: true
		};
	}

	async fetchPosts() {
		try {
			await this.props.readPosts(20);
		} catch (error) {
			console.log(error);
			this.setState({ errorMsg: error.message });
		}
	}

	async componentDidMount() {
		await this.fetchPosts();
		this.setState({ loading: false });
	}

	render() {
		return (
			<React.Fragment>
				<CreatePostModal
					visible={this.state.modal}
					onClose={() => this.setState({ modal: false })}
					createPost={this.props.createPost}
				/>

				<div
					style={{
						maxHeight: 'calc(100vh - 100px)',
						// width: 'calc(100vw - 400px)',
						padding: '50px',
						minWidth: '700px',
						display: 'flex',
						flexDirection: 'column'
					}}
				>
					<div
						style={{
							margin: '10px 0px',
							display: 'flex',
							flexDirection: 'row',
							justifyContent: 'space-between'
						}}
					>
						<h2 style={{ padding: '0px', margin: '0px' }}>Posts</h2>
						{this.state.loading ? (
							<div className="Theme-Loading-Spin-Div" />
						) : (
							<button className="Theme-Btn" onClick={() => this.setState({ modal: true })}>
								Create Post
							</button>
						)}
					</div>
					{this.props.posts.length === 0 ? this.state.loading ? (
						<h3 style={{ textAlign: 'center' }}>Loading...</h3>
					) : (
						<h3 style={{ textAlign: 'center' }}>No Posts Found</h3>
					) : (
						<PostList posts={this.props.posts} />
					)}
					<div
						style={{
							display: 'flex',
							flexDirection: 'row',
							justifyContent: 'space-between',
							alignItems: 'center',
							padding: '10px 0px'
						}}
					>
						<button
							className="Theme-Btn"
							onClick={async () => {
								if (this.state.loading === false) {
									this.setState({ loading: true });
									await this.fetchPosts();
									this.setState({ loading: false });
								}
							}}
						>
							More
							{this.state.loading && '..'}
						</button>
					</div>
				</div>
			</React.Fragment>
		);
	}
}

const mapStateToProps = ({ posts }) => ({ posts });

const mapDispatchToProps = (dispatch) => ({
	createPost: (...args) => dispatch(createPost(...args)),
	readPosts: (...args) => dispatch(readPosts(...args))
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
