import React, { useState } from 'react';
import { FaCheck } from 'react-icons/fa';

import DataEditBtns from './DataEditBtns';
import { InfoInput } from './PostDataInfo';

const PostDataSubPosts = (props) => {
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
			<h4 style={{ margin: '0px', padding: '0px' }}>Series, Sub-Posts</h4>
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
				onSave={() => {
					return new Promise(async (resolve, reject) => {
						try {
							await props.editPost({ title, description, formatted_date: formattedDate, thumbnail });
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

export default PostDataSubPosts;
