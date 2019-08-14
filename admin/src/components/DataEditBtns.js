import React, { useState } from 'react';

const DataEditBtns = (props) => {
	const [ errorMsg, setErrorMsg ] = useState('');
	const [ loading, setLoading ] = useState(false);

	return (
		<React.Fragment>
			{errorMsg && (
				<p style={{ fontSize: '14px', margin: '20px 0px 10px 0px', padding: '0px', color: '#ea4335' }}>
					{errorMsg}
				</p>
			)}
			{props.editable ? (
				<div style={{ marginTop: '10px' }}>
					<button
						style={{ marginLeft: '0px' }}
						className="Theme-Btn Theme-Btn-Grey"
						onClick={() => props.setEditable(false)}
					>
						Cancel
					</button>
					<button
						style={{ marginLeft: '10px' }}
						className="Theme-Btn"
						onClick={async () => {
							await setLoading(true);

							try {
								await props.onSave();
							} catch (error) {
								setErrorMsg(error.meessage);
							}

							setLoading(false);
						}}
					>
						{loading ? 'Saving..' : 'Save'}
					</button>
				</div>
			) : (
				<button style={{ marginTop: '10px' }} className="Theme-Btn" onClick={() => props.setEditable(true)}>
					Edit
				</button>
			)}
		</React.Fragment>
	);
};

export default DataEditBtns;
