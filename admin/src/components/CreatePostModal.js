import React, { useState } from 'react';
import Rodal from 'rodal';
import IDStrInput from './IDStrInput';

import 'rodal/lib/rodal.css';

const CreatePostModal = (props) => {
	const [ valid, setValid ] = useState(false);
	const [ idstr, setIdstr ] = useState('');
	const [ loading, setLoading ] = useState(false);
	const [ errorMsg, setErrorMsg ] = useState('');

	const closeModal = () => {
		setValid(false);
		setIdstr('');
		setLoading(false);
		setErrorMsg('');
		props.onClose();
	};

	return (
		<Rodal
			visible={props.visible}
			onClose={closeModal}
			showCloseButton={false}
			className="Rodal-ClassName"
			customStyles={{
				height: 'auto',
				width: 'auto',
				bottom: 'auto',
				top: 'auto',
				left: 'auto',
				right: 'auto',
				padding: '0px',
				backgroundColor: 'rgba(0,0,0,0)'
			}}
		>
			<div
				style={{
					// width: '240px',
					background: 'white',
					borderRadius: '4px',
					padding: '20px'
				}}
			>
				<h4
					style={{
						margin: '0px 0px 10px 0px',
						padding: '0px'
					}}
				>
					{props.text}
				</h4>
				<IDStrInput
					idstr={idstr}
					setIdstr={setIdstr}
					valid={valid}
					setValid={setValid}
					loading={loading}
					setLoading={setLoading}
					errorMsg={errorMsg}
					setErrorMsg={setErrorMsg}
				/>
				{!!errorMsg && (
					<p style={{ color: '#ea4335', fontSize: '13px', margin: '10px 0px 0px 0px', padding: '0px' }}>
						{errorMsg}
					</p>
				)}
				<button
					style={{ marginTop: '10px' }}
					disabled={!valid || loading}
					className="Theme-Btn"
					onClick={async () => {
						if (idstr.length > 0) {
							try {
								await props.createPost(idstr);
								closeModal();
								const el = document.querySelector('.PostList-Comp-Div');
								el.scrollTop = el.scrollHeight;
							} catch (error) {
								setErrorMsg(error.message);
							}
						}
					}}
				>
					{props.btnText}
				</button>
			</div>
		</Rodal>
	);
};

export default CreatePostModal;
