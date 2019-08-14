import React, { useEffect, useRef } from 'react';
import api from '../api';
import { FaCheck, FaTimes, FaFire } from 'react-icons/fa';

const IDStrInput = (props) => {
	const inputElm = useRef(null);

	const checkValid = async (text) => {
		try {
			const resp = await api.get('/post/idstr/available?idstr=' + text);
			console.log(text, resp.data);
			props.setValid(resp.data.available);
		} catch (error) {
			props.setErrorMsg(error.message);
		}
	};

	useEffect(() => {
		inputElm.current.focus();
	}, []);

	return (
		<div
			style={{
				display: 'flex',
				flexDirection: 'row',
				alignItems: 'center',
				width: '240px'
			}}
		>
			<input
				onChange={async (e) => {
					await props.setIdstr(e.target.value);

					if (props.idstr.length === 0) {
						props.setValid(false);
						props.setErrorMsg('');
					}

					if (!props.loading && props.idstr.length > 0) {
						await props.setLoading(true);
						setTimeout(async () => {
							await checkValid(props.idstr);
							props.setLoading(false);
						}, 1000);
					}
				}}
				ref={inputElm}
				placeholder="Unique ID-Str"
				style={{ width: '100%' }}
				className="Theme-Input"
				value={props.idstr}
			/>
			{props.loading ? (
				<div style={{ marginLeft: '10px' }} className="Theme-Loading-Spin-Div" />
			) : props.idstr.length > 0 ? !!props.errorMsg ? (
				<FaFire style={{ color: '#ea4335', marginLeft: '10px', fontSize: '20px' }} />
			) : props.valid ? (
				<FaCheck style={{ color: '#6cc644', marginLeft: '10px', fontSize: '20px' }} />
			) : (
				<FaTimes style={{ color: '#ea4335', marginLeft: '10px', fontSize: '20px' }} />
			) : (
				<div />
			)}
		</div>
	);
};

export default IDStrInput;
