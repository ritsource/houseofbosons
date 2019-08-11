import React, { useState } from 'react';

import DataEditBtns from './DataEditBtns';

const PostDataBool = (props) => {
	const [ editable, setEditable ] = useState(false);
	// const [ isFeatured, setIsFeatured ] = useState(props.is_featured);
	const [ isPublic, setIsPublic ] = useState(props.is_public);
	const [ isDeleted, setIsDeleted ] = useState(props.is_deleted);
	const [ isSeries, setIsSeries ] = useState(props.is_series);

	const resetState = () => {
		setIsPublic(props.is_public);
		setIsDeleted(props.is_deleted);
		setIsSeries(props.is_series);
	};

	return (
		<div style={{ margin: '20px 0px 0px 0px' }}>
			<h4 style={{ margin: '0px', padding: '0px' }}>Booleans</h4>
			<BoolInput
				text="Is Public"
				bool={isPublic}
				toggleBool={() => {
					if (editable) {
						setIsPublic(!isPublic);
					}
				}}
			/>
			<BoolInput
				text="Is Deleted"
				bool={isDeleted}
				toggleBool={() => {
					if (editable) {
						setIsDeleted(!isDeleted);
					}
				}}
			/>
			<BoolInput
				text="Is Series"
				bool={isSeries}
				toggleBool={() => {
					if (editable) {
						setIsSeries(!isSeries);
					}
				}}
			/>
			<DataEditBtns
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
	'-webkit-transition': '0.1s linear',
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
			margin: '10px 0px 0px 0px'
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
