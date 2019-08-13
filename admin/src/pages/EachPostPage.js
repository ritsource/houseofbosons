import React from 'react';
import { connect } from 'react-redux';

import { readSinglePost, editPost, deletePostPerm, deletePostTemp } from '../actions/active_post_actions';
import { readTopics, createTopic, editTopic, deleteTopic } from '../actions/topic_actions';

import PostData from '../components/PostData';

class EachPostPage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			loading: true
		};
	}

	async componentDidMount() {
		await this.props.readSinglePost(this.props.match.params.postid);
		this.setState({ loading: false });
	}

	componentWillUnmount() {}

	render() {
		return (
			<div
				style={{
					// maxHeight: 'calc(100vh - 100px)',
					padding: '50px',
					width: '700px',
					// border: '1px solid red',
					display: 'flex',
					flexDirection: 'column'
				}}
			>
				{this.state.loading ? (
					<h3>Loading...</h3>
				) : this.props.post ? (
					<React.Fragment>
						<h2
							style={{
								lineHeight: '28px',
								height: '28px',
								padding: '0px',
								margin: '0px',
								overflow: 'hidden',
								whiteSpace: 'nowrap',
								textOverflow: 'ellipsis',
								textAlign: 'center'
							}}
						>
							{this.props.post.title}
						</h2>
						<PostData
							post={this.props.post}
							editPost={this.props.editPost}
							deletePostPerm={this.props.deletePostPerm}
							deletePostTemp={this.props.deletePostTemp}
							topics={this.props.topics}
							readTopics={this.props.readTopics}
							createTopic={this.props.createTopic}
							deleteTopic={this.props.deleteTopic}
						/>
					</React.Fragment>
				) : (
					<h3>Something went wrong</h3>
				)}
			</div>
		);
	}
}

const mapStateToProps = ({ activePost, topics }) => ({ post: activePost, topics });

const mapDispatchToProps = (dispatch) => ({
	readSinglePost: (...args) => dispatch(readSinglePost(...args)),
	editPost: (...args) => dispatch(editPost(...args)),
	deletePostTemp: (...args) => dispatch(deletePostTemp(...args)),
	deletePostPerm: (...args) => dispatch(deletePostPerm(...args)),
	readTopics: (...args) => dispatch(readTopics(...args)),
	createTopic: (...args) => dispatch(createTopic(...args)),
	editTopic: (...args) => dispatch(editTopic(...args)),
	deleteTopic: (...args) => dispatch(deleteTopic(...args))
});

export default connect(mapStateToProps, mapDispatchToProps)(EachPostPage);
