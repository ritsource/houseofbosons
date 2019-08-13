import React, { useState } from 'react';
import { FaCheck } from 'react-icons/fa';

import DataEditBtns from './DataEditBtns';
import { InfoInput } from './PostDataInfo';
import { white } from 'ansi-colors';

const PostDataSrc = (props) => {
	const [ editable, setEditable ] = useState(false);
	const [ docType, setDocType ] = useState(props.post.doc_type);
	const [ mdSrc, setMdSrc ] = useState(props.post.md_src);
	const [ htmlSrc, setHtmlSrc ] = useState(props.post.html_src);

	const resetState = () => {
		setDocType(props.post.doc_type);
		setMdSrc(props.post.md_src);
		setHtmlSrc(props.post.html_src);
	};

	return (
		<div style={{ margin: '30px 0px 0px 0px' }}>
			<h4 style={{ margin: '0px', padding: '0px' }}>Meta Data</h4>

			<div style={{ marginTop: '10px', display: 'flex', alignItems: 'center' }}>
				<p style={{ margin: '0px', width: '127px', padding: '0px', fontSize: '14px' }}>Document Type</p>

				<div style={{ display: 'flex', alignItems: 'center' }}>
					<div style={{ display: 'flex', alignItems: 'center' }}>
						<DocTypeCheckbox
							onClick={() => {
								if (editable) {
									setDocType(0);
								}
							}}
							bool={docType === 0}
						/>
						<p style={{ fontSize: '14px', padding: '0px', margin: '0px 0px 0px 5px' }}>Markdown</p>
					</div>
					<div style={{ display: 'flex', alignItems: 'center', marginLeft: '10px' }}>
						<DocTypeCheckbox
							onClick={() => {
								if (editable) {
									setDocType(1);
								}
							}}
							bool={docType === 1}
						/>
						<p style={{ fontSize: '14px', padding: '0px', margin: '0px 0px 0px 5px' }}>HTML</p>
					</div>
				</div>
			</div>

			<InfoInput
				text={mdSrc}
				setText={setMdSrc}
				editable={editable}
				label="Markdown Source"
				placeholder="Markdown document source"
			/>
			<InfoInput
				text={htmlSrc}
				setText={setHtmlSrc}
				editable={editable}
				label="HTML Source"
				placeholder="HTML document source"
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

const checkboxstyle = {
	border: '1px solid #6441a5',
	display: 'flex',
	color: 'white',
	justifyContent: 'center',
	alignItems: 'center',
	padding: '3px',
	borderRadius: '4px',
	cursor: 'pointer'
};

const DocTypeCheckbox = (props) => (
	<div
		onClick={props.onClick}
		style={
			props.bool ? (
				{
					...checkboxstyle,
					background: '#6cc644',
					border: '1px solid #6cc644'
				}
			) : (
				{
					...checkboxstyle
				}
			)
		}
	>
		<FaCheck style={{ fontSize: '11px' }} />
	</div>
);

export default PostDataSrc;
