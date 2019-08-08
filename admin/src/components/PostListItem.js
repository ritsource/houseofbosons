import React from 'react';

const PostListItem = (props) => {
	return (
		<div>
			<p>{props.post.title}</p>
		</div>
	);
};

export default PostListItem;
