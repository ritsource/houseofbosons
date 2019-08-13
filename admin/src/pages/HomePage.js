import React from 'react';
import { connect } from 'react-redux';

import { createPost, readPosts } from '../actions/post_actions';
import { readTopics, createTopic, editTopic, deleteTopic } from '../actions/topic_actions';

import PostList from '../components/PostList';
import CreatePostModal from '../components/CreatePostModal';
import TopicManageModal from '../components/TopicManageModal';

class HomePage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			loading: true,
			errorMsg: false,
			idstrModal: false,
			topicVis: false
		};
	}

	async fetchPosts() {
		try {
			await this.props.readPosts(10);
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
					text="Create Post"
					btnText="Create"
					visible={this.state.idstrModal}
					onClose={() => this.setState({ idstrModal: false })}
					createPost={this.props.createPost}
				/>
				<TopicManageModal
					text="Manage Topics"
					forSelection={false}
					alreadySelected={[]}
					topics={this.props.topics}
					readTopics={this.props.readTopics}
					createTopic={this.props.createTopic}
					deleteTopic={this.props.deleteTopic}
					visible={this.state.topicVis}
					onClose={() => this.setState({ topicVis: false })}
				/>

				<div
					style={{
						minHeight: 'calc(100vh - 100px)',
						// maxHeight: 'calc(100vh - 100px)',
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
							<div>
								<button
									style={{ marginRight: '10px' }}
									className="Theme-Btn"
									onClick={() => this.setState({ topicVis: true })}
								>
									Topics
								</button>
								<button className="Theme-Btn" onClick={() => this.setState({ idstrModal: true })}>
									Create Post
								</button>
							</div>
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

const mapStateToProps = ({ posts, topics }) => ({ posts, topics });

const mapDispatchToProps = (dispatch) => ({
	createPost: (...args) => dispatch(createPost(...args)),
	readPosts: (...args) => dispatch(readPosts(...args)),
	readTopics: (...args) => dispatch(readTopics(...args)),
	createTopic: (...args) => dispatch(createTopic(...args)),
	editTopic: (...args) => dispatch(editTopic(...args)),
	deleteTopic: (...args) => dispatch(deleteTopic(...args))
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
