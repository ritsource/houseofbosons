import React from 'react';

const Header = () => {
	return (
		<div style={{ position: 'fixed', top: '25px', left: '50px' }}>
			<a href="/">
				<div
					style={{
						cursor: 'pointer',
						width: '20px',
						height: '20px',
						borderRadius: '50%',
						background: '#6441a5'
					}}
				/>
			</a>
		</div>
	);
};

export default Header;
