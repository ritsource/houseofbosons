import React, { useState, useEffect } from 'react';

import CreatePostModal from './CreatePostModal';
import PostDataBool from './PostDataBool';
import PostDataInfo from './PostDataInfo';
import PostDataSrc from './PostDataSrc';

const PostData = (props) => {
	const [ modal, setModal ] = useState(false);

	const editPost = (data) => {
		props.editPost(props.post._id, data);
	};

	return (
		<div style={{ marginTop: '10px' }}>
			<CreatePostModal
				text="New ID-String"
				btnText="Save"
				visible={modal}
				onClose={() => setModal(false)}
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
				<button style={{ marginLeft: '10px' }} className="Theme-Btn" onClick={() => setModal(true)}>
					Edit
				</button>
			</div>
			<PostDataInfo post={props.post} />
			<PostDataSrc post={props.post} />
			<PostDataBool post={props.post} />
		</div>
	);
};

export default PostData;
