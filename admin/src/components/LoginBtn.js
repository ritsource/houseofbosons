import React from 'react';

const LoginBtn = (props) => (
	<button
		style={{
			color: 'white',
			background: '#6441a5',
			border: '0px solid white',
			borderRadius: '4px',
			padding: '10px 20px',
			cursor: 'pointer'
		}}
	>
		{props.children}
	</button>
);

export default LoginBtn;
