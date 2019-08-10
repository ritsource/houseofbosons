import React, { useState } from 'react';
import Rodal from 'rodal';
import IDStrInput from './IDStrInput';

import 'rodal/lib/rodal.css';

const CreatePostModal = (props) => {
	const [ valid, setValid ] = useState(false);
	const [ idstr, setIdstr ] = useState('');
	const [ loading, setLoading ] = useState('');
	const [ errorMsg, setErrorMsg ] = useState('');

	return (
		<Rodal
			visible={props.visible}
			onClose={props.onClose}
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
					<p
						style={{
							color: 'red',
							fontSize: '13px',
							margin: '10px 0px 0px 0px',
							padding: '0px'
						}}
					>
						{errorMsg}
					</p>
				)}
				<button
					style={{ marginTop: '10px' }}
					disabled={!valid && loading}
					className="Theme-Btn"
					onClick={() => {
						if (idstr.length > 0) {
							props.createPost(idstr);
						}
					}}
				>
					Create
				</button>
			</div>
		</Rodal>
	);
};

export default CreatePostModal;
