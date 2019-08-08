import React from 'react';
import PostListItem from './PostListItem';

const PostList = (props) => {
	return (
		<div>
			{props.posts.map((post) => {
				return <PostListItem post={post} />;
			})}
		</div>
	);
};

export default PostList;
