import React, { useState, useEffect } from 'react';

import DataEditBtns from './DataEditBtns';

const PostDataBool = (props) => {
	const [ editable, setEditable ] = useState(false);
	// const [ isFeatured, setIsFeatured ] = useState(props.is_featured);
	const [ isPublic, setIsPublic ] = useState(props.post.is_public);
	const [ isDeleted, setIsDeleted ] = useState(props.post.is_deleted);
	const [ isSeries, setIsSeries ] = useState(props.post.is_series);

	useEffect(
		() => {
			console.log(props.post.is_series);
			// resetState();
		},
		[ props.post ]
	);

	const resetState = () => {
		setIsPublic(props.post.is_public);
		setIsDeleted(props.post.is_deleted);
		console.log('x', props.post);
		console.log('f:', props.post.is_series);

		setIsSeries(props.post.is_series);
	};

	return (
		<div style={{ margin: '30px 0px 0px 0px' }}>
			<h4 style={{ margin: '0px 0px 10px 0px', padding: '0px' }}>Booleans</h4>
			<BoolInput
				text="Is Public"
				bool={isPublic}
				toggleBool={() => setIsPublic(editable ? !isPublic : isPublic)}
			/>
			<BoolInput
				text="Is Deleted"
				bool={isDeleted}
				toggleBool={() => setIsDeleted(editable ? !isDeleted : isDeleted)}
			/>
			<BoolInput
				text="Is Series"
				bool={isSeries}
				toggleBool={() => setIsSeries(editable ? !isSeries : isSeries)}
			/>
			<DataEditBtns
				onSave={() => {
					return new Promise(async (resolve, reject) => {
						try {
							await props.editPost({ is_public: isPublic, is_deleted: isDeleted, is_series: isSeries });
							setEditable(false);
							resolve();
						} catch (error) {
							reject(error);
						}
					});
				}}
				editable={editable}
				setEditable={(...args) => {
					resetState();
					setEditable(...args);
				}}
			/>
		</div>
	);
};

const indiBtnUp = {
	background: '#6441a5',
	border: '1px solid #6441a5',
	height: '16px',
	WebkitTransition: '0.1s linear',
	transition: '0.1s linear'
};
const indiBtnBottom = {
	border: '1px solid #6441a5',
	width: '20px',
	height: '8px',
	marginBottom: '-14px',
	borderRadius: '4px'
};

const BoolInput = (props) => (
	<div
		style={{
			display: 'flex',
			flexDirection: 'row',
			justifyContent: 'space-between',
			alignItems: 'center',
			maxWidth: '200px',
			margin: '0px',
			padding: '8px 0px'
		}}
	>
		<p style={{ margin: '0px', padding: '0px', fontSize: '14px' }}>{props.text}</p>
		<div style={{ cursor: 'pointer' }} onClick={props.toggleBool}>
			<div
				style={
					props.bool ? (
						{
							...indiBtnBottom,
							background: '#6cc644',
							border: '1px solid #6cc644'
						}
					) : (
						{
							...indiBtnBottom
						}
					)
				}
			/>
			<div
				style={
					props.bool ? (
						{
							...indiBtnUp,
							marginLeft: '10px',
							width: '10px',
							borderRadius: '4px',
							background: '#6cc644',
							border: '1px solid #6cc644'
						}
					) : (
						{
							...indiBtnUp,
							width: '10px',
							borderRadius: '4px'
						}
					)
				}
			/>
		</div>
	</div>
);

export default PostDataBool;
