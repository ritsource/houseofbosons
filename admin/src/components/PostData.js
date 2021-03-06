import React, { useState } from 'react';

import CreatePostModal from './CreatePostModal';
import TopicManageModal from './TopicManageModal';
import PostDataBool from './PostDataBool';
import PostDataInfo from './PostDataInfo';
import PostDataSrc from './PostDataSrc';
import PostDataSubPosts from './PostDataSubPosts';

const PostData = (props) => {
	const [ idstrModal, setIstrModal ] = useState(false);
	const [ topicModal, setTopicModal ] = useState(false);

	const editPost = (data) => {
		props.editPost(props.post._id, { ...props.post, ...data });
	};

	return (
		<div style={{ marginTop: '10px' }}>
			<CreatePostModal
				text="New ID-String"
				btnText="Save"
				visible={idstrModal}
				onClose={() => setIstrModal(false)}
				createPost={(text) => {
					editPost({ id_str: text });
				}}
			/>

			<h4 style={{ margin: '10px 0px 0px 0px', padding: '0px' }}>ID-String</h4>
			<div className="Flex-Row" style={{ width: '100%', margin: '10px 0px' }}>
				<input
					style={{ width: '100%' }}
					className="Theme-Input"
					value={props.post.id_str}
					onChange={() => {}}
				/>
				<button style={{ marginLeft: '10px' }} className="Theme-Btn" onClick={() => setIstrModal(true)}>
					Edit
				</button>
			</div>
			<PostDataInfo post={props.post} editPost={async (data) => await props.editPost(props.post._id, data)} />
			<PostDataSrc post={props.post} editPost={async (data) => await props.editPost(props.post._id, data)} />

			{props.post.topics && (
				<React.Fragment>
					<TopicManageModal
						text="Select Topics"
						forSelection={true}
						onSelection={async (selected) => {
							const topics = selected.map((t) => t.title);
							return new Promise(async (resolve, reject) => {
								try {
									await editPost({ topics });
									resolve();
								} catch (error) {
									reject(error);
								}
							});
						}}
						alreadySelected={props.post.topics}
						readTopics={props.readTopics}
						createTopic={props.createTopic}
						deleteTopic={props.deleteTopic}
						topics={props.topics || []}
						visible={topicModal}
						onClose={() => setTopicModal(false)}
					/>

					<h4 style={{ margin: '30px 0px 0px 0px', padding: '0px' }}>ID-String</h4>
					<div className="Flex-Row" style={{ width: '100%', margin: '18px 0px' }}>
						<button
							style={{ marginRight: '10px' }}
							className="Theme-Btn"
							onClick={() => setTopicModal(true)}
						>
							Manage
						</button>
						<div>
							{props.post.topics.map((t, i) => {
								return (
									<button
										key={i}
										className="Theme-Btn"
										style={{ margin: '5px', background: '#0090C3', border: '1px solid #0090C3' }}
									>
										{t}
									</button>
								);
							})}
						</div>
					</div>
				</React.Fragment>
			)}
			<PostDataBool post={props.post} editPost={async (data) => await props.editPost(props.post._id, data)} />
			{/* <PostDataSubPosts post={props.post} editPost={async (data) => await props.editPost(props.post._id, data)} /> */}
		</div>
	);
};

export default PostData;
