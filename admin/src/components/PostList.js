import React from 'react';
import PostListItem from './PostListItem';

const PostList = (props) => {
	return (
		<div
			className="PostList-Comp-Div"
			style={{
				// overflow: 'auto'
			}}
		>
			{props.posts.map((post, i) => {
				return <PostListItem key={i} post={post} />;
			})}
		</div>
	);
};

export default PostList;
