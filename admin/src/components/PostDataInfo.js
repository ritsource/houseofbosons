import React, { useState } from 'react';

import DataEditBtns from './DataEditBtns';

const PostDataInfo = (props) => {
	const [ editable, setEditable ] = useState(false);
	const [ title, setTitle ] = useState(props.post.title);
	const [ description, setDescription ] = useState(props.post.description);
	const [ formattedDate, setFormattedDate ] = useState(props.post.formatted_date);
	const [ thumbnail, setThumbnail ] = useState(props.post.thumbnail);

	const resetState = () => {
		setTitle(props.post.title);
		setDescription(props.post.description);
		setFormattedDate(props.post.formatted_date);
		setThumbnail(props.post.thumbnail);
	};

	return (
		<div style={{ margin: '30px 0px 0px 0px' }}>
			<h4 style={{ margin: '0px', padding: '0px' }}>Meta Data</h4>
			<InfoInput text={title} setText={setTitle} editable={editable} label="Title" placeholder="Post title" />
			<InfoInput
				text={description}
				setText={setDescription}
				editable={editable}
				label="Description"
				placeholder="Description text"
			/>
			<InfoInput
				text={formattedDate}
				setText={setFormattedDate}
				editable={editable}
				label="Date"
				placeholder="Formatted date text"
			/>
			<InfoInput
				text={thumbnail}
				setText={setThumbnail}
				editable={editable}
				label="Thumbnail"
				placeholder="Thumbnail source link"
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

export const InfoInput = (props) => (
	<div style={{ marginTop: '10px', display: 'flex', alignItems: 'center' }}>
		<p style={{ margin: '0px', width: '160px', padding: '0px', fontSize: '14px' }}>{props.label}</p>
		<input
			placeholder={props.placeholder}
			style={{ width: '100%' }}
			className="Theme-Input"
			value={props.text}
			onChange={(e) => {
				if (props.editable) {
					props.setText(e.target.value);
				}
			}}
		/>
	</div>
);

export default PostDataInfo;
