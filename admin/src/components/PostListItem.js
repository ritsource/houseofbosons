import React from 'react';
import { Link } from 'react-router-dom';
import { FaLock, FaLockOpen, FaTrash } from 'react-icons/fa';

const PostListItem = (props) => {
	return (
		<div
			style={{
				display: 'flex',
				flexDirection: 'row',
				justifyContent: 'space-between',
				alignItems: 'center',
				// border: '1px solid red',
				padding: '8px 0px'
			}}
		>
			<Link to={'/post/' + props.post._id}>
				<p
					style={{
						fontSize: '14px',
						height: '16px',
						lineHeight: '16px',
						padding: '0px',
						margin: '0px',
						overflow: 'hidden',
						whiteSpace: 'nowrap',
						textOverflow: 'ellipsis',
						maxWidth: '500px'
					}}
				>
					{props.post.title}
				</p>
			</Link>
			<div>
				{props.post.is_deleted && (
					<FaTrash
						title="This post has been temporarily deleted"
						style={{ color: '#ea4335', margin: '0px 20px 0px 0px' }}
					/>
				)}
				{props.post.is_public ? (
					<FaLock
						title="Accessible with only admin credentials"
						style={{ color: '#6441a5', margin: '0px 20px 0px 0px' }}
					/>
				) : (
					<FaLockOpen
						title="Publically accessible"
						style={{ color: '#6cc644', margin: '0px 20px 0px 0px' }}
					/>
				)}
			</div>
		</div>
	);
};

export default PostListItem;
